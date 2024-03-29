// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/TcMits/wnc-final/ent/migrate"
	"github.com/google/uuid"

	"github.com/TcMits/wnc-final/ent/admin"
	"github.com/TcMits/wnc-final/ent/bankaccount"
	"github.com/TcMits/wnc-final/ent/contact"
	"github.com/TcMits/wnc-final/ent/customer"
	"github.com/TcMits/wnc-final/ent/debt"
	"github.com/TcMits/wnc-final/ent/employee"
	"github.com/TcMits/wnc-final/ent/partner"
	"github.com/TcMits/wnc-final/ent/transaction"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Admin is the client for interacting with the Admin builders.
	Admin *AdminClient
	// BankAccount is the client for interacting with the BankAccount builders.
	BankAccount *BankAccountClient
	// Contact is the client for interacting with the Contact builders.
	Contact *ContactClient
	// Customer is the client for interacting with the Customer builders.
	Customer *CustomerClient
	// Debt is the client for interacting with the Debt builders.
	Debt *DebtClient
	// Employee is the client for interacting with the Employee builders.
	Employee *EmployeeClient
	// Partner is the client for interacting with the Partner builders.
	Partner *PartnerClient
	// Transaction is the client for interacting with the Transaction builders.
	Transaction *TransactionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Admin = NewAdminClient(c.config)
	c.BankAccount = NewBankAccountClient(c.config)
	c.Contact = NewContactClient(c.config)
	c.Customer = NewCustomerClient(c.config)
	c.Debt = NewDebtClient(c.config)
	c.Employee = NewEmployeeClient(c.config)
	c.Partner = NewPartnerClient(c.config)
	c.Transaction = NewTransactionClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Admin:       NewAdminClient(cfg),
		BankAccount: NewBankAccountClient(cfg),
		Contact:     NewContactClient(cfg),
		Customer:    NewCustomerClient(cfg),
		Debt:        NewDebtClient(cfg),
		Employee:    NewEmployeeClient(cfg),
		Partner:     NewPartnerClient(cfg),
		Transaction: NewTransactionClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Admin:       NewAdminClient(cfg),
		BankAccount: NewBankAccountClient(cfg),
		Contact:     NewContactClient(cfg),
		Customer:    NewCustomerClient(cfg),
		Debt:        NewDebtClient(cfg),
		Employee:    NewEmployeeClient(cfg),
		Partner:     NewPartnerClient(cfg),
		Transaction: NewTransactionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Admin.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Admin.Use(hooks...)
	c.BankAccount.Use(hooks...)
	c.Contact.Use(hooks...)
	c.Customer.Use(hooks...)
	c.Debt.Use(hooks...)
	c.Employee.Use(hooks...)
	c.Partner.Use(hooks...)
	c.Transaction.Use(hooks...)
}

// AdminClient is a client for the Admin schema.
type AdminClient struct {
	config
}

