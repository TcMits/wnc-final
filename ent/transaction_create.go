// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/TcMits/wnc-final/ent/bankaccount"
	"github.com/TcMits/wnc-final/ent/debt"
	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TransactionCreate is the builder for creating a Transaction entity.
type TransactionCreate struct {
	config
	mutation *TransactionMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (tc *TransactionCreate) SetCreateTime(t time.Time) *TransactionCreate {
	tc.mutation.SetCreateTime(t)
	return tc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableCreateTime(t *time.Time) *TransactionCreate {
	if t != nil {
		tc.SetCreateTime(*t)
	}
	return tc
}

// SetUpdateTime sets the "update_time" field.
func (tc *TransactionCreate) SetUpdateTime(t time.Time) *TransactionCreate {
	tc.mutation.SetUpdateTime(t)
	return tc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableUpdateTime(t *time.Time) *TransactionCreate {
	if t != nil {
		tc.SetUpdateTime(*t)
	}
	return tc
}

// SetSourceTransactionID sets the "source_transaction_id" field.
func (tc *TransactionCreate) SetSourceTransactionID(u uuid.UUID) *TransactionCreate {
	tc.mutation.SetSourceTransactionID(u)
	return tc
}

// SetNillableSourceTransactionID sets the "source_transaction_id" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableSourceTransactionID(u *uuid.UUID) *TransactionCreate {
	if u != nil {
		tc.SetSourceTransactionID(*u)
	}
	return tc
}

// SetStatus sets the "status" field.
func (tc *TransactionCreate) SetStatus(t transaction.Status) *TransactionCreate {
	tc.mutation.SetStatus(t)
	return tc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableStatus(t *transaction.Status) *TransactionCreate {
	if t != nil {
		tc.SetStatus(*t)
	}
	return tc
}

// SetReceiverBankAccountNumber sets the "receiver_bank_account_number" field.
func (tc *TransactionCreate) SetReceiverBankAccountNumber(s string) *TransactionCreate {
	tc.mutation.SetReceiverBankAccountNumber(s)
	return tc
}

// SetReceiverBankName sets the "receiver_bank_name" field.
func (tc *TransactionCreate) SetReceiverBankName(s string) *TransactionCreate {
	tc.mutation.SetReceiverBankName(s)
	return tc
}

// SetReceiverName sets the "receiver_name" field.
func (tc *TransactionCreate) SetReceiverName(s string) *TransactionCreate {
	tc.mutation.SetReceiverName(s)
	return tc
}

// SetReceiverID sets the "receiver_id" field.
func (tc *TransactionCreate) SetReceiverID(u uuid.UUID) *TransactionCreate {
	tc.mutation.SetReceiverID(u)
	return tc
}

// SetNillableReceiverID sets the "receiver_id" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableReceiverID(u *uuid.UUID) *TransactionCreate {
	if u != nil {
		tc.SetReceiverID(*u)
	}
	return tc
}

// SetSenderBankAccountNumber sets the "sender_bank_account_number" field.
func (tc *TransactionCreate) SetSenderBankAccountNumber(s string) *TransactionCreate {
	tc.mutation.SetSenderBankAccountNumber(s)
	return tc
}

// SetSenderBankName sets the "sender_bank_name" field.
func (tc *TransactionCreate) SetSenderBankName(s string) *TransactionCreate {
	tc.mutation.SetSenderBankName(s)
	return tc
}

// SetSenderName sets the "sender_name" field.
func (tc *TransactionCreate) SetSenderName(s string) *TransactionCreate {
	tc.mutation.SetSenderName(s)
	return tc
}

// SetSenderID sets the "sender_id" field.
func (tc *TransactionCreate) SetSenderID(u uuid.UUID) *TransactionCreate {
	tc.mutation.SetSenderID(u)
	return tc
}

// SetNillableSenderID sets the "sender_id" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableSenderID(u *uuid.UUID) *TransactionCreate {
	if u != nil {
		tc.SetSenderID(*u)
	}
	return tc
}

// SetAmount sets the "amount" field.
func (tc *TransactionCreate) SetAmount(d decimal.Decimal) *TransactionCreate {
	tc.mutation.SetAmount(d)
	return tc
}

// SetTransactionType sets the "transaction_type" field.
func (tc *TransactionCreate) SetTransactionType(tt transaction.TransactionType) *TransactionCreate {
	tc.mutation.SetTransactionType(tt)
	return tc
}

// SetDescription sets the "description" field.
func (tc *TransactionCreate) SetDescription(s string) *TransactionCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableDescription(s *string) *TransactionCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TransactionCreate) SetID(u uuid.UUID) *TransactionCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableID(u *uuid.UUID) *TransactionCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// SetSourceTransaction sets the "source_transaction" edge to the Transaction entity.
func (tc *TransactionCreate) SetSourceTransaction(t *Transaction) *TransactionCreate {
	return tc.SetSourceTransactionID(t.ID)
}

