// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/TcMits/wnc-final/ent/employee"
	"github.com/TcMits/wnc-final/ent/predicate"
	"github.com/google/uuid"
)

// EmployeeQuery is the builder for querying Employee entities.
type EmployeeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Employee
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmployeeQuery builder.
func (eq *EmployeeQuery) Where(ps ...predicate.Employee) *EmployeeQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit adds a limit step to the query.
func (eq *EmployeeQuery) Limit(limit int) *EmployeeQuery {
	eq.limit = &limit
	return eq
}

// Offset adds an offset step to the query.
func (eq *EmployeeQuery) Offset(offset int) *EmployeeQuery {
	eq.offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *EmployeeQuery) Unique(unique bool) *EmployeeQuery {
	eq.unique = &unique
	return eq
}

// Order adds an order step to the query.
func (eq *EmployeeQuery) Order(o ...OrderFunc) *EmployeeQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// First returns the first Employee entity from the query.
// Returns a *NotFoundError when no Employee was found.
func (eq *EmployeeQuery) First(ctx context.Context) (*Employee, error) {
	nodes, err := eq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{employee.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EmployeeQuery) FirstX(ctx context.Context) *Employee {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Employee ID from the query.
// Returns a *NotFoundError when no Employee ID was found.
func (eq *EmployeeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = eq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{employee.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EmployeeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Employee entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Employee entity is found.
// Returns a *NotFoundError when no Employee entities are found.
func (eq *EmployeeQuery) Only(ctx context.Context) (*Employee, error) {
	nodes, err := eq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{employee.Label}
	default:
		return nil, &NotSingularError{employee.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EmployeeQuery) OnlyX(ctx context.Context) *Employee {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Employee ID in the query.
// Returns a *NotSingularError when more than one Employee ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *EmployeeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = eq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{employee.Label}
	default:
		err = &NotSingularError{employee.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EmployeeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Employees.
func (eq *EmployeeQuery) All(ctx context.Context) ([]*Employee, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return eq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (eq *EmployeeQuery) AllX(ctx context.Context) []*Employee {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Employee IDs.
func (eq *EmployeeQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := eq.Select(employee.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EmployeeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EmployeeQuery) Count(ctx context.Context) (int, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return eq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EmployeeQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EmployeeQuery) Exist(ctx context.Context) (bool, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return eq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EmployeeQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmployeeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EmployeeQuery) Clone() *EmployeeQuery {
	if eq == nil {
		return nil
	}
	return &EmployeeQuery{
		config:     eq.config,
		limit:      eq.limit,
		offset:     eq.offset,
		order:      append([]OrderFunc{}, eq.order...),
		predicates: append([]predicate.Employee{}, eq.predicates...),
		// clone intermediate query.
		sql:    eq.sql.Clone(),
		path:   eq.path,
		unique: eq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Employee.Query().
//		GroupBy(employee.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eq *EmployeeQuery) GroupBy(field string, fields ...string) *EmployeeGroupBy {
	grbuild := &EmployeeGroupBy{config: eq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eq.sqlQuery(ctx), nil
	}
	grbuild.label = employee.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Employee.Query().
//		Select(employee.FieldCreateTime).
//		Scan(ctx, &v)
func (eq *EmployeeQuery) Select(fields ...string) *EmployeeSelect {
	eq.fields = append(eq.fields, fields...)
	selbuild := &EmployeeSelect{EmployeeQuery: eq}
	selbuild.label = employee.Label
	selbuild.flds, selbuild.scan = &eq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a EmployeeSelect configured with the given aggregations.
func (eq *EmployeeQuery) Aggregate(fns ...AggregateFunc) *EmployeeSelect {
	return eq.Select().Aggregate(fns...)
}

func (eq *EmployeeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range eq.fields {
		if !employee.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EmployeeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Employee, error) {
	var (
		nodes = []*Employee{}
		_spec = eq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Employee).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Employee{config: eq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(eq.modifiers) > 0 {
		_spec.Modifiers = eq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (eq *EmployeeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	if len(eq.modifiers) > 0 {
		_spec.Modifiers = eq.modifiers
	}
	_spec.Node.Columns = eq.fields
	if len(eq.fields) > 0 {
		_spec.Unique = eq.unique != nil && *eq.unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EmployeeQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (eq *EmployeeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   employee.Table,
			Columns: employee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: employee.FieldID,
			},
		},
		From:   eq.sql,
		Unique: true,
	}
	if unique := eq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := eq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, employee.FieldID)
		for i := range fields {
			if fields[i] != employee.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EmployeeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(employee.Table)
	columns := eq.fields
	if len(columns) == 0 {
		columns = employee.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.unique != nil && *eq.unique {
		selector.Distinct()
	}
	for _, m := range eq.modifiers {
		m(selector)
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (eq *EmployeeQuery) ForUpdate(opts ...sql.LockOption) *EmployeeQuery {
	if eq.driver.Dialect() == dialect.Postgres {
		eq.Unique(false)
	}
	eq.modifiers = append(eq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return eq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (eq *EmployeeQuery) ForShare(opts ...sql.LockOption) *EmployeeQuery {
	if eq.driver.Dialect() == dialect.Postgres {
		eq.Unique(false)
	}
	eq.modifiers = append(eq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return eq
}

// EmployeeGroupBy is the group-by builder for Employee entities.
type EmployeeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EmployeeGroupBy) Aggregate(fns ...AggregateFunc) *EmployeeGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the group-by query and scans the result into the given value.
func (egb *EmployeeGroupBy) Scan(ctx context.Context, v any) error {
	query, err := egb.path(ctx)
	if err != nil {
		return err
	}
	egb.sql = query
	return egb.sqlScan(ctx, v)
}

func (egb *EmployeeGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range egb.fields {
		if !employee.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := egb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (egb *EmployeeGroupBy) sqlQuery() *sql.Selector {
	selector := egb.sql.Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(egb.fields)+len(egb.fns))
		for _, f := range egb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(egb.fields...)...)
}

// EmployeeSelect is the builder for selecting fields of Employee entities.
type EmployeeSelect struct {
	*EmployeeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (es *EmployeeSelect) Aggregate(fns ...AggregateFunc) *EmployeeSelect {
	es.fns = append(es.fns, fns...)
	return es
}

// Scan applies the selector query and scans the result into the given value.
func (es *EmployeeSelect) Scan(ctx context.Context, v any) error {
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	es.sql = es.EmployeeQuery.sqlQuery(ctx)
	return es.sqlScan(ctx, v)
}

func (es *EmployeeSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(es.fns))
	for _, fn := range es.fns {
		aggregation = append(aggregation, fn(es.sql))
	}
	switch n := len(*es.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		es.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		es.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := es.sql.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}