// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/TcMits/wnc-final/ent/bankaccount"
	"github.com/TcMits/wnc-final/ent/predicate"
)

// BankAccountDelete is the builder for deleting a BankAccount entity.
type BankAccountDelete struct {
	config
	hooks    []Hook
	mutation *BankAccountMutation
}

// Where appends a list predicates to the BankAccountDelete builder.
func (bad *BankAccountDelete) Where(ps ...predicate.BankAccount) *BankAccountDelete {
	bad.mutation.Where(ps...)
	return bad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bad *BankAccountDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bad.hooks) == 0 {
		affected, err = bad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BankAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bad.mutation = mutation
			affected, err = bad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bad.hooks) - 1; i >= 0; i-- {
			if bad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bad *BankAccountDelete) ExecX(ctx context.Context) int {
	n, err := bad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bad *BankAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: bankaccount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: bankaccount.FieldID,
			},
		},
	}
	if ps := bad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// BankAccountDeleteOne is the builder for deleting a single BankAccount entity.
type BankAccountDeleteOne struct {
	bad *BankAccountDelete
}

// Exec executes the deletion query.
func (bado *BankAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := bado.bad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bankaccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bado *BankAccountDeleteOne) ExecX(ctx context.Context) {
	bado.bad.ExecX(ctx)
}
