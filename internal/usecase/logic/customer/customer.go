package customer

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/password"
)

type (
	CustomerListUseCase struct {
		RepoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput]
	}
	CustomerUpdateUseCase struct {
		RepoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput]
	}
	CustomerCreateUseCase struct {
		RepoCreate repository.CreateModelRepository[*model.EmployeeCreateCustomerResp, *model.CustomerCreateInput]
	}
	CustomerValidateCreateUseCase struct {
		UC1 usecase.ICustomerListUseCase
	}
	CustomerGetFirstUseCase struct {
		UC1 usecase.ICustomerListUseCase
	}
	CustomerGetUserUseCase struct {
		UC1 usecase.ICustomerGetFirstUseCase
	}
	CustomerUseCase struct {
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
		usecase.ICustomerListUseCase
		usecase.ICustomerGetFirstUseCase
	}

	EmployeeCustomerUseCase struct {
		usecase.IEmployeeConfigUseCase
		usecase.IEmployeeGetUserUseCase
		usecase.IEmployeeCustomerCreateUseCase
		usecase.IEmployeeCustomerValidateCreateUseCase
		usecase.IEmployeeCustomerListUseCase
		usecase.IEmployeeCustomerGetFirstUseCase
		usecase.IIsNextUseCase[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput]
	}
)

func (useCase *CustomerGetUserUseCase) GetUser(ctx context.Context, input map[string]any) (any, error) {
	usernameAny, ok := input["username"]
	if !ok {
		return nil, usecase.ValidationError(fmt.Errorf("username is required"))
	}
	username, ok := usernameAny.(string)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("wrong type of username, expected type of string, not %T", username))
	}
	u, err := useCase.UC1.GetFirst(ctx, nil, &model.CustomerWhereInput{
		Or: []*model.CustomerWhereInput{
			{Username: &username}, {PhoneNumber: &username}, {Email: &username},
		},
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}
func (s *CustomerGetFirstUseCase) GetFirst(ctx context.Context, o *model.CustomerOrderInput, w *model.CustomerWhereInput) (*model.Customer, error) {
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

func (s *CustomerListUseCase) List(ctx context.Context, limit, offset *int, o *model.CustomerOrderInput, w *model.CustomerWhereInput) ([]*model.Customer, error) {
	entities, err := s.RepoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.customer.List: %s", err))
	}
	return entities, nil
}

func (s *CustomerUpdateUseCase) Update(ctx context.Context, m *model.Customer, i *model.CustomerUpdateInput) (*model.Customer, error) {
	return s.RepoUpdate.Update(ctx, m, i)
}

func (s *CustomerValidateCreateUseCase) isExist(ctx context.Context, i *model.CustomerWhereInput) (bool, error) {
	es, err := s.UC1.List(ctx, generic.GetPointer(1), generic.GetPointer(0), nil, i)
	if err != nil {
		return false, usecase.WrapError(fmt.Errorf("internal.usecase.logic.customer.customer.CustomerValidateCreateUseCase.isExist: %s", err))
	}
	return len(es) > 0, nil
}

func (s *CustomerValidateCreateUseCase) ValidateCreate(ctx context.Context, i *model.CustomerCreateInput) (*model.CustomerCreateInput, error) {
	var es bool
	var err error
	es, err = s.isExist(ctx, &model.CustomerWhereInput{
		Username: &i.Username,
	})
	if err != nil {
		return nil, err
	}
	if es {
		return nil, usecase.ValidationError(fmt.Errorf("username %s is in use", i.Username))
	}
	es, err = s.isExist(ctx, &model.CustomerWhereInput{
		Email: &i.Email,
	})
	if err != nil {
		return nil, err
	}
	if es {
		return nil, usecase.ValidationError(fmt.Errorf("email %s is in use", i.Email))
	}
	es, err = s.isExist(ctx, &model.CustomerWhereInput{
		PhoneNumber: &i.PhoneNumber,
	})
	if err != nil {
		return nil, err
	}
	if es {
		return nil, usecase.ValidationError(fmt.Errorf("phone number %s is in use", i.PhoneNumber))
	}
	hashPwd, err := password.GetHashPassword("12345678")
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.customer.customer.CustomerValidateCreateUseCase.ValidateCreate: %s", err))
	}
	i.Password = generic.GetPointer(hashPwd)
	i.IsActive = generic.GetPointer(true)
	return i, nil
}

func (s *CustomerCreateUseCase) Create(ctx context.Context, i *model.CustomerCreateInput) (*model.EmployeeCreateCustomerResp, error) {
	e, err := s.RepoCreate.Create(ctx, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.customer.customer.CustomerCreateUseCase.Create: %s", err))
	}
	return e, nil
}
