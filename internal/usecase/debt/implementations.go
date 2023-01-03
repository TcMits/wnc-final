package debt

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/ent/debt"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
)

func (s *CustomerDebtListUseCase) List(ctx context.Context, limit, offset *int, o *model.DebtOrderInput, w *model.DebtWhereInput) ([]*model.Debt, error) {
	entites, err := s.repoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.debt.implementations.CustomerDebtListUseCase.List: %s", err))
	}
	return entites, nil
}
func (s *CustomerDebtUpdateUseCase) Update(ctx context.Context, e *model.Debt, i *model.DebtUpdateInput) (*model.Debt, error) {
	e, err := s.repoUpdate.Update(ctx, e, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.debt.implementations.CustomerDebtUpdateUseCase.Update: %s", err))
	}
	return e, nil
}

func (s *CustomerDebtCreateUseCase) Create(ctx context.Context, i *model.DebtCreateInput) (*model.Debt, error) {
	entity, err := s.repoCreate.Create(ctx, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.debt.implementations.CustomerDebtCreateUseCase.Create: %s", err))
	}
	receiver, err := s.cGFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{
		HasBankAccountsWith: []*model.BankAccountWhereInput{{ID: entity.ReceiverID}},
	})
	if err != nil {
		return nil, err
	}
	err = s.taskExecutor.ExecuteTask(ctx, &task.DebtNotifyPayload{
		UserID: receiver.ID,
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.debt.implementations.CustomerDebtCreateUseCase.Create: %s", err))
	}
	return entity, nil
}

func (s *CustomerDebtValidateCreateInputUseCase) ValidateCreate(ctx context.Context, i *model.DebtCreateInput) (*model.DebtCreateInput, error) {
	user := usecase.GetUserAsCustomer(ctx)
	ownerBA, err := s.bAGFUC.GetFirst(ctx, nil, &model.BankAccountWhereInput{
		IsForPayment: generic.GetPointer(true),
		CustomerID:   generic.GetPointer(user.ID),
	})
	if err != nil {
		return nil, err
	}
	if ownerBA == nil {
		return nil, usecase.WrapError(fmt.Errorf("invalid owner"))
	}
	if !ownerBA.IsForPayment {
		return nil, usecase.WrapError(fmt.Errorf("owner not for payment"))
	}
	owner := user
	receiverBA, err := s.bAGFUC.GetFirst(ctx, nil, &model.BankAccountWhereInput{ID: generic.GetPointer(i.ReceiverID)})
	if err != nil {
		return nil, err
	}
	if receiverBA == nil {
		return nil, usecase.WrapError(fmt.Errorf("invalid receiver"))
	}
	if !receiverBA.IsForPayment {
		return nil, usecase.WrapError(fmt.Errorf("receiver not for payment"))
	}
	receiver, err := s.cGFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{ID: generic.GetPointer(receiverBA.CustomerID)})
	if err != nil {
		return nil, err
	}
	if receiver == nil {
		return nil, usecase.WrapError(fmt.Errorf("invalid receiver"))
	}
	i.Status = generic.GetPointer(debt.StatusPending)
	i.OwnerBankAccountNumber = ownerBA.AccountNumber
	i.OwnerID = ownerBA.ID
	i.OwnerBankName = *s.cfUC.GetProductOwnerName()
	i.OwnerName = owner.GetName()
	i.ReceiverBankAccountNumber = receiverBA.AccountNumber
	i.ReceiverBankName = *s.cfUC.GetProductOwnerName()
	i.ReceiverName = receiver.GetName()
	return i, nil
}

func (uc *CustomerDebtListMineUseCase) ListMine(ctx context.Context, limit, offset *int, o *model.DebtOrderInput, w *model.DebtWhereInput) ([]*model.Debt, error) {
	user := usecase.GetUserAsCustomer(ctx)
	if w == nil {
		w = new(model.DebtWhereInput)
	}
	w.Or = []*model.DebtWhereInput{
		{HasReceiverWith: []*model.BankAccountWhereInput{{CustomerID: &user.ID}}},
		{HasOwnerWith: []*model.BankAccountWhereInput{{CustomerID: &user.ID}}},
	}
	return uc.dLUC.List(ctx, limit, offset, o, w)
}

