package ent

import (
	"context"
	"fmt"
	"reflect"

	"github.com/Pallinder/go-randomdata"
	"github.com/TcMits/wnc-final/ent/bankaccount"
	"github.com/TcMits/wnc-final/ent/debt"
	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/password"
	"github.com/hyuti/factory-go/factory"
	"github.com/shopspring/decimal"
)

type (
	ClientCtxType string
	Opt           struct {
		Key   string
		Value any
	}
	IFactory[ModelType any] interface {
		Create(context.Context) (ModelType, error)
		CreateWithClient(context.Context, *Client) (ModelType, error)
	}

	Factory[ModelType any] struct {
		worker *factory.Factory
	}
)

const ClientCtxKey ClientCtxType = "client"

func (s *Factory[ModelType]) Create(ctx context.Context) (ModelType, error) {
	eAny, err := s.worker.CreateWithContext(ctx)
	if err != nil {
		return generic.Zero[ModelType](), err
	}
	e, ok := eAny.(ModelType)
	if !ok {
		return generic.Zero[ModelType](), fmt.Errorf("unexpected type %t", eAny)
	}
	return e, nil
}

func (s *Factory[ModelType]) CreateWithClient(ctx context.Context, client *Client) (ModelType, error) {
	EmbedClient(&ctx, client)
	return s.Create(ctx)
}

func getClient(ctx context.Context) (*Client, error) {
	client, ok := ctx.Value(ClientCtxKey).(*Client)
	if !ok || client == nil {
		return nil, fmt.Errorf("cannot find client in context")
	}
	return client, nil
}
func convertInputToOutput[ModelInputType, ModelType any](
	ctx context.Context,
	args factory.Args,
	factory *factory.Factory,
	saver func(context.Context, *Client, ModelInputType) (ModelType, error),
	opts ...Opt,
) error {
	optMap := make(map[string]any)
	for _, opt := range opts {
		optMap[opt.Key] = opt.Value
	}
	iAny, err := factory.CreateWithContextAndOption(ctx, optMap)
	if err != nil {
		return err
	}
	i, ok := iAny.(ModelInputType)
	if !ok {
		return fmt.Errorf("unexpected type %t", iAny)
	}
	client, err := getClient(ctx)
	if err != nil {
		return err
	}
	e, err := saver(ctx, client, i)
	if err != nil {
		return err
	}
	inst := args.Instance()
	dst := reflect.ValueOf(inst)
	src := reflect.ValueOf(e).Elem()
	dst.Elem().Set(src)
	return nil
}
func factoryTemplate[ModelType, ModelInputType any](
	model ModelType,
	f *factory.Factory,
	saver func(context.Context, *Client, ModelInputType) (ModelType, error),
	opts ...Opt,
) *factory.Factory {
	return factory.NewFactory(
		model,
	).OnCreate(func(a factory.Args) error {
		ctx := a.Context()
		err := convertInputToOutput(
			ctx,
			a,
			f,
			saver,
			opts...,
		)
		if err != nil {
			return err
		}
		return nil
	})
}
func EmbedClient(ctx *context.Context, v *Client) {
	c := *ctx
	client := c.Value(ClientCtxKey)
	if client == nil {
		*ctx = context.WithValue(*ctx, ClientCtxKey, v)
	}
}

var partnerFactory = factory.NewFactory(
	func() *PartnerCreateInput {
		pair, _ := password.GenerateRSAKeyPair()
		return &PartnerCreateInput{
			PublicKey:  pair.PublicKey,
			PrivateKey: pair.PrivateKey,
		}
	}(),
).Attr("APIKey", func(a factory.Args) (interface{}, error) {
	return randomdata.Alphanumeric(30), nil
}).Attr("SecretKey", func(a factory.Args) (interface{}, error) {
	return randomdata.Alphanumeric(30), nil
}).Attr("Name", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(randomdata.FullName(randomdata.RandomGender)), nil
})

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
	return fmt.Sprintf("+849%s%s", randomdata.Digits(7), s), nil
}).SeqString("Email", func(s string) (interface{}, error) {
	return fmt.Sprintf("%s%s", s, randomdata.Email()), nil
}).Attr("IsActive", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(true), nil
})

