package contact_test

import (
	"context"
	"testing"

	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/contact"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/stretchr/testify/require"
)

var authenticateCtx = usecase.AuthenticateCtx

func TestListUseCase(t *testing.T) {
	t.Parallel()
	c, _ := datastore.NewClientTestConnection(t)
	defer c.Close()
	ctx := context.Background()
	authenticateCtx(&ctx, c, nil)
	ent.CreateFakeContact(ctx, c, nil)
	uc := contact.NewCustomerContactListUseCase(repository.GetContactListRepository(c))
	l, o := 2, 0
	res, err := uc.List(ctx, &l, &o, nil, nil)
	require.Nil(t, err)
	require.Equal(t, 1, len(res))
}

func TestListMineUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerContactListMineUseCase)
	}{
		{
			name: "only my contacts",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerContactListMineUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				i := ent.ContactFactory(ctx, ent.Opt{Key: "OwnerID", Value: user.ID})
				ent.CreateFakeContact(ctx, c, i)
				l, o := 2, 0
				res, err := uc.ListMine(ctx, &l, &o, nil, nil)
				require.Nil(t, err)
				require.Equal(t, 1, len(res))
			},
		},
		{
			name: "my contacts and other",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerContactListMineUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				i := ent.ContactFactory(ctx, ent.Opt{Key: "OwnerID", Value: user.ID})
				ent.CreateFakeContact(ctx, c, i)
				ent.CreateFakeContact(ctx, c, nil)
				l, o := 2, 0
				res, err := uc.ListMine(ctx, &l, &o, nil, nil)
				require.Nil(t, err)
				require.Equal(t, 1, len(res))
			},
		},
		{
			name: "only other",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerContactListMineUseCase) {
				ent.CreateFakeContact(ctx, c, nil)
				l, o := 2, 0
				res, err := uc.ListMine(ctx, &l, &o, nil, nil)
				require.Nil(t, err)
				require.Equal(t, 0, len(res))
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
			uc := contact.NewCustomerContactListMineUseCase(
				repository.GetContactListRepository(c),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}
func TestGetFirstMineUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerContactGetFirstMineUseCase)
	}{
		{
			name: "only my contacts",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerContactGetFirstMineUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				i := ent.ContactFactory(ctx, ent.Opt{Key: "OwnerID", Value: user.ID})
				e, _ := ent.CreateFakeContact(ctx, c, i)
				res, err := uc.GetFirstMine(ctx, nil, nil)
				require.Nil(t, err)
				require.NotNil(t, res)
				require.Equal(t, res.ID, e.ID)
			},
		},
		{
			name: "my contacts and other",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerContactGetFirstMineUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				i := ent.ContactFactory(ctx, ent.Opt{Key: "OwnerID", Value: user.ID})
				e, _ := ent.CreateFakeContact(ctx, c, i)
				ent.CreateFakeContact(ctx, c, nil)
				res, err := uc.GetFirstMine(ctx, nil, nil)
				require.Nil(t, err)
				require.NotNil(t, res)
				require.Equal(t, res.ID, e.ID)
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
			uc := contact.NewCustomerContactGetFirstMineUseCase(
				repository.GetContactListRepository(c),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}

func TestCreateUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerContactCreateUseCase)
	}{
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerContactCreateUseCase) {
				i := ent.ContactFactory(ctx)
				res, err := uc.Create(ctx, i)
				require.Nil(t, err)
				require.NotNil(t, res)
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
			uc := contact.NewCustomerContactCreateUseCase(
				repository.GetContactCreateRepository(c),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}

func TestUpdateUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerContactUpdateUseCase)
	}{
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerContactUpdateUseCase) {
				e, _ := ent.CreateFakeContact(ctx, c, nil)
				res, err := uc.Update(ctx, e, &model.ContactUpdateInput{
					SuggestName: generic.GetPointer("foo"),
				})
				require.Nil(t, err)
				require.NotNil(t, res)
				require.Equal(t, res.SuggestName, "foo")
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
			uc := contact.NewCustomerContactUpdateUseCase(
				repository.GetContactUpdateRepository(c),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}

func TestValidateCreateUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerContactValidateCreateInputUseCase)
	}{
		{
			name: "duplicate constraint",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerContactValidateCreateInputUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				i := ent.ContactFactory(ctx, ent.Opt{Key: "OwnerID", Value: user.ID}, ent.Opt{Key: "BankName", Value: "foo"})
				ent.CreateFakeContact(ctx, c, i)
				_, err := uc.ValidateCreate(ctx, i)
				require.ErrorContains(t, err, "the account number of the bank already existed")
			},
		},
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerContactValidateCreateInputUseCase) {
				i := ent.ContactFactory(ctx, ent.Opt{Key: "BankName", Value: "foo"})
				res, err := uc.ValidateCreate(ctx, i)
				require.Nil(t, err)
				require.Equal(t, res.BankName, "foo")
				user := usecase.GetUserAsCustomer(ctx)
				require.Equal(t, res.OwnerID.String(), user.ID.String())
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
			uc := contact.NewCustomerContactValidateCreateInputUseCase(
				repository.GetContactListRepository(c),
				repository.GetCustomerListRepository(c),
				generic.GetPointer("foo"),
				generic.GetPointer("foo"),
				generic.GetPointer("foo"),
				"foo",
				"foo",
				"foo",
				"foo",
				"foo",
				"foo",
				"foo",
				"foo",
				"foo",
				"foo",
				generic.GetPointer(float64(1000)),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}

func TestValidateUpdateUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerContactValidateUpdateInputUseCase)
	}{
		{
			name: "duplicate constraint",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerContactValidateUpdateInputUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				i := ent.ContactFactory(ctx, ent.Opt{Key: "OwnerID", Value: user.ID}, ent.Opt{Key: "BankName", Value: "foo"})
				e1, _ := ent.CreateFakeContact(ctx, c, i)
				i = ent.ContactFactory(ctx, ent.Opt{Key: "OwnerID", Value: user.ID}, ent.Opt{Key: "BankName", Value: "foo"})
				e2, _ := ent.CreateFakeContact(ctx, c, i)
				_, err := uc.ValidateUpdate(ctx, e2, &model.ContactUpdateInput{
					AccountNumber: &e1.AccountNumber,
					BankName:      &e1.BankName,
				})
				require.ErrorContains(t, err, "the account number of the bank already existed")
			},
		},
		{
			name: "success with input has no fields",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerContactValidateUpdateInputUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				i := ent.ContactFactory(ctx, ent.Opt{Key: "OwnerID", Value: user.ID}, ent.Opt{Key: "BankName", Value: "foo"})
				e, _ := ent.CreateFakeContact(ctx, c, i)
				res, err := uc.ValidateUpdate(ctx, e, &model.ContactUpdateInput{})
				require.Nil(t, err)
				require.Nil(t, res.OwnerID)
				require.Equal(t, *res.BankName, e.BankName)
				require.Equal(t, *res.AccountNumber, e.AccountNumber)
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
			uc := contact.NewCustomerContactValidateUpdateInputUseCase(
				repository.GetContactListRepository(c),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}
