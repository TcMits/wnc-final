package ent

import (
	"context"
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/TcMits/wnc-final/ent/bankaccount"
	"github.com/TcMits/wnc-final/ent/debt"
	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/password"
	"github.com/bluele/factory-go/factory"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Opt struct {
	Key   string
	Value any
}

var customerFactory = factory.NewFactory(
	&CustomerCreateInput{},
).Attr("Password", func(a factory.Args) (interface{}, error) {
	pwd, err := password.GetHashPassword("123456789")
	return generic.GetPointer(pwd), err
}).SeqString("Username", func(s string) (interface{}, error) {
	return fmt.Sprintf("username%s", s), nil
}).Attr("FirstName", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(randomdata.FirstName(randomdata.RandomGender)), nil
}).Attr("LastName", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(randomdata.LastName()), nil
}).SeqString("PhoneNumber", func(s string) (interface{}, error) {
	return fmt.Sprintf("%s%s", randomdata.PhoneNumber(), s), nil
}).SeqString("Email", func(s string) (interface{}, error) {
	return fmt.Sprintf("%s%s", s, randomdata.Email()), nil
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
	return generic.GetPointer(fmt.Sprintf("%s%s", randomdata.Digits(10), s)), nil
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

var debtFactory = factory.NewFactory(
	&DebtCreateInput{},
).Attr("Status", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(debt.StatusPending), nil
}).Attr("Amount", func(a factory.Args) (interface{}, error) {
	return decimal.NewFromInt32(1), nil
}).Attr("Description", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(randomdata.Paragraph()), nil
})

func getClient(ctx context.Context) (*Client, error) {
	client, ok := ctx.Value("client").(*Client)
	if !ok {
		return nil, fmt.Errorf("cannot find client in context")
	}
	return client, nil
}
func EmbedClient(ctx *context.Context, v *Client) {
	*ctx = context.WithValue(*ctx, "client", v)
}

var contactFactory = factory.NewFactory(
	&ContactCreateInput{
		BankName: "Bank name",
	},
).SeqString("AccountNumber", func(s string) (interface{}, error) {
	return fmt.Sprintf("%s%s", randomdata.Digits(10), s), nil
}).SeqString("SuggestName", func(s string) (interface{}, error) {
	return randomdata.FullName(randomdata.RandomGender), nil
}).Attr("OwnerID", func(a factory.Args) (interface{}, error) {
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	owner, err := CreateFakeCustomer(a.Context(), client, nil)
	if err != nil {
		return nil, err
	}
	return owner.ID, nil
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
func DebtFactory() *DebtCreateInput {
	return debtFactory.MustCreate().(*DebtCreateInput)
}

func ContactFactory(ctx context.Context, opts ...Opt) *ContactCreateInput {
	optMap := make(map[string]any)
	for _, opt := range opts {
		optMap[opt.Key] = opt.Value
	}
	return contactFactory.MustCreateWithContextAndOption(ctx, optMap).(*ContactCreateInput)
}

func CreateFakeDebt(ctx context.Context, c *Client, i *DebtCreateInput) (*Debt, error) {
	if i == nil {
		i = DebtFactory()
	}
	if ctx == nil {
		ctx = context.Background()
	}
	var ent1, ent2 *BankAccount
	var err error
	if i.ReceiverID == generic.Zero[uuid.UUID]() {
		ent1, err = CreateFakeBankAccount(ctx, c, nil)
		if err != nil {
			return nil, err
		}
		i.ReceiverID = ent1.ID
	} else {
		ent1, err = c.BankAccount.Query().Where(bankaccount.ID(i.ReceiverID)).First(ctx)
		if err != nil {
			return nil, err
		}
	}
	if i.OwnerID == generic.Zero[uuid.UUID]() {
		ent2, err = CreateFakeBankAccount(ctx, c, nil)
		if err != nil {
			return nil, err
		}
		i.OwnerID = ent2.ID
	} else {
		ent2, err = c.BankAccount.Query().Where(bankaccount.ID(i.OwnerID)).First(ctx)
		if err != nil {
			return nil, err
		}
	}
	i.ReceiverBankAccountNumber = ent1.AccountNumber
	i.ReceiverName = NormalizeName(ent1.QueryCustomer().FirstX(ctx).FirstName, ent1.QueryCustomer().FirstX(ctx).LastName)
	i.OwnerBankAccountNumber = ent2.AccountNumber
	i.OwnerName = NormalizeName(ent2.QueryCustomer().FirstX(ctx).FirstName, ent2.QueryCustomer().FirstX(ctx).LastName)
	return c.Debt.Create().SetInput(i).Save(ctx)
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
func CreateFakeContact(ctx context.Context, c *Client, i *ContactCreateInput) (*Contact, error) {
	if i == nil {
		i = ContactFactory(ctx)
	}
	if ctx == nil {
		ctx = context.Background()
	}
	return c.Contact.Create().SetInput(i).Save(ctx)
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