// SetFeeTransactionID sets the "fee_transaction" edge to the Transaction entity by ID.
func (tc *TransactionCreate) SetFeeTransactionID(id uuid.UUID) *TransactionCreate {
	tc.mutation.SetFeeTransactionID(id)
	return tc
}

// SetNillableFeeTransactionID sets the "fee_transaction" edge to the Transaction entity by ID if the given value is not nil.
func (tc *TransactionCreate) SetNillableFeeTransactionID(id *uuid.UUID) *TransactionCreate {
	if id != nil {
		tc = tc.SetFeeTransactionID(*id)
	}
	return tc
}

// SetFeeTransaction sets the "fee_transaction" edge to the Transaction entity.
func (tc *TransactionCreate) SetFeeTransaction(t *Transaction) *TransactionCreate {
	return tc.SetFeeTransactionID(t.ID)
}

// SetReceiver sets the "receiver" edge to the BankAccount entity.
func (tc *TransactionCreate) SetReceiver(b *BankAccount) *TransactionCreate {
	return tc.SetReceiverID(b.ID)
}

// SetSender sets the "sender" edge to the BankAccount entity.
func (tc *TransactionCreate) SetSender(b *BankAccount) *TransactionCreate {
	return tc.SetSenderID(b.ID)
}

// SetDebtID sets the "debt" edge to the Debt entity by ID.
func (tc *TransactionCreate) SetDebtID(id uuid.UUID) *TransactionCreate {
	tc.mutation.SetDebtID(id)
	return tc
}

// SetNillableDebtID sets the "debt" edge to the Debt entity by ID if the given value is not nil.
func (tc *TransactionCreate) SetNillableDebtID(id *uuid.UUID) *TransactionCreate {
	if id != nil {
		tc = tc.SetDebtID(*id)
	}
	return tc
}

// SetDebt sets the "debt" edge to the Debt entity.
func (tc *TransactionCreate) SetDebt(d *Debt) *TransactionCreate {
	return tc.SetDebtID(d.ID)
}

// Mutation returns the TransactionMutation object of the builder.
func (tc *TransactionCreate) Mutation() *TransactionMutation {
	return tc.mutation
}