// NewAdminClient returns a client for the Admin from the given config.
func NewAdminClient(c config) *AdminClient {
	return &AdminClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `admin.Hooks(f(g(h())))`.
func (c *AdminClient) Use(hooks ...Hook) {
	c.hooks.Admin = append(c.hooks.Admin, hooks...)
}

// Create returns a builder for creating a Admin entity.
func (c *AdminClient) Create() *AdminCreate {
	mutation := newAdminMutation(c.config, OpCreate)
	return &AdminCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Admin entities.
func (c *AdminClient) CreateBulk(builders ...*AdminCreate) *AdminCreateBulk {
	return &AdminCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Admin.
func (c *AdminClient) Update() *AdminUpdate {
	mutation := newAdminMutation(c.config, OpUpdate)
	return &AdminUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AdminClient) UpdateOne(a *Admin) *AdminUpdateOne {
	mutation := newAdminMutation(c.config, OpUpdateOne, withAdmin(a))
	return &AdminUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AdminClient) UpdateOneID(id uuid.UUID) *AdminUpdateOne {
	mutation := newAdminMutation(c.config, OpUpdateOne, withAdminID(id))
	return &AdminUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Admin.
func (c *AdminClient) Delete() *AdminDelete {
	mutation := newAdminMutation(c.config, OpDelete)
	return &AdminDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AdminClient) DeleteOne(a *Admin) *AdminDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AdminClient) DeleteOneID(id uuid.UUID) *AdminDeleteOne {
	builder := c.Delete().Where(admin.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AdminDeleteOne{builder}
}

// Query returns a query builder for Admin.
func (c *AdminClient) Query() *AdminQuery {
	return &AdminQuery{
		config: c.config,
	}
}

// Get returns a Admin entity by its id.
func (c *AdminClient) Get(ctx context.Context, id uuid.UUID) (*Admin, error) {
	return c.Query().Where(admin.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AdminClient) GetX(ctx context.Context, id uuid.UUID) *Admin {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AdminClient) Hooks() []Hook {
	return c.hooks.Admin
}

// BankAccountClient is a client for the BankAccount schema.
type BankAccountClient struct {
	config
}

// NewBankAccountClient returns a client for the BankAccount from the given config.
func NewBankAccountClient(c config) *BankAccountClient {
	return &BankAccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `bankaccount.Hooks(f(g(h())))`.
func (c *BankAccountClient) Use(hooks ...Hook) {
	c.hooks.BankAccount = append(c.hooks.BankAccount, hooks...)
}

// Create returns a builder for creating a BankAccount entity.
func (c *BankAccountClient) Create() *BankAccountCreate {
	mutation := newBankAccountMutation(c.config, OpCreate)
	return &BankAccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BankAccount entities.
func (c *BankAccountClient) CreateBulk(builders ...*BankAccountCreate) *BankAccountCreateBulk {
	return &BankAccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BankAccount.
func (c *BankAccountClient) Update() *BankAccountUpdate {
	mutation := newBankAccountMutation(c.config, OpUpdate)
	return &BankAccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BankAccountClient) UpdateOne(ba *BankAccount) *BankAccountUpdateOne {
	mutation := newBankAccountMutation(c.config, OpUpdateOne, withBankAccount(ba))
	return &BankAccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BankAccountClient) UpdateOneID(id uuid.UUID) *BankAccountUpdateOne {
	mutation := newBankAccountMutation(c.config, OpUpdateOne, withBankAccountID(id))
	return &BankAccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BankAccount.
func (c *BankAccountClient) Delete() *BankAccountDelete {
	mutation := newBankAccountMutation(c.config, OpDelete)
	return &BankAccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BankAccountClient) DeleteOne(ba *BankAccount) *BankAccountDeleteOne {
	return c.DeleteOneID(ba.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BankAccountClient) DeleteOneID(id uuid.UUID) *BankAccountDeleteOne {
	builder := c.Delete().Where(bankaccount.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BankAccountDeleteOne{builder}
}

// Query returns a query builder for BankAccount.
func (c *BankAccountClient) Query() *BankAccountQuery {
	return &BankAccountQuery{
		config: c.config,
	}
}

// Get returns a BankAccount entity by its id.
func (c *BankAccountClient) Get(ctx context.Context, id uuid.UUID) (*BankAccount, error) {
	return c.Query().Where(bankaccount.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BankAccountClient) GetX(ctx context.Context, id uuid.UUID) *BankAccount {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCustomer queries the customer edge of a BankAccount.
func (c *BankAccountClient) QueryCustomer(ba *BankAccount) *CustomerQuery {
	query := &CustomerQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ba.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bankaccount.Table, bankaccount.FieldID, id),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, bankaccount.CustomerTable, bankaccount.CustomerColumn),
		)
		fromV = sqlgraph.Neighbors(ba.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySentTransaction queries the sent_transaction edge of a BankAccount.
func (c *BankAccountClient) QuerySentTransaction(ba *BankAccount) *TransactionQuery {
	query := &TransactionQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ba.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bankaccount.Table, bankaccount.FieldID, id),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, bankaccount.SentTransactionTable, bankaccount.SentTransactionColumn),
		)
		fromV = sqlgraph.Neighbors(ba.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReceivedTransaction queries the received_transaction edge of a BankAccount.
func (c *BankAccountClient) QueryReceivedTransaction(ba *BankAccount) *TransactionQuery {
	query := &TransactionQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ba.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bankaccount.Table, bankaccount.FieldID, id),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, bankaccount.ReceivedTransactionTable, bankaccount.ReceivedTransactionColumn),
		)
		fromV = sqlgraph.Neighbors(ba.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOwnedDebts queries the owned_debts edge of a BankAccount.
func (c *BankAccountClient) QueryOwnedDebts(ba *BankAccount) *DebtQuery {
	query := &DebtQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ba.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bankaccount.Table, bankaccount.FieldID, id),
			sqlgraph.To(debt.Table, debt.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, bankaccount.OwnedDebtsTable, bankaccount.OwnedDebtsColumn),
		)
		fromV = sqlgraph.Neighbors(ba.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReceivedDebts queries the received_debts edge of a BankAccount.
func (c *BankAccountClient) QueryReceivedDebts(ba *BankAccount) *DebtQuery {
	query := &DebtQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ba.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bankaccount.Table, bankaccount.FieldID, id),
			sqlgraph.To(debt.Table, debt.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, bankaccount.ReceivedDebtsTable, bankaccount.ReceivedDebtsColumn),
		)
		fromV = sqlgraph.Neighbors(ba.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BankAccountClient) Hooks() []Hook {
	return c.hooks.BankAccount
}

// ContactClient is a client for the Contact schema.
type ContactClient struct {
	config
}

// NewContactClient returns a client for the Contact from the given config.
func NewContactClient(c config) *ContactClient {
	return &ContactClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `contact.Hooks(f(g(h())))`.
func (c *ContactClient) Use(hooks ...Hook) {
	c.hooks.Contact = append(c.hooks.Contact, hooks...)
}

// Create returns a builder for creating a Contact entity.
func (c *ContactClient) Create() *ContactCreate {
	mutation := newContactMutation(c.config, OpCreate)
	return &ContactCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Contact entities.
func (c *ContactClient) CreateBulk(builders ...*ContactCreate) *ContactCreateBulk {
	return &ContactCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Contact.
func (c *ContactClient) Update() *ContactUpdate {
	mutation := newContactMutation(c.config, OpUpdate)
	return &ContactUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ContactClient) UpdateOne(co *Contact) *ContactUpdateOne {
	mutation := newContactMutation(c.config, OpUpdateOne, withContact(co))
	return &ContactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ContactClient) UpdateOneID(id uuid.UUID) *ContactUpdateOne {
	mutation := newContactMutation(c.config, OpUpdateOne, withContactID(id))
	return &ContactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Contact.
func (c *ContactClient) Delete() *ContactDelete {
	mutation := newContactMutation(c.config, OpDelete)
	return &ContactDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ContactClient) DeleteOne(co *Contact) *ContactDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ContactClient) DeleteOneID(id uuid.UUID) *ContactDeleteOne {
	builder := c.Delete().Where(contact.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ContactDeleteOne{builder}
}

// Query returns a query builder for Contact.
func (c *ContactClient) Query() *ContactQuery {
	return &ContactQuery{
		config: c.config,
	}
}

// Get returns a Contact entity by its id.
func (c *ContactClient) Get(ctx context.Context, id uuid.UUID) (*Contact, error) {
	return c.Query().Where(contact.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ContactClient) GetX(ctx context.Context, id uuid.UUID) *Contact {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Contact.
func (c *ContactClient) QueryOwner(co *Contact) *CustomerQuery {
	query := &CustomerQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(contact.Table, contact.FieldID, id),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, contact.OwnerTable, contact.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ContactClient) Hooks() []Hook {
	return c.hooks.Contact
}

// CustomerClient is a client for the Customer schema.
type CustomerClient struct {
	config
}

// NewCustomerClient returns a client for the Customer from the given config.
func NewCustomerClient(c config) *CustomerClient {
	return &CustomerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `customer.Hooks(f(g(h())))`.
func (c *CustomerClient) Use(hooks ...Hook) {
	c.hooks.Customer = append(c.hooks.Customer, hooks...)
}

// Create returns a builder for creating a Customer entity.
func (c *CustomerClient) Create() *CustomerCreate {
	mutation := newCustomerMutation(c.config, OpCreate)
	return &CustomerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Customer entities.
func (c *CustomerClient) CreateBulk(builders ...*CustomerCreate) *CustomerCreateBulk {
	return &CustomerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Customer.
func (c *CustomerClient) Update() *CustomerUpdate {
	mutation := newCustomerMutation(c.config, OpUpdate)
	return &CustomerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CustomerClient) UpdateOne(cu *Customer) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomer(cu))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CustomerClient) UpdateOneID(id uuid.UUID) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomerID(id))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Customer.
func (c *CustomerClient) Delete() *CustomerDelete {
	mutation := newCustomerMutation(c.config, OpDelete)
	return &CustomerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CustomerClient) DeleteOne(cu *Customer) *CustomerDeleteOne {
	return c.DeleteOneID(cu.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CustomerClient) DeleteOneID(id uuid.UUID) *CustomerDeleteOne {
	builder := c.Delete().Where(customer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CustomerDeleteOne{builder}
}

// Query returns a query builder for Customer.
func (c *CustomerClient) Query() *CustomerQuery {
	return &CustomerQuery{
		config: c.config,
	}
}

// Get returns a Customer entity by its id.
func (c *CustomerClient) Get(ctx context.Context, id uuid.UUID) (*Customer, error) {
	return c.Query().Where(customer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CustomerClient) GetX(ctx context.Context, id uuid.UUID) *Customer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBankAccounts queries the bank_accounts edge of a Customer.
func (c *CustomerClient) QueryBankAccounts(cu *Customer) *BankAccountQuery {
	query := &BankAccountQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, id),
			sqlgraph.To(bankaccount.Table, bankaccount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, customer.BankAccountsTable, customer.BankAccountsColumn),
		)
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryContacts queries the contacts edge of a Customer.
func (c *CustomerClient) QueryContacts(cu *Customer) *ContactQuery {
	query := &ContactQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, id),
			sqlgraph.To(contact.Table, contact.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, customer.ContactsTable, customer.ContactsColumn),
		)
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CustomerClient) Hooks() []Hook {
	return c.hooks.Customer
}

// DebtClient is a client for the Debt schema.
type DebtClient struct {
	config
}

// NewDebtClient returns a client for the Debt from the given config.
func NewDebtClient(c config) *DebtClient {
	return &DebtClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `debt.Hooks(f(g(h())))`.
func (c *DebtClient) Use(hooks ...Hook) {
	c.hooks.Debt = append(c.hooks.Debt, hooks...)
}

// Create returns a builder for creating a Debt entity.
func (c *DebtClient) Create() *DebtCreate {
	mutation := newDebtMutation(c.config, OpCreate)
	return &DebtCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Debt entities.
func (c *DebtClient) CreateBulk(builders ...*DebtCreate) *DebtCreateBulk {
	return &DebtCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Debt.
func (c *DebtClient) Update() *DebtUpdate {
	mutation := newDebtMutation(c.config, OpUpdate)
	return &DebtUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DebtClient) UpdateOne(d *Debt) *DebtUpdateOne {
	mutation := newDebtMutation(c.config, OpUpdateOne, withDebt(d))
	return &DebtUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DebtClient) UpdateOneID(id uuid.UUID) *DebtUpdateOne {
	mutation := newDebtMutation(c.config, OpUpdateOne, withDebtID(id))
	return &DebtUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Debt.
func (c *DebtClient) Delete() *DebtDelete {
	mutation := newDebtMutation(c.config, OpDelete)
	return &DebtDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DebtClient) DeleteOne(d *Debt) *DebtDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DebtClient) DeleteOneID(id uuid.UUID) *DebtDeleteOne {
	builder := c.Delete().Where(debt.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DebtDeleteOne{builder}
}

// Query returns a query builder for Debt.
func (c *DebtClient) Query() *DebtQuery {
	return &DebtQuery{
		config: c.config,
	}
}

// Get returns a Debt entity by its id.
func (c *DebtClient) Get(ctx context.Context, id uuid.UUID) (*Debt, error) {
	return c.Query().Where(debt.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DebtClient) GetX(ctx context.Context, id uuid.UUID) *Debt {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Debt.
func (c *DebtClient) QueryOwner(d *Debt) *BankAccountQuery {
	query := &BankAccountQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(debt.Table, debt.FieldID, id),
			sqlgraph.To(bankaccount.Table, bankaccount.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, debt.OwnerTable, debt.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReceiver queries the receiver edge of a Debt.
func (c *DebtClient) QueryReceiver(d *Debt) *BankAccountQuery {
	query := &BankAccountQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(debt.Table, debt.FieldID, id),
			sqlgraph.To(bankaccount.Table, bankaccount.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, debt.ReceiverTable, debt.ReceiverColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTransaction queries the transaction edge of a Debt.
func (c *DebtClient) QueryTransaction(d *Debt) *TransactionQuery {
	query := &TransactionQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(debt.Table, debt.FieldID, id),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, debt.TransactionTable, debt.TransactionColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DebtClient) Hooks() []Hook {
	return c.hooks.Debt
}

// EmployeeClient is a client for the Employee schema.
type EmployeeClient struct {
	config
}

// NewEmployeeClient returns a client for the Employee from the given config.
func NewEmployeeClient(c config) *EmployeeClient {
	return &EmployeeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `employee.Hooks(f(g(h())))`.
func (c *EmployeeClient) Use(hooks ...Hook) {
	c.hooks.Employee = append(c.hooks.Employee, hooks...)
}

// Create returns a builder for creating a Employee entity.
func (c *EmployeeClient) Create() *EmployeeCreate {
	mutation := newEmployeeMutation(c.config, OpCreate)
	return &EmployeeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Employee entities.
func (c *EmployeeClient) CreateBulk(builders ...*EmployeeCreate) *EmployeeCreateBulk {
	return &EmployeeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Employee.
func (c *EmployeeClient) Update() *EmployeeUpdate {
	mutation := newEmployeeMutation(c.config, OpUpdate)
	return &EmployeeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EmployeeClient) UpdateOne(e *Employee) *EmployeeUpdateOne {
	mutation := newEmployeeMutation(c.config, OpUpdateOne, withEmployee(e))
	return &EmployeeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EmployeeClient) UpdateOneID(id uuid.UUID) *EmployeeUpdateOne {
	mutation := newEmployeeMutation(c.config, OpUpdateOne, withEmployeeID(id))
	return &EmployeeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Employee.
func (c *EmployeeClient) Delete() *EmployeeDelete {
	mutation := newEmployeeMutation(c.config, OpDelete)
	return &EmployeeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EmployeeClient) DeleteOne(e *Employee) *EmployeeDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EmployeeClient) DeleteOneID(id uuid.UUID) *EmployeeDeleteOne {
	builder := c.Delete().Where(employee.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EmployeeDeleteOne{builder}
}

// Query returns a query builder for Employee.
func (c *EmployeeClient) Query() *EmployeeQuery {
	return &EmployeeQuery{
		config: c.config,
	}
}

// Get returns a Employee entity by its id.
func (c *EmployeeClient) Get(ctx context.Context, id uuid.UUID) (*Employee, error) {
	return c.Query().Where(employee.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EmployeeClient) GetX(ctx context.Context, id uuid.UUID) *Employee {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *EmployeeClient) Hooks() []Hook {
	return c.hooks.Employee
}

// PartnerClient is a client for the Partner schema.
type PartnerClient struct {
	config
}

// NewPartnerClient returns a client for the Partner from the given config.
func NewPartnerClient(c config) *PartnerClient {
	return &PartnerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `partner.Hooks(f(g(h())))`.
func (c *PartnerClient) Use(hooks ...Hook) {
	c.hooks.Partner = append(c.hooks.Partner, hooks...)
}

// Create returns a builder for creating a Partner entity.
func (c *PartnerClient) Create() *PartnerCreate {
	mutation := newPartnerMutation(c.config, OpCreate)
	return &PartnerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Partner entities.
func (c *PartnerClient) CreateBulk(builders ...*PartnerCreate) *PartnerCreateBulk {
	return &PartnerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Partner.
func (c *PartnerClient) Update() *PartnerUpdate {
	mutation := newPartnerMutation(c.config, OpUpdate)
	return &PartnerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PartnerClient) UpdateOne(pa *Partner) *PartnerUpdateOne {
	mutation := newPartnerMutation(c.config, OpUpdateOne, withPartner(pa))
	return &PartnerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PartnerClient) UpdateOneID(id uuid.UUID) *PartnerUpdateOne {
	mutation := newPartnerMutation(c.config, OpUpdateOne, withPartnerID(id))
	return &PartnerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Partner.
func (c *PartnerClient) Delete() *PartnerDelete {
	mutation := newPartnerMutation(c.config, OpDelete)
	return &PartnerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PartnerClient) DeleteOne(pa *Partner) *PartnerDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PartnerClient) DeleteOneID(id uuid.UUID) *PartnerDeleteOne {
	builder := c.Delete().Where(partner.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PartnerDeleteOne{builder}
}

// Query returns a query builder for Partner.
func (c *PartnerClient) Query() *PartnerQuery {
	return &PartnerQuery{
		config: c.config,
	}
}

// Get returns a Partner entity by its id.
func (c *PartnerClient) Get(ctx context.Context, id uuid.UUID) (*Partner, error) {
	return c.Query().Where(partner.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PartnerClient) GetX(ctx context.Context, id uuid.UUID) *Partner {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PartnerClient) Hooks() []Hook {
	return c.hooks.Partner
}

// TransactionClient is a client for the Transaction schema.
type TransactionClient struct {
	config
}

// NewTransactionClient returns a client for the Transaction from the given config.
func NewTransactionClient(c config) *TransactionClient {
	return &TransactionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `transaction.Hooks(f(g(h())))`.
func (c *TransactionClient) Use(hooks ...Hook) {
	c.hooks.Transaction = append(c.hooks.Transaction, hooks...)
}

// Create returns a builder for creating a Transaction entity.
func (c *TransactionClient) Create() *TransactionCreate {
	mutation := newTransactionMutation(c.config, OpCreate)
	return &TransactionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Transaction entities.
func (c *TransactionClient) CreateBulk(builders ...*TransactionCreate) *TransactionCreateBulk {
	return &TransactionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Transaction.
func (c *TransactionClient) Update() *TransactionUpdate {
	mutation := newTransactionMutation(c.config, OpUpdate)
	return &TransactionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TransactionClient) UpdateOne(t *Transaction) *TransactionUpdateOne {
	mutation := newTransactionMutation(c.config, OpUpdateOne, withTransaction(t))
	return &TransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TransactionClient) UpdateOneID(id uuid.UUID) *TransactionUpdateOne {
	mutation := newTransactionMutation(c.config, OpUpdateOne, withTransactionID(id))
	return &TransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Transaction.
func (c *TransactionClient) Delete() *TransactionDelete {
	mutation := newTransactionMutation(c.config, OpDelete)
	return &TransactionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TransactionClient) DeleteOne(t *Transaction) *TransactionDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TransactionClient) DeleteOneID(id uuid.UUID) *TransactionDeleteOne {
	builder := c.Delete().Where(transaction.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TransactionDeleteOne{builder}
}

// Query returns a query builder for Transaction.
func (c *TransactionClient) Query() *TransactionQuery {
	return &TransactionQuery{
		config: c.config,
	}
}

// Get returns a Transaction entity by its id.
func (c *TransactionClient) Get(ctx context.Context, id uuid.UUID) (*Transaction, error) {
	return c.Query().Where(transaction.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TransactionClient) GetX(ctx context.Context, id uuid.UUID) *Transaction {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySourceTransaction queries the source_transaction edge of a Transaction.
func (c *TransactionClient) QuerySourceTransaction(t *Transaction) *TransactionQuery {
	query := &TransactionQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, id),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, transaction.SourceTransactionTable, transaction.SourceTransactionColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFeeTransaction queries the fee_transaction edge of a Transaction.
func (c *TransactionClient) QueryFeeTransaction(t *Transaction) *TransactionQuery {
	query := &TransactionQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, id),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, transaction.FeeTransactionTable, transaction.FeeTransactionColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReceiver queries the receiver edge of a Transaction.
func (c *TransactionClient) QueryReceiver(t *Transaction) *BankAccountQuery {
	query := &BankAccountQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, id),
			sqlgraph.To(bankaccount.Table, bankaccount.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, transaction.ReceiverTable, transaction.ReceiverColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySender queries the sender edge of a Transaction.
func (c *TransactionClient) QuerySender(t *Transaction) *BankAccountQuery {
	query := &BankAccountQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, id),
			sqlgraph.To(bankaccount.Table, bankaccount.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, transaction.SenderTable, transaction.SenderColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDebt queries the debt edge of a Transaction.
func (c *TransactionClient) QueryDebt(t *Transaction) *DebtQuery {
	query := &DebtQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, id),
			sqlgraph.To(debt.Table, debt.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, transaction.DebtTable, transaction.DebtColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TransactionClient) Hooks() []Hook {
	return c.hooks.Transaction
}
