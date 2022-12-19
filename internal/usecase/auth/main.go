package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/me"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/error/wrapper"
	"github.com/TcMits/wnc-final/pkg/tool/jwt"
	"github.com/TcMits/wnc-final/pkg/tool/password"
)

type (
	CustomerLoginUseCase struct {
		usecase.ICustomerGetUserUseCase
		secretKey  *string
		refreshTTL time.Duration
		accessTTL  time.Duration
		repoList   repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput]
	}
	CustomerValidateLoginInputUseCase struct {
		usecase.ICustomerGetUserUseCase
		repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput]
	}
	CustomerRenewAccessTokenUseCase struct {
		usecase.ICustomerGetUserUseCase
		secretKey  *string
		accessTTL  time.Duration
		repoList   repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput]
		repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput]
	}
	CustomerLogoutUseCase struct {
		repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput]
	}
	CustomerAuthUseCase struct {
		usecase.ICustomerGetUserUseCase
		usecase.ICustomerConfigUseCase
		*CustomerLoginUseCase
		*CustomerValidateLoginInputUseCase
		*CustomerRenewAccessTokenUseCase
		*CustomerLogoutUseCase
	}
)

func NewCustomerAuthUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput],
	secretKey *string,
	refreshTTL time.Duration,
	accessTTL time.Duration,
) usecase.ICustomerAuthUseCase {
	gUUc := me.NewCustomerGetUserUseCase(repoList)
	uc := CustomerAuthUseCase{
		gUUc,
		config.NewCustomerConfigUseCase(secretKey),
		&CustomerLoginUseCase{
			ICustomerGetUserUseCase: gUUc,
			repoList:                repoList,
			secretKey:               secretKey,
			refreshTTL:              refreshTTL,
			accessTTL:               accessTTL,
		},
		&CustomerValidateLoginInputUseCase{
			ICustomerGetUserUseCase: gUUc,
			repoList:                repoList,
		},
		&CustomerRenewAccessTokenUseCase{
			ICustomerGetUserUseCase: gUUc,
			secretKey:               secretKey,
			accessTTL:               accessTTL,
			repoList:                repoList,
			repoUpdate:              repoUpdate,
		},
		&CustomerLogoutUseCase{
			repoUpdate: repoUpdate,
		},
	}
	return uc
}

func invalidateToken(
	ctx context.Context,
	repo repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput],
	user *model.Customer,
) (*model.Customer, error) {
	user, err := repo.Update(ctx, user, &model.CustomerUpdateInput{
		ClearJwtTokenKey: true,
	})
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	return user, nil
}

func (useCase *CustomerLoginUseCase) Login(ctx context.Context, input *model.CustomerLoginInput) (any, error) {
	entityAny, err := useCase.GetUser(ctx, map[string]any{"username": *input.Username})
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	entity := entityAny.(*model.Customer)
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	if !entity.IsActive {
		return nil, usecase.WrapError(fmt.Errorf("user is not active"))
	}
	payload := map[string]any{
		"username": entity.Username,
		"password": entity.Password,
		"jwt_key":  entity.JwtTokenKey,
	}
	tokenPair, err := jwt.NewTokenPair(
		*useCase.secretKey,
		payload,
		payload,
		useCase.accessTTL,
		useCase.refreshTTL,
	)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.auth.Login: %w", err))
	}
	return tokenPair, nil
}
func (useCase *CustomerValidateLoginInputUseCase) ValidateLoginInput(
	ctx context.Context,
	input *model.CustomerLoginInput,
) (*model.CustomerLoginInput, error) {
	entityAny, err := useCase.GetUser(ctx, map[string]any{"username": *input.Username})
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	entity := entityAny.(*model.Customer)
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	err = password.ValidatePassword(entity.Password, *input.Password)
	if err != nil {
		return nil, usecase.WrapError(wrapper.NewValidationError(fmt.Errorf("password is invalid")))
	}
	return input, nil
}

func (useCase *CustomerRenewAccessTokenUseCase) RenewToken(
	ctx context.Context,
	refreshToken *string,
) (any, error) {
	payload, err := jwt.ParseJWT(*refreshToken, *useCase.secretKey)
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	userAny, err := useCase.GetUser(ctx, map[string]any{"username": payload["username"]})
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	user := userAny.(*model.Customer)
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	_, err = invalidateToken(ctx, useCase.repoUpdate, user)
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	token, err := jwt.NewAccessToken(payload, *useCase.secretKey, useCase.accessTTL)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.auth.RenewToken: %w", err))
	}
	return &jwt.TokenPair{
		RefreshToken: refreshToken,
		AccessToken:  &token,
	}, nil
}

func (useCase *CustomerLogoutUseCase) Logout(
	ctx context.Context,
	user *model.Customer,
) error {
	_, err := invalidateToken(ctx, useCase.repoUpdate, user)
	if err != nil {
		return usecase.WrapError(err)
	}
	return nil
}
