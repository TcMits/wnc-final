package transaction

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
	"github.com/TcMits/wnc-final/pkg/tool/template"
	"github.com/shopspring/decimal"
)

func (uc *CustomerTransactionUpdateUseCase) Update(ctx context.Context, e *model.Transaction, i *model.TransactionUpdateInput) (*model.Transaction, error) {
	e, err := uc.repoUpdate.Update(ctx, e, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.transaction.implementations.CustomerTransactionUpdateUseCase.Update: %s", err))
	}
	return e, nil
}
func (uc *CustomerTransactionValidateConfirmInputUseCase) ValidateConfirmInput(ctx context.Context, e *model.Transaction, i *model.TransactionConfirmUseCaseInput) error {
	if e.Status == transaction.StatusDraft {
		pl, err := usecase.ParseToken(ctx, i.Token, *uc.cfUC.GetSecret())
		if err != nil {
			return usecase.ValidationError(fmt.Errorf("token expired"))
		}
		iFPBMAny, ok := pl["is_fee_paid_by_me"]
		if !ok {
			return usecase.ValidationError(fmt.Errorf("invalid token"))
		}
		_, ok = iFPBMAny.(bool)
		if !ok {
			return usecase.ValidationError(fmt.Errorf("invalid token"))
		}
		tkAny, ok := pl["token"]
		if !ok {
			return usecase.ValidationError(fmt.Errorf("invalid token"))
		}
		tk, ok := tkAny.(string)
		if !ok {
			return usecase.ValidationError(fmt.Errorf("invalid token"))
		}
		err = usecase.ValidateHashInfo(usecase.MakeOTPValue(ctx, i.Otp), tk)
		if err != nil {
			return usecase.ValidationError(fmt.Errorf("invalid token"))
		}
		return nil
	}
	return usecase.ValidationError(fmt.Errorf("cannot confirm %s transaction", e.Status))
}
func (uc *CustomerTransactionConfirmSuccessUseCase) ConfirmSuccess(ctx context.Context, e *model.Transaction, token *string) (*model.Transaction, error) {
	if e.TransactionType == transaction.TransactionTypeInternal {
		ni := &model.TransactionCreateInput{
			SourceTransactionID: &e.ID,
			Amount:              decimal.NewFromFloat(*uc.cfUC.GetFeeAmount()),
			Status:              generic.GetPointer(transaction.StatusSuccess),
			TransactionType:     transaction.TransactionTypeInternal,
			Description:         uc.cfUC.GetFeeDesc(),
		}
		pl, _ := usecase.ParseToken(ctx, *token, *uc.cfUC.GetSecret())
		iFPBMAny := pl["is_fee_paid_by_me"]
		isFeePaidByMe, _ := iFPBMAny.(bool)
		if isFeePaidByMe {
			ni.SenderID = *e.SenderID
			ni.SenderBankAccountNumber = e.SenderBankAccountNumber
			ni.SenderBankName = e.SenderBankName
		} else {
			ni.SenderID = *e.ReceiverID
			ni.SenderBankAccountNumber = e.ReceiverBankAccountNumber
			ni.SenderBankName = e.ReceiverBankName
		}
		e, err := uc.tCRepo.ConfirmSuccess(ctx, e, ni)
		if err != nil {
			return nil, usecase.WrapError(fmt.Errorf("internal.usecase.transaction.implementations.CustomerTransactionConfirmSuccessUseCase.ConfirmSuccess: %s", err))
		}
		return e, nil
	}
	return nil, usecase.WrapError(fmt.Errorf("unhandled external transaction case"))
}

