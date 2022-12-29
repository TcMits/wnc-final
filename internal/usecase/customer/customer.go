package customer

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	CustomerListUseCase struct {
		repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput]
	}
	CustomerUpdateUseCase struct {
		repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput]
	}
	CustomerGetFirstUseCase struct {
		gFUC usecase.ICustomerListUseCase
	}
)

func NewCustomerGetFirstUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerGetFirstUseCase {
	uc := &CustomerGetFirstUseCase{
		gFUC: repoList,
	}
	return uc
}

func NewCustomerUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput],
) usecase.ICustomerUpdateUseCase {
	uc := &CustomerUpdateUseCase{
		repoUpdate: repoUpdate,
	}
	return uc
}

func NewCustomerListUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerListUseCase {
	uc := &CustomerListUseCase{
		repoList: repoList,
	}
	return uc
}

func (uc *CustomerGetFirstUseCase) GetFirst(ctx context.Context, o *model.CustomerOrderInput, w *model.CustomerWhereInput) (*model.Customer, error) {
	l, of := 1, 0
	entities, err := uc.gFUC.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (uc *CustomerListUseCase) List(ctx context.Context, limit, offset *int, o *model.CustomerOrderInput, w *model.CustomerWhereInput) ([]*model.Customer, error) {
	entities, err := uc.repoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.customer.List: %s", err))
	}
	return entities, nil
}

func (uc *CustomerUpdateUseCase) Update(ctx context.Context, m *model.Customer, i *model.CustomerUpdateInput) (*model.Customer, error) {
	return uc.repoUpdate.Update(ctx, m, i)
}
