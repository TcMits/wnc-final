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
	"github.com/TcMits/wnc-final/pkg/infrastructure/backgroundserver"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	"github.com/TcMits/wnc-final/ent"
)

func TestListUseCase(t *testing.T) {
	t.Parallel()
	c, _ := datastore.NewClientTestConnection(t)
	defer c.Close()
	ctx := context.Background()
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
	mBA, _ := ent.CreateFakeBankAccount(ctx, c, nil)
	ctx = context.WithValue(ctx, "user", mBA.QueryCustomer().FirstX(ctx))
	txc1 := ent.TransactionFactory()
	txc2 := ent.TransactionFactory()
	txc1.SenderID = mBA.ID
	txc2.ReceiverID = &mBA.ID
	ent.CreateFakeTransaction(ctx, c, txc1)
	ent.CreateFakeTransaction(ctx, c, txc2)
	ent.CreateFakeTransaction(ctx, c, nil)
	uc := transaction.NewCustomerTransactionListMyTxcUseCase(repository.GetTransactionListRepository(c))
	l, o := 3, 0
	result, err := uc.ListMyTxc(ctx, &l, &o, nil, nil)
	require.Nil(t, err)
	require.Equal(t, 2, len(result))
}

func TestGetFirstMyTxcUseCase(t *testing.T) {
	t.Parallel()
	c, _ := datastore.NewClientTestConnection(t)
	defer c.Close()
	ctx := context.Background()
	mBA, _ := ent.CreateFakeBankAccount(ctx, c, nil)
	ctx = context.WithValue(ctx, "user", mBA.QueryCustomer().FirstX(ctx))
	txc1 := ent.TransactionFactory()
	txc1.SenderID = mBA.ID
	entity1, _ := ent.CreateFakeTransaction(ctx, c, txc1)
	ent.CreateFakeTransaction(ctx, c, nil)
	uc := transaction.NewCustomerTransactionGetFirstMyTxUseCase(repository.GetTransactionListRepository(c))
	result, err := uc.GetFirstMyTxc(ctx, nil, nil)
	require.Nil(t, err)
	require.Equal(t, entity1.ID, result.ID)
}

