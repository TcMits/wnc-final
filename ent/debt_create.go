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

// DebtCreate is the builder for creating a Debt entity.
type DebtCreate struct {
	config
	mutation *DebtMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (dc *DebtCreate) SetCreateTime(t time.Time) *DebtCreate {
	dc.mutation.SetCreateTime(t)
	return dc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (dc *DebtCreate) SetNillableCreateTime(t *time.Time) *DebtCreate {
	if t != nil {
		dc.SetCreateTime(*t)
	}
	return dc
}

// SetUpdateTime sets the "update_time" field.
func (dc *DebtCreate) SetUpdateTime(t time.Time) *DebtCreate {
	dc.mutation.SetUpdateTime(t)
	return dc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (dc *DebtCreate) SetNillableUpdateTime(t *time.Time) *DebtCreate {
	if t != nil {
		dc.SetUpdateTime(*t)
	}
	return dc
}

// SetOwnerBankAccountNumber sets the "owner_bank_account_number" field.
func (dc *DebtCreate) SetOwnerBankAccountNumber(s string) *DebtCreate {
	dc.mutation.SetOwnerBankAccountNumber(s)
	return dc
}

// SetOwnerBankName sets the "owner_bank_name" field.
func (dc *DebtCreate) SetOwnerBankName(s string) *DebtCreate {
	dc.mutation.SetOwnerBankName(s)
	return dc
}

// SetOwnerName sets the "owner_name" field.
func (dc *DebtCreate) SetOwnerName(s string) *DebtCreate {
	dc.mutation.SetOwnerName(s)
	return dc
}

// SetOwnerID sets the "owner_id" field.
func (dc *DebtCreate) SetOwnerID(u uuid.UUID) *DebtCreate {
	dc.mutation.SetOwnerID(u)
	return dc
}

// SetReceiverBankAccountNumber sets the "receiver_bank_account_number" field.
func (dc *DebtCreate) SetReceiverBankAccountNumber(s string) *DebtCreate {
	dc.mutation.SetReceiverBankAccountNumber(s)
	return dc
}

// SetReceiverBankName sets the "receiver_bank_name" field.
func (dc *DebtCreate) SetReceiverBankName(s string) *DebtCreate {
	dc.mutation.SetReceiverBankName(s)
	return dc
}

// SetReceiverName sets the "receiver_name" field.
func (dc *DebtCreate) SetReceiverName(s string) *DebtCreate {
	dc.mutation.SetReceiverName(s)
	return dc
}

// SetReceiverID sets the "receiver_id" field.
func (dc *DebtCreate) SetReceiverID(u uuid.UUID) *DebtCreate {
	dc.mutation.SetReceiverID(u)
	return dc
}

// SetTransactionID sets the "transaction_id" field.
func (dc *DebtCreate) SetTransactionID(u uuid.UUID) *DebtCreate {
	dc.mutation.SetTransactionID(u)
	return dc
}

// SetNillableTransactionID sets the "transaction_id" field if the given value is not nil.
func (dc *DebtCreate) SetNillableTransactionID(u *uuid.UUID) *DebtCreate {
	if u != nil {
		dc.SetTransactionID(*u)
	}
	return dc
}

// SetStatus sets the "status" field.
func (dc *DebtCreate) SetStatus(d debt.Status) *DebtCreate {
	dc.mutation.SetStatus(d)
	return dc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dc *DebtCreate) SetNillableStatus(d *debt.Status) *DebtCreate {
	if d != nil {
		dc.SetStatus(*d)
	}
	return dc
}

// SetDescription sets the "description" field.
func (dc *DebtCreate) SetDescription(s string) *DebtCreate {
	dc.mutation.SetDescription(s)
	return dc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dc *DebtCreate) SetNillableDescription(s *string) *DebtCreate {
	if s != nil {
		dc.SetDescription(*s)
	}
	return dc
}

// SetAmount sets the "amount" field.
func (dc *DebtCreate) SetAmount(d decimal.Decimal) *DebtCreate {
	dc.mutation.SetAmount(d)
	return dc
}

// SetID sets the "id" field.
func (dc *DebtCreate) SetID(u uuid.UUID) *DebtCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DebtCreate) SetNillableID(u *uuid.UUID) *DebtCreate {
	if u != nil {
		dc.SetID(*u)
	}
	return dc
}

// SetOwner sets the "owner" edge to the BankAccount entity.
func (dc *DebtCreate) SetOwner(b *BankAccount) *DebtCreate {
	return dc.SetOwnerID(b.ID)
}

// SetReceiver sets the "receiver" edge to the BankAccount entity.
func (dc *DebtCreate) SetReceiver(b *BankAccount) *DebtCreate {
	return dc.SetReceiverID(b.ID)
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (dc *DebtCreate) SetTransaction(t *Transaction) *DebtCreate {
	return dc.SetTransactionID(t.ID)
}

// Mutation returns the DebtMutation object of the builder.
func (dc *DebtCreate) Mutation() *DebtMutation {
	return dc.mutation
}

// Save creates the Debt in the database.
func (dc *DebtCreate) Save(ctx context.Context) (*Debt, error) {
	var (
		err  error
		node *Debt
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DebtMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Debt)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DebtMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DebtCreate) SaveX(ctx context.Context) *Debt {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DebtCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DebtCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DebtCreate) defaults() {
	if _, ok := dc.mutation.CreateTime(); !ok {
		v := debt.DefaultCreateTime()
		dc.mutation.SetCreateTime(v)
	}
	if _, ok := dc.mutation.UpdateTime(); !ok {
		v := debt.DefaultUpdateTime()
		dc.mutation.SetUpdateTime(v)
	}
	if _, ok := dc.mutation.Status(); !ok {
		v := debt.DefaultStatus
		dc.mutation.SetStatus(v)
	}
	if _, ok := dc.mutation.Description(); !ok {
		v := debt.DefaultDescription
		dc.mutation.SetDescription(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := debt.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DebtCreate) check() error {
	if _, ok := dc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Debt.create_time"`)}
	}
	if _, ok := dc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Debt.update_time"`)}
	}
	if _, ok := dc.mutation.OwnerBankAccountNumber(); !ok {
		return &ValidationError{Name: "owner_bank_account_number", err: errors.New(`ent: missing required field "Debt.owner_bank_account_number"`)}
	}
	if v, ok := dc.mutation.OwnerBankAccountNumber(); ok {
		if err := debt.OwnerBankAccountNumberValidator(v); err != nil {
			return &ValidationError{Name: "owner_bank_account_number", err: fmt.Errorf(`ent: validator failed for field "Debt.owner_bank_account_number": %w`, err)}
		}
	}
	if _, ok := dc.mutation.OwnerBankName(); !ok {
		return &ValidationError{Name: "owner_bank_name", err: errors.New(`ent: missing required field "Debt.owner_bank_name"`)}
	}
	if v, ok := dc.mutation.OwnerBankName(); ok {
		if err := debt.OwnerBankNameValidator(v); err != nil {
			return &ValidationError{Name: "owner_bank_name", err: fmt.Errorf(`ent: validator failed for field "Debt.owner_bank_name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.OwnerName(); !ok {
		return &ValidationError{Name: "owner_name", err: errors.New(`ent: missing required field "Debt.owner_name"`)}
	}
	if v, ok := dc.mutation.OwnerName(); ok {
		if err := debt.OwnerNameValidator(v); err != nil {
			return &ValidationError{Name: "owner_name", err: fmt.Errorf(`ent: validator failed for field "Debt.owner_name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "Debt.owner_id"`)}
	}
	if _, ok := dc.mutation.ReceiverBankAccountNumber(); !ok {
		return &ValidationError{Name: "receiver_bank_account_number", err: errors.New(`ent: missing required field "Debt.receiver_bank_account_number"`)}
	}
	if v, ok := dc.mutation.ReceiverBankAccountNumber(); ok {
		if err := debt.ReceiverBankAccountNumberValidator(v); err != nil {
			return &ValidationError{Name: "receiver_bank_account_number", err: fmt.Errorf(`ent: validator failed for field "Debt.receiver_bank_account_number": %w`, err)}
		}
	}
	if _, ok := dc.mutation.ReceiverBankName(); !ok {
		return &ValidationError{Name: "receiver_bank_name", err: errors.New(`ent: missing required field "Debt.receiver_bank_name"`)}
	}
	if v, ok := dc.mutation.ReceiverBankName(); ok {
		if err := debt.ReceiverBankNameValidator(v); err != nil {
			return &ValidationError{Name: "receiver_bank_name", err: fmt.Errorf(`ent: validator failed for field "Debt.receiver_bank_name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.ReceiverName(); !ok {
		return &ValidationError{Name: "receiver_name", err: errors.New(`ent: missing required field "Debt.receiver_name"`)}
	}
	if v, ok := dc.mutation.ReceiverName(); ok {
		if err := debt.ReceiverNameValidator(v); err != nil {
			return &ValidationError{Name: "receiver_name", err: fmt.Errorf(`ent: validator failed for field "Debt.receiver_name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.ReceiverID(); !ok {
		return &ValidationError{Name: "receiver_id", err: errors.New(`ent: missing required field "Debt.receiver_id"`)}
	}
	if _, ok := dc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Debt.status"`)}
	}
	if v, ok := dc.mutation.Status(); ok {
		if err := debt.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Debt.status": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Debt.amount"`)}
	}
	if _, ok := dc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Debt.owner"`)}
	}
	if _, ok := dc.mutation.ReceiverID(); !ok {
		return &ValidationError{Name: "receiver", err: errors.New(`ent: missing required edge "Debt.receiver"`)}
	}
	return nil
}

func (dc *DebtCreate) sqlSave(ctx context.Context) (*Debt, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
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

func (dc *DebtCreate) createSpec() (*Debt, *sqlgraph.CreateSpec) {
	var (
		_node = &Debt{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: debt.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: debt.FieldID,
			},
		}
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.CreateTime(); ok {
		_spec.SetField(debt.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := dc.mutation.UpdateTime(); ok {
		_spec.SetField(debt.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := dc.mutation.OwnerBankAccountNumber(); ok {
		_spec.SetField(debt.FieldOwnerBankAccountNumber, field.TypeString, value)
		_node.OwnerBankAccountNumber = value
	}
	if value, ok := dc.mutation.OwnerBankName(); ok {
		_spec.SetField(debt.FieldOwnerBankName, field.TypeString, value)
		_node.OwnerBankName = value
	}
	if value, ok := dc.mutation.OwnerName(); ok {
		_spec.SetField(debt.FieldOwnerName, field.TypeString, value)
		_node.OwnerName = value
	}
	if value, ok := dc.mutation.ReceiverBankAccountNumber(); ok {
		_spec.SetField(debt.FieldReceiverBankAccountNumber, field.TypeString, value)
		_node.ReceiverBankAccountNumber = value
	}
	if value, ok := dc.mutation.ReceiverBankName(); ok {
		_spec.SetField(debt.FieldReceiverBankName, field.TypeString, value)
		_node.ReceiverBankName = value
	}
	if value, ok := dc.mutation.ReceiverName(); ok {
		_spec.SetField(debt.FieldReceiverName, field.TypeString, value)
		_node.ReceiverName = value
	}
	if value, ok := dc.mutation.Status(); ok {
		_spec.SetField(debt.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := dc.mutation.Description(); ok {
		_spec.SetField(debt.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := dc.mutation.Amount(); ok {
		_spec.SetField(debt.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if nodes := dc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   debt.OwnerTable,
			Columns: []string{debt.OwnerColumn},
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
		_node.OwnerID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ReceiverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   debt.ReceiverTable,
			Columns: []string{debt.ReceiverColumn},
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
	if nodes := dc.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   debt.TransactionTable,
			Columns: []string{debt.TransactionColumn},
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
		_node.TransactionID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DebtCreateBulk is the builder for creating many Debt entities in bulk.
type DebtCreateBulk struct {
	config
	builders []*DebtCreate
}

// Save creates the Debt entities in the database.
func (dcb *DebtCreateBulk) Save(ctx context.Context) ([]*Debt, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Debt, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DebtMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DebtCreateBulk) SaveX(ctx context.Context) []*Debt {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DebtCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DebtCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}