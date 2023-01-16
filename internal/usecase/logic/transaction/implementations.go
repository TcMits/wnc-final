package transaction

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
	"github.com/TcMits/wnc-final/pkg/tool/password"
	"github.com/TcMits/wnc-final/pkg/tool/template"
	"github.com/shopspring/decimal"
)

func (s *CustomerTransactionUpdateUseCase) Update(ctx context.Context, e *model.Transaction, i *model.TransactionUpdateInput) (*model.Transaction, error) {
	e, err := s.RepoUpdate.Update(ctx, e, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.transaction.implementations.CustomerTransactionUpdateUseCase.Update: %s", err))
	}
	return e, nil
}
func (s *CustomerTransactionValidateConfirmInputUseCase) ValidateConfirmInput(ctx context.Context, e *model.Transaction, i *model.TransactionConfirmUseCaseInput) error {
	if e.Status == transaction.StatusDraft {
		pl, err := usecase.ParseToken(ctx, i.Token, *s.UC1.GetSecret())
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
func (s *CustomerTransactionConfirmSuccessUseCase) ConfirmSuccess(ctx context.Context, e *model.Transaction, token *string) (*model.Transaction, error) {
	if e.TransactionType == transaction.TransactionTypeInternal {
		ni := &model.TransactionCreateInput{
			SourceTransactionID: &e.ID,
			Amount:              decimal.NewFromFloat(*s.UC1.GetFeeAmount()),
			Status:              generic.GetPointer(transaction.StatusSuccess),
			TransactionType:     transaction.TransactionTypeInternal,
			Description:         s.UC1.GetFeeDesc(),
		}
		pl, _ := usecase.ParseToken(ctx, *token, *s.UC1.GetSecret())
		iFPBMAny := pl["is_fee_paid_by_me"]
		isFeePaidByMe, _ := iFPBMAny.(bool)
		if isFeePaidByMe {
			ni.SenderID = e.SenderID
			ni.SenderBankAccountNumber = e.SenderBankAccountNumber
			ni.SenderBankName = e.SenderBankName
		} else {
			ni.SenderID = e.ReceiverID
			ni.SenderBankAccountNumber = e.ReceiverBankAccountNumber
			ni.SenderBankName = e.ReceiverBankName
		}
		e, err := s.R.ConfirmSuccess(ctx, e, ni)
		if err != nil {
			return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.transaction.implementations.CustomerTransactionConfirmSuccessUseCase.ConfirmSuccess: %s", err))
		}
		return e, nil
	} else if e.TransactionType == transaction.TransactionTypeExternal {
		var ni *model.TransactionCreateInput
		pl, _ := usecase.ParseToken(ctx, *token, *s.UC1.GetSecret())
		iFPBMAny := pl["is_fee_paid_by_me"]
		isFeePaidByMe, _ := iFPBMAny.(bool)
		if isFeePaidByMe {
			ni = &model.TransactionCreateInput{
				SourceTransactionID:     &e.ID,
				Amount:                  decimal.NewFromFloat(*s.UC1.GetFeeAmount()),
				Status:                  generic.GetPointer(transaction.StatusSuccess),
				TransactionType:         transaction.TransactionTypeInternal,
				Description:             s.UC1.GetFeeDesc(),
				SenderID:                e.SenderID,
				SenderBankAccountNumber: e.SenderBankAccountNumber,
				SenderBankName:          e.SenderBankName,
			}
		}
		e, err := s.R.ConfirmSuccess(ctx, e, ni)
		if err != nil {
			return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.transaction.implementations.CustomerTransactionConfirmSuccessUseCase.ConfirmSuccess: %s", err))
		}
		return e, nil
	}
	return nil, usecase.ValidationError(fmt.Errorf("unexpected transaction type: %s", e.TransactionType.String()))
}

func (s *CustomerTransactionCreateUseCase) Create(ctx context.Context, i *model.TransactionCreateUseCaseInput) (*model.TransactionCreateResp, error) {
	entity, err := s.RepoCreate.Create(ctx, i.TransactionCreateInput)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.transaction.implementations.CustomerTransactionCreateUseCase.Create: %s", err))
	}
	otp := usecase.GenerateOTP(6)
	otpHashValue, err := usecase.GenerateHashInfo(usecase.MakeOTPValue(ctx, otp))
	if err != nil {
		return nil, err
	}
	tk, err := usecase.GenerateConfirmTxcToken(
		ctx,
		otpHashValue,
		*s.UC1.GetSecret(),
		i.IsFeePaidByMe,
		s.OtpTimeout,
	)
	if err != nil {
		return nil, err
	}
	user := usecase.GetUserAsCustomer(ctx)
	msg, err := template.RenderFileToStr(*s.TxcConfirmMailTemp, map[string]string{
		"otp":     otp,
		"name":    user.GetName(),
		"expires": fmt.Sprintf("%.0f", s.OtpTimeout.Minutes()),
	}, ctx)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.transaction.implementations.CustomerTransactionCreateUseCase.Create: %s", err))
	}
	err = s.TaskExecutor.ExecuteTask(ctx, &mail.EmailPayload{
		Subject: *s.TxcConfirmSubjectMail,
		Message: *msg,
		To:      []string{user.Email},
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.transaction.implementations.CustomerTransactionCreateUseCase.Create: %s", err))
	}
	return &model.TransactionCreateResp{
		Transaction: entity,
		Token:       tk,
	}, nil
}

func (s *CustomerTransactionValidateCreateInputUseCase) ValidateCreate(ctx context.Context, i *model.TransactionCreateUseCaseInput) (*model.TransactionCreateUseCaseInput, error) {
	user := usecase.GetUserAsCustomer(ctx)
	ba, err := s.UC2.GetFirst(ctx, nil, &model.BankAccountWhereInput{
		IsForPayment: generic.GetPointer(true),
		CustomerID:   generic.GetPointer(user.ID),
	})
	if err != nil {
		return nil, err
	}
	if ba == nil {
		return nil, usecase.ValidationError(fmt.Errorf("bank account sender is invalid"))
	}
	am, _ := i.Amount.Float64()
	if i.IsFeePaidByMe {
		if ok, err := ba.IsBalanceSufficient(am + *s.UC1.GetFeeAmount()); err != nil {
			return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.transaction.implementations.CustomerTransactionValidateCreateInputUseCase.Validate: %s", err))
		} else if !ok {
			return nil, usecase.ValidationError(fmt.Errorf("insufficient balance sender"))
		}
	} else if ok, err := ba.IsBalanceSufficient(am); err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.transaction.implementations.CustomerTransactionValidateCreateInputUseCase.Validate: %s", err))
	} else if !ok {
		return nil, usecase.ValidationError(fmt.Errorf("insufficient balance sender"))
	}
	i.Status = generic.GetPointer(transaction.StatusDraft)
	i.SenderID = &ba.ID
	i.SenderBankAccountNumber = ba.AccountNumber
	i.SenderBankName = *s.UC1.GetProductOwnerName()
	i.SenderName = user.GetName()
	if i.TransactionType == transaction.TransactionTypeInternal {
		baOther, err := s.UC2.GetFirst(ctx, nil, &model.BankAccountWhereInput{
			ID:           i.ReceiverID,
			IsForPayment: generic.GetPointer(true),
		})
		if err != nil {
			return nil, err
		}
		if baOther == nil {
			return nil, usecase.ValidationError(fmt.Errorf("bank account receiver is invalid"))
		}
		if !i.IsFeePaidByMe {
			if ok, err := baOther.IsBalanceSufficient(*s.UC1.GetFeeAmount()); err != nil {
				return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.transaction.implementations.CustomerTransactionValidateCreateInputUseCase.Validate: %s", err))
			} else if !ok {
				return nil, usecase.ValidationError(fmt.Errorf("insufficient balance receiver"))
			}
		}
		other, err := s.UC3.GetFirst(ctx, nil, &model.CustomerWhereInput{ID: &baOther.CustomerID})
		if err != nil {
			return nil, err
		}
		i.ReceiverBankAccountNumber = baOther.AccountNumber
		i.ReceiverBankName = *s.UC1.GetProductOwnerName()
		i.ReceiverName = other.GetName()
		i.TransactionType = transaction.TransactionTypeInternal
	} else {
		baOther, err := s.W1.Get(ctx, &model.WhereInputPartner{
			AccountNumber: i.ReceiverBankAccountNumber,
		})
		if err != nil {
			return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.transaction.implementations.CustomerTransactionValidateCreateInputUseCase.Validate: %s", err))
		}
		if baOther == nil {
			return nil, usecase.ValidationError(fmt.Errorf("bank account receiver is invalid"))
		}
		if !i.IsFeePaidByMe {
			am, _ := i.Amount.Abs().Float64()
			if am < *s.UC1.GetFeeAmount() {
				return nil, usecase.ValidationError(fmt.Errorf("insufficient balance sender"))
			}
			i.Amount = decimal.NewFromFloat(am - *s.UC1.GetFeeAmount())
		}
		i.ReceiverBankAccountNumber = baOther.AccountNumber
		i.ReceiverBankName = s.W1.GetName()
		i.ReceiverName = baOther.Name
		i.TransactionType = transaction.TransactionTypeExternal

		iP, err := s.W1.PreValidate(ctx, &model.TransactionCreateInputPartner{
			Amount:                    i.Amount,
			Description:               *i.Description,
			SenderName:                i.SenderName,
			SenderBankAccountNumber:   i.SenderBankAccountNumber,
			ReceiverBankAccountNumber: i.ReceiverBankAccountNumber,
		})
		if err != nil {
			return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.transaction.implementations.CustomerTransactionValidateCreateInputUseCase.Validate: %s", err))
		}
		err = s.W1.Validate(ctx, iP)
		if err != nil {
			return nil, usecase.ValidationError(fmt.Errorf("invalid data"))
		}
	}
	return i, nil
}

