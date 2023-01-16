package employee

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/google/uuid"
)

type (
	EmployeeGetUserUseCase struct {
		UC1 usecase.IEmployeeGetFirstUseCase
	}
	EmployeeListUseCase struct {
		RepoList repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput]
	}
	EmployeeUpdateUseCase struct {
		RepoUpdate repository.UpdateModelRepository[*model.Employee, *model.EmployeeUpdateInput]
	}
	EmployeeGetFirstUseCase struct {
		UC1 usecase.IEmployeeListUseCase
	}
)

func (s *EmployeeGetUserUseCase) GetUser(ctx context.Context, input map[string]any) (any, error) {
	usernameAny, ok := input["username"]
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("username is required"))
	}
	username, ok := usernameAny.(string)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("wrong type of username, expected type of string, not %T", username))
	}
	u, err := s.UC1.GetFirst(ctx, nil, &model.EmployeeWhereInput{
		Username: &username,
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *EmployeeGetFirstUseCase) GetFirst(ctx context.Context, o *model.EmployeeOrderInput, w *model.EmployeeWhereInput) (*model.Employee, error) {
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

func (s *EmployeeListUseCase) List(ctx context.Context, limit, offset *int, o *model.EmployeeOrderInput, w *model.EmployeeWhereInput) ([]*model.Employee, error) {
	entities, err := s.RepoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.employee.EmployeeListUseCase.List: %s", err))
	}
	return entities, nil
}

func (s *EmployeeUpdateUseCase) Update(ctx context.Context, e *model.Employee, i *model.EmployeeUpdateInput) (*model.Employee, error) {
	e, err := s.RepoUpdate.Update(ctx, e, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.employee.employee.EmployeeUpdateUseCase.Update: %s", err))
	}
	return e, nil
}

type (
	AdminEmployeeListUseCase struct {
		R repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput]
	}
	AdminEmployeeUpdateUseCase struct {
		R repository.UpdateModelRepository[*model.Employee, *model.EmployeeUpdateInput]
	}
	AdminEmployeeCreateUseCase struct {
		R repository.CreateModelRepository[*model.Employee, *model.EmployeeCreateInput]
	}
	AdminEmployeeDeleteUseCase struct {
		R repository.DeleteModelRepository[*model.Employee]
	}
	AdminEmployeeGetFirstUseCase struct {
		UC1 usecase.IAdminEmployeeListUseCase
	}
	AdminEmployeeValidateCreateUseCase struct {
		UC1 usecase.IAdminEmployeeGetFirstUseCase
	}
	AdminEmployeeValidateUpdateUseCase struct {
		UC1 usecase.IAdminEmployeeGetFirstUseCase
	}
	AdminEmployeeUseCase struct {
		usecase.IAdminConfigUseCase
		usecase.IAdminGetUserUseCase
		usecase.IAdminEmployeeListUseCase
		usecase.IAdminEmployeeGetFirstUseCase
		usecase.IAdminEmployeeCreateUseCase
		usecase.IAdminEmployeeValidateCreateUseCase
		usecase.IAdminEmployeeUpdateUseCase
		usecase.IAdminEmployeeValidateUpdateUseCase
		usecase.IAdminEmployeeDeleteUseCase
		usecase.IIsNextUseCase[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput]
	}
)

func (s *AdminEmployeeListUseCase) List(ctx context.Context, limit, offset *int, o *model.EmployeeOrderInput, w *model.EmployeeWhereInput) ([]*model.Employee, error) {
	entites, err := s.R.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.employee.implementations.AdminEmployeeListUseCase.List: %s", err))
	}
	return entites, nil
}
func (s *AdminEmployeeDeleteUseCase) Delete(ctx context.Context, e *model.Employee) error {
	err := s.R.Delete(ctx, e)
	if err != nil {
		return usecase.WrapError(fmt.Errorf("internal.usecase.logic.employee.implementations.AdminEmployeeDeleteUseCase.Delete: %s", err))
	}
	return nil
}
func (s *AdminEmployeeGetFirstUseCase) GetFirst(ctx context.Context, o *model.EmployeeOrderInput, w *model.EmployeeWhereInput) (*model.Employee, error) {
	entites, err := s.UC1.List(ctx, generic.GetPointer(1), generic.GetPointer(0), o, w)
	if err != nil {
		return nil, err
	}
	if len(entites) > 0 {
		return entites[0], nil
	}
	return nil, nil
}
func (s *AdminEmployeeCreateUseCase) Create(ctx context.Context, i *model.EmployeeCreateInput) (*model.Employee, error) {
	e, err := s.R.Create(ctx, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.employee.implementations.AdminEmployeeCreateUseCase.Create: %s", err))
	}
	return e, nil
}
func (s *AdminEmployeeUpdateUseCase) Update(ctx context.Context, e *model.Employee, i *model.EmployeeUpdateInput) (*model.Employee, error) {
	e, err := s.R.Update(ctx, e, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.employee.implementations.AdminEmployeeUpdateUseCase.Update: %s", err))
	}
	return e, nil
}
func (s *AdminEmployeeValidateCreateUseCase) ValidateCreate(ctx context.Context, i *model.EmployeeCreateInput) (*model.EmployeeCreateInput, error) {
	if i == nil {
		return nil, usecase.ValidationError(fmt.Errorf("invalid input"))
	}
	e, err := s.UC1.GetFirst(ctx, nil, &model.EmployeeWhereInput{
		Username: &i.Username,
	})
	if err != nil {
		return nil, err
	}
	if e != nil {
		return nil, usecase.ValidationError(fmt.Errorf("username already existed"))
	}
	i.IsActive = generic.GetPointer(true)
	return i, nil
}
func (s *AdminEmployeeValidateUpdateUseCase) ValidateUpdate(ctx context.Context, e *model.Employee, i *model.EmployeeUpdateInput) (*model.EmployeeUpdateInput, error) {
	if i == nil {
		return nil, usecase.ValidationError(fmt.Errorf("invalid input"))
	}
	if i.Username == nil || *i.Username == "" {
		i.Username = &e.Username
	}
	e, err := s.UC1.GetFirst(ctx, nil, &model.EmployeeWhereInput{
		Username: i.Username,
		IDNotIn:  []uuid.UUID{e.ID},
	})
	if err != nil {
		return nil, err
	}
	if e != nil {
		return nil, usecase.ValidationError(fmt.Errorf("username already existed"))
	}
	return i, nil
}
