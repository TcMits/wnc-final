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
	GetUserUseCase struct {
		repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput]
	}

	MeCustomerUseCase struct {
		usecase.ICustomerConfigUseCase
		*GetUserUseCase
	}
)

func NewMeCustomerUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	sk *string,
) usecase.IMeCustomerUseCase {
	uc := MeCustomerUseCase{
		config.NewCustomerConfigUseCase(sk),
		&GetUserUseCase{repoList: repoList},
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
		return eV, fmt.Errorf("internal.usecase.users.getUser: %w", err)
	}
	if len(entities) == 0 {
		var eV ModelType
		return eV, wrapper.NewNotFoundError(fmt.Errorf("entity does not exist"))
	}
	return entities[0], nil
}

func (useCase *GetUserUseCase) GetUser(ctx context.Context, input map[string]any) (any, error) {
	usernameAny := input["username"]
	username, ok := usernameAny.(string)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("wrong type of username, expected type of string, not %T", username))
	}
	u, err := getUser(ctx, useCase.repoList, &model.CustomerWhereInput{
		Username: &username,
	})
	return u, usecase.WrapError(err)
}