var employeeFactory = factory.NewFactory(
	&EmployeeCreateInput{},
).Attr("Password", func(a factory.Args) (interface{}, error) {
	pwd, err := password.GetHashPassword("123456789")
	return generic.GetPointer(pwd), err
}).SeqString("Username", func(s string) (interface{}, error) {
	return fmt.Sprintf("username%s", s), nil
}).Attr("FirstName", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(randomdata.FirstName(randomdata.RandomGender)), nil
}).Attr("LastName", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(randomdata.LastName()), nil
}).Attr("IsActive", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(true), nil
})
var adminFactory = factory.NewFactory(
	&AdminCreateInput{},
).Attr("Password", func(a factory.Args) (interface{}, error) {
	pwd, err := password.GetHashPassword("123456789")
	return generic.GetPointer(pwd), err
}).SeqString("Username", func(s string) (interface{}, error) {
	return fmt.Sprintf("username%s", s), nil
}).Attr("FirstName", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(randomdata.FirstName(randomdata.RandomGender)), nil
}).Attr("LastName", func(a factory.Args) (interface{}, error) {
	return generic.GetPointer(randomdata.LastName()), nil
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
	return generic.GetPointer(fmt.Sprintf("%s%s", randomdata.Digits(16), s)), nil
}).SubFactory("CustomerID", CustomerFactory(), func(i interface{}) (interface{}, error) {
	e, ok := i.(*Customer)
	if !ok {
		return nil, fmt.Errorf("unexpected type %t", i)
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
).SubFactory("SenderID", BankAccountFactory(), func(i interface{}) (interface{}, error) {
	e, ok := i.(*BankAccount)
	if !ok {
		return nil, fmt.Errorf("unexpected type %t", i)
	}
	return &e.ID, nil
}).Attr("SenderName", func(a factory.Args) (interface{}, error) {
	ins := a.Instance().(*TransactionCreateInput)
	sid := ins.SenderID
	client, err := getClient(a.Context())
	if err != nil {
		return nil, err
	}
	ba, err := client.BankAccount.Query().Where(bankaccount.ID(*sid)).First(a.Context())
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
	ba, err := client.BankAccount.Query().Where(bankaccount.ID(*sid)).First(a.Context())
	if err != nil {
		return nil, err
	}
	return ba.AccountNumber, nil
}).SubFactory("ReceiverID", BankAccountFactory(), func(i interface{}) (interface{}, error) {
	e, ok := i.(*BankAccount)
	if !ok {
		return nil, fmt.Errorf("unexpected type %t", i)
	}
	return &e.ID, nil

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
}).SubFactory("OwnerID", BankAccountFactory(Opt{"IsForPayment", generic.GetPointer(true)}), func(i interface{}) (interface{}, error) {
	e, ok := i.(*BankAccount)
	if !ok {
		return nil, fmt.Errorf("unexpected type %t", i)
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
}).SubFactory("ReceiverID", BankAccountFactory(Opt{"IsForPayment", generic.GetPointer(true)}), func(i interface{}) (interface{}, error) {
	e, ok := i.(*BankAccount)
	if !ok {
		return nil, fmt.Errorf("unexpected type %t", i)
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
).SubFactory("AccountNumber", BankAccountFactory(Opt{"IsForPayment", generic.GetPointer(true)}), func(i interface{}) (interface{}, error) {
	e, ok := i.(*BankAccount)
	if !ok {
		return nil, fmt.Errorf("unexpected type %t", i)
	}
	return e.AccountNumber, nil
}).SeqString("SuggestName", func(s string) (interface{}, error) {
	return randomdata.FullName(randomdata.RandomGender), nil
}).SubFactory("OwnerID", CustomerFactory(), func(i interface{}) (interface{}, error) {
	e, ok := i.(*Customer)
	if !ok {
		return nil, fmt.Errorf("unexpected type %t", i)
	}
	return e.ID, nil
})

func TransactionFactory(opts ...Opt) *factory.Factory {
	return factoryTemplate(
		new(Transaction),
		transactionFactory,
		func(ctx context.Context, client *Client, i *TransactionCreateInput) (*Transaction, error) {
			return client.Transaction.Create().SetInput(i).Save(ctx)
		},
		opts...,
	)
}
func MustTransactionFactory(opts ...Opt) IFactory[*Transaction] {
	return &Factory[*Transaction]{
		worker: TransactionFactory(opts...),
	}
}
func EmployeeFactory(opts ...Opt) *factory.Factory {
	return factoryTemplate(
		new(Employee),
		employeeFactory,
		func(ctx context.Context, client *Client, i *EmployeeCreateInput) (*Employee, error) {
			return client.Employee.Create().SetInput(i).Save(ctx)
		},
		opts...,
	)
}
func MustEmployeeFactory(opts ...Opt) IFactory[*Employee] {
	return &Factory[*Employee]{
		worker: EmployeeFactory(opts...),
	}
}
func AdminFactory(opts ...Opt) *factory.Factory {
	return factoryTemplate(
		new(Admin),
		adminFactory,
		func(ctx context.Context, client *Client, i *AdminCreateInput) (*Admin, error) {
			return client.Admin.Create().SetInput(i).Save(ctx)
		},
		opts...,
	)
}
func MustAdminFactory(opts ...Opt) IFactory[*Admin] {
	return &Factory[*Admin]{
		worker: AdminFactory(opts...),
	}
}
func DebtFactory(opts ...Opt) *factory.Factory {
	return factoryTemplate(
		new(Debt),
		debtFactory,
		func(ctx context.Context, client *Client, i *DebtCreateInput) (*Debt, error) {
			return client.Debt.Create().SetInput(i).Save(ctx)
		},
		opts...,
	)
}
func MustDebtFactory(opts ...Opt) IFactory[*Debt] {
	return &Factory[*Debt]{
		worker: DebtFactory(opts...),
	}
}

func ContactFactory(opts ...Opt) *factory.Factory {
	return factoryTemplate(
		new(Contact),
		contactFactory,
		func(ctx context.Context, client *Client, i *ContactCreateInput) (*Contact, error) {
			return client.Contact.Create().SetInput(i).Save(ctx)
		},
		opts...,
	)
}
func MustContactFactory(opts ...Opt) IFactory[*Contact] {
	return &Factory[*Contact]{
		worker: ContactFactory(opts...),
	}
}

func PartnerFactory(opts ...Opt) *factory.Factory {
	return factoryTemplate(
		new(Partner),
		partnerFactory,
		func(ctx context.Context, client *Client, i *PartnerCreateInput) (*Partner, error) {
			return client.Partner.Create().SetInput(i).Save(ctx)
		},
		opts...,
	)
}
func MustPartnerFactory(opts ...Opt) IFactory[*Partner] {
	return &Factory[*Partner]{
		worker: PartnerFactory(opts...),
	}
}
func CustomerFactory(opts ...Opt) *factory.Factory {
	return factoryTemplate(
		new(Customer),
		customerFactory,
		func(ctx context.Context, client *Client, i *CustomerCreateInput) (*Customer, error) {
			return client.Customer.Create().SetInput(i).Save(ctx)
		},
		opts...,
	)
}
func MustCustomerFactory(opts ...Opt) IFactory[*Customer] {
	return &Factory[*Customer]{
		worker: CustomerFactory(opts...),
	}
}

func BankAccountFactory(opts ...Opt) *factory.Factory {
	return factoryTemplate(
		new(BankAccount),
		bankAccountFactory,
		func(ctx context.Context, client *Client, i *BankAccountCreateInput) (*BankAccount, error) {
			return client.BankAccount.Create().SetInput(i).Save(ctx)
		},
		opts...,
	)
}

func MustBankAccountFactory(opts ...Opt) IFactory[*BankAccount] {
	return &Factory[*BankAccount]{
		worker: BankAccountFactory(opts...),
	}
}
