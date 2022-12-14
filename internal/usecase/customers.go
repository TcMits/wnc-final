package usecase

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/error/wrapper"
	"github.com/google/uuid"
)

type (
	CustomerCreateUseCase struct {
		repoCreate repository.CreateModelRepository[*model.Customer, *model.CustomerCreateInput]
	}
	CustomerListUseCase struct {
		repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput]
	}
	CustomerDetailUseCase struct {
		repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput]
	}
	CustomerDeleteUseCase struct {
		repoList   repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput]
		repoDelete repository.DeleteModelRepository[*model.Customer]
	}

	CustomerUseCase struct {
		*CustomerCreateUseCase
		*CustomerListUseCase
		*CustomerDetailUseCase
		*CustomerDeleteUseCase
	}
)

func NewCustomerUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	repoCreate repository.CreateModelRepository[*model.Customer, *model.CustomerCreateInput],
	repoDelete repository.DeleteModelRepository[*model.Customer],
) ICustomerUseCase {
	uc := CustomerUseCase{
		&CustomerCreateUseCase{repoCreate: repoCreate},
		&CustomerListUseCase{repoList: repoList},
		&CustomerDetailUseCase{repoList: repoList},
		&CustomerDeleteUseCase{repoList: repoList, repoDelete: repoDelete},
	}
	return uc
}
func (useCase *CustomerListUseCase) List(ctx context.Context, limit, offset *int) ([]*model.Customer, error) {
	entities, err := useCase.repoList.List(ctx, limit, offset, nil, nil)
	if err != nil {
		return nil, wrapError(fmt.Errorf("internal.usecase.customers.List: %w", err))
	}
	return entities, nil
}

func (useCase *CustomerCreateUseCase) Create(ctx context.Context, input *model.CustomerCreateInput) (*model.Customer, error) {
	entity, err := useCase.repoCreate.Create(ctx, input)
	if err != nil {
		return nil, wrapError(fmt.Errorf("internal.usecase.customers.Create: %w", err))
	}
	return entity, nil
}

func (useCase *CustomerDetailUseCase) Detail(ctx context.Context, id *uuid.UUID) (*model.Customer, error) {
	limit, offset := 1, 0
	entities, err := useCase.repoList.List(ctx, &limit, &offset, nil, &ent.CustomerWhereInput{ID: id})
	if err != nil {
		return nil, wrapError(fmt.Errorf("internal.usecase.customers.List: %w", err))
	}
	if len(entities) == 0 {
		return nil, wrapError(wrapper.NewNotFoundError(fmt.Errorf("entity does not exist")))
	}
	return entities[0], nil
}

func (useCase *CustomerDeleteUseCase) Delete(ctx context.Context, id *uuid.UUID) error {
	limit, offset := 1, 0
	entities, err := useCase.repoList.List(ctx, &limit, &offset, nil, &ent.CustomerWhereInput{ID: id})
	if err != nil {
		return wrapError(fmt.Errorf("internal.usecase.customers.Delete: %w", err))
	}
	if len(entities) == 0 {
		return wrapError(wrapper.NewNotFoundError(fmt.Errorf("entity does not exist")))
	}
	entity := entities[0]
	err = useCase.repoDelete.Delete(ctx, entity)
	if err != nil {
		return wrapError(fmt.Errorf("internal.usecase.customers.Delete: %w", err))
	}
	return nil
}
