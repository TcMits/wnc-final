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

// Transaction is the model entity for the Transaction schema.
type Transaction struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// SourceTransactionID holds the value of the "source_transaction_id" field.
	SourceTransactionID *uuid.UUID `json:"source_transaction_id,omitempty"`
	// Status holds the value of the "status" field.
	Status transaction.Status `json:"status,omitempty"`
	// ReceiverBankAccountNumber holds the value of the "receiver_bank_account_number" field.
	ReceiverBankAccountNumber string `json:"receiver_bank_account_number,omitempty"`
	// ReceiverBankName holds the value of the "receiver_bank_name" field.
	ReceiverBankName string `json:"receiver_bank_name,omitempty"`
	// ReceiverName holds the value of the "receiver_name" field.
	ReceiverName string `json:"receiver_name,omitempty"`
	// ReceiverID holds the value of the "receiver_id" field.
	ReceiverID *uuid.UUID `json:"receiver_id,omitempty"`
	// SenderBankAccountNumber holds the value of the "sender_bank_account_number" field.
	SenderBankAccountNumber string `json:"sender_bank_account_number,omitempty"`
	// SenderBankName holds the value of the "sender_bank_name" field.
	SenderBankName string `json:"sender_bank_name,omitempty"`
	// SenderName holds the value of the "sender_name" field.
	SenderName string `json:"sender_name,omitempty"`
	// SenderID holds the value of the "sender_id" field.
	SenderID *uuid.UUID `json:"sender_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
	// TransactionType holds the value of the "transaction_type" field.
	TransactionType transaction.TransactionType `json:"transaction_type,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransactionQuery when eager-loading is set.
	Edges TransactionEdges `json:"edges"`
}

