package customer

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/employee"
	"github.com/TcMits/wnc-final/internal/usecase/outliers"
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
	CustomerGetUserUseCase struct {
		gFUC usecase.ICustomerGetFirstUseCase
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

func NewCustomerGetUserUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerGetUserUseCase {
	uc := &CustomerGetUserUseCase{
		gFUC: NewCustomerGetFirstUseCase(repoList),
	}
	return uc
}
func NewCustomerUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerUseCase {
	return &CustomerUseCase{
		ICustomerConfigUseCase:   config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		ICustomerGetUserUseCase:  NewCustomerGetUserUseCase(repoList),
		ICustomerListUseCase:     NewCustomerListUseCase(repoList),
		ICustomerGetFirstUseCase: NewCustomerGetFirstUseCase(repoList),
	}
}

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
	repoIsNext repository.IIsNextModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	rle repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
	secretKey *string,
	prodOwnerName *string,
) usecase.IEmployeeCustomerUseCase {
	return &EmployeeCustomerUseCase{
		IEmployeeCustomerCreateUseCase:         NewCustomerCreateUseCase(repoCreate),
		IEmployeeCustomerValidateCreateUseCase: NewCustomerValidateCreateUseCase(repoList),
		IEmployeeConfigUseCase:                 config.NewEmployeeConfigUseCase(secretKey, prodOwnerName),
		IEmployeeCustomerListUseCase:           NewCustomerListUseCase(repoList),
		IIsNextUseCase:                         outliers.NewIsNextUseCase(repoIsNext),
		IEmployeeCustomerGetFirstUseCase:       NewCustomerGetFirstUseCase(repoList),
		IEmployeeGetUserUseCase:                employee.NewEmployeeGetUserUseCase(rle),
	}

}

func (useCase *CustomerGetUserUseCase) GetUser(ctx context.Context, input map[string]any) (any, error) {
	usernameAny, ok := input["username"]
	if !ok {
		return nil, usecase.ValidationError(fmt.Errorf("username is required"))
	}
	username, ok := usernameAny.(string)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("wrong type of username, expected type of string, not %T", username))
	}
	u, err := useCase.gFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{
		Or: []*model.CustomerWhereInput{
			{Username: &username}, {PhoneNumber: &username}, {Email: &username},
		},
	})
	if err != nil {
		return nil, err
	}
	return u, nil
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
	i.IsActive = generic.GetPointer(true)
	return i, nil
}

func (s *CustomerCreateUseCase) Create(ctx context.Context, i *model.CustomerCreateInput) (*model.Customer, error) {
	e, err := s.repoCreate.Create(ctx, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.customer.customer.CustomerCreateUseCase.Create: %s", err))
	}
	return e, nil
}
