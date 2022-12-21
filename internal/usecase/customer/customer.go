package customer

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	CustomerGetFirstUseCase struct {
		repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput]
	}
)

func NewCustomerGetFirstUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerGetFirstUseCase {
	uc := &CustomerGetFirstUseCase{
		repoList: repoList,
	}
	return uc
}

func (uc *CustomerGetFirstUseCase) GetFirst(ctx context.Context, o *model.CustomerOrderInput, w *model.CustomerWhereInput) (*model.Customer, error) {
	l, of := 1, 0
	entities, err := uc.repoList.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.customer.GetFirst: %s", err))
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}