func TestUpdateUseCase(t *testing.T) {
	t.Parallel()
	c, _ := datastore.NewClientTestConnection(t)
	defer c.Close()
	ctx := context.Background()
	entity1, _ := ent.CreateFakeTransaction(ctx, c, nil)
	user := entity1.QuerySender().FirstX(ctx).QueryCustomer().FirstX(ctx)
	ctx = context.WithValue(ctx, "user", user)
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
		setUp  func(*testing.T, context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerTransactionValidateCreateInputUseCase)
	}{
		{
			name:  "bank account sender does have draft transactions",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateCreateInputUseCase) {
				i1 := ent.BankAccountFactory()
				i1.IsForPayment = generic.GetPointer(true)
				ba, _ := ent.CreateFakeBankAccount(ctx, c, i1)
				user := ba.QueryCustomer().FirstX(ctx)
				ctx = context.WithValue(ctx, "user", user)
				i2 := ent.TransactionFactory()
				i2.SenderID = ba.ID
				ent.CreateFakeTransaction(ctx, c, i2)
				i3 := ent.TransactionFactory()
				i3.SenderID = ba.ID
				_, err := uc.Validate(ctx, i3, true)
				require.ErrorContains(t, err, "there is a draft transaction to be processed. Cannot create a new transaction")
			},
		},
		{
			name:  "insufficient balance from sender and fee paid by me",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateCreateInputUseCase) {
				i1 := ent.BankAccountFactory()
				i1.IsForPayment = generic.GetPointer(true)
				ba, _ := ent.CreateFakeBankAccount(ctx, c, i1)
				user := ba.QueryCustomer().FirstX(ctx)
				ctx = context.WithValue(ctx, "user", user)
				i2 := ent.TransactionFactory()
				i2.SenderID = ba.ID
				_, err := uc.Validate(ctx, i2, true)
				require.ErrorContains(t, err, "insufficient balance")
			},
		},
		{
			name:  "insufficient balance from sender and fee not paid by me",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateCreateInputUseCase) {
				i1 := ent.BankAccountFactory()
				i1.IsForPayment = generic.GetPointer(true)
				ba, _ := ent.CreateFakeBankAccount(ctx, c, i1)
				user := ba.QueryCustomer().FirstX(ctx)
				ctx = context.WithValue(ctx, "user", user)
				i2 := ent.TransactionFactory()
				i2.SenderID = ba.ID
				_, err := uc.Validate(ctx, i2, false)
				require.ErrorContains(t, err, "insufficient balance")
			},
		},
		{
			name:  "insufficient balance from receiver and fee not paid by me",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateCreateInputUseCase) {
				i1 := ent.BankAccountFactory()
				i2 := ent.BankAccountFactory()
				i1.IsForPayment = generic.GetPointer(true)
				i2.IsForPayment = generic.GetPointer(true)
				i1.CashIn = float64(100000)
				i1.CashOut = float64(1)
				sender, _ := ent.CreateFakeBankAccount(ctx, c, i1)
				receiver, _ := ent.CreateFakeBankAccount(ctx, c, i2)
				user := sender.QueryCustomer().FirstX(ctx)
				ctx = context.WithValue(ctx, "user", user)
				i3 := ent.TransactionFactory()
				i3.SenderID = sender.ID
				i3.ReceiverID = &receiver.ID
				_, err := uc.Validate(ctx, i3, false)
				require.ErrorContains(t, err, "insufficient balance")
			},
		},
		{
			name:  "success",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateCreateInputUseCase) {
				i1 := ent.BankAccountFactory()
				i2 := ent.BankAccountFactory()
				i1.IsForPayment = generic.GetPointer(true)
				i2.IsForPayment = generic.GetPointer(true)
				i1.CashIn = float64(100000)
				i1.CashOut = float64(1)
				i2.CashIn = float64(100000)
				i2.CashOut = float64(1)
				sender, _ := ent.CreateFakeBankAccount(ctx, c, i1)
				receiver, _ := ent.CreateFakeBankAccount(ctx, c, i2)
				user := sender.QueryCustomer().FirstX(ctx)
				ctx = context.WithValue(ctx, "user", user)
				i3 := ent.TransactionFactory()
				i3.SenderID = sender.ID
				i3.ReceiverID = &receiver.ID
				i3, err := uc.Validate(ctx, i3, true)
				require.Nil(t, err)
				require.Equal(t, entTxc.StatusDraft.String(), i3.Status.String())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := datastore.NewClientTestConnection(t)
			defer c.Close()
			ctx := context.Background()
			require.NoError(t, c.Schema.Create(ctx))
			tt.setUp(t, ctx, c)
			uc := transaction.NewCustomerTransactionValidateCreateInputUseCase(
				repository.GetTransactionListRepository(c),
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

func TestValidateConfirmInputUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerTransactionValidateConfirmInputUseCase)
	}{
		{
			name: "not draft transaction",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {
				user, _ := ent.CreateFakeCustomer(ctx, c, nil)
				ctx = context.WithValue(ctx, "user", user)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateConfirmInputUseCase) {
				i1 := ent.TransactionFactory()
				i1.Status = generic.GetPointer(entTxc.StatusSuccess)
				entity1, _ := ent.CreateFakeTransaction(ctx, c, i1)
				err := uc.ValidateConfirmInput(ctx, entity1, nil, nil)
				require.ErrorContains(t, err, fmt.Sprintf("cannot confirm %s transaction", entity1.Status))
			},
		},
		{
			name: "token invalid: not have field",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {
				user, _ := ent.CreateFakeCustomer(ctx, c, nil)
				ctx = context.WithValue(ctx, "user", user)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateConfirmInputUseCase) {
				tk, _ := usecase.GenerateConfirmTxcToken(
					ctx,
					map[string]any{},
					"foo",
					time.Minute*30,
				)
				entity1, _ := ent.CreateFakeTransaction(ctx, c, nil)
				err := uc.ValidateConfirmInput(ctx, entity1, nil, &tk)
				require.ErrorContains(t, err, "invalid token")
			},
		},
		{
			name: "token invalid: have field but invalid type",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {
				user, _ := ent.CreateFakeCustomer(ctx, c, nil)
				ctx = context.WithValue(ctx, "user", user)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateConfirmInputUseCase) {
				tk, _ := usecase.GenerateConfirmTxcToken(
					ctx,
					map[string]any{"is_fee_paid_by_me": "foo"},
					"foo",
					time.Minute*30,
				)
				entity1, _ := ent.CreateFakeTransaction(ctx, c, nil)
				err := uc.ValidateConfirmInput(ctx, entity1, nil, &tk)
				require.ErrorContains(t, err, "invalid token")
			},
		},
		{
			name: "success",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {
				user, _ := ent.CreateFakeCustomer(ctx, c, nil)
				ctx = context.WithValue(ctx, "user", user)
			},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionValidateConfirmInputUseCase) {
				otp := usecase.GenerateOTP(6)
				hashValue, _ := usecase.GenerateHashInfo(usecase.MakeOTPValue(ctx, otp))
				tk, _ := usecase.GenerateConfirmTxcToken(
					ctx,
					map[string]any{
						"is_fee_paid_by_me": true,
						"token":             hashValue,
					},
					"foo",
					time.Minute*30,
				)
				entity1, _ := ent.CreateFakeTransaction(ctx, c, nil)
				err := uc.ValidateConfirmInput(ctx, entity1, &otp, &tk)
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
			tt.setUp(t, ctx, c)
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
		setUp  func(*testing.T, context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client, usecase.ICustomerTransactionConfirmSuccessUseCase)
	}{
		{
			name:  "success and fee paid by me",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionConfirmSuccessUseCase) {
				i1 := ent.BankAccountFactory()
				i2 := ent.BankAccountFactory()
				i1.IsForPayment = generic.GetPointer(true)
				i2.IsForPayment = generic.GetPointer(true)
				i1.CashIn = float64(100000)
				i1.CashOut = float64(1)
				i2.CashIn = float64(100000)
				i2.CashOut = float64(1)
				sender, _ := ent.CreateFakeBankAccount(ctx, c, i1)
				receiver, _ := ent.CreateFakeBankAccount(ctx, c, i2)
				user := sender.QueryCustomer().FirstX(ctx)
				ctx = context.WithValue(ctx, "user", user)
				i3 := ent.TransactionFactory()
				i3.SenderID = sender.ID
				i3.ReceiverID = &receiver.ID
				entity1, _ := ent.CreateFakeTransaction(ctx, c, i3)
				tk, _ := usecase.GenerateConfirmTxcToken(
					ctx,
					map[string]any{"is_fee_paid_by_me": true},
					"foo",
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
			name:  "success and fee not paid by me",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client, uc usecase.ICustomerTransactionConfirmSuccessUseCase) {
				i1 := ent.BankAccountFactory()
				i2 := ent.BankAccountFactory()
				i1.IsForPayment = generic.GetPointer(true)
				i2.IsForPayment = generic.GetPointer(true)
				i1.CashIn = float64(100000)
				i1.CashOut = float64(1)
				i2.CashIn = float64(100000)
				i2.CashOut = float64(1)
				sender, _ := ent.CreateFakeBankAccount(ctx, c, i1)
				receiver, _ := ent.CreateFakeBankAccount(ctx, c, i2)
				user := sender.QueryCustomer().FirstX(ctx)
				ctx = context.WithValue(ctx, "user", user)
				i3 := ent.TransactionFactory()
				i3.SenderID = sender.ID
				i3.ReceiverID = &receiver.ID
				entity1, _ := ent.CreateFakeTransaction(ctx, c, i3)
				tk, _ := usecase.GenerateConfirmTxcToken(
					ctx,
					map[string]any{"is_fee_paid_by_me": false},
					"foo",
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
			tt.setUp(t, ctx, c)
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
	cfg, _ := config.NewConfig()
	c, _ := datastore.NewClientTestConnection(t)
	cTask := backgroundserver.NewClient(cfg.Redis.URL, cfg.Redis.Password, cfg.Redis.DB)
	exc := task.GetEmailTaskExecutor(cTask)
	defer c.Close()
	ctx := context.Background()
	i := ent.TransactionFactory()
	s, _ := ent.CreateFakeBankAccount(ctx, c, nil)
	r, _ := ent.CreateFakeBankAccount(ctx, c, nil)
	i.SenderID = s.ID
	i.ReceiverID = &r.ID
	ent.CreateFakeTransaction(ctx, c, i)
	uc := transaction.NewCustomerTransactionCreateUseCase(
		exc,
		repository.GetTransactionCreateRepository(c),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeAmount,
		&cfg.TransactionUseCase.FeeDesc,
	)
	_, err := uc.Create(ctx, i, true)
	require.Nil(t, err)
}
