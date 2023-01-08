package customer

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
)

type (
	CustomerListUseCase struct {
		repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput]
	}
	CustomerUpdateUseCase struct {
		repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput]
	}
	CustomerCreateUseCase struct {
		repoCreate repository.CreateModelRepository[*model.Customer, *model.CustomerCreateInput]
	}
	CustomerValidateCreateUseCase struct {
		cLUC usecase.ICustomerListUseCase
	}
	CustomerGetFirstUseCase struct {
		cLUC usecase.ICustomerListUseCase
	}

	EmployeeCustomerUseCase struct {
		usecase.IEmployeeConfigUseCase
		usecase.IEmployeeGetUserUseCase
		usecase.IEmployeeCustomerCreateUseCase
		usecase.IEmployeeCustomerValidateCreateUseCase
	}
)

func NewCustomerValidateCreateUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.IEmployeeCustomerValidateCreateUseCase {
	return &CustomerValidateCreateUseCase{
		cLUC: NewCustomerListUseCase(repoList),
	}
}

func NewCustomerCreateUseCase(
	repoCreate repository.CreateModelRepository[*model.Customer, *model.CustomerCreateInput],
) usecase.IEmployeeCustomerCreateUseCase {
	return &CustomerCreateUseCase{
		repoCreate: repoCreate,
	}
}

func NewCustomerGetFirstUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerGetFirstUseCase {
	uc := &CustomerGetFirstUseCase{
		cLUC: repoList,
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

func NewEmployeeCustomerUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	repoCreate repository.CreateModelRepository[*model.Customer, *model.CustomerCreateInput],
	rle repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
	secretKey *string,
	prodOwnerName *string,
) usecase.IEmployeeCustomerUseCase {
	return &EmployeeCustomerUseCase{
		IEmployeeCustomerCreateUseCase:         NewCustomerCreateUseCase(repoCreate),
		IEmployeeCustomerValidateCreateUseCase: NewCustomerValidateCreateUseCase(repoList),
		IEmployeeConfigUseCase:                 config.NewEmployeeConfigUseCase(secretKey, prodOwnerName),
	}
}

func (uc *CustomerGetFirstUseCase) GetFirst(ctx context.Context, o *model.CustomerOrderInput, w *model.CustomerWhereInput) (*model.Customer, error) {
	l, of := 1, 0
	entities, err := uc.cLUC.List(ctx, &l, &of, o, w)
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

func (s *CustomerValidateCreateUseCase) isExist(ctx context.Context, i *model.CustomerWhereInput) (bool, error) {
	es, err := s.cLUC.List(ctx, generic.GetPointer(1), generic.GetPointer(0), nil, i)
	if err != nil {
		return false, usecase.WrapError(fmt.Errorf("internal.usecase.customer.customer.CustomerValidateCreateUseCase.isExist: %s", err))
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
		return nil, usecase.WrapError(fmt.Errorf("username %s is in use", i.Username))
	}
	es, err = s.isExist(ctx, &model.CustomerWhereInput{
		Email: &i.Email,
	})
	if err != nil {
		return nil, err
	}
	if es {
		return nil, usecase.WrapError(fmt.Errorf("email %s is in use", i.Email))
	}
	es, err = s.isExist(ctx, &model.CustomerWhereInput{
		PhoneNumber: &i.PhoneNumber,
	})
	if err != nil {
		return nil, err
	}
	if es {
		return nil, usecase.WrapError(fmt.Errorf("phone number %s is in use", i.PhoneNumber))
	}
	i.Password = generic.GetPointer("12345678")
	return i, nil
}

func (s *CustomerCreateUseCase) Create(ctx context.Context, i *model.CustomerCreateInput) (*model.Customer, error) {
	return s.repoCreate.Create(ctx, i)
}
