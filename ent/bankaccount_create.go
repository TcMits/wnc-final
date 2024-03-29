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
	"github.com/TcMits/wnc-final/ent/customer"
	"github.com/TcMits/wnc-final/ent/debt"
	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/google/uuid"
)

// BankAccountCreate is the builder for creating a BankAccount entity.
type BankAccountCreate struct {
	config
	mutation *BankAccountMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (bac *BankAccountCreate) SetCreateTime(t time.Time) *BankAccountCreate {
	bac.mutation.SetCreateTime(t)
	return bac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (bac *BankAccountCreate) SetNillableCreateTime(t *time.Time) *BankAccountCreate {
	if t != nil {
		bac.SetCreateTime(*t)
	}
	return bac
}

// SetUpdateTime sets the "update_time" field.
func (bac *BankAccountCreate) SetUpdateTime(t time.Time) *BankAccountCreate {
	bac.mutation.SetUpdateTime(t)
	return bac
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (bac *BankAccountCreate) SetNillableUpdateTime(t *time.Time) *BankAccountCreate {
	if t != nil {
		bac.SetUpdateTime(*t)
	}
	return bac
}

// SetCustomerID sets the "customer_id" field.
func (bac *BankAccountCreate) SetCustomerID(u uuid.UUID) *BankAccountCreate {
	bac.mutation.SetCustomerID(u)
	return bac
}

// SetCashIn sets the "cash_in" field.
func (bac *BankAccountCreate) SetCashIn(f float64) *BankAccountCreate {
	bac.mutation.SetCashIn(f)
	return bac
}

// SetCashOut sets the "cash_out" field.
func (bac *BankAccountCreate) SetCashOut(f float64) *BankAccountCreate {
	bac.mutation.SetCashOut(f)
	return bac
}

// SetAccountNumber sets the "account_number" field.
func (bac *BankAccountCreate) SetAccountNumber(s string) *BankAccountCreate {
	bac.mutation.SetAccountNumber(s)
	return bac
}

// SetNillableAccountNumber sets the "account_number" field if the given value is not nil.
func (bac *BankAccountCreate) SetNillableAccountNumber(s *string) *BankAccountCreate {
	if s != nil {
		bac.SetAccountNumber(*s)
	}
	return bac
}

// SetIsForPayment sets the "is_for_payment" field.
func (bac *BankAccountCreate) SetIsForPayment(b bool) *BankAccountCreate {
	bac.mutation.SetIsForPayment(b)
	return bac
}

// SetNillableIsForPayment sets the "is_for_payment" field if the given value is not nil.
func (bac *BankAccountCreate) SetNillableIsForPayment(b *bool) *BankAccountCreate {
	if b != nil {
		bac.SetIsForPayment(*b)
	}
	return bac
}

// SetID sets the "id" field.
func (bac *BankAccountCreate) SetID(u uuid.UUID) *BankAccountCreate {
	bac.mutation.SetID(u)
	return bac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bac *BankAccountCreate) SetNillableID(u *uuid.UUID) *BankAccountCreate {
	if u != nil {
		bac.SetID(*u)
	}
	return bac
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (bac *BankAccountCreate) SetCustomer(c *Customer) *BankAccountCreate {
	return bac.SetCustomerID(c.ID)
}

// AddSentTransactionIDs adds the "sent_transaction" edge to the Transaction entity by IDs.
func (bac *BankAccountCreate) AddSentTransactionIDs(ids ...uuid.UUID) *BankAccountCreate {
	bac.mutation.AddSentTransactionIDs(ids...)
	return bac
}

// AddSentTransaction adds the "sent_transaction" edges to the Transaction entity.
func (bac *BankAccountCreate) AddSentTransaction(t ...*Transaction) *BankAccountCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bac.AddSentTransactionIDs(ids...)
}

// AddReceivedTransactionIDs adds the "received_transaction" edge to the Transaction entity by IDs.
func (bac *BankAccountCreate) AddReceivedTransactionIDs(ids ...uuid.UUID) *BankAccountCreate {
	bac.mutation.AddReceivedTransactionIDs(ids...)
	return bac
}

// AddReceivedTransaction adds the "received_transaction" edges to the Transaction entity.
func (bac *BankAccountCreate) AddReceivedTransaction(t ...*Transaction) *BankAccountCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bac.AddReceivedTransactionIDs(ids...)
}

// AddOwnedDebtIDs adds the "owned_debts" edge to the Debt entity by IDs.
func (bac *BankAccountCreate) AddOwnedDebtIDs(ids ...uuid.UUID) *BankAccountCreate {
	bac.mutation.AddOwnedDebtIDs(ids...)
	return bac
}

// AddOwnedDebts adds the "owned_debts" edges to the Debt entity.
func (bac *BankAccountCreate) AddOwnedDebts(d ...*Debt) *BankAccountCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return bac.AddOwnedDebtIDs(ids...)
}

// AddReceivedDebtIDs adds the "received_debts" edge to the Debt entity by IDs.
func (bac *BankAccountCreate) AddReceivedDebtIDs(ids ...uuid.UUID) *BankAccountCreate {
	bac.mutation.AddReceivedDebtIDs(ids...)
	return bac
}

// AddReceivedDebts adds the "received_debts" edges to the Debt entity.
func (bac *BankAccountCreate) AddReceivedDebts(d ...*Debt) *BankAccountCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return bac.AddReceivedDebtIDs(ids...)
}

