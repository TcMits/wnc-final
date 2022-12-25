package ent

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/ent/bankaccount"
	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/password"
	"github.com/bluele/factory-go/factory"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

var customerFactory = factory.NewFactory(
	&CustomerCreateInput{},
).Attr("Password", func(a factory.Args) (interface{}, error) {
	pwd, err := password.GetHashPassword("123456789")
	return generic.GetPointer(pwd), err
}).SeqString("Username", func(s string) (interface{}, error) {
	return fmt.Sprintf("username%s", s), nil
}).Attr("FirstName", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer("Foo"), nil
}).Attr("LastName", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer("Foo"), nil
}).SeqString("PhoneNumber", func(s string) (interface{}, error) {
	return fmt.Sprintf("+8492345678%s", s), nil
}).SeqString("Email", func(s string) (interface{}, error) {
	return fmt.Sprintf("user-%s-@gmail.com", s), nil
}).Attr("IsActive", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(true), nil
})

var bankAccountFactory = factory.NewFactory(
	&BankAccountCreateInput{},
).Attr("CashIn", func(a factory.Args) (interface{}, error) {
	return float64(1), nil
}).Attr("CashOut", func(a factory.Args) (interface{}, error) {
	return float64(1), nil
}).SeqString("AccountNumber", func(s string) (interface{}, error) {
	return generic.GetPointer(fmt.Sprintf("account-number-%s", s)), nil
}).Attr("IsForPayment", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(false), nil
})

var transactionFactory = factory.NewFactory(
	&TransactionCreateInput{},
).Attr("Status", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(transaction.StatusDraft), nil
}).Attr("Amount", func(a factory.Args) (interface{}, error) {
	return decimal.NewFromInt32(1), nil
}).Attr("TransactionType", func(a factory.Args) (interface{}, error) {
	return transaction.TransactionTypeInternal, nil
})

func TransactionFactory() *TransactionCreateInput {
	return transactionFactory.MustCreate().(*TransactionCreateInput)
}
func CustomerFactory() *CustomerCreateInput {
	return customerFactory.MustCreate().(*CustomerCreateInput)
}
func BankAccountFactory() *BankAccountCreateInput {
	return bankAccountFactory.MustCreate().(*BankAccountCreateInput)
}

func CreateFakeCustomer(ctx context.Context, c *Client, i *CustomerCreateInput) (*Customer, error) {
	if i == nil {
		i = CustomerFactory()
	}
	if ctx == nil {
		ctx = context.Background()
	}
	return c.Customer.Create().SetInput(i).Save(ctx)
}
func CreateFakeBankAccount(ctx context.Context, c *Client, i *BankAccountCreateInput) (*BankAccount, error) {
	if i == nil {
		i = BankAccountFactory()
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if i.CustomerID == generic.Zero[uuid.UUID]() {
		ent1, err := CreateFakeCustomer(ctx, c, nil)
		if err != nil {
			return nil, err
		}
		i.CustomerID = ent1.ID
	}
	return c.BankAccount.Create().SetInput(i).Save(ctx)
}

func CreateFakeTransaction(ctx context.Context, c *Client, i *TransactionCreateInput) (*Transaction, error) {
	if i == nil {
		i = TransactionFactory()
	}
	if ctx == nil {
		ctx = context.Background()
	}
	var ent1, ent2 *BankAccount
	var err error
	if i.ReceiverID == nil {
		ent1, err = CreateFakeBankAccount(ctx, c, nil)
		if err != nil {
			return nil, err
		}
		i.ReceiverID = generic.GetPointer(ent1.ID)
	} else {
		ent1, err = c.BankAccount.Query().Where(bankaccount.ID(*i.ReceiverID)).First(ctx)
		if err != nil {
			return nil, err
		}
	}
	if i.SenderID == generic.Zero[uuid.UUID]() {
		ent2, err = CreateFakeBankAccount(ctx, c, nil)
		if err != nil {
			return nil, err
		}
		i.SenderID = ent2.ID
	} else {
		ent2, err = c.BankAccount.Query().Where(bankaccount.ID(i.SenderID)).First(ctx)
		if err != nil {
			return nil, err
		}
	}
	i.ReceiverBankAccountNumber = ent1.AccountNumber
	i.ReceiverName = NormalizeName(ent1.QueryCustomer().FirstX(ctx).FirstName, ent1.QueryCustomer().FirstX(ctx).LastName)
	i.SenderBankAccountNumber = ent2.AccountNumber
	i.SenderName = NormalizeName(ent2.QueryCustomer().FirstX(ctx).FirstName, ent2.QueryCustomer().FirstX(ctx).LastName)
	return c.Transaction.Create().SetInput(i).Save(ctx)
}
