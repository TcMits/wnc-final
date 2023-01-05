package ent

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/TcMits/wnc-final/ent/bankaccount"
	"github.com/TcMits/wnc-final/ent/debt"
	"github.com/TcMits/wnc-final/ent/predicate"
	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Debt is the model entity for the Debt schema.
type Debt struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// OwnerBankAccountNumber holds the value of the "owner_bank_account_number" field.
	OwnerBankAccountNumber string `json:"owner_bank_account_number,omitempty"`
	// OwnerBankName holds the value of the "owner_bank_name" field.
	OwnerBankName string `json:"owner_bank_name,omitempty"`
	// OwnerName holds the value of the "owner_name" field.
	OwnerName string `json:"owner_name,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID *uuid.UUID `json:"owner_id,omitempty"`
	// ReceiverBankAccountNumber holds the value of the "receiver_bank_account_number" field.
	ReceiverBankAccountNumber string `json:"receiver_bank_account_number,omitempty"`
	// ReceiverBankName holds the value of the "receiver_bank_name" field.
	ReceiverBankName string `json:"receiver_bank_name,omitempty"`
	// ReceiverName holds the value of the "receiver_name" field.
	ReceiverName string `json:"receiver_name,omitempty"`
	// ReceiverID holds the value of the "receiver_id" field.
	ReceiverID *uuid.UUID `json:"receiver_id,omitempty"`
	// TransactionID holds the value of the "transaction_id" field.
	TransactionID *uuid.UUID `json:"transaction_id,omitempty"`
	// Status holds the value of the "status" field.
	Status debt.Status `json:"status,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DebtQuery when eager-loading is set.
	Edges DebtEdges `json:"edges"`
}