// Mutation returns the BankAccountMutation object of the builder.
func (bac *BankAccountCreate) Mutation() *BankAccountMutation {
	return bac.mutation
}

// Save creates the BankAccount in the database.
func (bac *BankAccountCreate) Save(ctx context.Context) (*BankAccount, error) {
	var (
		err  error
		node *BankAccount
	)
	bac.defaults()
	if len(bac.hooks) == 0 {
		if err = bac.check(); err != nil {
			return nil, err
		}
		node, err = bac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BankAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bac.check(); err != nil {
				return nil, err
			}
			bac.mutation = mutation
			if node, err = bac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bac.hooks) - 1; i >= 0; i-- {
			if bac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, bac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*BankAccount)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BankAccountMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bac *BankAccountCreate) SaveX(ctx context.Context) *BankAccount {
	v, err := bac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bac *BankAccountCreate) Exec(ctx context.Context) error {
	_, err := bac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bac *BankAccountCreate) ExecX(ctx context.Context) {
	if err := bac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bac *BankAccountCreate) defaults() {
	if _, ok := bac.mutation.CreateTime(); !ok {
		v := bankaccount.DefaultCreateTime()
		bac.mutation.SetCreateTime(v)
	}
	if _, ok := bac.mutation.UpdateTime(); !ok {
		v := bankaccount.DefaultUpdateTime()
		bac.mutation.SetUpdateTime(v)
	}
	if _, ok := bac.mutation.AccountNumber(); !ok {
		v := bankaccount.DefaultAccountNumber()
		bac.mutation.SetAccountNumber(v)
	}
	if _, ok := bac.mutation.IsForPayment(); !ok {
		v := bankaccount.DefaultIsForPayment
		bac.mutation.SetIsForPayment(v)
	}
	if _, ok := bac.mutation.ID(); !ok {
		v := bankaccount.DefaultID()
		bac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bac *BankAccountCreate) check() error {
	if _, ok := bac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "BankAccount.create_time"`)}
	}
	if _, ok := bac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "BankAccount.update_time"`)}
	}
	if _, ok := bac.mutation.CustomerID(); !ok {
		return &ValidationError{Name: "customer_id", err: errors.New(`ent: missing required field "BankAccount.customer_id"`)}
	}
	if _, ok := bac.mutation.CashIn(); !ok {
		return &ValidationError{Name: "cash_in", err: errors.New(`ent: missing required field "BankAccount.cash_in"`)}
	}
	if v, ok := bac.mutation.CashIn(); ok {
		if err := bankaccount.CashInValidator(v); err != nil {
			return &ValidationError{Name: "cash_in", err: fmt.Errorf(`ent: validator failed for field "BankAccount.cash_in": %w`, err)}
		}
	}
	if _, ok := bac.mutation.CashOut(); !ok {
		return &ValidationError{Name: "cash_out", err: errors.New(`ent: missing required field "BankAccount.cash_out"`)}
	}
	if v, ok := bac.mutation.CashOut(); ok {
		if err := bankaccount.CashOutValidator(v); err != nil {
			return &ValidationError{Name: "cash_out", err: fmt.Errorf(`ent: validator failed for field "BankAccount.cash_out": %w`, err)}
		}
	}
	if _, ok := bac.mutation.AccountNumber(); !ok {
		return &ValidationError{Name: "account_number", err: errors.New(`ent: missing required field "BankAccount.account_number"`)}
	}
	if v, ok := bac.mutation.AccountNumber(); ok {
		if err := bankaccount.AccountNumberValidator(v); err != nil {
			return &ValidationError{Name: "account_number", err: fmt.Errorf(`ent: validator failed for field "BankAccount.account_number": %w`, err)}
		}
	}
	if _, ok := bac.mutation.IsForPayment(); !ok {
		return &ValidationError{Name: "is_for_payment", err: errors.New(`ent: missing required field "BankAccount.is_for_payment"`)}
	}
	if _, ok := bac.mutation.CustomerID(); !ok {
		return &ValidationError{Name: "customer", err: errors.New(`ent: missing required edge "BankAccount.customer"`)}
	}
	return nil
}

