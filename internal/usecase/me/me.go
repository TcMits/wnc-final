package me

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/customer"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/password"
)

type (
	CustomerGetUserUseCase struct {
		gFUC usecase.ICustomerGetFirstUseCase
	}
	CustomerChangePasswordUseCase struct {
		cUUC usecase.ICustomerUpdateUseCase
	}
	CustomerValidateChangePasswordUseCase struct{}
	CustomerGetUserFromCtx                struct{}

	CustomerMeUseCase struct {
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
		usecase.ICustomerChangePasswordUseCase
		usecase.ICustomerValidateChangePasswordUseCase
		usecase.ICustomerGetUserFromCtxUseCase
	}
)

func NewCustomerChangePasswordUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput],
) usecase.ICustomerChangePasswordUseCase {
	return &CustomerChangePasswordUseCase{cUUC: customer.NewCustomerUpdateUseCase(repoUpdate)}
}
func NewCustomerValidateChangePasswordUseCase() usecase.ICustomerValidateChangePasswordUseCase {
	return &CustomerValidateChangePasswordUseCase{}
}
func NewCustomerGetUserFromCtxUserCase() usecase.ICustomerGetUserFromCtxUseCase {
	return &CustomerGetUserFromCtx{}
}

func NewCustomerGetUserUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerGetUserUseCase {
	uc := &CustomerGetUserUseCase{
		gFUC: customer.NewCustomerGetFirstUseCase(repoList),
	}
	return uc
}
func NewCustomerMeUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerMeUseCase {
	uc := &CustomerMeUseCase{
		ICustomerConfigUseCase:                 config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		ICustomerGetUserUseCase:                NewCustomerGetUserUseCase(repoList),
		ICustomerChangePasswordUseCase:         NewCustomerChangePasswordUseCase(repoUpdate),
		ICustomerValidateChangePasswordUseCase: NewCustomerValidateChangePasswordUseCase(),
		ICustomerGetUserFromCtxUseCase:         NewCustomerGetUserFromCtxUserCase(),
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

func (s *CustomerChangePasswordUseCase) ChangePassword(ctx context.Context, i *model.CustomerChangePasswordInput) (*model.Customer, error) {
	user := usecase.GetUserAsCustomer(ctx)
	user, err := s.cUUC.Update(ctx, user, &model.CustomerUpdateInput{
		ClearPassword: true,
		Password:      i.HashPwd,
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.auth.CustomerChangePasswordUseCase.ChangePassword: %w", err))
	}
	return user, nil
}

func (s *CustomerValidateChangePasswordUseCase) ValidateChangePassword(ctx context.Context, i *model.CustomerChangePasswordInput) (*model.CustomerChangePasswordInput, error) {
	user := usecase.GetUserAsCustomer(ctx)
	if err := password.ValidatePassword(user.Password, i.OldPassword); err != nil {
		return nil, usecase.WrapError(fmt.Errorf("old password is invalid"))
	}
	if i.Password == i.OldPassword {
		return nil, usecase.WrapError(fmt.Errorf("new password match old password is not allowed"))
	}
	if i.Password != i.ConfirmPassword {
		return nil, usecase.WrapError(fmt.Errorf("password not match"))
	}
	hashPwd, err := password.GetHashPassword(i.Password)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.me.CustomerValidateChangePasswordUseCase.ValidateChangePassword: %w", err))
	}
	i.HashPwd = &hashPwd
	return i, nil
}

func (s *CustomerGetUserFromCtx) GetUserFromCtx(ctx context.Context) (*model.Customer, error) {
	user := usecase.GetUserAsCustomer(ctx)
	return user, nil
}
