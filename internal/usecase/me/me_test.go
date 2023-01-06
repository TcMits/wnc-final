package me_test

import (
	"context"
	"testing"

	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/me"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/tool/password"
	"github.com/stretchr/testify/require"
)

func authenticateCtx(ctx *context.Context, c *ent.Client, user *model.Customer) {
	if user == nil {
		user, _ = ent.CreateFakeCustomer(*ctx, c, nil)
	}
	*ctx = context.WithValue(*ctx, string(usecase.UserCtxKey), user)
}

func TestChangePasswordUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerChangePasswordUseCase)
	}{
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerChangePasswordUseCase) {
				hashPwd, _ := password.GetHashPassword("foobaz")
				i := &model.CustomerChangePasswordInput{
					HashPwd: &hashPwd,
				}
				e, err := uc.ChangePassword(ctx, i)
				require.Nil(t, err)
				require.NotNil(t, e)
				err = password.ValidatePassword(e.Password, "foobaz")
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
			uc := me.NewCustomerChangePasswordUseCase(
				repository.GetCustomerUpdateRepository(c),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}

func TestValidateChangePassword(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerValidateChangePasswordUseCase)
	}{
		{
			name: "old password is invalid",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerValidateChangePasswordUseCase) {
				i := &model.CustomerChangePasswordInput{
					OldPassword: "foo",
				}
				_, err := uc.ValidateChangePassword(ctx, i)
				require.ErrorContains(t, err, "invalid")
			},
		},
		{
			name: "new password match old password is not allowed",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerValidateChangePasswordUseCase) {
				i := &model.CustomerChangePasswordInput{
					Password:    "123456789",
					OldPassword: "123456789",
				}
				_, err := uc.ValidateChangePassword(ctx, i)
				require.ErrorContains(t, err, "new password match old password is not allowed")
			},
		},
		{
			name: "password not match",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerValidateChangePasswordUseCase) {
				i := &model.CustomerChangePasswordInput{
					Password:        "foo",
					ConfirmPassword: "baz",
					OldPassword:     "123456789",
				}
				_, err := uc.ValidateChangePassword(ctx, i)
				require.ErrorContains(t, err, "password not match")
			},
		},
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerValidateChangePasswordUseCase) {
				i := &model.CustomerChangePasswordInput{
					Password:        "foo",
					ConfirmPassword: "foo",
					OldPassword:     "123456789",
				}
				res, err := uc.ValidateChangePassword(ctx, i)
				require.Nil(t, err)
				require.NotNil(t, res.HashPwd)
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
			uc := me.NewCustomerValidateChangePasswordUseCase()
			tt.expect(t, ctx, c, uc)
		})
	}
}