func (bac *BankAccountCreate) sqlSave(ctx context.Context) (*BankAccount, error) {
	_node, _spec := bac.createSpec()
	if err := sqlgraph.CreateNode(ctx, bac.driver, _spec); err != nil {
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

func (bac *BankAccountCreate) createSpec() (*BankAccount, *sqlgraph.CreateSpec) {
	var (
		_node = &BankAccount{config: bac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bankaccount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: bankaccount.FieldID,
			},
		}
	)
	if id, ok := bac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bac.mutation.CreateTime(); ok {
		_spec.SetField(bankaccount.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := bac.mutation.UpdateTime(); ok {
		_spec.SetField(bankaccount.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := bac.mutation.CashIn(); ok {
		_spec.SetField(bankaccount.FieldCashIn, field.TypeFloat64, value)
		_node.CashIn = value
	}
	if value, ok := bac.mutation.CashOut(); ok {
		_spec.SetField(bankaccount.FieldCashOut, field.TypeFloat64, value)
		_node.CashOut = value
	}
	if value, ok := bac.mutation.AccountNumber(); ok {
		_spec.SetField(bankaccount.FieldAccountNumber, field.TypeString, value)
		_node.AccountNumber = value
	}
	if value, ok := bac.mutation.IsForPayment(); ok {
		_spec.SetField(bankaccount.FieldIsForPayment, field.TypeBool, value)
		_node.IsForPayment = value
	}
	if nodes := bac.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   bankaccount.CustomerTable,
			Columns: []string{bankaccount.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CustomerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bac.mutation.SentTransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   bankaccount.SentTransactionTable,
			Columns: []string{bankaccount.SentTransactionColumn},
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
	if nodes := bac.mutation.ReceivedTransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   bankaccount.ReceivedTransactionTable,
			Columns: []string{bankaccount.ReceivedTransactionColumn},
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
	if nodes := bac.mutation.OwnedDebtsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   bankaccount.OwnedDebtsTable,
			Columns: []string{bankaccount.OwnedDebtsColumn},
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
	if nodes := bac.mutation.ReceivedDebtsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   bankaccount.ReceivedDebtsTable,
			Columns: []string{bankaccount.ReceivedDebtsColumn},
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

// BankAccountCreateBulk is the builder for creating many BankAccount entities in bulk.
type BankAccountCreateBulk struct {
	config
	builders []*BankAccountCreate
}

// Save creates the BankAccount entities in the database.
func (bacb *BankAccountCreateBulk) Save(ctx context.Context) ([]*BankAccount, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bacb.builders))
	nodes := make([]*BankAccount, len(bacb.builders))
	mutators := make([]Mutator, len(bacb.builders))
	for i := range bacb.builders {
		func(i int, root context.Context) {
			builder := bacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BankAccountMutation)
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
					_, err = mutators[i+1].Mutate(root, bacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bacb *BankAccountCreateBulk) SaveX(ctx context.Context) []*BankAccount {
	v, err := bacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bacb *BankAccountCreateBulk) Exec(ctx context.Context) error {
	_, err := bacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bacb *BankAccountCreateBulk) ExecX(ctx context.Context) {
	if err := bacb.Exec(ctx); err != nil {
		panic(err)
	}
}
