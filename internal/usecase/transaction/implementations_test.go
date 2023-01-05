package transaction_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/TcMits/wnc-final/config"
	entTxc "github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/transaction"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
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
	ent.EmbedClient(&ctx, c)
	ent.CreateFakeTransaction(ctx, c, nil)
	uc := transaction.NewCustomerTransactionListUseCase(repository.GetTransactionListRepository(c))
	l, o := 1, 0
	result, err := uc.List(ctx, &l, &o, nil, nil)
	require.Nil(t, err)
	require.Equal(t, 1, len(result))
}

func TestListMyTxcUseCase(t *testing.T) {
	t.Parallel()
	c, _ := datastore.NewClientTestConnection(t)
	defer c.Close()
	ctx := context.Background()
	ent.EmbedClient(&ctx, c)
	mBA, _ := ent.CreateFakeBankAccount(ctx, c, nil)
	authenticateCtx(&ctx, c, mBA.QueryCustomer().FirstX(ctx))
	ent.CreateFakeTransaction(ctx, c, nil, ent.Opt{Key: "SenderID", Value: mBA.ID})
	ent.CreateFakeTransaction(ctx, c, nil, ent.Opt{Key: "ReceiverID", Value: &mBA.ID})
	ent.CreateFakeTransaction(ctx, c, nil)
	uc := transaction.NewCustomerTransactionListMineUseCase(repository.GetTransactionListRepository(c))
	l, o := 3, 0
	result, err := uc.ListMine(ctx, &l, &o, nil, nil)
	require.Nil(t, err)
	require.Equal(t, 2, len(result))
}

func TestGetFirstMyTxcUseCase(t *testing.T) {
	t.Parallel()
	c, _ := datastore.NewClientTestConnection(t)
	defer c.Close()
	ctx := context.Background()
	ent.EmbedClient(&ctx, c)
	mBA, _ := ent.CreateFakeBankAccount(ctx, c, nil)
	authenticateCtx(&ctx, c, mBA.QueryCustomer().FirstX(ctx))
	entity1, _ := ent.CreateFakeTransaction(ctx, c, nil, ent.Opt{Key: "SenderID", Value: mBA.ID})
	ent.CreateFakeTransaction(ctx, c, nil)
	uc := transaction.NewCustomerTransactionGetFirstMineUseCase(repository.GetTransactionListRepository(c))
	result, err := uc.GetFirstMine(ctx, nil, nil)
	require.Nil(t, err)
	require.Equal(t, entity1.ID, result.ID)
}

func TestUpdateUseCase(t *testing.T) {
	t.Parallel()
	c, _ := datastore.NewClientTestConnection(t)
	defer c.Close()
	ctx := context.Background()
	ent.EmbedClient(&ctx, c)
	entity1, _ := ent.CreateFakeTransaction(ctx, c, nil)
	authenticateCtx(&ctx, c, entity1.QuerySender().FirstX(ctx).QueryCustomer().FirstX(ctx))
	uc := transaction.NewCustomerTransactionUpdateUseCase(repository.GetTransactionUpdateRepository(c))
	entity1, err := uc.Update(ctx, entity1, &model.TransactionUpdateInput{
		Status: generic.GetPointer(entTxc.StatusSuccess),
	})
	require.Nil(t, err)
	require.Equal(t, entity1.Status, entTxc.StatusSuccess)
}

func TestValidateCreateInputUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerTransactionValidateCreateInputUseCase)
	}{
		{
			name: "bank account sender does have draft transactions",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateCreateInputUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				ba, _ := ent.CreateFakeBankAccount(ctx, c, nil, ent.Opt{Key: "CustomerID", Value: user.ID}, ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)})
				ent.CreateFakeTransaction(ctx, c, nil, ent.Opt{Key: "SenderID", Value: ba.ID})
				i3 := ent.TransactionFactory(ctx, ent.Opt{Key: "SenderID", Value: ba.ID})
				i := &model.TransactionCreateUseCaseInput{
					TransactionCreateInput: i3,
					IsFeePaidByMe:          true,
				}
				_, err := uc.ValidateCreate(ctx, i)
				require.ErrorContains(t, err, "there is a draft transaction to be processed. Cannot create a new transaction")
			},
		},
		{
			name: "insufficient balance from sender and fee paid by me",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateCreateInputUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				ba, _ := ent.CreateFakeBankAccount(ctx, c, nil, ent.Opt{Key: "CustomerID", Value: user.ID}, ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)})
				i2 := ent.TransactionFactory(ctx, ent.Opt{Key: "SenderID", Value: ba.ID})
				i := &model.TransactionCreateUseCaseInput{
					TransactionCreateInput: i2,
					IsFeePaidByMe:          true,
				}
				_, err := uc.ValidateCreate(ctx, i)
				require.ErrorContains(t, err, "insufficient balance")
			},
		},
		{
			name: "insufficient balance from sender and fee not paid by me",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateCreateInputUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				ba, _ := ent.CreateFakeBankAccount(ctx, c, nil, ent.Opt{Key: "CustomerID", Value: user.ID}, ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)})
				i2 := ent.TransactionFactory(ctx, ent.Opt{Key: "SenderID", Value: ba.ID})
				i := &model.TransactionCreateUseCaseInput{
					TransactionCreateInput: i2,
					IsFeePaidByMe:          false,
				}
				_, err := uc.ValidateCreate(ctx, i)
				require.ErrorContains(t, err, "insufficient balance")
			},
		},
		{
			name: "insufficient balance from receiver and fee not paid by me",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateCreateInputUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				sender, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CashIn", Value: float64(100000)},
					ent.Opt{Key: "CashOut", Value: float64(1)},
					ent.Opt{Key: "CustomerID", Value: user.ID},
				)
				receiver, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
				)
				i1 := ent.TransactionFactory(ctx,
					ent.Opt{Key: "SenderID", Value: sender.ID},
					ent.Opt{Key: "ReceiverID", Value: &receiver.ID},
				)
				i := &model.TransactionCreateUseCaseInput{
					TransactionCreateInput: i1,
					IsFeePaidByMe:          false,
				}
				_, err := uc.ValidateCreate(ctx, i)
				require.ErrorContains(t, err, "insufficient balance")
			},
		},
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateCreateInputUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				sender, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CashIn", Value: float64(100000)},
					ent.Opt{Key: "CashOut", Value: float64(1)},
					ent.Opt{Key: "CustomerID", Value: user.ID},
				)
				receiver, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CashIn", Value: float64(100000)},
					ent.Opt{Key: "CashOut", Value: float64(1)},
				)
				i1 := ent.TransactionFactory(ctx,
					ent.Opt{Key: "SenderID", Value: sender.ID},
					ent.Opt{Key: "ReceiverID", Value: &receiver.ID},
				)
				i := &model.TransactionCreateUseCaseInput{
					TransactionCreateInput: i1,
					IsFeePaidByMe:          true,
				}
				i, err := uc.ValidateCreate(ctx, i)
				require.Nil(t, err)
				require.Equal(t, entTxc.StatusDraft.String(), i.TransactionCreateInput.Status.String())
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
			uc := transaction.NewCustomerTransactionValidateCreateInputUseCase(
				repository.GetTransactionListRepository(c),
				repository.GetBankAccountListRepository(c),
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

func TestValidateConfirmInputUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerTransactionValidateConfirmInputUseCase)
	}{
		{
			name: "not draft transaction",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateConfirmInputUseCase) {
				entity1, _ := ent.CreateFakeTransaction(ctx, c, nil, ent.Opt{Key: "Status", Value: generic.GetPointer(entTxc.StatusSuccess)})
				err := uc.ValidateConfirmInput(ctx, entity1, nil)
				require.ErrorContains(t, err, fmt.Sprintf("cannot confirm %s transaction", entity1.Status))
			},
		},
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateConfirmInputUseCase) {
				otp := usecase.GenerateOTP(6)
				hashValue, _ := usecase.GenerateHashInfo(usecase.MakeOTPValue(ctx, otp))
				tk, _ := usecase.GenerateConfirmTxcToken(
					ctx,
					hashValue,
					"foo",
					true,
					time.Minute*30,
				)
				entity1, _ := ent.CreateFakeTransaction(ctx, c, nil)

				err := uc.ValidateConfirmInput(ctx, entity1, &model.TransactionConfirmUseCaseInput{
					Token: tk,
					Otp:   otp,
				})
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
			uc := transaction.NewCustomerTransactionValidateConfirmInputUseCase(
				generic.GetPointer("foo"),
				generic.GetPointer("foo"),
				generic.GetPointer(float64(1000)),
				generic.GetPointer("foo"),
			)
			tt.expect(t, ctx, c, uc)
		})
	}
}

func TestConfirmSuccessUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, *context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerTransactionConfirmSuccessUseCase)
	}{
		{
			name: "success and fee paid by me",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionConfirmSuccessUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				sender, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CashIn", Value: float64(100000)},
					ent.Opt{Key: "CashOut", Value: float64(1)},
					ent.Opt{Key: "CustomerID", Value: user.ID},
				)
				receiver, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CashIn", Value: float64(100000)},
					ent.Opt{Key: "CashOut", Value: float64(1)},
				)
				entity1, _ := ent.CreateFakeTransaction(ctx, c, nil,
					ent.Opt{Key: "SenderID", Value: sender.ID},
					ent.Opt{Key: "ReceiverID", Value: &receiver.ID},
				)
				tk, _ := usecase.GenerateConfirmTxcToken(
					ctx,
					"foo",
					"foo",
					true,
					time.Minute*30,
				)
				entity1, err := uc.ConfirmSuccess(ctx, entity1, &tk)
				require.Nil(t, err)
				require.Equal(t, entTxc.StatusSuccess.String(), entity1.Status.String())

				oldBalance := sender.GetBalance()
				sender, _ = ent.RefreshBankAccountFromDB(ctx, c, sender)
				require.Less(t, sender.GetBalance(), oldBalance)

				oldBalance = receiver.GetBalance()
				receiver, _ = ent.RefreshBankAccountFromDB(ctx, c, receiver)
				require.Greater(t, receiver.GetBalance(), oldBalance)

				feeTxc := c.Transaction.Query().Where(entTxc.SourceTransactionID(entity1.ID)).FirstX(ctx)
				require.Equal(t, *feeTxc.SenderID, sender.ID)
				require.Equal(t, *feeTxc.SourceTransactionID, entity1.ID)
				require.Equal(t, feeTxc.Amount, decimal.NewFromFloat(float64(1000)))
			},
		},
		{
			name: "success and fee not paid by me",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionConfirmSuccessUseCase) {
				user := usecase.GetUserAsCustomer(ctx)
				sender, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CashIn", Value: float64(100000)},
					ent.Opt{Key: "CashOut", Value: float64(1)},
					ent.Opt{Key: "CustomerID", Value: user.ID},
				)
				receiver, _ := ent.CreateFakeBankAccount(ctx, c, nil,
					ent.Opt{Key: "IsForPayment", Value: generic.GetPointer(true)},
					ent.Opt{Key: "CashIn", Value: float64(100000)},
					ent.Opt{Key: "CashOut", Value: float64(1)},
				)
				entity1, _ := ent.CreateFakeTransaction(ctx, c, nil,
					ent.Opt{Key: "SenderID", Value: sender.ID},
					ent.Opt{Key: "ReceiverID", Value: &receiver.ID},
				)
				tk, _ := usecase.GenerateConfirmTxcToken(
					ctx,
					"foo",
					"foo",
					false,
					time.Minute*30,
				)
				entity1, err := uc.ConfirmSuccess(ctx, entity1, &tk)
				require.Nil(t, err)
				require.Equal(t, entTxc.StatusSuccess.String(), entity1.Status.String())

				oldBalance := sender.GetBalance()
				sender, _ = ent.RefreshBankAccountFromDB(ctx, c, sender)
				require.Less(t, sender.GetBalance(), oldBalance)

				oldBalance = receiver.GetBalance()
				receiver, _ = ent.RefreshBankAccountFromDB(ctx, c, receiver)
				require.Greater(t, receiver.GetBalance(), oldBalance)

				feeTxc := c.Transaction.Query().Where(entTxc.SourceTransactionID(entity1.ID)).FirstX(ctx)
				require.Equal(t, *feeTxc.SenderID, receiver.ID)
				require.Equal(t, *feeTxc.SourceTransactionID, entity1.ID)
				require.Equal(t, feeTxc.Amount, decimal.NewFromFloat(float64(1000)))
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
			uc := transaction.NewCustomerTransactionConfirmSuccessUseCase(
				repository.GetTransactionConfirmSuccessRepository(c),
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
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerTransactionCreateUseCase)
	}{
		{
			name: "success",
			setUp: func(t *testing.T, ctx *context.Context, c *ent.Client) {
				authenticateCtx(ctx, c, nil)
				ent.EmbedClient(ctx, c)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionCreateUseCase) {
				i := ent.TransactionFactory(ctx)
				_, err := uc.Create(ctx, &model.TransactionCreateUseCaseInput{
					TransactionCreateInput: i,
					IsFeePaidByMe:          true,
				})
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
			cfg, _ := config.NewConfigForTest()
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			taskExecutorMock := task.NewMockIExecuteTask[*mail.EmailPayload](mockCtl)
			taskExecutorMock.EXPECT().ExecuteTask(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			uc := transaction.NewCustomerTransactionCreateUseCase(
				taskExecutorMock,
				repository.GetTransactionCreateRepository(c),
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
