package debt

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/ent/debt"
	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
	"github.com/TcMits/wnc-final/pkg/tool/template"
)

func (s *CustomerDebtListUseCase) List(ctx context.Context, limit, offset *int, o *model.DebtOrderInput, w *model.DebtWhereInput) ([]*model.Debt, error) {
	entites, err := s.RepoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.debt.implementations.CustomerDebtListUseCase.List: %s", err))
	}
	return entites, nil
}
func (s *CustomerDebtUpdateUseCase) Update(ctx context.Context, e *model.Debt, i *model.DebtUpdateInput) (*model.Debt, error) {
	e, err := s.RepoUpdate.Update(ctx, e, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.debt.implementations.CustomerDebtUpdateUseCase.Update: %s", err))
	}
	return e, nil
}

func (s *CustomerDebtCreateUseCase) Create(ctx context.Context, i *model.DebtCreateInput) (*model.Debt, error) {
	entity, err := s.RepoCreate.Create(ctx, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.debt.implementations.CustomerDebtCreateUseCase.Create: %s", err))
	}
	receiver, err := s.UC1.GetFirst(ctx, nil, &model.CustomerWhereInput{
		HasBankAccountsWith: []*model.BankAccountWhereInput{{ID: entity.ReceiverID}},
	})
	if err != nil {
		return nil, err
	}
	err = s.TaskExecutor.ExecuteTask(ctx, &task.DebtNotifyPayload{
		UserID: receiver.ID,
		ID:     entity.ID,
		Event:  sse.DebtCreated,
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.debt.implementations.CustomerDebtCreateUseCase.Create: %s", err))
	}
	return entity, nil
}

func (s *CustomerDebtValidateCreateInputUseCase) ValidateCreate(ctx context.Context, i *model.DebtCreateInput) (*model.DebtCreateInput, error) {
	user := usecase.GetUserAsCustomer(ctx)
	ownerBA, err := s.UC2.GetFirst(ctx, nil, &model.BankAccountWhereInput{
		IsForPayment: generic.GetPointer(true),
		CustomerID:   generic.GetPointer(user.ID),
	})
	if err != nil {
		return nil, err
	}
	if ownerBA == nil {
		return nil, usecase.ValidationError(fmt.Errorf("invalid owner"))
	}
	owner := user
	receiverBA, err := s.UC2.GetFirst(ctx, nil, &model.BankAccountWhereInput{
		ID:           generic.GetPointer(i.ReceiverID),
		IsForPayment: generic.GetPointer(true),
	})
	if err != nil {
		return nil, err
	}
	if receiverBA == nil {
		return nil, usecase.ValidationError(fmt.Errorf("invalid receiver"))
	}
	receiver, err := s.UC3.GetFirst(ctx, nil, &model.CustomerWhereInput{ID: generic.GetPointer(receiverBA.CustomerID)})
	if err != nil {
		return nil, err
	}
	if receiver == nil {
		return nil, usecase.ValidationError(fmt.Errorf("invalid receiver"))
	}
	i.Status = generic.GetPointer(debt.StatusPending)
	i.OwnerBankAccountNumber = ownerBA.AccountNumber
	i.OwnerID = ownerBA.ID
	i.OwnerBankName = *s.UC1.GetProductOwnerName()
	i.OwnerName = owner.GetName()
	i.ReceiverBankAccountNumber = receiverBA.AccountNumber
	i.ReceiverBankName = *s.UC1.GetProductOwnerName()
	i.ReceiverName = receiver.GetName()
	return i, nil
}

func (s *CustomerDebtListMineUseCase) ListMine(ctx context.Context, limit, offset *int, o *model.DebtOrderInput, w *model.DebtWhereInput) ([]*model.Debt, error) {
	user := usecase.GetUserAsCustomer(ctx)
	if w == nil {
		w = new(model.DebtWhereInput)
	}
	w.Or = []*model.DebtWhereInput{
		{HasReceiverWith: []*model.BankAccountWhereInput{{CustomerID: &user.ID}}},
		{HasOwnerWith: []*model.BankAccountWhereInput{{CustomerID: &user.ID}}},
	}
	return s.UC1.List(ctx, limit, offset, o, w)
}

func (s *CustomerDebtGetFirstMineUseCase) GetFirstMine(ctx context.Context, o *model.DebtOrderInput, w *model.DebtWhereInput) (*model.Debt, error) {
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

func (s *CustomerDebtCancelUseCase) getTarget(ctx context.Context, e *model.Debt) (*model.Customer, error) {
	target, err := s.UC2.GetFirst(ctx, nil, &model.CustomerWhereInput{
		HasBankAccountsWith: []*model.BankAccountWhereInput{{ID: e.OwnerID}},
	})
	if err != nil {
		return nil, err
	}
	user := usecase.GetUserAsCustomer(ctx)
	if target.ID == user.ID {
		target, err = s.UC2.GetFirst(ctx, nil, &model.CustomerWhereInput{
			HasBankAccountsWith: []*model.BankAccountWhereInput{{ID: e.ReceiverID}},
		})
		if err != nil {
			return nil, err
		}
	}
	return target, nil
}
func (s *CustomerDebtValidateCancelUseCase) ValidateCancel(ctx context.Context, e *model.Debt, i *model.DebtUpdateInput) (*model.DebtUpdateInput, error) {
	user := usecase.GetUserAsCustomer(ctx)
	owner, err := s.UC1.GetFirst(ctx, nil, &model.CustomerWhereInput{
		HasBankAccountsWith: []*model.BankAccountWhereInput{{ID: e.OwnerID}}})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.debt.implementations.CustomerDebtValidateCancelUseCase.ValidateCancel: %s", err))
	}
	if owner.ID != user.ID {
		if e.Status.String() != debt.StatusPending.String() {
			return nil, usecase.ValidationError(fmt.Errorf("cannot cancel %s debt", e.Status.String()))
		}
	}
	if i == nil {
		i = new(model.DebtUpdateInput)
	}
	i.Status = generic.GetPointer(debt.StatusCancelled)
	return i, nil
}
func (s *CustomerDebtCancelUseCase) Cancel(ctx context.Context, e *model.Debt, i *model.DebtUpdateInput) (*model.Debt, error) {
	e, err := s.UC1.Update(ctx, e, i)
	if err != nil {
		return nil, err
	}
	target, err := s.getTarget(ctx, e)
	if err != nil {
		return nil, err
	}
	err = s.TaskExecutor.ExecuteTask(ctx, &task.DebtNotifyPayload{
		UserID: target.ID,
		ID:     e.ID,
		Event:  sse.DebtCanceled,
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.debt.implementations.CustomerDebtCancelUseCase.Cancel: %s", err))
	}
	return e, nil
}
func (s *CustomerDebtValidateFulfillUseCase) ValidateFulfill(ctx context.Context, e *model.Debt) error {
	if e.Status == debt.StatusPending {
		owner, err := s.UC1.GetFirst(ctx, nil, &model.CustomerWhereInput{HasBankAccountsWith: []*model.BankAccountWhereInput{{ID: e.OwnerID}}})
		if err != nil {
			return usecase.WrapError(fmt.Errorf("internal.usecase.logic.debt.implementations.CustomerDebtValidateFulfillUseCase.ValidateFulfill: %s", err))
		}
		if owner == nil {
			return usecase.ValidationError(fmt.Errorf("invalid owner"))
		}
		user := usecase.GetUserAsCustomer(ctx)
		if user.ID != owner.ID {
			bA, err := s.UC2.GetFirst(ctx, nil, &model.BankAccountWhereInput{
				ID: e.ReceiverID,
			})
			if err != nil {
				return usecase.WrapError(fmt.Errorf("internal.usecase.logic.debt.implementations.CustomerDebtValidateFulfillUseCase.ValidateFulfill: %s", err))
			}
			aM, _ := e.Amount.Float64()
			ok, err := bA.IsBalanceSufficient(aM)
			if err != nil {
				return usecase.WrapError(fmt.Errorf("internal.usecase.logic.debt.implementations.CustomerDebtValidateFulfillUseCase.ValidateFulfill: %s", err))
			}
			if ok {
				return nil
			}
			return usecase.ValidationError(fmt.Errorf("insufficient ballence"))
		}
		return usecase.ValidationError(fmt.Errorf("cannot fulfill debt which you created"))
	}
	return usecase.ValidationError(fmt.Errorf("cannot fulfill %s debt", e.Status.String()))
}
func (s *CustomerDebtFulfillUseCase) Fulfill(ctx context.Context, e *model.Debt) (*model.DebtFulfillResp, error) {
	otp := usecase.GenerateOTP(6)
	fmt.Println(otp)
	otpHashValue, err := usecase.GenerateHashInfo(usecase.MakeOTPValue(ctx, otp, e.ID.String()))
	if err != nil {
		return nil, err
	}
	tk, err := usecase.GenerateFulfillToken(
		ctx,
		otpHashValue,
		*s.UC1.GetSecret(),
		s.OtpTimeout,
	)
	if err != nil {
		return nil, err
	}
	user := usecase.GetUserAsCustomer(ctx)
	msg, err := template.RenderFileToStr(*s.DebtFulfillMailTemp, map[string]string{
		"otp":     otp,
		"name":    user.GetName(),
		"expires": fmt.Sprintf("%.0f", s.OtpTimeout.Minutes()),
	}, ctx)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.debt.implementations.CustomerDebtFulfillUseCase.Fulfill: %s", err))
	}
	err = s.TaskExecutor.ExecuteTask(ctx, &mail.EmailPayload{
		Subject: *s.DebtFulfillSubjectMail,
		Message: *msg,
		To:      []string{user.Email},
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.debt.implementations.CustomerDebtFulfillUseCase.Fulfill: %s", err))
	}
	return &model.DebtFulfillResp{
		Token: tk,
	}, nil
}

func (s *CustomerDebtValidateFulfillWithTokenUseCase) ValidateFulfillWithToken(ctx context.Context, e *model.Debt, i *model.DebtFulfillWithTokenInput) (*model.DebtFulfillWithTokenInput, error) {
	pl, err := usecase.ParseToken(ctx, i.Token, *s.UC1.GetSecret())
	if err != nil {
		return nil, usecase.ValidationError(fmt.Errorf("token expired"))
	}
	tkAny, ok := pl["token"]
	if !ok {
		return nil, usecase.ValidationError(fmt.Errorf("invalid token"))
	}
	tk, ok := tkAny.(string)
	if !ok {
		return nil, usecase.ValidationError(fmt.Errorf("invalid token"))
	}
	err = usecase.ValidateHashInfo(usecase.MakeOTPValue(ctx, i.Otp, e.ID.String()), tk)
	if err != nil {
		return nil, usecase.ValidationError(fmt.Errorf("invalid otp"))
	}
	return i, nil
}

func (s *CustomerDebtFulfillWithTokenUseCase) getOwner(ctx context.Context, e *model.Debt) (*model.Customer, error) {
	owner, err := s.UC1.GetFirst(ctx, nil, &model.CustomerWhereInput{
		HasBankAccountsWith: []*model.BankAccountWhereInput{{ID: e.OwnerID}},
	})
	if err != nil {
		return nil, err
	}
	return owner, nil
}
func (s *CustomerDebtFulfillWithTokenUseCase) FulfillWithToken(ctx context.Context, e *model.Debt, i *model.DebtFulfillWithTokenInput) (*model.Debt, error) {
	e, err := s.RepoFulfill.Fulfill(ctx, e, nil)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.debt.implementations.CustomerDebtFulfillWithTokenUseCase.FulfillWithToken: %s", err))
	}
	owner, err := s.getOwner(ctx, e)
	if err != nil {
		return nil, err
	}
	err = s.TaskExecutor.ExecuteTask(ctx, &task.DebtNotifyPayload{
		UserID: owner.ID,
		ID:     e.ID,
		Event:  sse.DebtFulfilled,
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.debt.implementations.CustomerDebtFulfillWithTokenUseCase.FulfillWithToken: %s", err))
	}
	return e, nil
}

func (s *CustomerDebtIsNextUseCase) IsNext(ctx context.Context, limit, offset int, o *model.DebtOrderInput, w *model.DebtWhereInput) (bool, error) {
	user := usecase.GetUserAsCustomer(ctx)
	if w == nil {
		w = new(model.DebtWhereInput)
	}
	w.Or = []*model.DebtWhereInput{
		{HasReceiverWith: []*model.BankAccountWhereInput{{CustomerID: &user.ID}}},
		{HasOwnerWith: []*model.BankAccountWhereInput{{CustomerID: &user.ID}}},
	}
	return s.UC1.IsNext(ctx, limit, offset, o, w)
}
