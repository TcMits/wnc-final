package debt_test

import (
	"context"
	"testing"
	"time"

	"github.com/TcMits/wnc-final/config"
	"github.com/TcMits/wnc-final/ent"
	entDebt "github.com/TcMits/wnc-final/ent/debt"
	entTxc "github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/debt"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var authenticateCtx = usecase.AuthenticateCtx

func TestListUseCase(t *testing.T) {
	t.Parallel()
	c, _ := datastore.NewClientTestConnection(t)
	defer c.Close()
	ctx := context.Background()
	ent.EmbedClient(&ctx, c)
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
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtValidateCreateInputUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				ownerBA, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CustomerID", Value: user.ID},
				)
				receiverBA, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
				)
				i := ent.DebtFactory(ctx, ent.Opt{Key: "OwnerID", Value: ownerBA.ID}, ent.Opt{Key: "ReceiverID", Value: receiverBA.ID})
				i, err := uc.ValidateCreate(ctx, i)
				require.Nil(t, err)
				require.Equal(t, i.Status.String(), entDebt.StatusPending.String())
				require.Equal(t, i.OwnerBankAccountNumber, ownerBA.AccountNumber)
			},
		},
		{
			name: "back account receiver not for payment",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtValidateCreateInputUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				ownerBA, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "CustomerID", Value: user.ID},
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
				)
				receiverBA, _ := ent.CreateFakeBankAccount(ctx, c, nil)
				i := ent.DebtFactory(ctx,
					ent.Opt{Key: "OwnerID", Value: ownerBA.ID},
					ent.Opt{Key: "ReceiverID", Value: receiverBA.ID},
				)
				_, err := uc.ValidateCreate(ctx, i)
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
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtCreateUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				ba, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "CustomerID", Value: user.ID},
				)
				i := ent.DebtFactory(ctx,
					ent.Opt{Key: "OwnerID", Value: ba.ID},
				)
				_, err := uc.Create(ctx, i)
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
			taskExecutorMock := task.NewMockIExecuteTask[*task.DebtNotifyPayload](mockCtl)
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

func TestValidateCancelUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerDebtValidateCancelUseCase)
	}{
		{
			name: "not pending debt",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtValidateCancelUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				receiverBA, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CustomerID", Value: user.ID},
				)
				e1, _ := ent.CreateFakeDebt(ctx, c, nil,
					ent.Opt{Key: "ReceiverID", Value: receiverBA.ID},
					ent.Opt{Key: "Status", Value: generic.GetPointer(entDebt.StatusCancelled)},
				)
				_, err := uc.ValidateCancel(ctx, e1, nil)
				require.ErrorContains(t, err, "cannot cancel")
			},
		},
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtValidateCancelUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				receiverBA, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CustomerID", Value: user.ID},
				)
				e1, _ := ent.CreateFakeDebt(ctx, c, nil,
					ent.Opt{Key: "ReceiverID", Value: receiverBA.ID},
				)
				res, err := uc.ValidateCancel(ctx, e1, nil)
				require.Nil(t, err)
				require.NotNil(t, res)
				require.Equal(t, res.Status.String(), entDebt.StatusCancelled.String())
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
			uc := debt.NewCustomerDebtValidateCancelUseCase(
				repository.GetCustomerListRepository(c),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}

func TestCancelUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerDebtCancelUseCase)
	}{
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtCancelUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				receiverBA, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CustomerID", Value: user.ID},
				)
				e1, _ := ent.CreateFakeDebt(ctx, c, nil,
					ent.Opt{Key: "ReceiverID", Value: receiverBA.ID},
				)
				i3 := &model.DebtUpdateInput{
					Status: generic.GetPointer(entDebt.StatusCancelled),
				}
				res, err := uc.Cancel(ctx, e1, i3)
				require.Nil(t, err)
				require.NotNil(t, res)
				require.Equal(t, res.Status.String(), entDebt.StatusCancelled.String())
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
			taskExecutorMock := task.NewMockIExecuteTask[*task.DebtNotifyPayload](mockCtl)
			taskExecutorMock.EXPECT().ExecuteTask(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			uc := debt.NewCustomerDebtCancelUseCase(
				repository.GetDebtUpdateRepository(c),
				taskExecutorMock,
				repository.GetCustomerListRepository(c),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}

func TestValidateFulfillUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerDebtValidateFulfillUseCase)
	}{
		{
			name: "not pending debt",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtValidateFulfillUseCase) {
				e1, _ := ent.CreateFakeDebt(ctx, c, nil,
					ent.Opt{Key: "Status", Value: generic.GetPointer(entDebt.StatusCancelled)},
				)
				err := uc.ValidateFulfill(ctx, e1)
				require.ErrorContains(t, err, "cannot fulfill")
			},
		},
		{
			name: "fulfill owned debt",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtValidateFulfillUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				ownerBA, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CustomerID", Value: user.ID},
				)
				e1, _ := ent.CreateFakeDebt(ctx, c, nil,
					ent.Opt{Key: "OwnerID", Value: ownerBA.ID},
				)
				err := uc.ValidateFulfill(ctx, e1)
				require.ErrorContains(t, err, "cannot fulfill debt which you created")
			},
		},
		{
			name: "insufficient ballence",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtValidateFulfillUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				receiverBA, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CustomerID", Value: user.ID},
				)
				e1, _ := ent.CreateFakeDebt(ctx, c, nil,
					ent.Opt{Key: "ReceiverID", Value: receiverBA.ID},
				)
				err := uc.ValidateFulfill(ctx, e1)
				require.ErrorContains(t, err, "insufficient ballence")
			},
		},
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtValidateFulfillUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				receiverBA, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CustomerID", Value: user.ID},
					ent.Opt{Key: "CashIn", Value: float64(1000)},
				)
				e1, _ := ent.CreateFakeDebt(ctx, c, nil,
					ent.Opt{Key: "ReceiverID", Value: receiverBA.ID},
				)
				err := uc.ValidateFulfill(ctx, e1)
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
			taskExecutorMock := task.NewMockIExecuteTask[*task.DebtNotifyPayload](mockCtl)
			taskExecutorMock.EXPECT().ExecuteTask(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			uc := debt.NewCustomerDebtValidateFulfillUseCase(
				repository.GetCustomerListRepository(c),
				repository.GetBankAccountListRepository(c),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}

func TestFulfillWithTokenUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerDebtFulfillWithTokenUseCase)
	}{
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtFulfillWithTokenUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				receiverBA, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CustomerID", Value: user.ID},
					ent.Opt{Key: "CashIn", Value: float64(1000)},
				)
				e1, _ := ent.CreateFakeDebt(ctx, c, nil,
					ent.Opt{Key: "ReceiverID", Value: receiverBA.ID},
				)
				ownerBA := e1.QueryOwner().FirstX(ctx)
				oldBalanceOwner := ownerBA.GetBalance()
				i := &model.DebtFulfillWithTokenInput{
					Token: "foo",
					Otp:   "foo",
				}
				res, err := uc.FulfillWithToken(ctx, e1, i)
				require.Nil(t, err)
				require.NotNil(t, res)
				require.Equal(t, res.Status.String(), entDebt.StatusFulfilled.String())

				oldBalanceRecv := receiverBA.GetBalance()
				receiverBA, _ = ent.RefreshBankAccountFromDB(ctx, c, receiverBA)
				require.Less(t, receiverBA.GetBalance(), oldBalanceRecv)

				ownerBA, _ = ent.RefreshBankAccountFromDB(ctx, c, ownerBA)
				require.Greater(t, ownerBA.GetBalance(), oldBalanceOwner)

				e2 := c.Transaction.Query().Where(entTxc.ID(*res.TransactionID)).FirstX(ctx)
				require.NotNil(t, e2)
				require.Equal(t, *e2.SenderID, receiverBA.ID)
				require.Equal(t, *e2.ReceiverID, ownerBA.ID)
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
			taskExecutorMock := task.NewMockIExecuteTask[*task.DebtNotifyPayload](mockCtl)
			taskExecutorMock.EXPECT().ExecuteTask(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			uc := debt.NewCustomerDebtFulfillWithTokenUseCase(
				repository.GetDebtFulfillRepository(c),
				repository.GetCustomerListRepository(c),
				taskExecutorMock,
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}

func TestFulfillUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerDebtFulfillUseCase)
	}{
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtFulfillUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				receiverBA, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CustomerID", Value: user.ID},
					ent.Opt{Key: "CashIn", Value: float64(1000)},
				)
				e1, _ := ent.CreateFakeDebt(ctx, c, nil,
					ent.Opt{Key: "ReceiverID", Value: receiverBA.ID},
				)
				res, err := uc.Fulfill(ctx, e1)
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
			tt.setUp(t, &ctx, c)
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			taskExecutorMock := task.NewMockIExecuteTask[*mail.EmailPayload](mockCtl)
			taskExecutorMock.EXPECT().ExecuteTask(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			uc := debt.NewCustomerDebtFulfillUseCase(
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

func TestValidateFulfillWithTokenUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerDebtValidateFulfillWithTokenUseCase)
	}{
		{
			name: "token expired",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtValidateFulfillWithTokenUseCase) {
				tk, _ := usecase.GenerateFulfillToken(ctx, "foo", "bar", time.Minute*30)
				e, _ := ent.CreateFakeDebt(ctx, c, nil)
				i := &model.DebtFulfillWithTokenInput{
					Token: tk,
					Otp:   "foo",
				}
				_, err := uc.ValidateFulfillWithToken(ctx, e, i)
				require.ErrorContains(t, err, "token expired")
			},
		},
		{
			name: "invalid otp",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtValidateFulfillWithTokenUseCase) {
				e, _ := ent.CreateFakeDebt(ctx, c, nil)
				otpHash, _ := usecase.GenerateHashInfo(usecase.MakeOTPValue(ctx, "foo", e.ID.String()))
				tk, _ := usecase.GenerateFulfillToken(ctx, otpHash, "foo", time.Minute*30)
				i := &model.DebtFulfillWithTokenInput{
					Token: tk,
					Otp:   "bar",
				}
				_, err := uc.ValidateFulfillWithToken(ctx, e, i)
				require.ErrorContains(t, err, "invalid otp")
			},
		},
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerDebtValidateFulfillWithTokenUseCase) {
				e, _ := ent.CreateFakeDebt(ctx, c, nil)
				otpHash, _ := usecase.GenerateHashInfo(usecase.MakeOTPValue(ctx, "foo", e.ID.String()))
				tk, _ := usecase.GenerateFulfillToken(ctx, otpHash, "foo", time.Minute*30)
				i := &model.DebtFulfillWithTokenInput{
					Token: tk,
					Otp:   "foo",
				}
				res, err := uc.ValidateFulfillWithToken(ctx, e, i)
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
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			taskExecutorMock := task.NewMockIExecuteTask[*mail.EmailPayload](mockCtl)
			taskExecutorMock.EXPECT().ExecuteTask(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			uc := debt.NewCustomerDebtValidateFulfillWithTokenUseCase(
				generic.GetPointer("foo"),
				generic.GetPointer("foo"),
				generic.GetPointer("foo"),
				generic.GetPointer(float64(1000)),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}
