package debt

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/ent/debt"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
)

func (s *CustomerDebtListUseCase) List(ctx context.Context, limit, offset *int, o *model.DebtOrderInput, w *model.DebtWhereInput) ([]*model.Debt, error) {
	return s.repoList.List(ctx, limit, offset, o, w)
}

func (s *CustomerDebtCreateUseCase) Create(ctx context.Context, i *model.DebtCreateInput) (*model.Debt, error) {
	return s.repoCreate.Create(ctx, i)
}

func (s *CustomerDebtValidateCreateInputUseCase) Validate(ctx context.Context, i *model.DebtCreateInput) (*model.DebtCreateInput, error) {
	ownerBA, err := s.bAGFUC.GetFirst(ctx, nil, &model.BankAccountWhereInput{ID: generic.GetPointer(i.OwnerID)})
	if err != nil {
		return nil, err
	}
	if ownerBA == nil {
		return nil, usecase.WrapError(fmt.Errorf("invalid owner"))
	}
	owner, err := s.cGFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{ID: generic.GetPointer(ownerBA.CustomerID)})
	if err != nil {
		return nil, err
	}
	if owner == nil {
		return nil, usecase.WrapError(fmt.Errorf("invalid owner"))
	}
	receiverBA, err := s.bAGFUC.GetFirst(ctx, nil, &model.BankAccountWhereInput{ID: generic.GetPointer(i.ReceiverID)})
	if err != nil {
		return nil, err
	}
	if receiverBA == nil {
		return nil, usecase.WrapError(fmt.Errorf("invalid receiver"))
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
	i.OwnerBankName = *s.cfUC.GetProductOwnerName()
	i.OwnerName = owner.GetName()
	i.ReceiverBankAccountNumber = receiverBA.AccountNumber
	i.ReceiverBankName = *s.cfUC.GetProductOwnerName()
	i.ReceiverName = receiver.GetName()
	return i, nil
}
