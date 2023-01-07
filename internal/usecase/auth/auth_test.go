package auth_test

import (
	"context"
	"testing"
	"time"

	"github.com/TcMits/wnc-final/config"
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/auth"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestForgetPasswordUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerForgetPasswordUseCase)
	}{
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerForgetPasswordUseCase) {
				user, _ := ent.CreateFakeCustomer(ctx, c, nil)
				i := &model.CustomerForgetPasswordInput{
					User:  user,
					Email: user.Email,
				}
				res, err := uc.ForgetPassword(ctx, i)
				require.Nil(t, err)
				require.NotNil(t, res)
				require.NotEqual(t, res.Token, "")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, _ := config.NewConfigForTest()
			c, _ := datastore.NewClientTestConnection(t)
			defer c.Close()
			ctx := context.Background()
			require.NoError(t, c.Schema.Create(ctx))
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			taskExecutorMock := task.NewMockIExecuteTask[*mail.EmailPayload](mockCtl)
			taskExecutorMock.EXPECT().ExecuteTask(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			tt.setUp(t, &ctx, c)
			uc := auth.NewCustomerForgetPasswordUseCase(
				taskExecutorMock,
				&cfg.App.SecretKey,
				&cfg.App.Name,
				&cfg.TransactionUseCase.FeeDesc,
				&cfg.Mail.ConfirmEmailSubject,
				&cfg.Mail.ConfirmEmailTemplate,
				&cfg.TransactionUseCase.FeeAmount,
				cfg.Mail.OTPTimeout,
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}

func TestValidateForgetPasswordUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerValidateForgetPasswordUsecase)
	}{
		{
			name: "user does not exist",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerValidateForgetPasswordUsecase) {
				i := &model.CustomerForgetPasswordInput{
					Email: "foo@gmail.com",
				}
				_, err := uc.ValidateForgetPassword(ctx, i)
				require.ErrorContains(t, err, "user does not exist")
			},
		},
		{
			name:  "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerValidateForgetPasswordUsecase) {
				email := "foo@gmail.com"
				ent.CreateFakeCustomer(ctx, c, nil, ent.Opt{Key: "Email", Value: email})
				res, err := uc.ValidateForgetPassword(ctx, &model.CustomerForgetPasswordInput{Email: email})
				require.Nil(t, err)
				require.NotNil(t, res)
				require.NotNil(t, res.User)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := datastore.NewClientTestConnection(t)
			defer c.Close()
			ctx := context.Background()
			require.NoError(t, c.Schema.Create(ctx))
			tt.setUp(t, &ctx, c)
			uc := auth.NewCustomerValidateForgetPasswordUseCase(
				repository.GetCustomerListRepository(c),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}

func TestChangePasswordWithTokenUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerChangePasswordWithTokenUseCase)
	}{
		{
			name:  "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerChangePasswordWithTokenUseCase) {
				user, _ := ent.CreateFakeCustomer(ctx, c, nil)
				i := &model.CustomerChangePasswordWithTokenInput{
					User:    user,
					HashPwd: generic.GetPointer("foo"),
				}
				err := uc.ChangePasswordWithToken(ctx, i)
				require.Nil(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := datastore.NewClientTestConnection(t)
			defer c.Close()
			ctx := context.Background()
			require.NoError(t, c.Schema.Create(ctx))
			tt.setUp(t, &ctx, c)
			uc := auth.NewCustomerChangePasswordWithTokenUseCase(
				repository.GetCustomerUpdateRepository(c),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}

func TestValidateChangePasswordWithToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerValidateChangePasswordWithTokenUseCase)
	}{
		{
			name: "user does not exist",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerValidateChangePasswordWithTokenUseCase) {
				tk, _ := usecase.GenerateForgetPwdToken(
					ctx,
					"foo",
					"foo@gmail.com",
					"foo",
					time.Minute*30,
				)
				i := &model.CustomerChangePasswordWithTokenInput{
					Token: tk,
				}
				_, err := uc.ValidateChangePasswordWithToken(ctx, i)
				require.ErrorContains(t, err, "user does not exist")
			},
		},
		{
			name: "otp invalid",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerValidateChangePasswordWithTokenUseCase) {
				user, _ := ent.CreateFakeCustomer(ctx, c, nil, ent.Opt{Key: "Email", Value: "foo@gmail.com"})
				otpHashValue, _ := usecase.GenerateHashInfo(usecase.MakeOTPValue(usecase.EmbedUser(ctx, user), "barfoo"))
				tk, _ := usecase.GenerateForgetPwdToken(
					ctx,
					otpHashValue,
					"foo@gmail.com",
					"foo",
					time.Minute*30,
				)
				i := &model.CustomerChangePasswordWithTokenInput{
					Token: tk,
					Otp:   "foobar",
				}
				_, err := uc.ValidateChangePasswordWithToken(ctx, i)
				require.ErrorContains(t, err, "otp invalid")
			},
		},
		{
			name: "password not match",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerValidateChangePasswordWithTokenUseCase) {
				user, _ := ent.CreateFakeCustomer(ctx, c, nil, ent.Opt{Key: "Email", Value: "foo@gmail.com"})
				otpHashValue, _ := usecase.GenerateHashInfo(usecase.MakeOTPValue(usecase.EmbedUser(ctx, user), "barfoo"))
				tk, _ := usecase.GenerateForgetPwdToken(
					ctx,
					otpHashValue,
					"foo@gmail.com",
					"foo",
					time.Minute*30,
				)
				i := &model.CustomerChangePasswordWithTokenInput{
					Token:           tk,
					Otp:             "barfoo",
					Password:        "foo",
					ConfirmPassword: "bar",
				}
				_, err := uc.ValidateChangePasswordWithToken(ctx, i)
				require.ErrorContains(t, err, "password not match")
			},
		},
		{
			name: "new password match old password is not allowed",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerValidateChangePasswordWithTokenUseCase) {
				user, _ := ent.CreateFakeCustomer(ctx, c, nil, ent.Opt{Key: "Email", Value: "foo@gmail.com"})
				otpHashValue, _ := usecase.GenerateHashInfo(usecase.MakeOTPValue(usecase.EmbedUser(ctx, user), "barfoo"))
				tk, _ := usecase.GenerateForgetPwdToken(
					ctx,
					otpHashValue,
					"foo@gmail.com",
					"foo",
					time.Minute*30,
				)
				i := &model.CustomerChangePasswordWithTokenInput{
					Token:           tk,
					Otp:             "barfoo",
					Password:        "123456789",
					ConfirmPassword: "123456789",
				}
				_, err := uc.ValidateChangePasswordWithToken(ctx, i)
				require.ErrorContains(t, err, "new password match old password is not allowed")
			},
		},
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerValidateChangePasswordWithTokenUseCase) {
				user, _ := ent.CreateFakeCustomer(ctx, c, nil, ent.Opt{Key: "Email", Value: "foo@gmail.com"})
				otpHashValue, _ := usecase.GenerateHashInfo(usecase.MakeOTPValue(usecase.EmbedUser(ctx, user), "barfoo"))
				tk, _ := usecase.GenerateForgetPwdToken(
					ctx,
					otpHashValue,
					"foo@gmail.com",
					"foo",
					time.Minute*30,
				)
				i := &model.CustomerChangePasswordWithTokenInput{
					Token:           tk,
					Otp:             "barfoo",
					Password:        "foobar",
					ConfirmPassword: "foobar",
				}
				res, err := uc.ValidateChangePasswordWithToken(ctx, i)
				require.Nil(t, err)
				require.NotNil(t, res)
				require.Equal(t, res.User.ID.String(), user.ID.String())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := datastore.NewClientTestConnection(t)
			defer c.Close()
			ctx := context.Background()
			require.NoError(t, c.Schema.Create(ctx))
			tt.setUp(t, &ctx, c)
			uc := auth.NewCustomerValidateChangePasswordWithTokenUseCase(
				repository.GetCustomerListRepository(c),
				generic.GetPointer("foo"),
				generic.GetPointer("foo"),
				generic.GetPointer("foo"),
				generic.GetPointer(float64(1000)),
			)
			tt.expect(t, ctx, c, uc)
		})
	}

}