// TransactionEdges holds the relations/edges for other nodes in the graph.
type TransactionEdges struct {
	// SourceTransaction holds the value of the source_transaction edge.
	SourceTransaction *Transaction `json:"source_transaction,omitempty"`
	// FeeTransaction holds the value of the fee_transaction edge.
	FeeTransaction *Transaction `json:"fee_transaction,omitempty"`
	// Receiver holds the value of the receiver edge.
	Receiver *BankAccount `json:"receiver,omitempty"`
	// Sender holds the value of the sender edge.
	Sender *BankAccount `json:"sender,omitempty"`
	// Debt holds the value of the debt edge.
	Debt *Debt `json:"debt,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// SourceTransactionOrErr returns the SourceTransaction value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) SourceTransactionOrErr() (*Transaction, error) {
	if e.loadedTypes[0] {
		if e.SourceTransaction == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: transaction.Label}
		}
		return e.SourceTransaction, nil
	}
	return nil, &NotLoadedError{edge: "source_transaction"}
}

// FeeTransactionOrErr returns the FeeTransaction value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) FeeTransactionOrErr() (*Transaction, error) {
	if e.loadedTypes[1] {
		if e.FeeTransaction == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: transaction.Label}
		}
		return e.FeeTransaction, nil
	}
	return nil, &NotLoadedError{edge: "fee_transaction"}
}

// ReceiverOrErr returns the Receiver value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) ReceiverOrErr() (*BankAccount, error) {
	if e.loadedTypes[2] {
		if e.Receiver == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: bankaccount.Label}
		}
		return e.Receiver, nil
	}
	return nil, &NotLoadedError{edge: "receiver"}
}

// SenderOrErr returns the Sender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) SenderOrErr() (*BankAccount, error) {
	if e.loadedTypes[3] {
		if e.Sender == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: bankaccount.Label}
		}
		return e.Sender, nil
	}
	return nil, &NotLoadedError{edge: "sender"}
}

// DebtOrErr returns the Debt value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) DebtOrErr() (*Debt, error) {
	if e.loadedTypes[4] {
		if e.Debt == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: debt.Label}
		}
		return e.Debt, nil
	}
	return nil, &NotLoadedError{edge: "debt"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transaction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transaction.FieldSourceTransactionID, transaction.FieldReceiverID, transaction.FieldSenderID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case transaction.FieldAmount:
			values[i] = new(decimal.Decimal)
		case transaction.FieldStatus, transaction.FieldReceiverBankAccountNumber, transaction.FieldReceiverBankName, transaction.FieldReceiverName, transaction.FieldSenderBankAccountNumber, transaction.FieldSenderBankName, transaction.FieldSenderName, transaction.FieldTransactionType, transaction.FieldDescription:
			values[i] = new(sql.NullString)
		case transaction.FieldCreateTime, transaction.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case transaction.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Transaction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transaction fields.
func (t *Transaction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case transaction.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				t.CreateTime = value.Time
			}
		case transaction.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				t.UpdateTime = value.Time
			}
		case transaction.FieldSourceTransactionID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field source_transaction_id", values[i])
			} else if value.Valid {
				t.SourceTransactionID = new(uuid.UUID)
				*t.SourceTransactionID = *value.S.(*uuid.UUID)
			}
		case transaction.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = transaction.Status(value.String)
			}
		case transaction.FieldReceiverBankAccountNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receiver_bank_account_number", values[i])
			} else if value.Valid {
				t.ReceiverBankAccountNumber = value.String
			}
		case transaction.FieldReceiverBankName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receiver_bank_name", values[i])
			} else if value.Valid {
				t.ReceiverBankName = value.String
			}
		case transaction.FieldReceiverName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receiver_name", values[i])
			} else if value.Valid {
				t.ReceiverName = value.String
			}
		case transaction.FieldReceiverID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field receiver_id", values[i])
			} else if value.Valid {
				t.ReceiverID = new(uuid.UUID)
				*t.ReceiverID = *value.S.(*uuid.UUID)
			}
		case transaction.FieldSenderBankAccountNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender_bank_account_number", values[i])
			} else if value.Valid {
				t.SenderBankAccountNumber = value.String
			}
		case transaction.FieldSenderBankName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender_bank_name", values[i])
			} else if value.Valid {
				t.SenderBankName = value.String
			}
		case transaction.FieldSenderName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender_name", values[i])
			} else if value.Valid {
				t.SenderName = value.String
			}
		case transaction.FieldSenderID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field sender_id", values[i])
			} else if value.Valid {
				t.SenderID = new(uuid.UUID)
				*t.SenderID = *value.S.(*uuid.UUID)
			}
		case transaction.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				t.Amount = *value
			}
		case transaction.FieldTransactionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_type", values[i])
			} else if value.Valid {
				t.TransactionType = transaction.TransactionType(value.String)
			}
		case transaction.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		}
	}
	return nil
}

// QuerySourceTransaction queries the "source_transaction" edge of the Transaction entity.
func (t *Transaction) QuerySourceTransaction() *TransactionQuery {
	return (&TransactionClient{config: t.config}).QuerySourceTransaction(t)
}

// QueryFeeTransaction queries the "fee_transaction" edge of the Transaction entity.
func (t *Transaction) QueryFeeTransaction() *TransactionQuery {
	return (&TransactionClient{config: t.config}).QueryFeeTransaction(t)
}

// QueryReceiver queries the "receiver" edge of the Transaction entity.
func (t *Transaction) QueryReceiver() *BankAccountQuery {
	return (&TransactionClient{config: t.config}).QueryReceiver(t)
}

// QuerySender queries the "sender" edge of the Transaction entity.
func (t *Transaction) QuerySender() *BankAccountQuery {
	return (&TransactionClient{config: t.config}).QuerySender(t)
}

// QueryDebt queries the "debt" edge of the Transaction entity.
func (t *Transaction) QueryDebt() *DebtQuery {
	return (&TransactionClient{config: t.config}).QueryDebt(t)
}

// Update returns a builder for updating this Transaction.
// Note that you need to call Transaction.Unwrap() before calling this method if this Transaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transaction) Update() *TransactionUpdateOne {
	return (&TransactionClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Transaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transaction) Unwrap() *Transaction {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transaction is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transaction) String() string {
	var builder strings.Builder
	builder.WriteString("Transaction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("create_time=")
	builder.WriteString(t.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(t.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := t.SourceTransactionID; v != nil {
		builder.WriteString("source_transaction_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", ")
	builder.WriteString("receiver_bank_account_number=")
	builder.WriteString(t.ReceiverBankAccountNumber)
	builder.WriteString(", ")
	builder.WriteString("receiver_bank_name=")
	builder.WriteString(t.ReceiverBankName)
	builder.WriteString(", ")
	builder.WriteString("receiver_name=")
	builder.WriteString(t.ReceiverName)
	builder.WriteString(", ")
	if v := t.ReceiverID; v != nil {
		builder.WriteString("receiver_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("sender_bank_account_number=")
	builder.WriteString(t.SenderBankAccountNumber)
	builder.WriteString(", ")
	builder.WriteString("sender_bank_name=")
	builder.WriteString(t.SenderBankName)
	builder.WriteString(", ")
	builder.WriteString("sender_name=")
	builder.WriteString(t.SenderName)
	builder.WriteString(", ")
	if v := t.SenderID; v != nil {
		builder.WriteString("sender_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", t.Amount))
	builder.WriteString(", ")
	builder.WriteString("transaction_type=")
	builder.WriteString(fmt.Sprintf("%v", t.TransactionType))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(t.Description)
	builder.WriteByte(')')
	return builder.String()
}

type TransactionCreateRepository struct {
	client   *Client
	isAtomic bool
}

func NewTransactionCreateRepository(
	client *Client,
	isAtomic bool,
) *TransactionCreateRepository {
	return &TransactionCreateRepository{
		client:   client,
		isAtomic: isAtomic,
	}
}

// using in Tx
func (r *TransactionCreateRepository) CreateWithClient(
	ctx context.Context, client *Client, input *TransactionCreateInput,
) (*Transaction, error) {
	instance, err := client.Transaction.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (r *TransactionCreateRepository) Create(
	ctx context.Context, input *TransactionCreateInput,
) (*Transaction, error) {
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

type TransactionDeleteRepository struct {
	client   *Client
	isAtomic bool
}

func NewTransactionDeleteRepository(
	client *Client,
	isAtomic bool,
) *TransactionDeleteRepository {
	return &TransactionDeleteRepository{
		client:   client,
		isAtomic: isAtomic,
	}
}

// using in Tx
func (r *TransactionDeleteRepository) DeleteWithClient(
	ctx context.Context, client *Client, instance *Transaction,
) error {
	err := client.Transaction.DeleteOne(instance).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionDeleteRepository) Delete(
	ctx context.Context, instance *Transaction,
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

// TransactionCreateInput represents a mutation input for creating transactions.
type TransactionCreateInput struct {
	CreateTime                *time.Time                  `json:"create_time,omitempty" form:"create_time"`
	UpdateTime                *time.Time                  `json:"update_time,omitempty" form:"update_time"`
	Status                    *transaction.Status         `json:"status,omitempty" form:"status"`
	ReceiverID                *uuid.UUID                  `json:"receiver_id,omitempty" form:"receiver_id"`
	ReceiverBankAccountNumber string                      `json:"receiver_bank_account_number,omitempty" form:"receiver_bank_account_number"`
	ReceiverBankName          string                      `json:"receiver_bank_name,omitempty" form:"receiver_bank_name"`
	ReceiverName              string                      `json:"receiver_name,omitempty" form:"receiver_name"`
	SenderID                  *uuid.UUID                  `json:"sender_id,omitempty" form:"sender_id"`
	SenderBankAccountNumber   string                      `json:"sender_bank_account_number,omitempty" form:"sender_bank_account_number"`
	SenderBankName            string                      `json:"sender_bank_name,omitempty" form:"sender_bank_name"`
	SenderName                string                      `json:"sender_name,omitempty" form:"sender_name"`
	Amount                    decimal.Decimal             `json:"amount,omitempty" form:"amount"`
	TransactionType           transaction.TransactionType `json:"transaction_type,omitempty" form:"transaction_type"`
	Description               *string                     `json:"description,omitempty" form:"description"`
	SourceTransactionID       *uuid.UUID                  `json:"source_transaction_id,omitempty" form:"source_transaction_id"`
	FeeTransactionID          *uuid.UUID                  `json:"fee_transaction_id,omitempty" form:"fee_transaction_id"`
	DebtID                    *uuid.UUID                  `json:"debt_id,omitempty" form:"debt_id"`
}

// Mutate applies the TransactionCreateInput on the TransactionCreate builder.
func (i *TransactionCreateInput) Mutate(m *TransactionMutation) {
	if v := i.CreateTime; v != nil {
		m.SetCreateTime(*v)
	}
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	m.SetReceiverBankAccountNumber(i.ReceiverBankAccountNumber)
	m.SetReceiverBankName(i.ReceiverBankName)
	m.SetReceiverName(i.ReceiverName)
	m.SetSenderBankAccountNumber(i.SenderBankAccountNumber)
	m.SetSenderBankName(i.SenderBankName)
	m.SetSenderName(i.SenderName)
	m.SetAmount(i.Amount)
	m.SetTransactionType(i.TransactionType)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if v := i.SourceTransactionID; v != nil {
		m.SetSourceTransactionID(*v)
	}
	if v := i.FeeTransactionID; v != nil {
		m.SetFeeTransactionID(*v)
	}
	if v := i.ReceiverID; v != nil {
		m.SetReceiverID(*v)
	}
	if v := i.SenderID; v != nil {
		m.SetSenderID(*v)
	}
	if v := i.DebtID; v != nil {
		m.SetDebtID(*v)
	}
}

// SetInput applies the change-set in the TransactionCreateInput on the create builder.
func (c *TransactionCreate) SetInput(i *TransactionCreateInput) *TransactionCreate {
	i.Mutate(c.Mutation())
	return c
}

// TransactionUpdateInput represents a mutation input for updating transactions.
type TransactionUpdateInput struct {
	ID                        uuid.UUID
	UpdateTime                *time.Time                   `json:"update_time,omitempty" form:"update_time"`
	Status                    *transaction.Status          `json:"status,omitempty" form:"status"`
	ReceiverBankAccountNumber *string                      `json:"receiver_bank_account_number,omitempty" form:"receiver_bank_account_number"`
	ReceiverBankName          *string                      `json:"receiver_bank_name,omitempty" form:"receiver_bank_name"`
	ReceiverName              *string                      `json:"receiver_name,omitempty" form:"receiver_name"`
	SenderBankAccountNumber   *string                      `json:"sender_bank_account_number,omitempty" form:"sender_bank_account_number"`
	SenderBankName            *string                      `json:"sender_bank_name,omitempty" form:"sender_bank_name"`
	SenderName                *string                      `json:"sender_name,omitempty" form:"sender_name"`
	Amount                    *decimal.Decimal             `json:"amount,omitempty" form:"amount"`
	TransactionType           *transaction.TransactionType `json:"transaction_type,omitempty" form:"transaction_type"`
	Description               *string                      `json:"description,omitempty" form:"description"`
	ClearDescription          bool
	SourceTransactionID       *uuid.UUID `json:"source_transaction_id,omitempty" form:"source_transaction_id"`
	ClearSourceTransaction    bool
	FeeTransactionID          *uuid.UUID `json:"fee_transaction_id,omitempty" form:"fee_transaction_id"`
	ClearFeeTransaction       bool
	ReceiverID                *uuid.UUID `json:"receiver_id,omitempty" form:"receiver_id"`
	ClearReceiver             bool
	SenderID                  *uuid.UUID `json:"sender_id,omitempty" form:"sender_id"`
	ClearSender               bool
	DebtID                    *uuid.UUID `json:"debt_id,omitempty" form:"debt_id"`
	ClearDebt                 bool
}

// Mutate applies the TransactionUpdateInput on the TransactionMutation.
func (i *TransactionUpdateInput) Mutate(m *TransactionMutation) {
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
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
	if v := i.SenderBankAccountNumber; v != nil {
		m.SetSenderBankAccountNumber(*v)
	}
	if v := i.SenderBankName; v != nil {
		m.SetSenderBankName(*v)
	}
	if v := i.SenderName; v != nil {
		m.SetSenderName(*v)
	}
	if v := i.Amount; v != nil {
		m.SetAmount(*v)
	}
	if v := i.TransactionType; v != nil {
		m.SetTransactionType(*v)
	}
	if i.ClearDescription {
		m.ClearDescription()
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if i.ClearSourceTransaction {
		m.ClearSourceTransaction()
	}
	if v := i.SourceTransactionID; v != nil {
		m.SetSourceTransactionID(*v)
	}
	if i.ClearFeeTransaction {
		m.ClearFeeTransaction()
	}
	if v := i.FeeTransactionID; v != nil {
		m.SetFeeTransactionID(*v)
	}
	if i.ClearReceiver {
		m.ClearReceiver()
	}
	if v := i.ReceiverID; v != nil {
		m.SetReceiverID(*v)
	}
	if i.ClearSender {
		m.ClearSender()
	}
	if v := i.SenderID; v != nil {
		m.SetSenderID(*v)
	}
	if i.ClearDebt {
		m.ClearDebt()
	}
	if v := i.DebtID; v != nil {
		m.SetDebtID(*v)
	}
}

// SetInput applies the change-set in the TransactionUpdateInput on the update builder.
func (u *TransactionUpdate) SetInput(i *TransactionUpdateInput) *TransactionUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the TransactionUpdateInput on the update-one builder.
func (u *TransactionUpdateOne) SetInput(i *TransactionUpdateInput) *TransactionUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

type TransactionReadRepository struct {
	client *Client
}

func NewTransactionReadRepository(
	client *Client,
) *TransactionReadRepository {
	return &TransactionReadRepository{
		client: client,
	}
}

func (r *TransactionReadRepository) prepareQuery(
	client *Client, limit *int, offset *int, o *TransactionOrderInput, w *TransactionWhereInput,
) (*TransactionQuery, error) {
	var err error
	q := r.client.Transaction.Query()
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
func (r *TransactionReadRepository) GetWithClient(
	ctx context.Context, client *Client, w *TransactionWhereInput, forUpdate bool,
) (*Transaction, error) {
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
func (r *TransactionReadRepository) ListWithClient(
	ctx context.Context, client *Client, limit *int, offset *int, o *TransactionOrderInput, w *TransactionWhereInput, forUpdate bool,
) ([]*Transaction, error) {
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

func (r *TransactionReadRepository) Count(ctx context.Context, w *TransactionWhereInput) (int, error) {
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

func (r *TransactionReadRepository) Get(ctx context.Context, w *TransactionWhereInput) (*Transaction, error) {
	return r.GetWithClient(ctx, r.client, w, false)
}

func (r *TransactionReadRepository) List(
	ctx context.Context, limit *int, offset *int, o *TransactionOrderInput, w *TransactionWhereInput,
) ([]*Transaction, error) {
	return r.ListWithClient(ctx, r.client, limit, offset, o, w, false)
}

type TransactionSerializer struct {
	columns map[string]func(context.Context, *Transaction) any
}

func NewTransactionSerializer(customColumns map[string]func(context.Context, *Transaction) any, columns ...string) *TransactionSerializer {
	columnsMap := map[string]func(context.Context, *Transaction) any{}
	for _, col := range columns {
		switch col {

		case transaction.FieldID:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.ID
			}

		case transaction.FieldCreateTime:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.CreateTime
			}

		case transaction.FieldUpdateTime:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.UpdateTime
			}

		case transaction.FieldSourceTransactionID:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.SourceTransactionID
			}

		case transaction.FieldStatus:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.Status
			}

		case transaction.FieldReceiverBankAccountNumber:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.ReceiverBankAccountNumber
			}

		case transaction.FieldReceiverBankName:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.ReceiverBankName
			}

		case transaction.FieldReceiverName:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.ReceiverName
			}

		case transaction.FieldReceiverID:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.ReceiverID
			}

		case transaction.FieldSenderBankAccountNumber:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.SenderBankAccountNumber
			}

		case transaction.FieldSenderBankName:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.SenderBankName
			}

		case transaction.FieldSenderName:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.SenderName
			}

		case transaction.FieldSenderID:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.SenderID
			}

		case transaction.FieldAmount:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.Amount
			}

		case transaction.FieldTransactionType:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.TransactionType
			}

		case transaction.FieldDescription:
			columnsMap[col] = func(ctx context.Context, t *Transaction) any {
				return t.Description
			}

		default:
			panic(fmt.Sprintf("Unexpect column %s", col))
		}
	}

	for k, serializeFunc := range customColumns {
		columnsMap[k] = serializeFunc
	}

	return &TransactionSerializer{
		columns: columnsMap,
	}
}

func (s *TransactionSerializer) Serialize(ctx context.Context, t *Transaction) map[string]any {
	result := make(map[string]any, len(s.columns))
	for col, serializeFunc := range s.columns {
		result[col] = serializeFunc(ctx, t)
	}
	return result
}

type TransactionUpdateRepository struct {
	client   *Client
	isAtomic bool
}

func NewTransactionUpdateRepository(
	client *Client,
	isAtomic bool,
) *TransactionUpdateRepository {
	return &TransactionUpdateRepository{
		client:   client,
		isAtomic: isAtomic,
	}
}

// using in Tx
func (r *TransactionUpdateRepository) UpdateWithClient(
	ctx context.Context, client *Client, instance *Transaction, input *TransactionUpdateInput,
) (*Transaction, error) {
	newInstance, err := client.Transaction.UpdateOne(instance).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	return newInstance, nil
}

func (r *TransactionUpdateRepository) Update(
	ctx context.Context, instance *Transaction, input *TransactionUpdateInput,
) (*Transaction, error) {
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

// TransactionWhereInput represents a where input for filtering Transaction queries.
type TransactionWhereInput struct {
	Predicates []predicate.Transaction  `json:"-"`
	Not        *TransactionWhereInput   `json:"not,omitempty"`
	Or         []*TransactionWhereInput `json:"or,omitempty"`
	And        []*TransactionWhereInput `json:"and,omitempty"`

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

	// "source_transaction_id" field predicates.
	SourceTransactionID       *uuid.UUID  `json:"source_transaction_id,omitempty" form:"source_transaction_id" param:"source_transaction_id" url:"source_transaction_id"`
	SourceTransactionIDNEQ    *uuid.UUID  `json:"source_transaction_id_neq,omitempty" form:"source_transaction_id_neq" param:"source_transaction_id_neq" url:"source_transaction_id_neq"`
	SourceTransactionIDIn     []uuid.UUID `json:"source_transaction_id_in,omitempty" form:"source_transaction_id_in" param:"source_transaction_id_in" url:"source_transaction_id_in"`
	SourceTransactionIDNotIn  []uuid.UUID `json:"source_transaction_id_not_in,omitempty" form:"source_transaction_id_not_in" param:"source_transaction_id_not_in" url:"source_transaction_id_not_in"`
	SourceTransactionIDIsNil  bool        `json:"source_transaction_id_is_nil,omitempty" form:"source_transaction_id_is_nil" param:"source_transaction_id_is_nil" url:"source_transaction_id_is_nil"`
	SourceTransactionIDNotNil bool        `json:"source_transaction_id_not_nil,omitempty" form:"source_transaction_id_not_nil" param:"source_transaction_id_not_nil" url:"source_transaction_id_not_nil"`

	// "status" field predicates.
	Status      *transaction.Status  `json:"status,omitempty" form:"status" param:"status" url:"status"`
	StatusNEQ   *transaction.Status  `json:"status_neq,omitempty" form:"status_neq" param:"status_neq" url:"status_neq"`
	StatusIn    []transaction.Status `json:"status_in,omitempty" form:"status_in" param:"status_in" url:"status_in"`
	StatusNotIn []transaction.Status `json:"status_not_in,omitempty" form:"status_not_in" param:"status_not_in" url:"status_not_in"`

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
	ReceiverID       *uuid.UUID  `json:"receiver_id,omitempty" form:"receiver_id" param:"receiver_id" url:"receiver_id"`
	ReceiverIDNEQ    *uuid.UUID  `json:"receiver_id_neq,omitempty" form:"receiver_id_neq" param:"receiver_id_neq" url:"receiver_id_neq"`
	ReceiverIDIn     []uuid.UUID `json:"receiver_id_in,omitempty" form:"receiver_id_in" param:"receiver_id_in" url:"receiver_id_in"`
	ReceiverIDNotIn  []uuid.UUID `json:"receiver_id_not_in,omitempty" form:"receiver_id_not_in" param:"receiver_id_not_in" url:"receiver_id_not_in"`
	ReceiverIDIsNil  bool        `json:"receiver_id_is_nil,omitempty" form:"receiver_id_is_nil" param:"receiver_id_is_nil" url:"receiver_id_is_nil"`
	ReceiverIDNotNil bool        `json:"receiver_id_not_nil,omitempty" form:"receiver_id_not_nil" param:"receiver_id_not_nil" url:"receiver_id_not_nil"`

	// "sender_bank_account_number" field predicates.
	SenderBankAccountNumber             *string  `json:"sender_bank_account_number,omitempty" form:"sender_bank_account_number" param:"sender_bank_account_number" url:"sender_bank_account_number"`
	SenderBankAccountNumberNEQ          *string  `json:"sender_bank_account_number_neq,omitempty" form:"sender_bank_account_number_neq" param:"sender_bank_account_number_neq" url:"sender_bank_account_number_neq"`
	SenderBankAccountNumberIn           []string `json:"sender_bank_account_number_in,omitempty" form:"sender_bank_account_number_in" param:"sender_bank_account_number_in" url:"sender_bank_account_number_in"`
	SenderBankAccountNumberNotIn        []string `json:"sender_bank_account_number_not_in,omitempty" form:"sender_bank_account_number_not_in" param:"sender_bank_account_number_not_in" url:"sender_bank_account_number_not_in"`
	SenderBankAccountNumberGT           *string  `json:"sender_bank_account_number_gt,omitempty" form:"sender_bank_account_number_gt" param:"sender_bank_account_number_gt" url:"sender_bank_account_number_gt"`
	SenderBankAccountNumberGTE          *string  `json:"sender_bank_account_number_gte,omitempty" form:"sender_bank_account_number_gte" param:"sender_bank_account_number_gte" url:"sender_bank_account_number_gte"`
	SenderBankAccountNumberLT           *string  `json:"sender_bank_account_number_lt,omitempty" form:"sender_bank_account_number_lt" param:"sender_bank_account_number_lt" url:"sender_bank_account_number_lt"`
	SenderBankAccountNumberLTE          *string  `json:"sender_bank_account_number_lte,omitempty" form:"sender_bank_account_number_lte" param:"sender_bank_account_number_lte" url:"sender_bank_account_number_lte"`
	SenderBankAccountNumberContains     *string  `json:"sender_bank_account_number_contains,omitempty" form:"sender_bank_account_number_contains" param:"sender_bank_account_number_contains" url:"sender_bank_account_number_contains"`
	SenderBankAccountNumberHasPrefix    *string  `json:"sender_bank_account_number_has_prefix,omitempty" form:"sender_bank_account_number_has_prefix" param:"sender_bank_account_number_has_prefix" url:"sender_bank_account_number_has_prefix"`
	SenderBankAccountNumberHasSuffix    *string  `json:"sender_bank_account_number_has_suffix,omitempty" form:"sender_bank_account_number_has_suffix" param:"sender_bank_account_number_has_suffix" url:"sender_bank_account_number_has_suffix"`
	SenderBankAccountNumberEqualFold    *string  `json:"sender_bank_account_number_equal_fold,omitempty" form:"sender_bank_account_number_equal_fold" param:"sender_bank_account_number_equal_fold" url:"sender_bank_account_number_equal_fold"`
	SenderBankAccountNumberContainsFold *string  `json:"sender_bank_account_number_contains_fold,omitempty" form:"sender_bank_account_number_contains_fold" param:"sender_bank_account_number_contains_fold" url:"sender_bank_account_number_contains_fold"`

	// "sender_bank_name" field predicates.
	SenderBankName             *string  `json:"sender_bank_name,omitempty" form:"sender_bank_name" param:"sender_bank_name" url:"sender_bank_name"`
	SenderBankNameNEQ          *string  `json:"sender_bank_name_neq,omitempty" form:"sender_bank_name_neq" param:"sender_bank_name_neq" url:"sender_bank_name_neq"`
	SenderBankNameIn           []string `json:"sender_bank_name_in,omitempty" form:"sender_bank_name_in" param:"sender_bank_name_in" url:"sender_bank_name_in"`
	SenderBankNameNotIn        []string `json:"sender_bank_name_not_in,omitempty" form:"sender_bank_name_not_in" param:"sender_bank_name_not_in" url:"sender_bank_name_not_in"`
	SenderBankNameGT           *string  `json:"sender_bank_name_gt,omitempty" form:"sender_bank_name_gt" param:"sender_bank_name_gt" url:"sender_bank_name_gt"`
	SenderBankNameGTE          *string  `json:"sender_bank_name_gte,omitempty" form:"sender_bank_name_gte" param:"sender_bank_name_gte" url:"sender_bank_name_gte"`
	SenderBankNameLT           *string  `json:"sender_bank_name_lt,omitempty" form:"sender_bank_name_lt" param:"sender_bank_name_lt" url:"sender_bank_name_lt"`
	SenderBankNameLTE          *string  `json:"sender_bank_name_lte,omitempty" form:"sender_bank_name_lte" param:"sender_bank_name_lte" url:"sender_bank_name_lte"`
	SenderBankNameContains     *string  `json:"sender_bank_name_contains,omitempty" form:"sender_bank_name_contains" param:"sender_bank_name_contains" url:"sender_bank_name_contains"`
	SenderBankNameHasPrefix    *string  `json:"sender_bank_name_has_prefix,omitempty" form:"sender_bank_name_has_prefix" param:"sender_bank_name_has_prefix" url:"sender_bank_name_has_prefix"`
	SenderBankNameHasSuffix    *string  `json:"sender_bank_name_has_suffix,omitempty" form:"sender_bank_name_has_suffix" param:"sender_bank_name_has_suffix" url:"sender_bank_name_has_suffix"`
	SenderBankNameEqualFold    *string  `json:"sender_bank_name_equal_fold,omitempty" form:"sender_bank_name_equal_fold" param:"sender_bank_name_equal_fold" url:"sender_bank_name_equal_fold"`
	SenderBankNameContainsFold *string  `json:"sender_bank_name_contains_fold,omitempty" form:"sender_bank_name_contains_fold" param:"sender_bank_name_contains_fold" url:"sender_bank_name_contains_fold"`

	// "sender_name" field predicates.
	SenderName             *string  `json:"sender_name,omitempty" form:"sender_name" param:"sender_name" url:"sender_name"`
	SenderNameNEQ          *string  `json:"sender_name_neq,omitempty" form:"sender_name_neq" param:"sender_name_neq" url:"sender_name_neq"`
	SenderNameIn           []string `json:"sender_name_in,omitempty" form:"sender_name_in" param:"sender_name_in" url:"sender_name_in"`
	SenderNameNotIn        []string `json:"sender_name_not_in,omitempty" form:"sender_name_not_in" param:"sender_name_not_in" url:"sender_name_not_in"`
	SenderNameGT           *string  `json:"sender_name_gt,omitempty" form:"sender_name_gt" param:"sender_name_gt" url:"sender_name_gt"`
	SenderNameGTE          *string  `json:"sender_name_gte,omitempty" form:"sender_name_gte" param:"sender_name_gte" url:"sender_name_gte"`
	SenderNameLT           *string  `json:"sender_name_lt,omitempty" form:"sender_name_lt" param:"sender_name_lt" url:"sender_name_lt"`
	SenderNameLTE          *string  `json:"sender_name_lte,omitempty" form:"sender_name_lte" param:"sender_name_lte" url:"sender_name_lte"`
	SenderNameContains     *string  `json:"sender_name_contains,omitempty" form:"sender_name_contains" param:"sender_name_contains" url:"sender_name_contains"`
	SenderNameHasPrefix    *string  `json:"sender_name_has_prefix,omitempty" form:"sender_name_has_prefix" param:"sender_name_has_prefix" url:"sender_name_has_prefix"`
	SenderNameHasSuffix    *string  `json:"sender_name_has_suffix,omitempty" form:"sender_name_has_suffix" param:"sender_name_has_suffix" url:"sender_name_has_suffix"`
	SenderNameEqualFold    *string  `json:"sender_name_equal_fold,omitempty" form:"sender_name_equal_fold" param:"sender_name_equal_fold" url:"sender_name_equal_fold"`
	SenderNameContainsFold *string  `json:"sender_name_contains_fold,omitempty" form:"sender_name_contains_fold" param:"sender_name_contains_fold" url:"sender_name_contains_fold"`

	// "sender_id" field predicates.
	SenderID       *uuid.UUID  `json:"sender_id,omitempty" form:"sender_id" param:"sender_id" url:"sender_id"`
	SenderIDNEQ    *uuid.UUID  `json:"sender_id_neq,omitempty" form:"sender_id_neq" param:"sender_id_neq" url:"sender_id_neq"`
	SenderIDIn     []uuid.UUID `json:"sender_id_in,omitempty" form:"sender_id_in" param:"sender_id_in" url:"sender_id_in"`
	SenderIDNotIn  []uuid.UUID `json:"sender_id_not_in,omitempty" form:"sender_id_not_in" param:"sender_id_not_in" url:"sender_id_not_in"`
	SenderIDIsNil  bool        `json:"sender_id_is_nil,omitempty" form:"sender_id_is_nil" param:"sender_id_is_nil" url:"sender_id_is_nil"`
	SenderIDNotNil bool        `json:"sender_id_not_nil,omitempty" form:"sender_id_not_nil" param:"sender_id_not_nil" url:"sender_id_not_nil"`

	// "amount" field predicates.
	Amount      *decimal.Decimal  `json:"amount,omitempty" form:"amount" param:"amount" url:"amount"`
	AmountNEQ   *decimal.Decimal  `json:"amount_neq,omitempty" form:"amount_neq" param:"amount_neq" url:"amount_neq"`
	AmountIn    []decimal.Decimal `json:"amount_in,omitempty" form:"amount_in" param:"amount_in" url:"amount_in"`
	AmountNotIn []decimal.Decimal `json:"amount_not_in,omitempty" form:"amount_not_in" param:"amount_not_in" url:"amount_not_in"`
	AmountGT    *decimal.Decimal  `json:"amount_gt,omitempty" form:"amount_gt" param:"amount_gt" url:"amount_gt"`
	AmountGTE   *decimal.Decimal  `json:"amount_gte,omitempty" form:"amount_gte" param:"amount_gte" url:"amount_gte"`
	AmountLT    *decimal.Decimal  `json:"amount_lt,omitempty" form:"amount_lt" param:"amount_lt" url:"amount_lt"`
	AmountLTE   *decimal.Decimal  `json:"amount_lte,omitempty" form:"amount_lte" param:"amount_lte" url:"amount_lte"`

	// "transaction_type" field predicates.
	TransactionType      *transaction.TransactionType  `json:"transaction_type,omitempty" form:"transaction_type" param:"transaction_type" url:"transaction_type"`
	TransactionTypeNEQ   *transaction.TransactionType  `json:"transaction_type_neq,omitempty" form:"transaction_type_neq" param:"transaction_type_neq" url:"transaction_type_neq"`
	TransactionTypeIn    []transaction.TransactionType `json:"transaction_type_in,omitempty" form:"transaction_type_in" param:"transaction_type_in" url:"transaction_type_in"`
	TransactionTypeNotIn []transaction.TransactionType `json:"transaction_type_not_in,omitempty" form:"transaction_type_not_in" param:"transaction_type_not_in" url:"transaction_type_not_in"`

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

	// "source_transaction" edge predicates.
	HasSourceTransaction     *bool                    `json:"has_source_transaction,omitempty" form:"has_source_transaction" param:"has_source_transaction" url:"has_source_transaction"`
	HasSourceTransactionWith []*TransactionWhereInput `json:"has_source_transaction_with,omitempty" form:"has_source_transaction_with" param:"has_source_transaction_with" url:"has_source_transaction_with"`

	// "fee_transaction" edge predicates.
	HasFeeTransaction     *bool                    `json:"has_fee_transaction,omitempty" form:"has_fee_transaction" param:"has_fee_transaction" url:"has_fee_transaction"`
	HasFeeTransactionWith []*TransactionWhereInput `json:"has_fee_transaction_with,omitempty" form:"has_fee_transaction_with" param:"has_fee_transaction_with" url:"has_fee_transaction_with"`

	// "receiver" edge predicates.
	HasReceiver     *bool                    `json:"has_receiver,omitempty" form:"has_receiver" param:"has_receiver" url:"has_receiver"`
	HasReceiverWith []*BankAccountWhereInput `json:"has_receiver_with,omitempty" form:"has_receiver_with" param:"has_receiver_with" url:"has_receiver_with"`

	// "sender" edge predicates.
	HasSender     *bool                    `json:"has_sender,omitempty" form:"has_sender" param:"has_sender" url:"has_sender"`
	HasSenderWith []*BankAccountWhereInput `json:"has_sender_with,omitempty" form:"has_sender_with" param:"has_sender_with" url:"has_sender_with"`

	// "debt" edge predicates.
	HasDebt     *bool             `json:"has_debt,omitempty" form:"has_debt" param:"has_debt" url:"has_debt"`
	HasDebtWith []*DebtWhereInput `json:"has_debt_with,omitempty" form:"has_debt_with" param:"has_debt_with" url:"has_debt_with"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *TransactionWhereInput) AddPredicates(predicates ...predicate.Transaction) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the TransactionWhereInput filter on the TransactionQuery builder.
