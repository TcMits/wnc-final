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
	return s.repoList.List(ctx, limit, offset, o, w)
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
	err = s.taskExecutor.ExecuteTask(ctx, &task.DebtCreateNotifyPayload{
		UserID: receiver.ID,
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.debt.implementations.CustomerDebtCreateUseCase.Create: %s", err))
	}
	return entity, nil
}

func (s *CustomerDebtValidateCreateInputUseCase) Validate(ctx context.Context, i *model.DebtCreateInput) (*model.DebtCreateInput, error) {
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