func (s *CustomerTransactionListUseCase) List(ctx context.Context, limit, offset *int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) ([]*model.Transaction, error) {
	entites, err := s.RepoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.bankaccount.implementations.CustomerTransactionListUseCase.List: %s", err))
	}
	return entites, nil
}

func (s *CustomerTransactionListMineUseCase) ListMine(ctx context.Context, limit, offset *int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) ([]*model.Transaction, error) {
	user := usecase.GetUserAsCustomer(ctx)
	if w == nil {
		w = new(model.TransactionWhereInput)
	}
	w.Or = []*model.TransactionWhereInput{
		{HasReceiverWith: []*model.BankAccountWhereInput{{CustomerID: &user.ID}}},
		{HasSenderWith: []*model.BankAccountWhereInput{{CustomerID: &user.ID}}},
	}
	return s.UC1.List(ctx, limit, offset, o, w)
}

func (s *CustomerTransactionGetFirstMineUseCase) GetFirstMine(ctx context.Context, o *model.TransactionOrderInput, w *model.TransactionWhereInput) (*model.Transaction, error) {
	l, of := 1, 0
	entities, err := s.UC1.ListMine(ctx, &l, &of, o, w)
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
	return s.UC1.IsNext(ctx, limit, offset, o, w)
}

// employee
func (s *EmployeeTransactionListUseCase) List(ctx context.Context, limit, offset *int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) ([]*model.Transaction, error) {
	entites, err := s.RepoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.bankaccount.implementations.EmployeeTransactionListUseCase.List: %s", err))
	}
	return entites, nil
}
func (s *EmployeeTransactionGetFirstUseCase) GetFirst(ctx context.Context, o *model.TransactionOrderInput, w *model.TransactionWhereInput) (*model.Transaction, error) {
	l, of := 1, 0
	entities, err := s.UC1.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (s *EmployeeTransactionIsNextUseCase) IsNext(ctx context.Context, limit, offset int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) (bool, error) {
	return s.UC1.IsNext(ctx, limit, offset, o, w)
}

// admin
func (s *AdminTransactionListUseCase) List(ctx context.Context, limit, offset *int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) ([]*model.Transaction, error) {
	entites, err := s.RepoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.bankaccount.implementations.AdminTransactionListUseCase.List: %s", err))
	}
	return entites, nil
}
func (s *AdminTransactionGetFirstUseCase) GetFirst(ctx context.Context, o *model.TransactionOrderInput, w *model.TransactionWhereInput) (*model.Transaction, error) {
	l, of := 1, 0
	entities, err := s.UC1.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (s *AdminTransactionIsNextUseCase) IsNext(ctx context.Context, limit, offset int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) (bool, error) {
	return s.UC1.IsNext(ctx, limit, offset, o, w)
}

func (s *PartnerTransactionValidateCreateInputUseCase) ValidateCreate(ctx context.Context, i *model.PartnerTransactionCreateInput) (*model.PartnerTransactionCreateInput, error) {
	ba, err := s.UC1.GetFirst(ctx, nil, &model.BankAccountWhereInput{
		IsForPayment:  generic.GetPointer(true),
		AccountNumber: &i.ReceiverBankAccountNumber,
	})
	if err != nil {
		return nil, err
	}
	if ba == nil {
		return nil, usecase.ValidationError(fmt.Errorf("account number is invalid"))
	}
	i.TransactionType = transaction.TransactionTypeExternal
	i.Status = generic.GetPointer(transaction.StatusDraft)
	receiver, err := s.UC3.GetFirst(ctx, nil, &model.CustomerWhereInput{ID: generic.GetPointer(ba.CustomerID)})
	if err != nil {
		return nil, err
	}
	data, err := template.RenderToStr(s.Layout, map[string]string{
		"receiver_bank_account_number": i.ReceiverBankAccountNumber,
		"sender_bank_account_number":   i.SenderBankAccountNumber,
		"sender_name":                  i.SenderName,
		"amount":                       i.Amount.String(),
		"description":                  *i.Description,
	}, ctx)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.transaction.implementations.PartnerTransactionValidateCreateInputUseCase.ValidateCreate: %s", err))
	}
	user := usecase.GetUserAsPartner(ctx)
	err = password.ValidateHashData(ctx, *data, user.SecretKey, i.Token)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("data is modified"))
	}
	err = password.VerifySignature(ctx, i.Signature, i.Token, user.PublicKey)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("verify signature failed"))
	}
	i.ReceiverID = generic.GetPointer(ba.ID)
	i.ReceiverName = receiver.GetName()
	i.ReceiverBankName = *s.UC2.GetProductOwnerName()
	i.SenderBankName = user.Name
	return i, nil
}
func (s *PartnerTransactionCreateUseCase) Create(ctx context.Context, i *model.PartnerTransactionCreateInput) (*model.Transaction, error) {
	e, err := s.Repo.Create(ctx, i.TransactionCreateInput)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.transaction.implementations.PartnerTransactionCreateUseCase.Create: %s", err))
	}
	return e, nil
}