func (i *TransactionWhereInput) Filter(q *TransactionQuery) (*TransactionQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyTransactionWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyTransactionWhereInput is returned in case the TransactionWhereInput is empty.
var ErrEmptyTransactionWhereInput = errors.New("ent: empty predicate TransactionWhereInput")

// P returns a predicate for filtering transactions.
// An error is returned if the input is empty or invalid.
func (i *TransactionWhereInput) P() (predicate.Transaction, error) {
	var predicates []predicate.Transaction
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, transaction.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Transaction, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, transaction.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Transaction, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, transaction.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, transaction.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, transaction.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, transaction.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, transaction.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, transaction.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, transaction.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, transaction.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, transaction.IDLTE(*i.IDLTE))
	}
	if i.CreateTime != nil {
		predicates = append(predicates, transaction.CreateTimeEQ(*i.CreateTime))
	}
	if i.CreateTimeNEQ != nil {
		predicates = append(predicates, transaction.CreateTimeNEQ(*i.CreateTimeNEQ))
	}
	if len(i.CreateTimeIn) > 0 {
		predicates = append(predicates, transaction.CreateTimeIn(i.CreateTimeIn...))
	}
	if len(i.CreateTimeNotIn) > 0 {
		predicates = append(predicates, transaction.CreateTimeNotIn(i.CreateTimeNotIn...))
	}
	if i.CreateTimeGT != nil {
		predicates = append(predicates, transaction.CreateTimeGT(*i.CreateTimeGT))
	}
	if i.CreateTimeGTE != nil {
		predicates = append(predicates, transaction.CreateTimeGTE(*i.CreateTimeGTE))
	}
	if i.CreateTimeLT != nil {
		predicates = append(predicates, transaction.CreateTimeLT(*i.CreateTimeLT))
	}
	if i.CreateTimeLTE != nil {
		predicates = append(predicates, transaction.CreateTimeLTE(*i.CreateTimeLTE))
	}
	if i.UpdateTime != nil {
		predicates = append(predicates, transaction.UpdateTimeEQ(*i.UpdateTime))
	}
	if i.UpdateTimeNEQ != nil {
		predicates = append(predicates, transaction.UpdateTimeNEQ(*i.UpdateTimeNEQ))
	}
	if len(i.UpdateTimeIn) > 0 {
		predicates = append(predicates, transaction.UpdateTimeIn(i.UpdateTimeIn...))
	}
	if len(i.UpdateTimeNotIn) > 0 {
		predicates = append(predicates, transaction.UpdateTimeNotIn(i.UpdateTimeNotIn...))
	}
	if i.UpdateTimeGT != nil {
		predicates = append(predicates, transaction.UpdateTimeGT(*i.UpdateTimeGT))
	}
	if i.UpdateTimeGTE != nil {
		predicates = append(predicates, transaction.UpdateTimeGTE(*i.UpdateTimeGTE))
	}
	if i.UpdateTimeLT != nil {
		predicates = append(predicates, transaction.UpdateTimeLT(*i.UpdateTimeLT))
	}
	if i.UpdateTimeLTE != nil {
		predicates = append(predicates, transaction.UpdateTimeLTE(*i.UpdateTimeLTE))
	}
	if i.SourceTransactionID != nil {
		predicates = append(predicates, transaction.SourceTransactionIDEQ(*i.SourceTransactionID))
	}
	if i.SourceTransactionIDNEQ != nil {
		predicates = append(predicates, transaction.SourceTransactionIDNEQ(*i.SourceTransactionIDNEQ))
	}
	if len(i.SourceTransactionIDIn) > 0 {
		predicates = append(predicates, transaction.SourceTransactionIDIn(i.SourceTransactionIDIn...))
	}
	if len(i.SourceTransactionIDNotIn) > 0 {
		predicates = append(predicates, transaction.SourceTransactionIDNotIn(i.SourceTransactionIDNotIn...))
	}
	if i.SourceTransactionIDIsNil {
		predicates = append(predicates, transaction.SourceTransactionIDIsNil())
	}
	if i.SourceTransactionIDNotNil {
		predicates = append(predicates, transaction.SourceTransactionIDNotNil())
	}
	if i.Status != nil {
		predicates = append(predicates, transaction.StatusEQ(*i.Status))
	}
	if i.StatusNEQ != nil {
		predicates = append(predicates, transaction.StatusNEQ(*i.StatusNEQ))
	}
	if len(i.StatusIn) > 0 {
		predicates = append(predicates, transaction.StatusIn(i.StatusIn...))
	}
	if len(i.StatusNotIn) > 0 {
		predicates = append(predicates, transaction.StatusNotIn(i.StatusNotIn...))
	}
	if i.ReceiverBankAccountNumber != nil {
		predicates = append(predicates, transaction.ReceiverBankAccountNumberEQ(*i.ReceiverBankAccountNumber))
	}
	if i.ReceiverBankAccountNumberNEQ != nil {
		predicates = append(predicates, transaction.ReceiverBankAccountNumberNEQ(*i.ReceiverBankAccountNumberNEQ))
	}
	if len(i.ReceiverBankAccountNumberIn) > 0 {
		predicates = append(predicates, transaction.ReceiverBankAccountNumberIn(i.ReceiverBankAccountNumberIn...))
	}
	if len(i.ReceiverBankAccountNumberNotIn) > 0 {
		predicates = append(predicates, transaction.ReceiverBankAccountNumberNotIn(i.ReceiverBankAccountNumberNotIn...))
	}
	if i.ReceiverBankAccountNumberGT != nil {
		predicates = append(predicates, transaction.ReceiverBankAccountNumberGT(*i.ReceiverBankAccountNumberGT))
	}
	if i.ReceiverBankAccountNumberGTE != nil {
		predicates = append(predicates, transaction.ReceiverBankAccountNumberGTE(*i.ReceiverBankAccountNumberGTE))
	}
	if i.ReceiverBankAccountNumberLT != nil {
		predicates = append(predicates, transaction.ReceiverBankAccountNumberLT(*i.ReceiverBankAccountNumberLT))
	}
	if i.ReceiverBankAccountNumberLTE != nil {
		predicates = append(predicates, transaction.ReceiverBankAccountNumberLTE(*i.ReceiverBankAccountNumberLTE))
	}
	if i.ReceiverBankAccountNumberContains != nil {
		predicates = append(predicates, transaction.ReceiverBankAccountNumberContains(*i.ReceiverBankAccountNumberContains))
	}
	if i.ReceiverBankAccountNumberHasPrefix != nil {
		predicates = append(predicates, transaction.ReceiverBankAccountNumberHasPrefix(*i.ReceiverBankAccountNumberHasPrefix))
	}
	if i.ReceiverBankAccountNumberHasSuffix != nil {
		predicates = append(predicates, transaction.ReceiverBankAccountNumberHasSuffix(*i.ReceiverBankAccountNumberHasSuffix))
	}
	if i.ReceiverBankAccountNumberEqualFold != nil {
		predicates = append(predicates, transaction.ReceiverBankAccountNumberEqualFold(*i.ReceiverBankAccountNumberEqualFold))
	}
	if i.ReceiverBankAccountNumberContainsFold != nil {
		predicates = append(predicates, transaction.ReceiverBankAccountNumberContainsFold(*i.ReceiverBankAccountNumberContainsFold))
	}
	if i.ReceiverBankName != nil {
		predicates = append(predicates, transaction.ReceiverBankNameEQ(*i.ReceiverBankName))
	}
	if i.ReceiverBankNameNEQ != nil {
		predicates = append(predicates, transaction.ReceiverBankNameNEQ(*i.ReceiverBankNameNEQ))
	}
	if len(i.ReceiverBankNameIn) > 0 {
		predicates = append(predicates, transaction.ReceiverBankNameIn(i.ReceiverBankNameIn...))
	}
	if len(i.ReceiverBankNameNotIn) > 0 {
		predicates = append(predicates, transaction.ReceiverBankNameNotIn(i.ReceiverBankNameNotIn...))
	}
	if i.ReceiverBankNameGT != nil {
		predicates = append(predicates, transaction.ReceiverBankNameGT(*i.ReceiverBankNameGT))
	}
	if i.ReceiverBankNameGTE != nil {
		predicates = append(predicates, transaction.ReceiverBankNameGTE(*i.ReceiverBankNameGTE))
	}
	if i.ReceiverBankNameLT != nil {
		predicates = append(predicates, transaction.ReceiverBankNameLT(*i.ReceiverBankNameLT))
	}
	if i.ReceiverBankNameLTE != nil {
		predicates = append(predicates, transaction.ReceiverBankNameLTE(*i.ReceiverBankNameLTE))
	}
	if i.ReceiverBankNameContains != nil {
		predicates = append(predicates, transaction.ReceiverBankNameContains(*i.ReceiverBankNameContains))
	}
	if i.ReceiverBankNameHasPrefix != nil {
		predicates = append(predicates, transaction.ReceiverBankNameHasPrefix(*i.ReceiverBankNameHasPrefix))
	}
	if i.ReceiverBankNameHasSuffix != nil {
		predicates = append(predicates, transaction.ReceiverBankNameHasSuffix(*i.ReceiverBankNameHasSuffix))
	}
	if i.ReceiverBankNameEqualFold != nil {
		predicates = append(predicates, transaction.ReceiverBankNameEqualFold(*i.ReceiverBankNameEqualFold))
	}
	if i.ReceiverBankNameContainsFold != nil {
		predicates = append(predicates, transaction.ReceiverBankNameContainsFold(*i.ReceiverBankNameContainsFold))
	}
	if i.ReceiverName != nil {
		predicates = append(predicates, transaction.ReceiverNameEQ(*i.ReceiverName))
	}
	if i.ReceiverNameNEQ != nil {
		predicates = append(predicates, transaction.ReceiverNameNEQ(*i.ReceiverNameNEQ))
	}
	if len(i.ReceiverNameIn) > 0 {
		predicates = append(predicates, transaction.ReceiverNameIn(i.ReceiverNameIn...))
	}
	if len(i.ReceiverNameNotIn) > 0 {
		predicates = append(predicates, transaction.ReceiverNameNotIn(i.ReceiverNameNotIn...))
	}
	if i.ReceiverNameGT != nil {
		predicates = append(predicates, transaction.ReceiverNameGT(*i.ReceiverNameGT))
	}
	if i.ReceiverNameGTE != nil {
		predicates = append(predicates, transaction.ReceiverNameGTE(*i.ReceiverNameGTE))
	}
	if i.ReceiverNameLT != nil {
		predicates = append(predicates, transaction.ReceiverNameLT(*i.ReceiverNameLT))
	}
	if i.ReceiverNameLTE != nil {
		predicates = append(predicates, transaction.ReceiverNameLTE(*i.ReceiverNameLTE))
	}
	if i.ReceiverNameContains != nil {
		predicates = append(predicates, transaction.ReceiverNameContains(*i.ReceiverNameContains))
	}
	if i.ReceiverNameHasPrefix != nil {
		predicates = append(predicates, transaction.ReceiverNameHasPrefix(*i.ReceiverNameHasPrefix))
	}
	if i.ReceiverNameHasSuffix != nil {
		predicates = append(predicates, transaction.ReceiverNameHasSuffix(*i.ReceiverNameHasSuffix))
	}
	if i.ReceiverNameEqualFold != nil {
		predicates = append(predicates, transaction.ReceiverNameEqualFold(*i.ReceiverNameEqualFold))
	}
	if i.ReceiverNameContainsFold != nil {
		predicates = append(predicates, transaction.ReceiverNameContainsFold(*i.ReceiverNameContainsFold))
	}
	if i.ReceiverID != nil {
		predicates = append(predicates, transaction.ReceiverIDEQ(*i.ReceiverID))
	}
	if i.ReceiverIDNEQ != nil {
		predicates = append(predicates, transaction.ReceiverIDNEQ(*i.ReceiverIDNEQ))
	}
	if len(i.ReceiverIDIn) > 0 {
		predicates = append(predicates, transaction.ReceiverIDIn(i.ReceiverIDIn...))
	}
	if len(i.ReceiverIDNotIn) > 0 {
		predicates = append(predicates, transaction.ReceiverIDNotIn(i.ReceiverIDNotIn...))
	}
	if i.ReceiverIDIsNil {
		predicates = append(predicates, transaction.ReceiverIDIsNil())
	}
	if i.ReceiverIDNotNil {
		predicates = append(predicates, transaction.ReceiverIDNotNil())
	}
	if i.SenderBankAccountNumber != nil {
		predicates = append(predicates, transaction.SenderBankAccountNumberEQ(*i.SenderBankAccountNumber))
	}
	if i.SenderBankAccountNumberNEQ != nil {
		predicates = append(predicates, transaction.SenderBankAccountNumberNEQ(*i.SenderBankAccountNumberNEQ))
	}
	if len(i.SenderBankAccountNumberIn) > 0 {
		predicates = append(predicates, transaction.SenderBankAccountNumberIn(i.SenderBankAccountNumberIn...))
	}
	if len(i.SenderBankAccountNumberNotIn) > 0 {
		predicates = append(predicates, transaction.SenderBankAccountNumberNotIn(i.SenderBankAccountNumberNotIn...))
	}
	if i.SenderBankAccountNumberGT != nil {
		predicates = append(predicates, transaction.SenderBankAccountNumberGT(*i.SenderBankAccountNumberGT))
	}
	if i.SenderBankAccountNumberGTE != nil {
		predicates = append(predicates, transaction.SenderBankAccountNumberGTE(*i.SenderBankAccountNumberGTE))
	}
	if i.SenderBankAccountNumberLT != nil {
		predicates = append(predicates, transaction.SenderBankAccountNumberLT(*i.SenderBankAccountNumberLT))
	}
	if i.SenderBankAccountNumberLTE != nil {
		predicates = append(predicates, transaction.SenderBankAccountNumberLTE(*i.SenderBankAccountNumberLTE))
	}
	if i.SenderBankAccountNumberContains != nil {
		predicates = append(predicates, transaction.SenderBankAccountNumberContains(*i.SenderBankAccountNumberContains))
	}
	if i.SenderBankAccountNumberHasPrefix != nil {
		predicates = append(predicates, transaction.SenderBankAccountNumberHasPrefix(*i.SenderBankAccountNumberHasPrefix))
	}
	if i.SenderBankAccountNumberHasSuffix != nil {
		predicates = append(predicates, transaction.SenderBankAccountNumberHasSuffix(*i.SenderBankAccountNumberHasSuffix))
	}
	if i.SenderBankAccountNumberEqualFold != nil {
		predicates = append(predicates, transaction.SenderBankAccountNumberEqualFold(*i.SenderBankAccountNumberEqualFold))
	}
	if i.SenderBankAccountNumberContainsFold != nil {
		predicates = append(predicates, transaction.SenderBankAccountNumberContainsFold(*i.SenderBankAccountNumberContainsFold))
	}
	if i.SenderBankName != nil {
		predicates = append(predicates, transaction.SenderBankNameEQ(*i.SenderBankName))
	}
	if i.SenderBankNameNEQ != nil {
		predicates = append(predicates, transaction.SenderBankNameNEQ(*i.SenderBankNameNEQ))
	}
	if len(i.SenderBankNameIn) > 0 {
		predicates = append(predicates, transaction.SenderBankNameIn(i.SenderBankNameIn...))
	}
	if len(i.SenderBankNameNotIn) > 0 {
		predicates = append(predicates, transaction.SenderBankNameNotIn(i.SenderBankNameNotIn...))
	}
	if i.SenderBankNameGT != nil {
		predicates = append(predicates, transaction.SenderBankNameGT(*i.SenderBankNameGT))
	}
	if i.SenderBankNameGTE != nil {
		predicates = append(predicates, transaction.SenderBankNameGTE(*i.SenderBankNameGTE))
	}
	if i.SenderBankNameLT != nil {
		predicates = append(predicates, transaction.SenderBankNameLT(*i.SenderBankNameLT))
	}
	if i.SenderBankNameLTE != nil {
		predicates = append(predicates, transaction.SenderBankNameLTE(*i.SenderBankNameLTE))
	}
	if i.SenderBankNameContains != nil {
		predicates = append(predicates, transaction.SenderBankNameContains(*i.SenderBankNameContains))
	}
	if i.SenderBankNameHasPrefix != nil {
		predicates = append(predicates, transaction.SenderBankNameHasPrefix(*i.SenderBankNameHasPrefix))
	}
	if i.SenderBankNameHasSuffix != nil {
		predicates = append(predicates, transaction.SenderBankNameHasSuffix(*i.SenderBankNameHasSuffix))
	}
	if i.SenderBankNameEqualFold != nil {
		predicates = append(predicates, transaction.SenderBankNameEqualFold(*i.SenderBankNameEqualFold))
	}
	if i.SenderBankNameContainsFold != nil {
		predicates = append(predicates, transaction.SenderBankNameContainsFold(*i.SenderBankNameContainsFold))
	}
	if i.SenderName != nil {
		predicates = append(predicates, transaction.SenderNameEQ(*i.SenderName))
	}
	if i.SenderNameNEQ != nil {
		predicates = append(predicates, transaction.SenderNameNEQ(*i.SenderNameNEQ))
	}
	if len(i.SenderNameIn) > 0 {
		predicates = append(predicates, transaction.SenderNameIn(i.SenderNameIn...))
	}
	if len(i.SenderNameNotIn) > 0 {
		predicates = append(predicates, transaction.SenderNameNotIn(i.SenderNameNotIn...))
	}
	if i.SenderNameGT != nil {
		predicates = append(predicates, transaction.SenderNameGT(*i.SenderNameGT))
	}
	if i.SenderNameGTE != nil {
		predicates = append(predicates, transaction.SenderNameGTE(*i.SenderNameGTE))
	}
	if i.SenderNameLT != nil {
		predicates = append(predicates, transaction.SenderNameLT(*i.SenderNameLT))
	}
	if i.SenderNameLTE != nil {
		predicates = append(predicates, transaction.SenderNameLTE(*i.SenderNameLTE))
	}
	if i.SenderNameContains != nil {
		predicates = append(predicates, transaction.SenderNameContains(*i.SenderNameContains))
	}
	if i.SenderNameHasPrefix != nil {
		predicates = append(predicates, transaction.SenderNameHasPrefix(*i.SenderNameHasPrefix))
	}
	if i.SenderNameHasSuffix != nil {
		predicates = append(predicates, transaction.SenderNameHasSuffix(*i.SenderNameHasSuffix))
	}
	if i.SenderNameEqualFold != nil {
		predicates = append(predicates, transaction.SenderNameEqualFold(*i.SenderNameEqualFold))
	}
	if i.SenderNameContainsFold != nil {
		predicates = append(predicates, transaction.SenderNameContainsFold(*i.SenderNameContainsFold))
	}
	if i.SenderID != nil {
		predicates = append(predicates, transaction.SenderIDEQ(*i.SenderID))
	}
	if i.SenderIDNEQ != nil {
		predicates = append(predicates, transaction.SenderIDNEQ(*i.SenderIDNEQ))
	}
	if len(i.SenderIDIn) > 0 {
		predicates = append(predicates, transaction.SenderIDIn(i.SenderIDIn...))
	}
	if len(i.SenderIDNotIn) > 0 {
		predicates = append(predicates, transaction.SenderIDNotIn(i.SenderIDNotIn...))
	}
	if i.SenderIDIsNil {
		predicates = append(predicates, transaction.SenderIDIsNil())
	}
	if i.SenderIDNotNil {
		predicates = append(predicates, transaction.SenderIDNotNil())
	}
	if i.Amount != nil {
		predicates = append(predicates, transaction.AmountEQ(*i.Amount))
	}
	if i.AmountNEQ != nil {
		predicates = append(predicates, transaction.AmountNEQ(*i.AmountNEQ))
	}
	if len(i.AmountIn) > 0 {
		predicates = append(predicates, transaction.AmountIn(i.AmountIn...))
	}
	if len(i.AmountNotIn) > 0 {
		predicates = append(predicates, transaction.AmountNotIn(i.AmountNotIn...))
	}
	if i.AmountGT != nil {
		predicates = append(predicates, transaction.AmountGT(*i.AmountGT))
	}
	if i.AmountGTE != nil {
		predicates = append(predicates, transaction.AmountGTE(*i.AmountGTE))
	}
	if i.AmountLT != nil {
		predicates = append(predicates, transaction.AmountLT(*i.AmountLT))
	}
	if i.AmountLTE != nil {
		predicates = append(predicates, transaction.AmountLTE(*i.AmountLTE))
	}
	if i.TransactionType != nil {
		predicates = append(predicates, transaction.TransactionTypeEQ(*i.TransactionType))
	}
	if i.TransactionTypeNEQ != nil {
		predicates = append(predicates, transaction.TransactionTypeNEQ(*i.TransactionTypeNEQ))
	}
	if len(i.TransactionTypeIn) > 0 {
		predicates = append(predicates, transaction.TransactionTypeIn(i.TransactionTypeIn...))
	}
	if len(i.TransactionTypeNotIn) > 0 {
		predicates = append(predicates, transaction.TransactionTypeNotIn(i.TransactionTypeNotIn...))
	}
	if i.Description != nil {
		predicates = append(predicates, transaction.DescriptionEQ(*i.Description))
	}
	if i.DescriptionNEQ != nil {
		predicates = append(predicates, transaction.DescriptionNEQ(*i.DescriptionNEQ))
	}
	if len(i.DescriptionIn) > 0 {
		predicates = append(predicates, transaction.DescriptionIn(i.DescriptionIn...))
	}
	if len(i.DescriptionNotIn) > 0 {
		predicates = append(predicates, transaction.DescriptionNotIn(i.DescriptionNotIn...))
	}
	if i.DescriptionGT != nil {
		predicates = append(predicates, transaction.DescriptionGT(*i.DescriptionGT))
	}
	if i.DescriptionGTE != nil {
		predicates = append(predicates, transaction.DescriptionGTE(*i.DescriptionGTE))
	}
	if i.DescriptionLT != nil {
		predicates = append(predicates, transaction.DescriptionLT(*i.DescriptionLT))
	}
	if i.DescriptionLTE != nil {
		predicates = append(predicates, transaction.DescriptionLTE(*i.DescriptionLTE))
	}
	if i.DescriptionContains != nil {
		predicates = append(predicates, transaction.DescriptionContains(*i.DescriptionContains))
	}
	if i.DescriptionHasPrefix != nil {
		predicates = append(predicates, transaction.DescriptionHasPrefix(*i.DescriptionHasPrefix))
	}
	if i.DescriptionHasSuffix != nil {
		predicates = append(predicates, transaction.DescriptionHasSuffix(*i.DescriptionHasSuffix))
	}
	if i.DescriptionIsNil {
		predicates = append(predicates, transaction.DescriptionIsNil())
	}
	if i.DescriptionNotNil {
		predicates = append(predicates, transaction.DescriptionNotNil())
	}
	if i.DescriptionEqualFold != nil {
		predicates = append(predicates, transaction.DescriptionEqualFold(*i.DescriptionEqualFold))
	}
	if i.DescriptionContainsFold != nil {
		predicates = append(predicates, transaction.DescriptionContainsFold(*i.DescriptionContainsFold))
	}

	if i.HasSourceTransaction != nil {
		p := transaction.HasSourceTransaction()
		if !*i.HasSourceTransaction {
			p = transaction.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasSourceTransactionWith) > 0 {
		with := make([]predicate.Transaction, 0, len(i.HasSourceTransactionWith))
		for _, w := range i.HasSourceTransactionWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasSourceTransactionWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, transaction.HasSourceTransactionWith(with...))
	}
	if i.HasFeeTransaction != nil {
		p := transaction.HasFeeTransaction()
		if !*i.HasFeeTransaction {
			p = transaction.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasFeeTransactionWith) > 0 {
		with := make([]predicate.Transaction, 0, len(i.HasFeeTransactionWith))
		for _, w := range i.HasFeeTransactionWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasFeeTransactionWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, transaction.HasFeeTransactionWith(with...))
	}
	if i.HasReceiver != nil {
		p := transaction.HasReceiver()
		if !*i.HasReceiver {
			p = transaction.Not(p)
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
		predicates = append(predicates, transaction.HasReceiverWith(with...))
	}
	if i.HasSender != nil {
		p := transaction.HasSender()
		if !*i.HasSender {
			p = transaction.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasSenderWith) > 0 {
		with := make([]predicate.BankAccount, 0, len(i.HasSenderWith))
		for _, w := range i.HasSenderWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasSenderWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, transaction.HasSenderWith(with...))
	}
	if i.HasDebt != nil {
		p := transaction.HasDebt()
		if !*i.HasDebt {
			p = transaction.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasDebtWith) > 0 {
		with := make([]predicate.Debt, 0, len(i.HasDebtWith))
		for _, w := range i.HasDebtWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasDebtWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, transaction.HasDebtWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyTransactionWhereInput
	case 1:
		return predicates[0], nil
	default:
		return transaction.And(predicates...), nil
	}
}

// Transactions is a parsable slice of Transaction.
type Transactions []*Transaction

func (t Transactions) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
