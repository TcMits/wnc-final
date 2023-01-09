package employee

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	EmployeeListUseCase struct {
		repoList repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput]
	}
	EmployeeUpdateUseCase struct {
		repoUpdate repository.UpdateModelRepository[*model.Employee, *model.EmployeeUpdateInput]
	}
	EmployeeGetFirstUseCase struct {
		eLUC usecase.IEmployeeListUseCase
	}
)

func NewEmployeeGetFirstUseCase(
	repoList repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
) usecase.IEmployeeGetFirstUseCase {
	uc := &EmployeeGetFirstUseCase{
		eLUC: repoList,
	}
	return uc
}

func NewEmployeeListUseCase(
	repoList repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
) usecase.IEmployeeListUseCase {
	uc := &EmployeeListUseCase{
		repoList: repoList,
	}
	return uc
}
func NewEmployeeUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Employee, *model.EmployeeUpdateInput],
) usecase.IEmployeeUpdateUseCase {
	return &EmployeeUpdateUseCase{
		repoUpdate: repoUpdate,
	}
}

func (uc *EmployeeGetFirstUseCase) GetFirst(ctx context.Context, o *model.EmployeeOrderInput, w *model.EmployeeWhereInput) (*model.Employee, error) {
	l, of := 1, 0
	entities, err := uc.eLUC.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (uc *EmployeeListUseCase) List(ctx context.Context, limit, offset *int, o *model.EmployeeOrderInput, w *model.EmployeeWhereInput) ([]*model.Employee, error) {
	entities, err := uc.repoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.employee.EmployeeListUseCase.List: %s", err))
	}
	return entities, nil
}

func (s *EmployeeUpdateUseCase) Update(ctx context.Context, e *model.Employee, i *model.EmployeeUpdateInput) (*model.Employee, error) {
	e, err := s.repoUpdate.Update(ctx, e, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.employee.employee.EmployeeUpdateUseCase.Update: %s", err))
	}
	return e, nil
}
