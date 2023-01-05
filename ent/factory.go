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
	"github.com/shopspring/decimal"
)

type Opt struct {
	Key   string
	Value any
}

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
	&BankAccountCreateInput{
		CashIn:       float64(1),
		CashOut:      float64(1),
		IsForPayment: generic.GetPointer(false),
	},
).SeqString("AccountNumber", func(s string) (interface{}, error) {
	return generic.GetPointer(fmt.Sprintf("%s%s", randomdata.Digits(10), s)), nil
}).Attr("CustomerID", func(a factory.Args) (interface{}, error) {
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	e, err := CreateFakeCustomer(a.Context(), client, nil)
	if err != nil {
		return nil, err
	}
	return e.ID, nil
})

var transactionFactory = factory.NewFactory(
	&TransactionCreateInput{
		SenderBankName:   "Sender bank",
		ReceiverBankName: "Receiver bank",
		Status:           generic.GetPointer(transaction.StatusDraft),
		Amount:           decimal.NewFromInt32(1),
		TransactionType:  transaction.TransactionTypeInternal,
	},
).Attr("SenderID", func(a factory.Args) (interface{}, error) {
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	e, err := CreateFakeBankAccount(a.Context(), client, nil, Opt{"IsForPayment", generic.GetPointer(true)})
	fmt.Println(e.ID)
	if err != nil {
		return nil, err
	}
	return e.ID, nil
}).Attr("SenderName", func(a factory.Args) (interface{}, error) {
	ins := a.Instance().(*TransactionCreateInput)
	sid := ins.SenderID
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	ba, err := client.BankAccount.Query().Where(bankaccount.ID(sid)).First(a.Context())
	if err != nil {
		return nil, err
	}
	user, err := ba.QueryCustomer().First(a.Context())
	return user.GetName(), err
}).Attr("SenderBankAccountNumber", func(a factory.Args) (interface{}, error) {
	ins := a.Instance().(*TransactionCreateInput)
	sid := ins.SenderID
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	ba, err := client.BankAccount.Query().Where(bankaccount.ID(sid)).First(a.Context())
	if err != nil {
		return nil, err
	}
	return ba.AccountNumber, nil
}).Attr("ReceiverID", func(a factory.Args) (interface{}, error) {
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	e, err := CreateFakeBankAccount(a.Context(), client, nil, Opt{"IsForPayment", generic.GetPointer(true)})
	if err != nil {
		return nil, err
	}
	return generic.GetPointer(e.ID), nil
}).Attr("ReceiverName", func(a factory.Args) (interface{}, error) {
	ins := a.Instance().(*TransactionCreateInput)
	sid := *ins.ReceiverID
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	ba, err := client.BankAccount.Query().Where(bankaccount.ID(sid)).First(a.Context())
	if err != nil {
		return nil, err
	}
	user, err := ba.QueryCustomer().First(a.Context())
	return user.GetName(), err
}).Attr("ReceiverBankAccountNumber", func(a factory.Args) (interface{}, error) {
	ins := a.Instance().(*TransactionCreateInput)
	sid := ins.ReceiverID
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	ba, err := client.BankAccount.Query().Where(bankaccount.ID(*sid)).First(a.Context())
	if err != nil {
		return nil, err
	}
	return ba.AccountNumber, nil
})

