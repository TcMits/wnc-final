package me

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/customer"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	CustomerGetUserUseCase struct {
		cGetter usecase.ICustomerGetFirstUseCase
	}

	CustomerMeUseCase struct {
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
	}
)

func NewCustomerGetUserUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerGetUserUseCase {
	uc := &CustomerGetUserUseCase{
		cGetter: customer.NewCustomerGetFirstUseCase(repoList),
	}
	return uc
}
func NewCustomerMeUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	sk *string,
) usecase.ICustomerMeUseCase {
	uc := CustomerMeUseCase{
		config.NewCustomerConfigUseCase(sk),
		NewCustomerGetUserUseCase(repoList),
	}
	return uc
}

func (useCase *CustomerGetUserUseCase) GetUser(ctx context.Context, input map[string]any) (any, error) {
	usernameAny, ok := input["username"]
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("username is required"))
	}
	username, ok := usernameAny.(string)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("wrong type of username, expected type of string, not %T", username))
	}
	u, err := useCase.cGetter.GetFirst(ctx, nil, &model.CustomerWhereInput{
		Username: &username,
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}