func (uc *CustomerDebtGetFirstMineUseCase) GetFirstMine(ctx context.Context, o *model.DebtOrderInput, w *model.DebtWhereInput) (*model.Debt, error) {
	l, of := 1, 0
	entities, err := uc.dLMUC.ListMine(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (uc *CustomerDebtCancelUseCase) getTarget(ctx context.Context, e *model.Debt) (*model.Customer, error) {
	target, err := uc.cGFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{
		HasBankAccountsWith: []*model.BankAccountWhereInput{{ID: e.OwnerID}},
	})
	if err != nil {
		return nil, err
	}
	user := usecase.GetUserAsCustomer(ctx)
	if target.ID == user.ID {
		target, err = uc.cGFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{
			HasBankAccountsWith: []*model.BankAccountWhereInput{{ID: e.ReceiverID}},
		})
		if err != nil {
			return nil, err
		}
	}
	return target, nil
}
func (uc *CustomerDebtValidateCancelUseCase) ValidateCancel(ctx context.Context, e *model.Debt, i *model.DebtUpdateInput) (*model.DebtUpdateInput, error) {
	user := usecase.GetUserAsCustomer(ctx)
	owner, err := uc.cGFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{
		HasBankAccountsWith: []*model.BankAccountWhereInput{{ID: e.OwnerID}}})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.debt.implementations.CustomerDebtValidateCancelUseCase.ValidateCancel: %s", err))
	}
	if owner.ID != user.ID {
		if e.Status.String() != debt.StatusPending.String() {
			return nil, usecase.WrapError(fmt.Errorf("cannot cancel %s debt", e.Status.String()))
		}
	}
	if i == nil {
		i = new(model.DebtUpdateInput)
	}
	i.Status = generic.GetPointer(debt.StatusCancelled)
	return i, nil
}
func (uc *CustomerDebtCancelUseCase) Cancel(ctx context.Context, e *model.Debt, i *model.DebtUpdateInput) (*model.Debt, error) {
	e, err := uc.dUUc.Update(ctx, e, i)
	if err != nil {
		return nil, err
	}
	target, err := uc.getTarget(ctx, e)
	if err != nil {
		return nil, err
	}
	err = uc.taskExecutor.ExecuteTask(ctx, &task.DebtNotifyPayload{
		UserID: target.ID,
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.debt.implementations.CustomerDebtCancelUseCase.Cancel: %s", err))
	}
	return e, nil
}
func (uc *CustomerDebtValidateFulfillUseCase) ValidateFulfill(ctx context.Context, e *model.Debt, i *model.DebtUpdateInput) (*model.DebtUpdateInput, error) {
	if e.Status == debt.StatusPending {
		owner, err := uc.cGFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{HasBankAccountsWith: []*model.BankAccountWhereInput{{ID: e.OwnerID}}})
		if err != nil {
			return nil, usecase.WrapError(fmt.Errorf("internal.usecase.debt.implementations.CustomerDebtValidateFulfillUseCase.ValidateFulfill: %s", err))
		}
		if owner == nil {
			return nil, usecase.WrapError(fmt.Errorf("invalid owner"))
		}
		user := usecase.GetUserAsCustomer(ctx)
		if user.ID != owner.ID {
			bA, err := uc.bAGFUC.GetFirst(ctx, nil, &model.BankAccountWhereInput{
				ID: e.ReceiverID,
			})
			if err != nil {
				return nil, usecase.WrapError(fmt.Errorf("internal.usecase.debt.implementations.CustomerDebtValidateFulfillUseCase.ValidateFulfill: %s", err))
			}
			ok, err := bA.IsBalanceSufficient(e.Amount.Abs().InexactFloat64())
			if err != nil {
				return nil, usecase.WrapError(fmt.Errorf("internal.usecase.debt.implementations.CustomerDebtValidateFulfillUseCase.ValidateFulfill: %s", err))
			}
			if ok {
				if i == nil {
					i = new(model.DebtUpdateInput)
				}
				i.Status = generic.GetPointer(debt.StatusFulfilled)
				return i, nil
			}
			return nil, usecase.WrapError(fmt.Errorf("insufficient ballence"))
		}
		return nil, usecase.WrapError(fmt.Errorf("cannot fulfill debt which you created"))
	}
	return nil, usecase.WrapError(fmt.Errorf("cannot fulfill %s debt", e.Status.String()))
}
func (uc *CustomerDebtFulfillUseCase) getOwner(ctx context.Context, e *model.Debt) (*model.Customer, error) {
	owner, err := uc.cGFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{
		HasBankAccountsWith: []*model.BankAccountWhereInput{{ID: e.OwnerID}},
	})
	if err != nil {
		return nil, err
	}
	return owner, nil
}
func (uc *CustomerDebtFulfillUseCase) Fulfill(ctx context.Context, e *model.Debt, i *model.DebtUpdateInput) (*model.Debt, error) {
	e, err := uc.repoFulfill.Fulfill(ctx, e, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.debt.implementations.CustomerDebtFulfillUseCase.Fulfill: %s", err))
	}
	owner, err := uc.getOwner(ctx, e)
	if err != nil {
		return nil, err
	}
	err = uc.taskExecutor.ExecuteTask(ctx, &task.DebtNotifyPayload{
		UserID: owner.ID,
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.debt.implementations.CustomerDebtFulfillUseCase.Fulfill: %s", err))
	}
	return e, nil
}
