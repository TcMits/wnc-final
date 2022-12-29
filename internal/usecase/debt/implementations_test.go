package debt_test

import (
	"context"
	"testing"

	"github.com/TcMits/wnc-final/ent"
	entDebt "github.com/TcMits/wnc-final/ent/debt"
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/debt"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func authenticateCtx(ctx *context.Context, c *ent.Client, user *model.Customer) {
	if user == nil {
		user, _ = ent.CreateFakeCustomer(*ctx, c, nil)
	}
	*ctx = context.WithValue(*ctx, usecase.UserCtxKey, middleware.UserCtxKey)
	*ctx = context.WithValue(*ctx, middleware.UserCtxKey, user)
}

func TestListUseCase(t *testing.T) {
	t.Parallel()
	c, _ := datastore.NewClientTestConnection(t)
	defer c.Close()
	ctx := context.Background()
	ent.CreateFakeDebt(ctx, c, nil)
	uc := debt.NewCustomerDebtListUseCase(repository.GetDebtListRepository(c))
	l, o := 1, 0
	result, err := uc.List(ctx, &l, &o, nil, nil)
	require.Nil(t, err)
	require.Equal(t, 1, len(result))
}

func TestValidateCreateInputUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerDebtValidateCreateInputUseCase)
	}{
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtValidateCreateInputUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				i1 := ent.BankAccountFactory()
				i1.IsForPayment = generic.GetPointer(true)
				i1.CustomerID = user.ID
				i2 := ent.BankAccountFactory()
				i2.IsForPayment = generic.GetPointer(true)
				ownerBA, _ := ent.CreateFakeBankAccount(ctx, c, i1)
				receiverBA, _ := ent.CreateFakeBankAccount(ctx, c, i2)
				i3 := ent.DebtFactory()
				i3.OwnerID = ownerBA.ID
				i3.ReceiverID = receiverBA.ID
				i3, err := uc.Validate(ctx, i3)
				require.Nil(t, err)
				require.Equal(t, i3.Status.String(), entDebt.StatusPending.String())
				require.Equal(t, i3.OwnerBankAccountNumber, ownerBA.AccountNumber)
			},
		},
		{
			name: "back account receiver not for payment",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtValidateCreateInputUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				i1 := ent.BankAccountFactory()
				i1.CustomerID = user.ID
				i1.IsForPayment = generic.GetPointer(true)
				i2 := ent.BankAccountFactory()
				ownerBA, _ := ent.CreateFakeBankAccount(ctx, c, i1)
				receiverBA, _ := ent.CreateFakeBankAccount(ctx, c, i2)
				i3 := ent.DebtFactory()
				i3.OwnerID = ownerBA.ID
				i3.ReceiverID = receiverBA.ID
				_, err := uc.Validate(ctx, i3)
				require.ErrorContains(t, err, "receiver not for payment")
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
			uc := debt.NewCustomerDebtValidateCreateInputUseCase(
				repository.GetBankAccountListRepository(c),
				repository.GetCustomerListRepository(c),
				generic.GetPointer("foo"),
				generic.GetPointer("foo"),
				generic.GetPointer(float64(1000)),
				generic.GetPointer("foo"),
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
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerDebtCreateUseCase)
	}{
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtCreateUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				i := ent.BankAccountFactory()
				i.CustomerID = user.ID
				ba, _ := ent.CreateFakeBankAccount(ctx, c, i)
				i1 := ent.DebtFactory()
				i1.OwnerID = ba.ID
				ent.CreateFakeDebt(ctx, c, i1)
				_, err := uc.Create(ctx, i1)
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
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			taskExecutorMock := task.NewMockIExecuteTask[*task.DebtCreateNotifyPayload](mockCtl)
			taskExecutorMock.EXPECT().ExecuteTask(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			uc := debt.NewCustomerDebtCreateUseCase(
				repository.GetDebtCreateRepository(c),
				repository.GetCustomerListRepository(c),
				taskExecutorMock,
				generic.GetPointer("foo"),
				generic.GetPointer("foo"),
				generic.GetPointer(float64(1000)),
				generic.GetPointer("foo"),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}