func (uc *CustomerTransactionCreateUseCase) Create(ctx context.Context, i *model.TransactionCreateUseCaseInput) (*model.TransactionCreateResp, error) {
	entity, err := uc.repoCreate.Create(ctx, i.TransactionCreateInput)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.transaction.implementations.CustomerTransactionCreateUseCase.Create: %s", err))
	}
	otp := usecase.GenerateOTP(6)
	otpHashValue, err := usecase.GenerateHashInfo(usecase.MakeOTPValue(ctx, otp))
	if err != nil {
		return nil, err
	}
	tk, err := usecase.GenerateConfirmTxcToken(
		ctx,
		otpHashValue,
		*uc.cfUC.GetSecret(),
		i.IsFeePaidByMe,
		uc.otpTimeout,
	)
	if err != nil {
		return nil, err
	}
	user := usecase.GetUserAsCustomer(ctx)
	msg, err := template.RenderToStr(*uc.txcConfirmMailTemp, map[string]string{
		"otp":     otp,
		"name":    user.GetName(),
		"expires": fmt.Sprintf("%.0f", uc.otpTimeout.Minutes()),
	}, ctx)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.transaction.implementations.CustomerTransactionCreateUseCase.Create: %s", err))
	}
	err = uc.taskExecutor.ExecuteTask(ctx, &mail.EmailPayload{
		Subject: *uc.txcConfirmSubjectMail,
		Message: *msg,
		To:      []string{user.Email},
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.transaction.implementations.CustomerTransactionCreateUseCase.Create: %s", err))
	}
	return &model.TransactionCreateResp{
		Transaction: entity,
		Token:       tk,
	}, nil
}

func (uc *CustomerTransactionValidateCreateInputUseCase) doesHaveDraftTxc(ctx context.Context, i *model.TransactionCreateInput) error {
	user := usecase.GetUserAsCustomer(ctx)
	entities, err := uc.tLUC.List(ctx, generic.GetPointer(1), generic.GetPointer(0), nil, &model.TransactionWhereInput{
		HasSenderWith: []*model.BankAccountWhereInput{{CustomerID: &user.ID}},
		Status:        generic.GetPointer(transaction.StatusDraft),
	})
	if err != nil {
		return usecase.WrapError(fmt.Errorf("internal.usecase.transaction.implementations.CustomerTransactionValidateCreateInputUseCase.doesHaveDraftTxc: %s", err))
	}
	if len(entities) > 0 {
		return usecase.ValidationError(fmt.Errorf("there is a draft transaction to be processed. Cannot create a new transaction"))
	}
	return nil
}

func (uc *CustomerTransactionValidateCreateInputUseCase) ValidateCreate(ctx context.Context, i *model.TransactionCreateUseCaseInput) (*model.TransactionCreateUseCaseInput, error) {
	user := usecase.GetUserAsCustomer(ctx)
	ba, err := uc.bAGFUC.GetFirst(ctx, nil, &model.BankAccountWhereInput{
		IsForPayment: generic.GetPointer(true),
		CustomerID:   generic.GetPointer(user.ID),
	})
	if err != nil {
		return nil, err
	}
	if ba == nil {
		return nil, usecase.ValidationError(fmt.Errorf("bank account sender is invalid"))
	}
	err = uc.doesHaveDraftTxc(ctx, i.TransactionCreateInput)
	if err != nil {
		return nil, err
	}
	if i.IsFeePaidByMe {
		if ok, err := ba.IsBalanceSufficient(i.Amount.Abs().InexactFloat64() + *uc.cfUC.GetFeeAmount()); err != nil {
			return nil, usecase.WrapError(fmt.Errorf("internal.usecase.transaction.implementations.CustomerTransactionValidateCreateInputUseCase.Validate: %s", err))
		} else if !ok {
			return nil, usecase.ValidationError(fmt.Errorf("insufficient balance sender"))
		}
	} else if ok, err := ba.IsBalanceSufficient(i.Amount.Abs().InexactFloat64()); err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.transaction.implementations.CustomerTransactionValidateCreateInputUseCase.Validate: %s", err))
	} else if !ok {
		return nil, usecase.ValidationError(fmt.Errorf("insufficient balance sender"))
	}
	if i.TransactionType == transaction.TransactionTypeInternal {
		baOther, err := uc.bAGFUC.GetFirst(ctx, nil, &model.BankAccountWhereInput{ID: i.ReceiverID})
		if err != nil {
			return nil, err
		}
		if baOther == nil {
			return nil, usecase.ValidationError(fmt.Errorf("bank account receiver is invalid"))
		}
		if !baOther.IsForPayment {
			return nil, usecase.ValidationError(fmt.Errorf("bank account receiver is not for payment"))
		}
		if !i.IsFeePaidByMe {
			if ok, err := baOther.IsBalanceSufficient(*uc.cfUC.GetFeeAmount()); err != nil {
				return nil, usecase.WrapError(fmt.Errorf("internal.usecase.transaction.implementations.CustomerTransactionValidateCreateInputUseCase.Validate: %s", err))
			} else if !ok {
				return nil, usecase.ValidationError(fmt.Errorf("insufficient balance receiver"))
			}
		}
		other, err := uc.cGFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{ID: &baOther.CustomerID})
		if err != nil {
			return nil, err
		}
		i.ReceiverBankAccountNumber = baOther.AccountNumber
		i.ReceiverBankName = *uc.cfUC.GetProductOwnerName()
		i.ReceiverName = other.GetName()
	}
	i.Status = generic.GetPointer(transaction.StatusDraft)
	i.SenderID = ba.ID
	i.SenderBankAccountNumber = ba.AccountNumber
	i.SenderBankName = *uc.cfUC.GetProductOwnerName()
	i.SenderName = user.GetName()
	return i, nil
}