// Save creates the Transaction in the database.
func (tc *TransactionCreate) Save(ctx context.Context) (*Transaction, error) {
	var (
		err  error
		node *Transaction
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Transaction)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TransactionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransactionCreate) SaveX(ctx context.Context) *Transaction {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TransactionCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TransactionCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TransactionCreate) defaults() {
	if _, ok := tc.mutation.CreateTime(); !ok {
		v := transaction.DefaultCreateTime()
		tc.mutation.SetCreateTime(v)
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		v := transaction.DefaultUpdateTime()
		tc.mutation.SetUpdateTime(v)
	}
	if _, ok := tc.mutation.Status(); !ok {
		v := transaction.DefaultStatus
		tc.mutation.SetStatus(v)
	}
	if _, ok := tc.mutation.Description(); !ok {
		v := transaction.DefaultDescription
		tc.mutation.SetDescription(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := transaction.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TransactionCreate) check() error {
	if _, ok := tc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Transaction.create_time"`)}
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Transaction.update_time"`)}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Transaction.status"`)}
	}
	if v, ok := tc.mutation.Status(); ok {
		if err := transaction.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Transaction.status": %w`, err)}
		}
	}
	if _, ok := tc.mutation.ReceiverBankAccountNumber(); !ok {
		return &ValidationError{Name: "receiver_bank_account_number", err: errors.New(`ent: missing required field "Transaction.receiver_bank_account_number"`)}
	}
	if v, ok := tc.mutation.ReceiverBankAccountNumber(); ok {
		if err := transaction.ReceiverBankAccountNumberValidator(v); err != nil {
			return &ValidationError{Name: "receiver_bank_account_number", err: fmt.Errorf(`ent: validator failed for field "Transaction.receiver_bank_account_number": %w`, err)}
		}
	}
	if _, ok := tc.mutation.ReceiverBankName(); !ok {
		return &ValidationError{Name: "receiver_bank_name", err: errors.New(`ent: missing required field "Transaction.receiver_bank_name"`)}
	}
	if v, ok := tc.mutation.ReceiverBankName(); ok {
		if err := transaction.ReceiverBankNameValidator(v); err != nil {
			return &ValidationError{Name: "receiver_bank_name", err: fmt.Errorf(`ent: validator failed for field "Transaction.receiver_bank_name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.ReceiverName(); !ok {
		return &ValidationError{Name: "receiver_name", err: errors.New(`ent: missing required field "Transaction.receiver_name"`)}
	}
	if v, ok := tc.mutation.ReceiverName(); ok {
		if err := transaction.ReceiverNameValidator(v); err != nil {
			return &ValidationError{Name: "receiver_name", err: fmt.Errorf(`ent: validator failed for field "Transaction.receiver_name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.SenderBankAccountNumber(); !ok {
		return &ValidationError{Name: "sender_bank_account_number", err: errors.New(`ent: missing required field "Transaction.sender_bank_account_number"`)}
	}
	if v, ok := tc.mutation.SenderBankAccountNumber(); ok {
		if err := transaction.SenderBankAccountNumberValidator(v); err != nil {
			return &ValidationError{Name: "sender_bank_account_number", err: fmt.Errorf(`ent: validator failed for field "Transaction.sender_bank_account_number": %w`, err)}
		}
	}
	if _, ok := tc.mutation.SenderBankName(); !ok {
		return &ValidationError{Name: "sender_bank_name", err: errors.New(`ent: missing required field "Transaction.sender_bank_name"`)}
	}
	if v, ok := tc.mutation.SenderBankName(); ok {
		if err := transaction.SenderBankNameValidator(v); err != nil {
			return &ValidationError{Name: "sender_bank_name", err: fmt.Errorf(`ent: validator failed for field "Transaction.sender_bank_name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.SenderName(); !ok {
		return &ValidationError{Name: "sender_name", err: errors.New(`ent: missing required field "Transaction.sender_name"`)}
	}
	if v, ok := tc.mutation.SenderName(); ok {
		if err := transaction.SenderNameValidator(v); err != nil {
			return &ValidationError{Name: "sender_name", err: fmt.Errorf(`ent: validator failed for field "Transaction.sender_name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Transaction.amount"`)}
	}
	if _, ok := tc.mutation.TransactionType(); !ok {
		return &ValidationError{Name: "transaction_type", err: errors.New(`ent: missing required field "Transaction.transaction_type"`)}
	}
	if v, ok := tc.mutation.TransactionType(); ok {
		if err := transaction.TransactionTypeValidator(v); err != nil {
			return &ValidationError{Name: "transaction_type", err: fmt.Errorf(`ent: validator failed for field "Transaction.transaction_type": %w`, err)}
		}
	}
	return nil
}

func (tc *TransactionCreate) sqlSave(ctx context.Context) (*Transaction, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (tc *TransactionCreate) createSpec() (*Transaction, *sqlgraph.CreateSpec) {
	var (
		_node = &Transaction{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: transaction.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.CreateTime(); ok {
		_spec.SetField(transaction.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := tc.mutation.UpdateTime(); ok {
		_spec.SetField(transaction.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(transaction.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.ReceiverBankAccountNumber(); ok {
		_spec.SetField(transaction.FieldReceiverBankAccountNumber, field.TypeString, value)
		_node.ReceiverBankAccountNumber = value
	}
	if value, ok := tc.mutation.ReceiverBankName(); ok {
		_spec.SetField(transaction.FieldReceiverBankName, field.TypeString, value)
		_node.ReceiverBankName = value
	}
	if value, ok := tc.mutation.ReceiverName(); ok {
		_spec.SetField(transaction.FieldReceiverName, field.TypeString, value)
		_node.ReceiverName = value
	}
	if value, ok := tc.mutation.SenderBankAccountNumber(); ok {
		_spec.SetField(transaction.FieldSenderBankAccountNumber, field.TypeString, value)
		_node.SenderBankAccountNumber = value
	}
	if value, ok := tc.mutation.SenderBankName(); ok {
		_spec.SetField(transaction.FieldSenderBankName, field.TypeString, value)
		_node.SenderBankName = value
	}
	if value, ok := tc.mutation.SenderName(); ok {
		_spec.SetField(transaction.FieldSenderName, field.TypeString, value)
		_node.SenderName = value
	}
	if value, ok := tc.mutation.Amount(); ok {
		_spec.SetField(transaction.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if value, ok := tc.mutation.TransactionType(); ok {
		_spec.SetField(transaction.FieldTransactionType, field.TypeEnum, value)
		_node.TransactionType = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(transaction.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := tc.mutation.SourceTransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   transaction.SourceTransactionTable,
			Columns: []string{transaction.SourceTransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SourceTransactionID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.FeeTransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   transaction.FeeTransactionTable,
			Columns: []string{transaction.FeeTransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ReceiverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   transaction.ReceiverTable,
			Columns: []string{transaction.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: bankaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ReceiverID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.SenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   transaction.SenderTable,
			Columns: []string{transaction.SenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: bankaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SenderID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.DebtIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   transaction.DebtTable,
			Columns: []string{transaction.DebtColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: debt.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TransactionCreateBulk is the builder for creating many Transaction entities in bulk.
type TransactionCreateBulk struct {
	config
	builders []*TransactionCreate
}

// Save creates the Transaction entities in the database.
func (tcb *TransactionCreateBulk) Save(ctx context.Context) ([]*Transaction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transaction, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TransactionCreateBulk) SaveX(ctx context.Context) []*Transaction {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TransactionCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TransactionCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
