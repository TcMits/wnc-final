package me

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/password"
)

type (
	CustomerChangePasswordUseCase struct {
		UC1 usecase.ICustomerUpdateUseCase
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

func (s *CustomerChangePasswordUseCase) ChangePassword(ctx context.Context, i *model.CustomerChangePasswordInput) (*model.Customer, error) {
	user := usecase.GetUserAsCustomer(ctx)
	user, err := s.UC1.Update(ctx, user, &model.CustomerUpdateInput{
		ClearPassword: true,
		Password:      i.HashPwd,
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.auth.CustomerChangePasswordUseCase.ChangePassword: %w", err))
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
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.me.CustomerValidateChangePasswordUseCase.ValidateChangePassword: %w", err))
	}
	i.HashPwd = &hashPwd
	return i, nil
}

func (s *CustomerGetUserFromCtx) GetUserFromCtx(ctx context.Context) (*model.Customer, error) {
	user := usecase.GetUserAsCustomer(ctx)
	return user, nil
}

// employee
type (
	EmployeeGetUserFromCtx struct{}

	EmployeeMeUseCase struct {
		usecase.IEmployeeConfigUseCase
		usecase.IEmployeeGetUserUseCase
		usecase.IEmployeeGetUserFromCtxUseCase
	}
)

func (s *EmployeeGetUserFromCtx) GetUserFromCtx(ctx context.Context) (*model.Employee, error) {
	user := usecase.GetUserAsEmployee(ctx)
	return user, nil
}

// admin
type (
	AdminGetUserFromCtx struct{}

	AdminMeUseCase struct {
		usecase.IAdminConfigUseCase
		usecase.IAdminGetUserUseCase
		usecase.IAdminGetUserFromCtxUseCase
	}
)

func (s *AdminGetUserFromCtx) GetUserFromCtx(ctx context.Context) (*model.Admin, error) {
	user := usecase.GetUserAsAdmin(ctx)
	return user, nil
}