// DebtEdges holds the relations/edges for other nodes in the graph.
type DebtEdges struct {
	// Owner holds the value of the owner edge.
	Owner *BankAccount `json:"owner,omitempty"`
	// Receiver holds the value of the receiver edge.
	Receiver *BankAccount `json:"receiver,omitempty"`
	// Transaction holds the value of the transaction edge.
	Transaction *Transaction `json:"transaction,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DebtEdges) OwnerOrErr() (*BankAccount, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: bankaccount.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// ReceiverOrErr returns the Receiver value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DebtEdges) ReceiverOrErr() (*BankAccount, error) {
	if e.loadedTypes[1] {
		if e.Receiver == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: bankaccount.Label}
		}
		return e.Receiver, nil
	}
	return nil, &NotLoadedError{edge: "receiver"}
}

// TransactionOrErr returns the Transaction value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DebtEdges) TransactionOrErr() (*Transaction, error) {
	if e.loadedTypes[2] {
		if e.Transaction == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: transaction.Label}
		}
		return e.Transaction, nil
	}
	return nil, &NotLoadedError{edge: "transaction"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Debt) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case debt.FieldOwnerID, debt.FieldReceiverID, debt.FieldTransactionID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case debt.FieldAmount:
			values[i] = new(decimal.Decimal)
		case debt.FieldOwnerBankAccountNumber, debt.FieldOwnerBankName, debt.FieldOwnerName, debt.FieldReceiverBankAccountNumber, debt.FieldReceiverBankName, debt.FieldReceiverName, debt.FieldStatus, debt.FieldDescription:
			values[i] = new(sql.NullString)
		case debt.FieldCreateTime, debt.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case debt.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Debt", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Debt fields.
func (d *Debt) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case debt.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case debt.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				d.CreateTime = value.Time
			}
		case debt.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				d.UpdateTime = value.Time
			}
		case debt.FieldOwnerBankAccountNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_bank_account_number", values[i])
			} else if value.Valid {
				d.OwnerBankAccountNumber = value.String
			}
		case debt.FieldOwnerBankName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_bank_name", values[i])
			} else if value.Valid {
				d.OwnerBankName = value.String
			}
		case debt.FieldOwnerName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_name", values[i])
			} else if value.Valid {
				d.OwnerName = value.String
			}
		case debt.FieldOwnerID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				d.OwnerID = new(uuid.UUID)
				*d.OwnerID = *value.S.(*uuid.UUID)
			}
		case debt.FieldReceiverBankAccountNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receiver_bank_account_number", values[i])
			} else if value.Valid {
				d.ReceiverBankAccountNumber = value.String
			}
		case debt.FieldReceiverBankName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receiver_bank_name", values[i])
			} else if value.Valid {
				d.ReceiverBankName = value.String
			}
		case debt.FieldReceiverName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receiver_name", values[i])
			} else if value.Valid {
				d.ReceiverName = value.String
			}
		case debt.FieldReceiverID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field receiver_id", values[i])
			} else if value.Valid {
				d.ReceiverID = new(uuid.UUID)
				*d.ReceiverID = *value.S.(*uuid.UUID)
			}
		case debt.FieldTransactionID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id", values[i])
			} else if value.Valid {
				d.TransactionID = new(uuid.UUID)
				*d.TransactionID = *value.S.(*uuid.UUID)
			}
		case debt.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				d.Status = debt.Status(value.String)
			}
		case debt.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				d.Description = value.String
			}
		case debt.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				d.Amount = *value
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Debt entity.
func (d *Debt) QueryOwner() *BankAccountQuery {
	return (&DebtClient{config: d.config}).QueryOwner(d)
}

// QueryReceiver queries the "receiver" edge of the Debt entity.
func (d *Debt) QueryReceiver() *BankAccountQuery {
	return (&DebtClient{config: d.config}).QueryReceiver(d)
}

// QueryTransaction queries the "transaction" edge of the Debt entity.
func (d *Debt) QueryTransaction() *TransactionQuery {
	return (&DebtClient{config: d.config}).QueryTransaction(d)
}

// Update returns a builder for updating this Debt.
// Note that you need to call Debt.Unwrap() before calling this method if this Debt
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Debt) Update() *DebtUpdateOne {
	return (&DebtClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Debt entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Debt) Unwrap() *Debt {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Debt is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Debt) String() string {
	var builder strings.Builder
	builder.WriteString("Debt(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("create_time=")
	builder.WriteString(d.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(d.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("owner_bank_account_number=")
	builder.WriteString(d.OwnerBankAccountNumber)
	builder.WriteString(", ")
	builder.WriteString("owner_bank_name=")
	builder.WriteString(d.OwnerBankName)
	builder.WriteString(", ")
	builder.WriteString("owner_name=")
	builder.WriteString(d.OwnerName)
	builder.WriteString(", ")
	if v := d.OwnerID; v != nil {
		builder.WriteString("owner_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("receiver_bank_account_number=")
	builder.WriteString(d.ReceiverBankAccountNumber)
	builder.WriteString(", ")
	builder.WriteString("receiver_bank_name=")
	builder.WriteString(d.ReceiverBankName)
	builder.WriteString(", ")
	builder.WriteString("receiver_name=")
	builder.WriteString(d.ReceiverName)
	builder.WriteString(", ")
	if v := d.ReceiverID; v != nil {
		builder.WriteString("receiver_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := d.TransactionID; v != nil {
		builder.WriteString("transaction_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", d.Status))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(d.Description)
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", d.Amount))
	builder.WriteByte(')')
	return builder.String()
}

type DebtCreateRepository struct {
	client   *Client
	isAtomic bool
}

func NewDebtCreateRepository(
	client *Client,
	isAtomic bool,
) *DebtCreateRepository {
	return &DebtCreateRepository{
		client:   client,
		isAtomic: isAtomic,
	}
}

// using in Tx
func (r *DebtCreateRepository) CreateWithClient(
	ctx context.Context, client *Client, input *DebtCreateInput,
) (*Debt, error) {
	instance, err := client.Debt.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (r *DebtCreateRepository) Create(
	ctx context.Context, input *DebtCreateInput,
) (*Debt, error) {
	if !r.isAtomic {
		return r.CreateWithClient(ctx, r.client, input)
	}
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	instance, err := r.CreateWithClient(ctx, tx.Client(), input)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction: %w", err)
	}
	return instance, nil
}

type DebtDeleteRepository struct {
	client   *Client
	isAtomic bool
}

func NewDebtDeleteRepository(
	client *Client,
	isAtomic bool,
) *DebtDeleteRepository {
	return &DebtDeleteRepository{
		client:   client,
		isAtomic: isAtomic,
	}
}

// using in Tx
func (r *DebtDeleteRepository) DeleteWithClient(
	ctx context.Context, client *Client, instance *Debt,
) error {
	err := client.Debt.DeleteOne(instance).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *DebtDeleteRepository) Delete(
	ctx context.Context, instance *Debt,
) error {
	if !r.isAtomic {
		return r.DeleteWithClient(ctx, r.client, instance)
	}
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return err
	}
	err = r.DeleteWithClient(ctx, tx.Client(), instance)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

// DebtCreateInput represents a mutation input for creating debts.
type DebtCreateInput struct {
	CreateTime                *time.Time      `json:"create_time,omitempty" form:"create_time"`
	UpdateTime                *time.Time      `json:"update_time,omitempty" form:"update_time"`
	OwnerID                   uuid.UUID       `json:"owner_id,omitempty" form:"owner_id"`
	OwnerBankAccountNumber    string          `json:"owner_bank_account_number,omitempty" form:"owner_bank_account_number"`
	OwnerBankName             string          `json:"owner_bank_name,omitempty" form:"owner_bank_name"`
	OwnerName                 string          `json:"owner_name,omitempty" form:"owner_name"`
	ReceiverID                uuid.UUID       `json:"receiver_id,omitempty" form:"receiver_id"`
	ReceiverBankAccountNumber string          `json:"receiver_bank_account_number,omitempty" form:"receiver_bank_account_number"`
	ReceiverBankName          string          `json:"receiver_bank_name,omitempty" form:"receiver_bank_name"`
	ReceiverName              string          `json:"receiver_name,omitempty" form:"receiver_name"`
	Status                    *debt.Status    `json:"status,omitempty" form:"status"`
	Description               *string         `json:"description,omitempty" form:"description"`
	Amount                    decimal.Decimal `json:"amount,omitempty" form:"amount"`
	TransactionID             *uuid.UUID      `json:"transaction_id,omitempty" form:"transaction_id"`
}

// Mutate applies the DebtCreateInput on the DebtCreate builder.
func (i *DebtCreateInput) Mutate(m *DebtMutation) {
	if v := i.CreateTime; v != nil {
		m.SetCreateTime(*v)
	}
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	m.SetOwnerBankAccountNumber(i.OwnerBankAccountNumber)
	m.SetOwnerBankName(i.OwnerBankName)
	m.SetOwnerName(i.OwnerName)
	m.SetReceiverBankAccountNumber(i.ReceiverBankAccountNumber)
	m.SetReceiverBankName(i.ReceiverBankName)
	m.SetReceiverName(i.ReceiverName)
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	m.SetAmount(i.Amount)
	m.SetOwnerID(i.OwnerID)
	m.SetReceiverID(i.ReceiverID)
	if v := i.TransactionID; v != nil {
		m.SetTransactionID(*v)
	}
}

// SetInput applies the change-set in the DebtCreateInput on the create builder.
func (c *DebtCreate) SetInput(i *DebtCreateInput) *DebtCreate {
	i.Mutate(c.Mutation())
	return c
}

// DebtUpdateInput represents a mutation input for updating debts.
type DebtUpdateInput struct {
	ID                        uuid.UUID
	UpdateTime                *time.Time   `json:"update_time,omitempty" form:"update_time"`
	OwnerBankAccountNumber    *string      `json:"owner_bank_account_number,omitempty" form:"owner_bank_account_number"`
	OwnerBankName             *string      `json:"owner_bank_name,omitempty" form:"owner_bank_name"`
	OwnerName                 *string      `json:"owner_name,omitempty" form:"owner_name"`
	ReceiverBankAccountNumber *string      `json:"receiver_bank_account_number,omitempty" form:"receiver_bank_account_number"`
	ReceiverBankName          *string      `json:"receiver_bank_name,omitempty" form:"receiver_bank_name"`
	ReceiverName              *string      `json:"receiver_name,omitempty" form:"receiver_name"`
	Status                    *debt.Status `json:"status,omitempty" form:"status"`
	Description               *string      `json:"description,omitempty" form:"description"`
	ClearDescription          bool
	Amount                    *decimal.Decimal `json:"amount,omitempty" form:"amount"`
	OwnerID                   *uuid.UUID       `json:"owner_id,omitempty" form:"owner_id"`
	ClearOwner                bool
	ReceiverID                *uuid.UUID `json:"receiver_id,omitempty" form:"receiver_id"`
	ClearReceiver             bool
	TransactionID             *uuid.UUID `json:"transaction_id,omitempty" form:"transaction_id"`
	ClearTransaction          bool
}

// Mutate applies the DebtUpdateInput on the DebtMutation.
func (i *DebtUpdateInput) Mutate(m *DebtMutation) {
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	if v := i.OwnerBankAccountNumber; v != nil {
		m.SetOwnerBankAccountNumber(*v)
	}
	if v := i.OwnerBankName; v != nil {
		m.SetOwnerBankName(*v)
	}
	if v := i.OwnerName; v != nil {
		m.SetOwnerName(*v)
	}
	if v := i.ReceiverBankAccountNumber; v != nil {
		m.SetReceiverBankAccountNumber(*v)
	}
	if v := i.ReceiverBankName; v != nil {
		m.SetReceiverBankName(*v)
	}
	if v := i.ReceiverName; v != nil {
		m.SetReceiverName(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearDescription {
		m.ClearDescription()
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if v := i.Amount; v != nil {
		m.SetAmount(*v)
	}
	if i.ClearOwner {
		m.ClearOwner()
	}
	if v := i.OwnerID; v != nil {
		m.SetOwnerID(*v)
	}
	if i.ClearReceiver {
		m.ClearReceiver()
	}
	if v := i.ReceiverID; v != nil {
		m.SetReceiverID(*v)
	}
	if i.ClearTransaction {
		m.ClearTransaction()
	}
	if v := i.TransactionID; v != nil {
		m.SetTransactionID(*v)
	}
}

// SetInput applies the change-set in the DebtUpdateInput on the update builder.
func (u *DebtUpdate) SetInput(i *DebtUpdateInput) *DebtUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the DebtUpdateInput on the update-one builder.
func (u *DebtUpdateOne) SetInput(i *DebtUpdateInput) *DebtUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

type DebtReadRepository struct {
	client *Client
}

func NewDebtReadRepository(
	client *Client,
) *DebtReadRepository {
	return &DebtReadRepository{
		client: client,
	}
}

func (r *DebtReadRepository) prepareQuery(
	client *Client, limit *int, offset *int, o *DebtOrderInput, w *DebtWhereInput,
) (*DebtQuery, error) {
	var err error
	q := r.client.Debt.Query()
	if limit != nil {
		q = q.Limit(*limit)
	}
	if offset != nil {
		q = q.Offset(*offset)
	}
	if o != nil {
		q = o.Order(q)
	}
	if w != nil {
		q, err = w.Filter(q)
		if err != nil {
			return nil, err
		}
	}
	return q, nil
}

// using in Tx
func (r *DebtReadRepository) GetWithClient(
	ctx context.Context, client *Client, w *DebtWhereInput, forUpdate bool,
) (*Debt, error) {
	q, err := r.prepareQuery(client, nil, nil, nil, w)
	if err != nil {
		return nil, err
	}
	if forUpdate {
		q = q.ForUpdate()
	}
	instance, err := q.Only(ctx)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// using in Tx
func (r *DebtReadRepository) ListWithClient(
	ctx context.Context, client *Client, limit *int, offset *int, o *DebtOrderInput, w *DebtWhereInput, forUpdate bool,
) ([]*Debt, error) {
	q, err := r.prepareQuery(client, limit, offset, o, w)
	if err != nil {
		return nil, err
	}
	if forUpdate {
		q = q.ForUpdate()
	}
	instances, err := q.All(ctx)
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func (r *DebtReadRepository) Count(ctx context.Context, w *DebtWhereInput) (int, error) {
	q, err := r.prepareQuery(r.client, nil, nil, nil, w)
	if err != nil {
		return 0, err
	}
	count, err := q.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *DebtReadRepository) Get(ctx context.Context, w *DebtWhereInput) (*Debt, error) {
	return r.GetWithClient(ctx, r.client, w, false)
}

func (r *DebtReadRepository) List(
	ctx context.Context, limit *int, offset *int, o *DebtOrderInput, w *DebtWhereInput,
) ([]*Debt, error) {
	return r.ListWithClient(ctx, r.client, limit, offset, o, w, false)
}

type DebtSerializer struct {
	columns map[string]func(context.Context, *Debt) any
}

func NewDebtSerializer(customColumns map[string]func(context.Context, *Debt) any, columns ...string) *DebtSerializer {
	columnsMap := map[string]func(context.Context, *Debt) any{}
	for _, col := range columns {
		switch col {

		case debt.FieldID:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.ID
			}

		case debt.FieldCreateTime:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.CreateTime
			}

		case debt.FieldUpdateTime:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.UpdateTime
			}

		case debt.FieldOwnerBankAccountNumber:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.OwnerBankAccountNumber
			}

		case debt.FieldOwnerBankName:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.OwnerBankName
			}

		case debt.FieldOwnerName:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.OwnerName
			}

		case debt.FieldOwnerID:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.OwnerID
			}

		case debt.FieldReceiverBankAccountNumber:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.ReceiverBankAccountNumber
			}

		case debt.FieldReceiverBankName:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.ReceiverBankName
			}

		case debt.FieldReceiverName:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.ReceiverName
			}

		case debt.FieldReceiverID:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.ReceiverID
			}

		case debt.FieldTransactionID:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.TransactionID
			}

		case debt.FieldStatus:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.Status
			}

		case debt.FieldDescription:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.Description
			}

		case debt.FieldAmount:
			columnsMap[col] = func(ctx context.Context, d *Debt) any {
				return d.Amount
			}

		default:
			panic(fmt.Sprintf("Unexpect column %s", col))
		}
	}

	for k, serializeFunc := range customColumns {
		columnsMap[k] = serializeFunc
	}

	return &DebtSerializer{
		columns: columnsMap,
	}
}

func (s *DebtSerializer) Serialize(ctx context.Context, d *Debt) map[string]any {
	result := make(map[string]any, len(s.columns))
	for col, serializeFunc := range s.columns {
		result[col] = serializeFunc(ctx, d)
	}
	return result
}

type DebtUpdateRepository struct {
	client   *Client
	isAtomic bool
}

func NewDebtUpdateRepository(
	client *Client,
	isAtomic bool,
) *DebtUpdateRepository {
	return &DebtUpdateRepository{
		client:   client,
		isAtomic: isAtomic,
	}
}

// using in Tx
func (r *DebtUpdateRepository) UpdateWithClient(
	ctx context.Context, client *Client, instance *Debt, input *DebtUpdateInput,
) (*Debt, error) {
	newInstance, err := client.Debt.UpdateOne(instance).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	return newInstance, nil
}

func (r *DebtUpdateRepository) Update(
	ctx context.Context, instance *Debt, input *DebtUpdateInput,
) (*Debt, error) {
	if !r.isAtomic {
		return r.UpdateWithClient(ctx, r.client, instance, input)
	}
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	instance, err = r.UpdateWithClient(ctx, tx.Client(), instance, input)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("committing transaction: %w", err)
	}
	return instance, nil
}

// DebtWhereInput represents a where input for filtering Debt queries.
type DebtWhereInput struct {
	Predicates []predicate.Debt  `json:"-"`
	Not        *DebtWhereInput   `json:"not,omitempty"`
	Or         []*DebtWhereInput `json:"or,omitempty"`
	And        []*DebtWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *uuid.UUID  `json:"id,omitempty" form:"id" param:"id" url:"id"`
	IDNEQ   *uuid.UUID  `json:"id_neq,omitempty" form:"id_neq" param:"id_neq" url:"id_neq"`
	IDIn    []uuid.UUID `json:"id_in,omitempty" form:"id_in" param:"id_in" url:"id_in"`
	IDNotIn []uuid.UUID `json:"id_not_in,omitempty" form:"id_not_in" param:"id_not_in" url:"id_not_in"`
	IDGT    *uuid.UUID  `json:"id_gt,omitempty" form:"id_gt" param:"id_gt" url:"id_gt"`
	IDGTE   *uuid.UUID  `json:"id_gte,omitempty" form:"id_gte" param:"id_gte" url:"id_gte"`
	IDLT    *uuid.UUID  `json:"id_lt,omitempty" form:"id_lt" param:"id_lt" url:"id_lt"`
	IDLTE   *uuid.UUID  `json:"id_lte,omitempty" form:"id_lte" param:"id_lte" url:"id_lte"`

	// "create_time" field predicates.
	CreateTime      *time.Time  `json:"create_time,omitempty" form:"create_time" param:"create_time" url:"create_time"`
	CreateTimeNEQ   *time.Time  `json:"create_time_neq,omitempty" form:"create_time_neq" param:"create_time_neq" url:"create_time_neq"`
	CreateTimeIn    []time.Time `json:"create_time_in,omitempty" form:"create_time_in" param:"create_time_in" url:"create_time_in"`
	CreateTimeNotIn []time.Time `json:"create_time_not_in,omitempty" form:"create_time_not_in" param:"create_time_not_in" url:"create_time_not_in"`
	CreateTimeGT    *time.Time  `json:"create_time_gt,omitempty" form:"create_time_gt" param:"create_time_gt" url:"create_time_gt"`
	CreateTimeGTE   *time.Time  `json:"create_time_gte,omitempty" form:"create_time_gte" param:"create_time_gte" url:"create_time_gte"`
	CreateTimeLT    *time.Time  `json:"create_time_lt,omitempty" form:"create_time_lt" param:"create_time_lt" url:"create_time_lt"`
	CreateTimeLTE   *time.Time  `json:"create_time_lte,omitempty" form:"create_time_lte" param:"create_time_lte" url:"create_time_lte"`

	// "update_time" field predicates.
	UpdateTime      *time.Time  `json:"update_time,omitempty" form:"update_time" param:"update_time" url:"update_time"`
	UpdateTimeNEQ   *time.Time  `json:"update_time_neq,omitempty" form:"update_time_neq" param:"update_time_neq" url:"update_time_neq"`
	UpdateTimeIn    []time.Time `json:"update_time_in,omitempty" form:"update_time_in" param:"update_time_in" url:"update_time_in"`
	UpdateTimeNotIn []time.Time `json:"update_time_not_in,omitempty" form:"update_time_not_in" param:"update_time_not_in" url:"update_time_not_in"`
	UpdateTimeGT    *time.Time  `json:"update_time_gt,omitempty" form:"update_time_gt" param:"update_time_gt" url:"update_time_gt"`
	UpdateTimeGTE   *time.Time  `json:"update_time_gte,omitempty" form:"update_time_gte" param:"update_time_gte" url:"update_time_gte"`
	UpdateTimeLT    *time.Time  `json:"update_time_lt,omitempty" form:"update_time_lt" param:"update_time_lt" url:"update_time_lt"`
	UpdateTimeLTE   *time.Time  `json:"update_time_lte,omitempty" form:"update_time_lte" param:"update_time_lte" url:"update_time_lte"`

	// "owner_bank_account_number" field predicates.
	OwnerBankAccountNumber             *string  `json:"owner_bank_account_number,omitempty" form:"owner_bank_account_number" param:"owner_bank_account_number" url:"owner_bank_account_number"`
	OwnerBankAccountNumberNEQ          *string  `json:"owner_bank_account_number_neq,omitempty" form:"owner_bank_account_number_neq" param:"owner_bank_account_number_neq" url:"owner_bank_account_number_neq"`
	OwnerBankAccountNumberIn           []string `json:"owner_bank_account_number_in,omitempty" form:"owner_bank_account_number_in" param:"owner_bank_account_number_in" url:"owner_bank_account_number_in"`
	OwnerBankAccountNumberNotIn        []string `json:"owner_bank_account_number_not_in,omitempty" form:"owner_bank_account_number_not_in" param:"owner_bank_account_number_not_in" url:"owner_bank_account_number_not_in"`
	OwnerBankAccountNumberGT           *string  `json:"owner_bank_account_number_gt,omitempty" form:"owner_bank_account_number_gt" param:"owner_bank_account_number_gt" url:"owner_bank_account_number_gt"`
	OwnerBankAccountNumberGTE          *string  `json:"owner_bank_account_number_gte,omitempty" form:"owner_bank_account_number_gte" param:"owner_bank_account_number_gte" url:"owner_bank_account_number_gte"`
	OwnerBankAccountNumberLT           *string  `json:"owner_bank_account_number_lt,omitempty" form:"owner_bank_account_number_lt" param:"owner_bank_account_number_lt" url:"owner_bank_account_number_lt"`
	OwnerBankAccountNumberLTE          *string  `json:"owner_bank_account_number_lte,omitempty" form:"owner_bank_account_number_lte" param:"owner_bank_account_number_lte" url:"owner_bank_account_number_lte"`
	OwnerBankAccountNumberContains     *string  `json:"owner_bank_account_number_contains,omitempty" form:"owner_bank_account_number_contains" param:"owner_bank_account_number_contains" url:"owner_bank_account_number_contains"`
	OwnerBankAccountNumberHasPrefix    *string  `json:"owner_bank_account_number_has_prefix,omitempty" form:"owner_bank_account_number_has_prefix" param:"owner_bank_account_number_has_prefix" url:"owner_bank_account_number_has_prefix"`
	OwnerBankAccountNumberHasSuffix    *string  `json:"owner_bank_account_number_has_suffix,omitempty" form:"owner_bank_account_number_has_suffix" param:"owner_bank_account_number_has_suffix" url:"owner_bank_account_number_has_suffix"`
	OwnerBankAccountNumberEqualFold    *string  `json:"owner_bank_account_number_equal_fold,omitempty" form:"owner_bank_account_number_equal_fold" param:"owner_bank_account_number_equal_fold" url:"owner_bank_account_number_equal_fold"`
	OwnerBankAccountNumberContainsFold *string  `json:"owner_bank_account_number_contains_fold,omitempty" form:"owner_bank_account_number_contains_fold" param:"owner_bank_account_number_contains_fold" url:"owner_bank_account_number_contains_fold"`

	// "owner_bank_name" field predicates.
	OwnerBankName             *string  `json:"owner_bank_name,omitempty" form:"owner_bank_name" param:"owner_bank_name" url:"owner_bank_name"`
	OwnerBankNameNEQ          *string  `json:"owner_bank_name_neq,omitempty" form:"owner_bank_name_neq" param:"owner_bank_name_neq" url:"owner_bank_name_neq"`
	OwnerBankNameIn           []string `json:"owner_bank_name_in,omitempty" form:"owner_bank_name_in" param:"owner_bank_name_in" url:"owner_bank_name_in"`
	OwnerBankNameNotIn        []string `json:"owner_bank_name_not_in,omitempty" form:"owner_bank_name_not_in" param:"owner_bank_name_not_in" url:"owner_bank_name_not_in"`
	OwnerBankNameGT           *string  `json:"owner_bank_name_gt,omitempty" form:"owner_bank_name_gt" param:"owner_bank_name_gt" url:"owner_bank_name_gt"`
	OwnerBankNameGTE          *string  `json:"owner_bank_name_gte,omitempty" form:"owner_bank_name_gte" param:"owner_bank_name_gte" url:"owner_bank_name_gte"`
	OwnerBankNameLT           *string  `json:"owner_bank_name_lt,omitempty" form:"owner_bank_name_lt" param:"owner_bank_name_lt" url:"owner_bank_name_lt"`
	OwnerBankNameLTE          *string  `json:"owner_bank_name_lte,omitempty" form:"owner_bank_name_lte" param:"owner_bank_name_lte" url:"owner_bank_name_lte"`
	OwnerBankNameContains     *string  `json:"owner_bank_name_contains,omitempty" form:"owner_bank_name_contains" param:"owner_bank_name_contains" url:"owner_bank_name_contains"`
	OwnerBankNameHasPrefix    *string  `json:"owner_bank_name_has_prefix,omitempty" form:"owner_bank_name_has_prefix" param:"owner_bank_name_has_prefix" url:"owner_bank_name_has_prefix"`
	OwnerBankNameHasSuffix    *string  `json:"owner_bank_name_has_suffix,omitempty" form:"owner_bank_name_has_suffix" param:"owner_bank_name_has_suffix" url:"owner_bank_name_has_suffix"`
	OwnerBankNameEqualFold    *string  `json:"owner_bank_name_equal_fold,omitempty" form:"owner_bank_name_equal_fold" param:"owner_bank_name_equal_fold" url:"owner_bank_name_equal_fold"`
	OwnerBankNameContainsFold *string  `json:"owner_bank_name_contains_fold,omitempty" form:"owner_bank_name_contains_fold" param:"owner_bank_name_contains_fold" url:"owner_bank_name_contains_fold"`

	// "owner_name" field predicates.
	OwnerName             *string  `json:"owner_name,omitempty" form:"owner_name" param:"owner_name" url:"owner_name"`
	OwnerNameNEQ          *string  `json:"owner_name_neq,omitempty" form:"owner_name_neq" param:"owner_name_neq" url:"owner_name_neq"`
	OwnerNameIn           []string `json:"owner_name_in,omitempty" form:"owner_name_in" param:"owner_name_in" url:"owner_name_in"`
	OwnerNameNotIn        []string `json:"owner_name_not_in,omitempty" form:"owner_name_not_in" param:"owner_name_not_in" url:"owner_name_not_in"`
	OwnerNameGT           *string  `json:"owner_name_gt,omitempty" form:"owner_name_gt" param:"owner_name_gt" url:"owner_name_gt"`
	OwnerNameGTE          *string  `json:"owner_name_gte,omitempty" form:"owner_name_gte" param:"owner_name_gte" url:"owner_name_gte"`
	OwnerNameLT           *string  `json:"owner_name_lt,omitempty" form:"owner_name_lt" param:"owner_name_lt" url:"owner_name_lt"`
	OwnerNameLTE          *string  `json:"owner_name_lte,omitempty" form:"owner_name_lte" param:"owner_name_lte" url:"owner_name_lte"`
	OwnerNameContains     *string  `json:"owner_name_contains,omitempty" form:"owner_name_contains" param:"owner_name_contains" url:"owner_name_contains"`
	OwnerNameHasPrefix    *string  `json:"owner_name_has_prefix,omitempty" form:"owner_name_has_prefix" param:"owner_name_has_prefix" url:"owner_name_has_prefix"`
	OwnerNameHasSuffix    *string  `json:"owner_name_has_suffix,omitempty" form:"owner_name_has_suffix" param:"owner_name_has_suffix" url:"owner_name_has_suffix"`
	OwnerNameEqualFold    *string  `json:"owner_name_equal_fold,omitempty" form:"owner_name_equal_fold" param:"owner_name_equal_fold" url:"owner_name_equal_fold"`
	OwnerNameContainsFold *string  `json:"owner_name_contains_fold,omitempty" form:"owner_name_contains_fold" param:"owner_name_contains_fold" url:"owner_name_contains_fold"`

	// "owner_id" field predicates.
	OwnerID      *uuid.UUID  `json:"owner_id,omitempty" form:"owner_id" param:"owner_id" url:"owner_id"`
	OwnerIDNEQ   *uuid.UUID  `json:"owner_id_neq,omitempty" form:"owner_id_neq" param:"owner_id_neq" url:"owner_id_neq"`
	OwnerIDIn    []uuid.UUID `json:"owner_id_in,omitempty" form:"owner_id_in" param:"owner_id_in" url:"owner_id_in"`
	OwnerIDNotIn []uuid.UUID `json:"owner_id_not_in,omitempty" form:"owner_id_not_in" param:"owner_id_not_in" url:"owner_id_not_in"`

	// "receiver_bank_account_number" field predicates.
	ReceiverBankAccountNumber             *string  `json:"receiver_bank_account_number,omitempty" form:"receiver_bank_account_number" param:"receiver_bank_account_number" url:"receiver_bank_account_number"`
	ReceiverBankAccountNumberNEQ          *string  `json:"receiver_bank_account_number_neq,omitempty" form:"receiver_bank_account_number_neq" param:"receiver_bank_account_number_neq" url:"receiver_bank_account_number_neq"`
	ReceiverBankAccountNumberIn           []string `json:"receiver_bank_account_number_in,omitempty" form:"receiver_bank_account_number_in" param:"receiver_bank_account_number_in" url:"receiver_bank_account_number_in"`
	ReceiverBankAccountNumberNotIn        []string `json:"receiver_bank_account_number_not_in,omitempty" form:"receiver_bank_account_number_not_in" param:"receiver_bank_account_number_not_in" url:"receiver_bank_account_number_not_in"`
	ReceiverBankAccountNumberGT           *string  `json:"receiver_bank_account_number_gt,omitempty" form:"receiver_bank_account_number_gt" param:"receiver_bank_account_number_gt" url:"receiver_bank_account_number_gt"`
	ReceiverBankAccountNumberGTE          *string  `json:"receiver_bank_account_number_gte,omitempty" form:"receiver_bank_account_number_gte" param:"receiver_bank_account_number_gte" url:"receiver_bank_account_number_gte"`
	ReceiverBankAccountNumberLT           *string  `json:"receiver_bank_account_number_lt,omitempty" form:"receiver_bank_account_number_lt" param:"receiver_bank_account_number_lt" url:"receiver_bank_account_number_lt"`
	ReceiverBankAccountNumberLTE          *string  `json:"receiver_bank_account_number_lte,omitempty" form:"receiver_bank_account_number_lte" param:"receiver_bank_account_number_lte" url:"receiver_bank_account_number_lte"`
	ReceiverBankAccountNumberContains     *string  `json:"receiver_bank_account_number_contains,omitempty" form:"receiver_bank_account_number_contains" param:"receiver_bank_account_number_contains" url:"receiver_bank_account_number_contains"`
	ReceiverBankAccountNumberHasPrefix    *string  `json:"receiver_bank_account_number_has_prefix,omitempty" form:"receiver_bank_account_number_has_prefix" param:"receiver_bank_account_number_has_prefix" url:"receiver_bank_account_number_has_prefix"`
	ReceiverBankAccountNumberHasSuffix    *string  `json:"receiver_bank_account_number_has_suffix,omitempty" form:"receiver_bank_account_number_has_suffix" param:"receiver_bank_account_number_has_suffix" url:"receiver_bank_account_number_has_suffix"`
	ReceiverBankAccountNumberEqualFold    *string  `json:"receiver_bank_account_number_equal_fold,omitempty" form:"receiver_bank_account_number_equal_fold" param:"receiver_bank_account_number_equal_fold" url:"receiver_bank_account_number_equal_fold"`
	ReceiverBankAccountNumberContainsFold *string  `json:"receiver_bank_account_number_contains_fold,omitempty" form:"receiver_bank_account_number_contains_fold" param:"receiver_bank_account_number_contains_fold" url:"receiver_bank_account_number_contains_fold"`

	// "receiver_bank_name" field predicates.
	ReceiverBankName             *string  `json:"receiver_bank_name,omitempty" form:"receiver_bank_name" param:"receiver_bank_name" url:"receiver_bank_name"`
	ReceiverBankNameNEQ          *string  `json:"receiver_bank_name_neq,omitempty" form:"receiver_bank_name_neq" param:"receiver_bank_name_neq" url:"receiver_bank_name_neq"`
	ReceiverBankNameIn           []string `json:"receiver_bank_name_in,omitempty" form:"receiver_bank_name_in" param:"receiver_bank_name_in" url:"receiver_bank_name_in"`
	ReceiverBankNameNotIn        []string `json:"receiver_bank_name_not_in,omitempty" form:"receiver_bank_name_not_in" param:"receiver_bank_name_not_in" url:"receiver_bank_name_not_in"`
	ReceiverBankNameGT           *string  `json:"receiver_bank_name_gt,omitempty" form:"receiver_bank_name_gt" param:"receiver_bank_name_gt" url:"receiver_bank_name_gt"`
	ReceiverBankNameGTE          *string  `json:"receiver_bank_name_gte,omitempty" form:"receiver_bank_name_gte" param:"receiver_bank_name_gte" url:"receiver_bank_name_gte"`
	ReceiverBankNameLT           *string  `json:"receiver_bank_name_lt,omitempty" form:"receiver_bank_name_lt" param:"receiver_bank_name_lt" url:"receiver_bank_name_lt"`
	ReceiverBankNameLTE          *string  `json:"receiver_bank_name_lte,omitempty" form:"receiver_bank_name_lte" param:"receiver_bank_name_lte" url:"receiver_bank_name_lte"`
	ReceiverBankNameContains     *string  `json:"receiver_bank_name_contains,omitempty" form:"receiver_bank_name_contains" param:"receiver_bank_name_contains" url:"receiver_bank_name_contains"`
	ReceiverBankNameHasPrefix    *string  `json:"receiver_bank_name_has_prefix,omitempty" form:"receiver_bank_name_has_prefix" param:"receiver_bank_name_has_prefix" url:"receiver_bank_name_has_prefix"`
	ReceiverBankNameHasSuffix    *string  `json:"receiver_bank_name_has_suffix,omitempty" form:"receiver_bank_name_has_suffix" param:"receiver_bank_name_has_suffix" url:"receiver_bank_name_has_suffix"`
	ReceiverBankNameEqualFold    *string  `json:"receiver_bank_name_equal_fold,omitempty" form:"receiver_bank_name_equal_fold" param:"receiver_bank_name_equal_fold" url:"receiver_bank_name_equal_fold"`
	ReceiverBankNameContainsFold *string  `json:"receiver_bank_name_contains_fold,omitempty" form:"receiver_bank_name_contains_fold" param:"receiver_bank_name_contains_fold" url:"receiver_bank_name_contains_fold"`

	// "receiver_name" field predicates.
	ReceiverName             *string  `json:"receiver_name,omitempty" form:"receiver_name" param:"receiver_name" url:"receiver_name"`
	ReceiverNameNEQ          *string  `json:"receiver_name_neq,omitempty" form:"receiver_name_neq" param:"receiver_name_neq" url:"receiver_name_neq"`
	ReceiverNameIn           []string `json:"receiver_name_in,omitempty" form:"receiver_name_in" param:"receiver_name_in" url:"receiver_name_in"`
	ReceiverNameNotIn        []string `json:"receiver_name_not_in,omitempty" form:"receiver_name_not_in" param:"receiver_name_not_in" url:"receiver_name_not_in"`
	ReceiverNameGT           *string  `json:"receiver_name_gt,omitempty" form:"receiver_name_gt" param:"receiver_name_gt" url:"receiver_name_gt"`
	ReceiverNameGTE          *string  `json:"receiver_name_gte,omitempty" form:"receiver_name_gte" param:"receiver_name_gte" url:"receiver_name_gte"`
	ReceiverNameLT           *string  `json:"receiver_name_lt,omitempty" form:"receiver_name_lt" param:"receiver_name_lt" url:"receiver_name_lt"`
	ReceiverNameLTE          *string  `json:"receiver_name_lte,omitempty" form:"receiver_name_lte" param:"receiver_name_lte" url:"receiver_name_lte"`
	ReceiverNameContains     *string  `json:"receiver_name_contains,omitempty" form:"receiver_name_contains" param:"receiver_name_contains" url:"receiver_name_contains"`
	ReceiverNameHasPrefix    *string  `json:"receiver_name_has_prefix,omitempty" form:"receiver_name_has_prefix" param:"receiver_name_has_prefix" url:"receiver_name_has_prefix"`
	ReceiverNameHasSuffix    *string  `json:"receiver_name_has_suffix,omitempty" form:"receiver_name_has_suffix" param:"receiver_name_has_suffix" url:"receiver_name_has_suffix"`
	ReceiverNameEqualFold    *string  `json:"receiver_name_equal_fold,omitempty" form:"receiver_name_equal_fold" param:"receiver_name_equal_fold" url:"receiver_name_equal_fold"`
	ReceiverNameContainsFold *string  `json:"receiver_name_contains_fold,omitempty" form:"receiver_name_contains_fold" param:"receiver_name_contains_fold" url:"receiver_name_contains_fold"`

	// "receiver_id" field predicates.
	ReceiverID      *uuid.UUID  `json:"receiver_id,omitempty" form:"receiver_id" param:"receiver_id" url:"receiver_id"`
	ReceiverIDNEQ   *uuid.UUID  `json:"receiver_id_neq,omitempty" form:"receiver_id_neq" param:"receiver_id_neq" url:"receiver_id_neq"`
	ReceiverIDIn    []uuid.UUID `json:"receiver_id_in,omitempty" form:"receiver_id_in" param:"receiver_id_in" url:"receiver_id_in"`
	ReceiverIDNotIn []uuid.UUID `json:"receiver_id_not_in,omitempty" form:"receiver_id_not_in" param:"receiver_id_not_in" url:"receiver_id_not_in"`

	// "transaction_id" field predicates.
	TransactionID       *uuid.UUID  `json:"transaction_id,omitempty" form:"transaction_id" param:"transaction_id" url:"transaction_id"`
	TransactionIDNEQ    *uuid.UUID  `json:"transaction_id_neq,omitempty" form:"transaction_id_neq" param:"transaction_id_neq" url:"transaction_id_neq"`
	TransactionIDIn     []uuid.UUID `json:"transaction_id_in,omitempty" form:"transaction_id_in" param:"transaction_id_in" url:"transaction_id_in"`
	TransactionIDNotIn  []uuid.UUID `json:"transaction_id_not_in,omitempty" form:"transaction_id_not_in" param:"transaction_id_not_in" url:"transaction_id_not_in"`
	TransactionIDIsNil  bool        `json:"transaction_id_is_nil,omitempty" form:"transaction_id_is_nil" param:"transaction_id_is_nil" url:"transaction_id_is_nil"`
	TransactionIDNotNil bool        `json:"transaction_id_not_nil,omitempty" form:"transaction_id_not_nil" param:"transaction_id_not_nil" url:"transaction_id_not_nil"`

	// "status" field predicates.
	Status      *debt.Status  `json:"status,omitempty" form:"status" param:"status" url:"status"`
	StatusNEQ   *debt.Status  `json:"status_neq,omitempty" form:"status_neq" param:"status_neq" url:"status_neq"`
	StatusIn    []debt.Status `json:"status_in,omitempty" form:"status_in" param:"status_in" url:"status_in"`
	StatusNotIn []debt.Status `json:"status_not_in,omitempty" form:"status_not_in" param:"status_not_in" url:"status_not_in"`

	// "description" field predicates.
	Description             *string  `json:"description,omitempty" form:"description" param:"description" url:"description"`
	DescriptionNEQ          *string  `json:"description_neq,omitempty" form:"description_neq" param:"description_neq" url:"description_neq"`
	DescriptionIn           []string `json:"description_in,omitempty" form:"description_in" param:"description_in" url:"description_in"`
	DescriptionNotIn        []string `json:"description_not_in,omitempty" form:"description_not_in" param:"description_not_in" url:"description_not_in"`
	DescriptionGT           *string  `json:"description_gt,omitempty" form:"description_gt" param:"description_gt" url:"description_gt"`
	DescriptionGTE          *string  `json:"description_gte,omitempty" form:"description_gte" param:"description_gte" url:"description_gte"`
	DescriptionLT           *string  `json:"description_lt,omitempty" form:"description_lt" param:"description_lt" url:"description_lt"`
	DescriptionLTE          *string  `json:"description_lte,omitempty" form:"description_lte" param:"description_lte" url:"description_lte"`
	DescriptionContains     *string  `json:"description_contains,omitempty" form:"description_contains" param:"description_contains" url:"description_contains"`
	DescriptionHasPrefix    *string  `json:"description_has_prefix,omitempty" form:"description_has_prefix" param:"description_has_prefix" url:"description_has_prefix"`
	DescriptionHasSuffix    *string  `json:"description_has_suffix,omitempty" form:"description_has_suffix" param:"description_has_suffix" url:"description_has_suffix"`
	DescriptionIsNil        bool     `json:"description_is_nil,omitempty" form:"description_is_nil" param:"description_is_nil" url:"description_is_nil"`
	DescriptionNotNil       bool     `json:"description_not_nil,omitempty" form:"description_not_nil" param:"description_not_nil" url:"description_not_nil"`
	DescriptionEqualFold    *string  `json:"description_equal_fold,omitempty" form:"description_equal_fold" param:"description_equal_fold" url:"description_equal_fold"`
	DescriptionContainsFold *string  `json:"description_contains_fold,omitempty" form:"description_contains_fold" param:"description_contains_fold" url:"description_contains_fold"`

	// "amount" field predicates.
	Amount      *decimal.Decimal  `json:"amount,omitempty" form:"amount" param:"amount" url:"amount"`
	AmountNEQ   *decimal.Decimal  `json:"amount_neq,omitempty" form:"amount_neq" param:"amount_neq" url:"amount_neq"`
	AmountIn    []decimal.Decimal `json:"amount_in,omitempty" form:"amount_in" param:"amount_in" url:"amount_in"`
	AmountNotIn []decimal.Decimal `json:"amount_not_in,omitempty" form:"amount_not_in" param:"amount_not_in" url:"amount_not_in"`
	AmountGT    *decimal.Decimal  `json:"amount_gt,omitempty" form:"amount_gt" param:"amount_gt" url:"amount_gt"`
	AmountGTE   *decimal.Decimal  `json:"amount_gte,omitempty" form:"amount_gte" param:"amount_gte" url:"amount_gte"`
	AmountLT    *decimal.Decimal  `json:"amount_lt,omitempty" form:"amount_lt" param:"amount_lt" url:"amount_lt"`
	AmountLTE   *decimal.Decimal  `json:"amount_lte,omitempty" form:"amount_lte" param:"amount_lte" url:"amount_lte"`

	// "owner" edge predicates.
	HasOwner     *bool                    `json:"has_owner,omitempty" form:"has_owner" param:"has_owner" url:"has_owner"`
	HasOwnerWith []*BankAccountWhereInput `json:"has_owner_with,omitempty" form:"has_owner_with" param:"has_owner_with" url:"has_owner_with"`

	// "receiver" edge predicates.
	HasReceiver     *bool                    `json:"has_receiver,omitempty" form:"has_receiver" param:"has_receiver" url:"has_receiver"`
	HasReceiverWith []*BankAccountWhereInput `json:"has_receiver_with,omitempty" form:"has_receiver_with" param:"has_receiver_with" url:"has_receiver_with"`

	// "transaction" edge predicates.
	HasTransaction     *bool                    `json:"has_transaction,omitempty" form:"has_transaction" param:"has_transaction" url:"has_transaction"`
	HasTransactionWith []*TransactionWhereInput `json:"has_transaction_with,omitempty" form:"has_transaction_with" param:"has_transaction_with" url:"has_transaction_with"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *DebtWhereInput) AddPredicates(predicates ...predicate.Debt) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the DebtWhereInput filter on the DebtQuery builder.