var debtFactory = factory.NewFactory(
	&DebtCreateInput{
		Status:           generic.GetPointer(debt.StatusPending),
		Amount:           decimal.NewFromInt32(1),
		OwnerBankName:    "Owner bank name",
		ReceiverBankName: "Receiver bank name",
	},
).Attr("Description", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(randomdata.Paragraph()), nil
}).Attr("OwnerID", func(a factory.Args) (interface{}, error) {
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	e, err := CreateFakeBankAccount(a.Context(), client, nil, Opt{"IsForPayment", generic.GetPointer(true)})
	if err != nil {
		return nil, err
	}
	return e.ID, nil
}).Attr("OwnerName", func(a factory.Args) (interface{}, error) {
	ins := a.Instance().(*DebtCreateInput)
	id := ins.OwnerID
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	e, err := client.BankAccount.Query().Where(bankaccount.ID(id)).First(a.Context())
	if err != nil {
		return nil, err
	}
	e1, err := e.QueryCustomer().First(a.Context())
	if err != nil {
		return nil, err
	}
	return e1.GetName(), nil
}).Attr("OwnerBankAccountNumber", func(a factory.Args) (interface{}, error) {
	ins := a.Instance().(*DebtCreateInput)
	id := ins.OwnerID
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	e, err := client.BankAccount.Query().Where(bankaccount.ID(id)).First(a.Context())
	if err != nil {
		return nil, err
	}
	return e.AccountNumber, nil
}).Attr("ReceiverID", func(a factory.Args) (interface{}, error) {
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	e, err := CreateFakeBankAccount(a.Context(), client, nil, Opt{"IsForPayment", generic.GetPointer(true)})
	if err != nil {
		return nil, err
	}
	return e.ID, nil
}).Attr("ReceiverName", func(a factory.Args) (interface{}, error) {
	ins := a.Instance().(*DebtCreateInput)
	id := ins.ReceiverID
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	e, err := client.BankAccount.Query().Where(bankaccount.ID(id)).First(a.Context())
	if err != nil {
		return nil, err
	}
	e1, err := e.QueryCustomer().First(a.Context())
	if err != nil {
		return nil, err
	}
	return e1.GetName(), nil
}).Attr("ReceiverBankAccountNumber", func(a factory.Args) (interface{}, error) {
	ins := a.Instance().(*DebtCreateInput)
	id := ins.ReceiverID
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	e, err := client.BankAccount.Query().Where(bankaccount.ID(id)).First(a.Context())
	if err != nil {
		return nil, err
	}
	return e.AccountNumber, nil
})

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

func TransactionFactory(ctx context.Context, opts ...Opt) *TransactionCreateInput {
	optMap := make(map[string]any)
	for _, opt := range opts {
		optMap[opt.Key] = opt.Value
	}
	return transactionFactory.MustCreateWithContextAndOption(ctx, optMap).(*TransactionCreateInput)
}
func CustomerFactory(ctx context.Context, opts ...Opt) *CustomerCreateInput {
	optMap := make(map[string]any)
	for _, opt := range opts {
		optMap[opt.Key] = opt.Value
	}
	return customerFactory.MustCreateWithContextAndOption(ctx, optMap).(*CustomerCreateInput)
}
func BankAccountFactory(ctx context.Context, opts ...Opt) *BankAccountCreateInput {
	optMap := make(map[string]any)
	for _, opt := range opts {
		optMap[opt.Key] = opt.Value
	}
	return bankAccountFactory.MustCreateWithContextAndOption(ctx, optMap).(*BankAccountCreateInput)
}
func DebtFactory(ctx context.Context, opts ...Opt) *DebtCreateInput {
	optMap := make(map[string]any)
	for _, opt := range opts {
		optMap[opt.Key] = opt.Value
	}
	return debtFactory.MustCreateWithContextAndOption(ctx, optMap).(*DebtCreateInput)
}

func ContactFactory(ctx context.Context, opts ...Opt) *ContactCreateInput {
	optMap := make(map[string]any)
	for _, opt := range opts {
		optMap[opt.Key] = opt.Value
	}
	return contactFactory.MustCreateWithContextAndOption(ctx, optMap).(*ContactCreateInput)
}

func CreateFakeDebt(ctx context.Context, c *Client, i *DebtCreateInput, opts ...Opt) (*Debt, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if i == nil {
		i = DebtFactory(ctx, opts...)
	}
	return c.Debt.Create().SetInput(i).Save(ctx)
}

func CreateFakeCustomer(ctx context.Context, c *Client, i *CustomerCreateInput, opts ...Opt) (*Customer, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if i == nil {
		i = CustomerFactory(ctx, opts...)
	}
	return c.Customer.Create().SetInput(i).Save(ctx)
}
func CreateFakeBankAccount(ctx context.Context, c *Client, i *BankAccountCreateInput, opts ...Opt) (*BankAccount, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if i == nil {
		i = BankAccountFactory(ctx, opts...)
	}
	return c.BankAccount.Create().SetInput(i).Save(ctx)
}
func CreateFakeContact(ctx context.Context, c *Client, i *ContactCreateInput) (*Contact, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if i == nil {
		i = ContactFactory(ctx)
	}
	return c.Contact.Create().SetInput(i).Save(ctx)
}

func CreateFakeTransaction(ctx context.Context, c *Client, i *TransactionCreateInput, opts ...Opt) (*Transaction, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if i == nil {
		i = TransactionFactory(ctx, opts...)
	}
	return c.Transaction.Create().SetInput(i).Save(ctx)
}
