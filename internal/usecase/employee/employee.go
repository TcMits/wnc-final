package employee

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/outliers"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/google/uuid"
)

type (
	EmployeeGetUserUseCase struct {
		gFUC usecase.IEmployeeGetFirstUseCase
	}
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

func NewEmployeeGetUserUseCase(
	repoList repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
) usecase.IEmployeeGetUserUseCase {
	uc := &EmployeeGetUserUseCase{
		gFUC: NewEmployeeGetFirstUseCase(repoList),
	}
	return uc
}

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
func (s *EmployeeGetUserUseCase) GetUser(ctx context.Context, input map[string]any) (any, error) {
	usernameAny, ok := input["username"]
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("username is required"))
	}
	username, ok := usernameAny.(string)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("wrong type of username, expected type of string, not %T", username))
	}
	u, err := s.gFUC.GetFirst(ctx, nil, &model.EmployeeWhereInput{
		Username: &username,
	})
	if err != nil {
		return nil, err
	}
	return u, nil
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

type (
	AdminEmployeeListUseCase struct {
		r repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput]
	}
	AdminEmployeeUpdateUseCase struct {
		r repository.UpdateModelRepository[*model.Employee, *model.EmployeeUpdateInput]
	}
	AdminEmployeeCreateUseCase struct {
		r repository.CreateModelRepository[*model.Employee, *model.EmployeeCreateInput]
	}
	AdminEmployeeDeleteUseCase struct {
		r repository.DeleteModelRepository[*model.Employee]
	}
	AdminEmployeeGetFirstUseCase struct {
		uc1 usecase.IAdminEmployeeListUseCase
	}
	AdminEmployeeValidateCreateUseCase struct {
		uc1 usecase.IAdminEmployeeGetFirstUseCase
	}
	AdminEmployeeValidateUpdateUseCase struct {
		uc1 usecase.IAdminEmployeeGetFirstUseCase
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
	entites, err := s.r.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.employee.implementations.AdminEmployeeListUseCase.List: %s", err))
	}
	return entites, nil
}
func (s *AdminEmployeeDeleteUseCase) Delete(ctx context.Context, e *model.Employee) error {
	err := s.r.Delete(ctx, e)
	if err != nil {
		return usecase.WrapError(fmt.Errorf("internal.usecase.employee.implementations.AdminEmployeeDeleteUseCase.Delete: %s", err))
	}
	return nil
}
func (s *AdminEmployeeGetFirstUseCase) GetFirst(ctx context.Context, o *model.EmployeeOrderInput, w *model.EmployeeWhereInput) (*model.Employee, error) {
	entites, err := s.uc1.List(ctx, generic.GetPointer(1), generic.GetPointer(0), o, w)
	if err != nil {
		return nil, err
	}
	if len(entites) > 0 {
		return entites[0], nil
	}
	return nil, nil
}
func (s *AdminEmployeeCreateUseCase) Create(ctx context.Context, i *model.EmployeeCreateInput) (*model.Employee, error) {
	e, err := s.r.Create(ctx, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.employee.implementations.AdminEmployeeCreateUseCase.Create: %s", err))
	}
	return e, nil
}
func (s *AdminEmployeeUpdateUseCase) Update(ctx context.Context, e *model.Employee, i *model.EmployeeUpdateInput) (*model.Employee, error) {
	e, err := s.r.Update(ctx, e, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.employee.implementations.AdminEmployeeUpdateUseCase.Update: %s", err))
	}
	return e, nil
}
func (s *AdminEmployeeValidateCreateUseCase) ValidateCreate(ctx context.Context, i *model.EmployeeCreateInput) (*model.EmployeeCreateInput, error) {
	if i == nil {
		return nil, usecase.ValidationError(fmt.Errorf("invalid input"))
	}
	e, err := s.uc1.GetFirst(ctx, nil, &model.EmployeeWhereInput{
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
	e, err := s.uc1.GetFirst(ctx, nil, &model.EmployeeWhereInput{
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

func NewAdminEmployeeUseCase(
	r1 repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
	r2 repository.CreateModelRepository[*model.Employee, *model.EmployeeCreateInput],
	r3 repository.UpdateModelRepository[*model.Employee, *model.EmployeeUpdateInput],
	r4 repository.DeleteModelRepository[*model.Employee],
	r5 repository.IIsNextModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
	secretKey,
	prodOwnerName *string,
) usecase.IAdminEmployeeUseCase {
	return &AdminEmployeeUseCase{
		IAdminConfigUseCase:                 config.NewAdminConfigUseCase(secretKey, prodOwnerName),
		IAdminGetUserUseCase:                NewEmployeeGetUserUseCase(r1),
		IAdminEmployeeListUseCase:           &AdminEmployeeListUseCase{r: r1},
		IAdminEmployeeGetFirstUseCase:       &AdminEmployeeGetFirstUseCase{uc1: &AdminEmployeeListUseCase{r1}},
		IAdminEmployeeCreateUseCase:         &AdminEmployeeCreateUseCase{r: r2},
		IAdminEmployeeValidateCreateUseCase: &AdminEmployeeValidateCreateUseCase{uc1: &AdminEmployeeGetFirstUseCase{uc1: &AdminEmployeeListUseCase{r1}}},
		IAdminEmployeeUpdateUseCase:         &AdminEmployeeUpdateUseCase{r: r3},
		IAdminEmployeeValidateUpdateUseCase: &AdminEmployeeValidateUpdateUseCase{uc1: &AdminEmployeeGetFirstUseCase{uc1: &AdminEmployeeListUseCase{r1}}},
		IAdminEmployeeDeleteUseCase:         &AdminEmployeeDeleteUseCase{r4},
		IIsNextUseCase:                      outliers.NewIsNextUseCase(r5),
	}
}