func (uc *CustomerTransactionListUseCase) List(ctx context.Context, limit, offset *int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) ([]*model.Transaction, error) {
	entites, err := uc.repoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.bankaccount.implementations.CustomerTransactionListUseCase.List: %s", err))
	}
	return entites, nil
}

func (uc *CustomerTransactionListMineUseCase) ListMine(ctx context.Context, limit, offset *int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) ([]*model.Transaction, error) {
	user := usecase.GetUserAsCustomer(ctx)
	if w == nil {
		w = new(model.TransactionWhereInput)
	}
	w.Or = []*model.TransactionWhereInput{
		{HasReceiverWith: []*model.BankAccountWhereInput{{CustomerID: &user.ID}}},
		{HasSenderWith: []*model.BankAccountWhereInput{{CustomerID: &user.ID}}},
	}
	return uc.tLUC.List(ctx, limit, offset, o, w)
}

func (uc *CustomerTransactionGetFirstMineUseCase) GetFirstMine(ctx context.Context, o *model.TransactionOrderInput, w *model.TransactionWhereInput) (*model.Transaction, error) {
	l, of := 1, 0
	entities, err := uc.tLMTUC.ListMine(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (s *CustomerTransactionIsNextUseCase) IsNext(ctx context.Context, limit, offset int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) (bool, error) {
	user := usecase.GetUserAsCustomer(ctx)
	if w == nil {
		w = new(model.TransactionWhereInput)
	}
	w.Or = []*model.TransactionWhereInput{
		{HasReceiverWith: []*model.BankAccountWhereInput{{CustomerID: &user.ID}}},
		{HasSenderWith: []*model.BankAccountWhereInput{{CustomerID: &user.ID}}},
	}
	return s.iNUC.IsNext(ctx, limit, offset, o, w)
}

// employee
func (s *EmployeeTransactionListUseCase) List(ctx context.Context, limit, offset *int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) ([]*model.Transaction, error) {
	entites, err := s.repoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.bankaccount.implementations.EmployeeTransactionListUseCase.List: %s", err))
	}
	return entites, nil
}
func (s *EmployeeTransactionGetFirstUseCase) GetFirst(ctx context.Context, o *model.TransactionOrderInput, w *model.TransactionWhereInput) (*model.Transaction, error) {
	l, of := 1, 0
	entities, err := s.tLTUC.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (s *EmployeeTransactionIsNextUseCase) IsNext(ctx context.Context, limit, offset int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) (bool, error) {
	return s.iNUC.IsNext(ctx, limit, offset, o, w)
}

// admin
func (s *AdminTransactionListUseCase) List(ctx context.Context, limit, offset *int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) ([]*model.Transaction, error) {
	entites, err := s.repoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.bankaccount.implementations.AdminTransactionListUseCase.List: %s", err))
	}
	return entites, nil
}
func (s *AdminTransactionGetFirstUseCase) GetFirst(ctx context.Context, o *model.TransactionOrderInput, w *model.TransactionWhereInput) (*model.Transaction, error) {
	l, of := 1, 0
	entities, err := s.tLTUC.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (s *AdminTransactionIsNextUseCase) IsNext(ctx context.Context, limit, offset int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) (bool, error) {
	return s.iNUC.IsNext(ctx, limit, offset, o, w)
}