func (i *DebtWhereInput) Filter(q *DebtQuery) (*DebtQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyDebtWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyDebtWhereInput is returned in case the DebtWhereInput is empty.
var ErrEmptyDebtWhereInput = errors.New("ent: empty predicate DebtWhereInput")

// P returns a predicate for filtering debts.
// An error is returned if the input is empty or invalid.
func (i *DebtWhereInput) P() (predicate.Debt, error) {
	var predicates []predicate.Debt
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, debt.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Debt, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, debt.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Debt, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, debt.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, debt.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, debt.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, debt.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, debt.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, debt.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, debt.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, debt.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, debt.IDLTE(*i.IDLTE))
	}
	if i.CreateTime != nil {
		predicates = append(predicates, debt.CreateTimeEQ(*i.CreateTime))
	}
	if i.CreateTimeNEQ != nil {
		predicates = append(predicates, debt.CreateTimeNEQ(*i.CreateTimeNEQ))
	}
	if len(i.CreateTimeIn) > 0 {
		predicates = append(predicates, debt.CreateTimeIn(i.CreateTimeIn...))
	}
	if len(i.CreateTimeNotIn) > 0 {
		predicates = append(predicates, debt.CreateTimeNotIn(i.CreateTimeNotIn...))
	}
	if i.CreateTimeGT != nil {
		predicates = append(predicates, debt.CreateTimeGT(*i.CreateTimeGT))
	}
	if i.CreateTimeGTE != nil {
		predicates = append(predicates, debt.CreateTimeGTE(*i.CreateTimeGTE))
	}
	if i.CreateTimeLT != nil {
		predicates = append(predicates, debt.CreateTimeLT(*i.CreateTimeLT))
	}
	if i.CreateTimeLTE != nil {
		predicates = append(predicates, debt.CreateTimeLTE(*i.CreateTimeLTE))
	}
	if i.UpdateTime != nil {
		predicates = append(predicates, debt.UpdateTimeEQ(*i.UpdateTime))
	}
	if i.UpdateTimeNEQ != nil {
		predicates = append(predicates, debt.UpdateTimeNEQ(*i.UpdateTimeNEQ))
	}
	if len(i.UpdateTimeIn) > 0 {
		predicates = append(predicates, debt.UpdateTimeIn(i.UpdateTimeIn...))
	}
	if len(i.UpdateTimeNotIn) > 0 {
		predicates = append(predicates, debt.UpdateTimeNotIn(i.UpdateTimeNotIn...))
	}
	if i.UpdateTimeGT != nil {
		predicates = append(predicates, debt.UpdateTimeGT(*i.UpdateTimeGT))
	}
	if i.UpdateTimeGTE != nil {
		predicates = append(predicates, debt.UpdateTimeGTE(*i.UpdateTimeGTE))
	}
	if i.UpdateTimeLT != nil {
		predicates = append(predicates, debt.UpdateTimeLT(*i.UpdateTimeLT))
	}
	if i.UpdateTimeLTE != nil {
		predicates = append(predicates, debt.UpdateTimeLTE(*i.UpdateTimeLTE))
	}
	if i.OwnerBankAccountNumber != nil {
		predicates = append(predicates, debt.OwnerBankAccountNumberEQ(*i.OwnerBankAccountNumber))
	}
	if i.OwnerBankAccountNumberNEQ != nil {
		predicates = append(predicates, debt.OwnerBankAccountNumberNEQ(*i.OwnerBankAccountNumberNEQ))
	}
	if len(i.OwnerBankAccountNumberIn) > 0 {
		predicates = append(predicates, debt.OwnerBankAccountNumberIn(i.OwnerBankAccountNumberIn...))
	}
	if len(i.OwnerBankAccountNumberNotIn) > 0 {
		predicates = append(predicates, debt.OwnerBankAccountNumberNotIn(i.OwnerBankAccountNumberNotIn...))
	}
	if i.OwnerBankAccountNumberGT != nil {
		predicates = append(predicates, debt.OwnerBankAccountNumberGT(*i.OwnerBankAccountNumberGT))
	}
	if i.OwnerBankAccountNumberGTE != nil {
		predicates = append(predicates, debt.OwnerBankAccountNumberGTE(*i.OwnerBankAccountNumberGTE))
	}
	if i.OwnerBankAccountNumberLT != nil {
		predicates = append(predicates, debt.OwnerBankAccountNumberLT(*i.OwnerBankAccountNumberLT))
	}
	if i.OwnerBankAccountNumberLTE != nil {
		predicates = append(predicates, debt.OwnerBankAccountNumberLTE(*i.OwnerBankAccountNumberLTE))
	}
	if i.OwnerBankAccountNumberContains != nil {
		predicates = append(predicates, debt.OwnerBankAccountNumberContains(*i.OwnerBankAccountNumberContains))
	}
	if i.OwnerBankAccountNumberHasPrefix != nil {
		predicates = append(predicates, debt.OwnerBankAccountNumberHasPrefix(*i.OwnerBankAccountNumberHasPrefix))
	}
	if i.OwnerBankAccountNumberHasSuffix != nil {
		predicates = append(predicates, debt.OwnerBankAccountNumberHasSuffix(*i.OwnerBankAccountNumberHasSuffix))
	}
	if i.OwnerBankAccountNumberEqualFold != nil {
		predicates = append(predicates, debt.OwnerBankAccountNumberEqualFold(*i.OwnerBankAccountNumberEqualFold))
	}
	if i.OwnerBankAccountNumberContainsFold != nil {
		predicates = append(predicates, debt.OwnerBankAccountNumberContainsFold(*i.OwnerBankAccountNumberContainsFold))
	}
	if i.OwnerBankName != nil {
		predicates = append(predicates, debt.OwnerBankNameEQ(*i.OwnerBankName))
	}
	if i.OwnerBankNameNEQ != nil {
		predicates = append(predicates, debt.OwnerBankNameNEQ(*i.OwnerBankNameNEQ))
	}
	if len(i.OwnerBankNameIn) > 0 {
		predicates = append(predicates, debt.OwnerBankNameIn(i.OwnerBankNameIn...))
	}
	if len(i.OwnerBankNameNotIn) > 0 {
		predicates = append(predicates, debt.OwnerBankNameNotIn(i.OwnerBankNameNotIn...))
	}
	if i.OwnerBankNameGT != nil {
		predicates = append(predicates, debt.OwnerBankNameGT(*i.OwnerBankNameGT))
	}
	if i.OwnerBankNameGTE != nil {
		predicates = append(predicates, debt.OwnerBankNameGTE(*i.OwnerBankNameGTE))
	}
	if i.OwnerBankNameLT != nil {
		predicates = append(predicates, debt.OwnerBankNameLT(*i.OwnerBankNameLT))
	}
	if i.OwnerBankNameLTE != nil {
		predicates = append(predicates, debt.OwnerBankNameLTE(*i.OwnerBankNameLTE))
	}
	if i.OwnerBankNameContains != nil {
		predicates = append(predicates, debt.OwnerBankNameContains(*i.OwnerBankNameContains))
	}
	if i.OwnerBankNameHasPrefix != nil {
		predicates = append(predicates, debt.OwnerBankNameHasPrefix(*i.OwnerBankNameHasPrefix))
	}
	if i.OwnerBankNameHasSuffix != nil {
		predicates = append(predicates, debt.OwnerBankNameHasSuffix(*i.OwnerBankNameHasSuffix))
	}
	if i.OwnerBankNameEqualFold != nil {
		predicates = append(predicates, debt.OwnerBankNameEqualFold(*i.OwnerBankNameEqualFold))
	}
	if i.OwnerBankNameContainsFold != nil {
		predicates = append(predicates, debt.OwnerBankNameContainsFold(*i.OwnerBankNameContainsFold))
	}
	if i.OwnerName != nil {
		predicates = append(predicates, debt.OwnerNameEQ(*i.OwnerName))
	}
	if i.OwnerNameNEQ != nil {
		predicates = append(predicates, debt.OwnerNameNEQ(*i.OwnerNameNEQ))
	}
	if len(i.OwnerNameIn) > 0 {
		predicates = append(predicates, debt.OwnerNameIn(i.OwnerNameIn...))
	}
	if len(i.OwnerNameNotIn) > 0 {
		predicates = append(predicates, debt.OwnerNameNotIn(i.OwnerNameNotIn...))
	}
	if i.OwnerNameGT != nil {
		predicates = append(predicates, debt.OwnerNameGT(*i.OwnerNameGT))
	}
	if i.OwnerNameGTE != nil {
		predicates = append(predicates, debt.OwnerNameGTE(*i.OwnerNameGTE))
	}
	if i.OwnerNameLT != nil {
		predicates = append(predicates, debt.OwnerNameLT(*i.OwnerNameLT))
	}
	if i.OwnerNameLTE != nil {
		predicates = append(predicates, debt.OwnerNameLTE(*i.OwnerNameLTE))
	}
	if i.OwnerNameContains != nil {
		predicates = append(predicates, debt.OwnerNameContains(*i.OwnerNameContains))
	}
	if i.OwnerNameHasPrefix != nil {
		predicates = append(predicates, debt.OwnerNameHasPrefix(*i.OwnerNameHasPrefix))
	}
	if i.OwnerNameHasSuffix != nil {
		predicates = append(predicates, debt.OwnerNameHasSuffix(*i.OwnerNameHasSuffix))
	}
	if i.OwnerNameEqualFold != nil {
		predicates = append(predicates, debt.OwnerNameEqualFold(*i.OwnerNameEqualFold))
	}
	if i.OwnerNameContainsFold != nil {
		predicates = append(predicates, debt.OwnerNameContainsFold(*i.OwnerNameContainsFold))
	}
	if i.OwnerID != nil {
		predicates = append(predicates, debt.OwnerIDEQ(*i.OwnerID))
	}
	if i.OwnerIDNEQ != nil {
		predicates = append(predicates, debt.OwnerIDNEQ(*i.OwnerIDNEQ))
	}
	if len(i.OwnerIDIn) > 0 {
		predicates = append(predicates, debt.OwnerIDIn(i.OwnerIDIn...))
	}
	if len(i.OwnerIDNotIn) > 0 {
		predicates = append(predicates, debt.OwnerIDNotIn(i.OwnerIDNotIn...))
	}
	if i.ReceiverBankAccountNumber != nil {
		predicates = append(predicates, debt.ReceiverBankAccountNumberEQ(*i.ReceiverBankAccountNumber))
	}
	if i.ReceiverBankAccountNumberNEQ != nil {
		predicates = append(predicates, debt.ReceiverBankAccountNumberNEQ(*i.ReceiverBankAccountNumberNEQ))
	}
	if len(i.ReceiverBankAccountNumberIn) > 0 {
		predicates = append(predicates, debt.ReceiverBankAccountNumberIn(i.ReceiverBankAccountNumberIn...))
	}
	if len(i.ReceiverBankAccountNumberNotIn) > 0 {
		predicates = append(predicates, debt.ReceiverBankAccountNumberNotIn(i.ReceiverBankAccountNumberNotIn...))
	}
	if i.ReceiverBankAccountNumberGT != nil {
		predicates = append(predicates, debt.ReceiverBankAccountNumberGT(*i.ReceiverBankAccountNumberGT))
	}
	if i.ReceiverBankAccountNumberGTE != nil {
		predicates = append(predicates, debt.ReceiverBankAccountNumberGTE(*i.ReceiverBankAccountNumberGTE))
	}
	if i.ReceiverBankAccountNumberLT != nil {
		predicates = append(predicates, debt.ReceiverBankAccountNumberLT(*i.ReceiverBankAccountNumberLT))
	}
	if i.ReceiverBankAccountNumberLTE != nil {
		predicates = append(predicates, debt.ReceiverBankAccountNumberLTE(*i.ReceiverBankAccountNumberLTE))
	}
	if i.ReceiverBankAccountNumberContains != nil {
		predicates = append(predicates, debt.ReceiverBankAccountNumberContains(*i.ReceiverBankAccountNumberContains))
	}
	if i.ReceiverBankAccountNumberHasPrefix != nil {
		predicates = append(predicates, debt.ReceiverBankAccountNumberHasPrefix(*i.ReceiverBankAccountNumberHasPrefix))
	}
	if i.ReceiverBankAccountNumberHasSuffix != nil {
		predicates = append(predicates, debt.ReceiverBankAccountNumberHasSuffix(*i.ReceiverBankAccountNumberHasSuffix))
	}
	if i.ReceiverBankAccountNumberEqualFold != nil {
		predicates = append(predicates, debt.ReceiverBankAccountNumberEqualFold(*i.ReceiverBankAccountNumberEqualFold))
	}
	if i.ReceiverBankAccountNumberContainsFold != nil {
		predicates = append(predicates, debt.ReceiverBankAccountNumberContainsFold(*i.ReceiverBankAccountNumberContainsFold))
	}
	if i.ReceiverBankName != nil {
		predicates = append(predicates, debt.ReceiverBankNameEQ(*i.ReceiverBankName))
	}
	if i.ReceiverBankNameNEQ != nil {
		predicates = append(predicates, debt.ReceiverBankNameNEQ(*i.ReceiverBankNameNEQ))
	}
	if len(i.ReceiverBankNameIn) > 0 {
		predicates = append(predicates, debt.ReceiverBankNameIn(i.ReceiverBankNameIn...))
	}
	if len(i.ReceiverBankNameNotIn) > 0 {
		predicates = append(predicates, debt.ReceiverBankNameNotIn(i.ReceiverBankNameNotIn...))
	}
	if i.ReceiverBankNameGT != nil {
		predicates = append(predicates, debt.ReceiverBankNameGT(*i.ReceiverBankNameGT))
	}
	if i.ReceiverBankNameGTE != nil {
		predicates = append(predicates, debt.ReceiverBankNameGTE(*i.ReceiverBankNameGTE))
	}
	if i.ReceiverBankNameLT != nil {
		predicates = append(predicates, debt.ReceiverBankNameLT(*i.ReceiverBankNameLT))
	}
	if i.ReceiverBankNameLTE != nil {
		predicates = append(predicates, debt.ReceiverBankNameLTE(*i.ReceiverBankNameLTE))
	}
	if i.ReceiverBankNameContains != nil {
		predicates = append(predicates, debt.ReceiverBankNameContains(*i.ReceiverBankNameContains))
	}
	if i.ReceiverBankNameHasPrefix != nil {
		predicates = append(predicates, debt.ReceiverBankNameHasPrefix(*i.ReceiverBankNameHasPrefix))
	}
	if i.ReceiverBankNameHasSuffix != nil {
		predicates = append(predicates, debt.ReceiverBankNameHasSuffix(*i.ReceiverBankNameHasSuffix))
	}
	if i.ReceiverBankNameEqualFold != nil {
		predicates = append(predicates, debt.ReceiverBankNameEqualFold(*i.ReceiverBankNameEqualFold))
	}
	if i.ReceiverBankNameContainsFold != nil {
		predicates = append(predicates, debt.ReceiverBankNameContainsFold(*i.ReceiverBankNameContainsFold))
	}
	if i.ReceiverName != nil {
		predicates = append(predicates, debt.ReceiverNameEQ(*i.ReceiverName))
	}
	if i.ReceiverNameNEQ != nil {
		predicates = append(predicates, debt.ReceiverNameNEQ(*i.ReceiverNameNEQ))
	}
	if len(i.ReceiverNameIn) > 0 {
		predicates = append(predicates, debt.ReceiverNameIn(i.ReceiverNameIn...))
	}
	if len(i.ReceiverNameNotIn) > 0 {
		predicates = append(predicates, debt.ReceiverNameNotIn(i.ReceiverNameNotIn...))
	}
	if i.ReceiverNameGT != nil {
		predicates = append(predicates, debt.ReceiverNameGT(*i.ReceiverNameGT))
	}
	if i.ReceiverNameGTE != nil {
		predicates = append(predicates, debt.ReceiverNameGTE(*i.ReceiverNameGTE))
	}
	if i.ReceiverNameLT != nil {
		predicates = append(predicates, debt.ReceiverNameLT(*i.ReceiverNameLT))
	}
	if i.ReceiverNameLTE != nil {
		predicates = append(predicates, debt.ReceiverNameLTE(*i.ReceiverNameLTE))
	}
	if i.ReceiverNameContains != nil {
		predicates = append(predicates, debt.ReceiverNameContains(*i.ReceiverNameContains))
	}
	if i.ReceiverNameHasPrefix != nil {
		predicates = append(predicates, debt.ReceiverNameHasPrefix(*i.ReceiverNameHasPrefix))
	}
	if i.ReceiverNameHasSuffix != nil {
		predicates = append(predicates, debt.ReceiverNameHasSuffix(*i.ReceiverNameHasSuffix))
	}
	if i.ReceiverNameEqualFold != nil {
		predicates = append(predicates, debt.ReceiverNameEqualFold(*i.ReceiverNameEqualFold))
	}
	if i.ReceiverNameContainsFold != nil {
		predicates = append(predicates, debt.ReceiverNameContainsFold(*i.ReceiverNameContainsFold))
	}
	if i.ReceiverID != nil {
		predicates = append(predicates, debt.ReceiverIDEQ(*i.ReceiverID))
	}
	if i.ReceiverIDNEQ != nil {
		predicates = append(predicates, debt.ReceiverIDNEQ(*i.ReceiverIDNEQ))
	}
	if len(i.ReceiverIDIn) > 0 {
		predicates = append(predicates, debt.ReceiverIDIn(i.ReceiverIDIn...))
	}
	if len(i.ReceiverIDNotIn) > 0 {
		predicates = append(predicates, debt.ReceiverIDNotIn(i.ReceiverIDNotIn...))
	}
	if i.TransactionID != nil {
		predicates = append(predicates, debt.TransactionIDEQ(*i.TransactionID))
	}
	if i.TransactionIDNEQ != nil {
		predicates = append(predicates, debt.TransactionIDNEQ(*i.TransactionIDNEQ))
	}
	if len(i.TransactionIDIn) > 0 {
		predicates = append(predicates, debt.TransactionIDIn(i.TransactionIDIn...))
	}
	if len(i.TransactionIDNotIn) > 0 {
		predicates = append(predicates, debt.TransactionIDNotIn(i.TransactionIDNotIn...))
	}
	if i.TransactionIDIsNil {
		predicates = append(predicates, debt.TransactionIDIsNil())
	}
	if i.TransactionIDNotNil {
		predicates = append(predicates, debt.TransactionIDNotNil())
	}
	if i.Status != nil {
		predicates = append(predicates, debt.StatusEQ(*i.Status))
	}
	if i.StatusNEQ != nil {
		predicates = append(predicates, debt.StatusNEQ(*i.StatusNEQ))
	}
	if len(i.StatusIn) > 0 {
		predicates = append(predicates, debt.StatusIn(i.StatusIn...))
	}
	if len(i.StatusNotIn) > 0 {
		predicates = append(predicates, debt.StatusNotIn(i.StatusNotIn...))
	}
	if i.Description != nil {
		predicates = append(predicates, debt.DescriptionEQ(*i.Description))
	}
	if i.DescriptionNEQ != nil {
		predicates = append(predicates, debt.DescriptionNEQ(*i.DescriptionNEQ))
	}
	if len(i.DescriptionIn) > 0 {
		predicates = append(predicates, debt.DescriptionIn(i.DescriptionIn...))
	}
	if len(i.DescriptionNotIn) > 0 {
		predicates = append(predicates, debt.DescriptionNotIn(i.DescriptionNotIn...))
	}
	if i.DescriptionGT != nil {
		predicates = append(predicates, debt.DescriptionGT(*i.DescriptionGT))
	}
	if i.DescriptionGTE != nil {
		predicates = append(predicates, debt.DescriptionGTE(*i.DescriptionGTE))
	}
	if i.DescriptionLT != nil {
		predicates = append(predicates, debt.DescriptionLT(*i.DescriptionLT))
	}
	if i.DescriptionLTE != nil {
		predicates = append(predicates, debt.DescriptionLTE(*i.DescriptionLTE))
	}
	if i.DescriptionContains != nil {
		predicates = append(predicates, debt.DescriptionContains(*i.DescriptionContains))
	}
	if i.DescriptionHasPrefix != nil {
		predicates = append(predicates, debt.DescriptionHasPrefix(*i.DescriptionHasPrefix))
	}
	if i.DescriptionHasSuffix != nil {
		predicates = append(predicates, debt.DescriptionHasSuffix(*i.DescriptionHasSuffix))
	}
	if i.DescriptionIsNil {
		predicates = append(predicates, debt.DescriptionIsNil())
	}
	if i.DescriptionNotNil {
		predicates = append(predicates, debt.DescriptionNotNil())
	}
	if i.DescriptionEqualFold != nil {
		predicates = append(predicates, debt.DescriptionEqualFold(*i.DescriptionEqualFold))
	}
	if i.DescriptionContainsFold != nil {
		predicates = append(predicates, debt.DescriptionContainsFold(*i.DescriptionContainsFold))
	}
	if i.Amount != nil {
		predicates = append(predicates, debt.AmountEQ(*i.Amount))
	}
	if i.AmountNEQ != nil {
		predicates = append(predicates, debt.AmountNEQ(*i.AmountNEQ))
	}
	if len(i.AmountIn) > 0 {
		predicates = append(predicates, debt.AmountIn(i.AmountIn...))
	}
	if len(i.AmountNotIn) > 0 {
		predicates = append(predicates, debt.AmountNotIn(i.AmountNotIn...))
	}
	if i.AmountGT != nil {
		predicates = append(predicates, debt.AmountGT(*i.AmountGT))
	}
	if i.AmountGTE != nil {
		predicates = append(predicates, debt.AmountGTE(*i.AmountGTE))
	}
	if i.AmountLT != nil {
		predicates = append(predicates, debt.AmountLT(*i.AmountLT))
	}
	if i.AmountLTE != nil {
		predicates = append(predicates, debt.AmountLTE(*i.AmountLTE))
	}

	if i.HasOwner != nil {
		p := debt.HasOwner()
		if !*i.HasOwner {
			p = debt.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasOwnerWith) > 0 {
		with := make([]predicate.BankAccount, 0, len(i.HasOwnerWith))
		for _, w := range i.HasOwnerWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasOwnerWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, debt.HasOwnerWith(with...))
	}
	if i.HasReceiver != nil {
		p := debt.HasReceiver()
		if !*i.HasReceiver {
			p = debt.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasReceiverWith) > 0 {
		with := make([]predicate.BankAccount, 0, len(i.HasReceiverWith))
		for _, w := range i.HasReceiverWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasReceiverWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, debt.HasReceiverWith(with...))
	}
	if i.HasTransaction != nil {
		p := debt.HasTransaction()
		if !*i.HasTransaction {
			p = debt.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasTransactionWith) > 0 {
		with := make([]predicate.Transaction, 0, len(i.HasTransactionWith))
		for _, w := range i.HasTransactionWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasTransactionWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, debt.HasTransactionWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyDebtWhereInput
	case 1:
		return predicates[0], nil
	default:
		return debt.And(predicates...), nil
	}
}

// Debts is a parsable slice of Debt.
type Debts []*Debt

func (d Debts) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
