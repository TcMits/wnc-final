package me

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/error/wrapper"
)

type (
	CustomerGetUserUseCase struct {
		repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput]
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
		repoList: repoList,
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

func getUser[ModelType, ModelOrderInput, ModelWhereInput any](
	ctx context.Context,
	repo repository.ListModelRepository[ModelType, ModelOrderInput, ModelWhereInput],
	wInput ModelWhereInput,
) (ModelType, error) {
	limit, offset := 1, 0
	var oInput ModelOrderInput
	entities, err := repo.List(ctx, &limit, &offset, oInput, wInput)
	if err != nil {
		var eV ModelType
		return eV, fmt.Errorf("internal.usecase.me.getUser: %w", err)
	}
	if len(entities) == 0 {
		var eV ModelType
		return eV, wrapper.NewNotFoundError(fmt.Errorf("entity does not exist"))
	}
	return entities[0], nil
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
	u, err := getUser(ctx, useCase.repoList, &model.CustomerWhereInput{
		Username: &username,
	})
	return u, usecase.WrapError(err)
}
