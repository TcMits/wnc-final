// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/TcMits/wnc-final/ent/contact"
	"github.com/TcMits/wnc-final/ent/customer"
	"github.com/TcMits/wnc-final/ent/predicate"
	"github.com/google/uuid"
)

// Contact is the model entity for the Contact schema.
type Contact struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID uuid.UUID `json:"owner_id,omitempty"`
	// AccountNumber holds the value of the "account_number" field.
	AccountNumber string `json:"account_number,omitempty"`
	// SuggestName holds the value of the "suggest_name" field.
	SuggestName string `json:"suggest_name,omitempty"`
	// BankName holds the value of the "bank_name" field.
	BankName string `json:"bank_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContactQuery when eager-loading is set.
	Edges ContactEdges `json:"edges"`
}

// ContactEdges holds the relations/edges for other nodes in the graph.
type ContactEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Customer `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContactEdges) OwnerOrErr() (*Customer, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contact) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case contact.FieldAccountNumber, contact.FieldSuggestName, contact.FieldBankName:
			values[i] = new(sql.NullString)
		case contact.FieldCreateTime, contact.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case contact.FieldID, contact.FieldOwnerID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Contact", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contact fields.
func (c *Contact) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contact.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case contact.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case contact.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = value.Time
			}
		case contact.FieldOwnerID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value != nil {
				c.OwnerID = *value
			}
		case contact.FieldAccountNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_number", values[i])
			} else if value.Valid {
				c.AccountNumber = value.String
			}
		case contact.FieldSuggestName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field suggest_name", values[i])
			} else if value.Valid {
				c.SuggestName = value.String
			}
		case contact.FieldBankName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bank_name", values[i])
			} else if value.Valid {
				c.BankName = value.String
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Contact entity.
func (c *Contact) QueryOwner() *CustomerQuery {
	return (&ContactClient{config: c.config}).QueryOwner(c)
}

// Update returns a builder for updating this Contact.
// Note that you need to call Contact.Unwrap() before calling this method if this Contact
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contact) Update() *ContactUpdateOne {
	return (&ContactClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Contact entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Contact) Unwrap() *Contact {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Contact is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contact) String() string {
	var builder strings.Builder
	builder.WriteString("Contact(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(fmt.Sprintf("%v", c.OwnerID))
	builder.WriteString(", ")
	builder.WriteString("account_number=")
	builder.WriteString(c.AccountNumber)
	builder.WriteString(", ")
	builder.WriteString("suggest_name=")
	builder.WriteString(c.SuggestName)
	builder.WriteString(", ")
	builder.WriteString("bank_name=")
	builder.WriteString(c.BankName)
	builder.WriteByte(')')
	return builder.String()
}

type ContactCreateRepository struct {
	client   *Client
	isAtomic bool
}

func NewContactCreateRepository(
	client *Client,
	isAtomic bool,
) *ContactCreateRepository {
	return &ContactCreateRepository{
		client:   client,
		isAtomic: isAtomic,
	}
}

// using in Tx
func (r *ContactCreateRepository) CreateWithClient(
	ctx context.Context, client *Client, input *ContactCreateInput,
) (*Contact, error) {
	instance, err := client.Contact.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (r *ContactCreateRepository) Create(
	ctx context.Context, input *ContactCreateInput,
) (*Contact, error) {
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

type ContactDeleteRepository struct {
	client   *Client
	isAtomic bool
}

func NewContactDeleteRepository(
	client *Client,
	isAtomic bool,
) *ContactDeleteRepository {
	return &ContactDeleteRepository{
		client:   client,
		isAtomic: isAtomic,
	}
}

// using in Tx
func (r *ContactDeleteRepository) DeleteWithClient(
	ctx context.Context, client *Client, instance *Contact,
) error {
	err := client.Contact.DeleteOne(instance).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContactDeleteRepository) Delete(
	ctx context.Context, instance *Contact,
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

// ContactCreateInput represents a mutation input for creating contacts.
type ContactCreateInput struct {
	CreateTime    *time.Time `json:"create_time,omitempty" form:"create_time"`
	UpdateTime    *time.Time `json:"update_time,omitempty" form:"update_time"`
	AccountNumber string     `json:"account_number,omitempty" form:"account_number"`
	SuggestName   string     `json:"suggest_name,omitempty" form:"suggest_name"`
	BankName      string     `json:"bank_name,omitempty" form:"bank_name"`
	OwnerID       uuid.UUID  `json:"owner_id,omitempty" form:"owner_id"`
}

// Mutate applies the ContactCreateInput on the ContactCreate builder.
func (i *ContactCreateInput) Mutate(m *ContactMutation) {
	if v := i.CreateTime; v != nil {
		m.SetCreateTime(*v)
	}
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	m.SetAccountNumber(i.AccountNumber)
	m.SetSuggestName(i.SuggestName)
	m.SetBankName(i.BankName)
	m.SetOwnerID(i.OwnerID)
}

// SetInput applies the change-set in the ContactCreateInput on the create builder.
func (c *ContactCreate) SetInput(i *ContactCreateInput) *ContactCreate {
	i.Mutate(c.Mutation())
	return c
}

// ContactUpdateInput represents a mutation input for updating contacts.
type ContactUpdateInput struct {
	ID            uuid.UUID
	UpdateTime    *time.Time `json:"update_time,omitempty" form:"update_time"`
	AccountNumber *string    `json:"account_number,omitempty" form:"account_number"`
	SuggestName   *string    `json:"suggest_name,omitempty" form:"suggest_name"`
	BankName      *string    `json:"bank_name,omitempty" form:"bank_name"`
	OwnerID       *uuid.UUID `json:"owner_id,omitempty" form:"owner_id"`
	ClearOwner    bool
}

// Mutate applies the ContactUpdateInput on the ContactMutation.
func (i *ContactUpdateInput) Mutate(m *ContactMutation) {
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	if v := i.AccountNumber; v != nil {
		m.SetAccountNumber(*v)
	}
	if v := i.SuggestName; v != nil {
		m.SetSuggestName(*v)
	}
	if v := i.BankName; v != nil {
		m.SetBankName(*v)
	}
	if i.ClearOwner {
		m.ClearOwner()
	}
	if v := i.OwnerID; v != nil {
		m.SetOwnerID(*v)
	}
}

// SetInput applies the change-set in the ContactUpdateInput on the update builder.
func (u *ContactUpdate) SetInput(i *ContactUpdateInput) *ContactUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the ContactUpdateInput on the update-one builder.
func (u *ContactUpdateOne) SetInput(i *ContactUpdateInput) *ContactUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

type ContactReadRepository struct {
	client *Client
}

func NewContactReadRepository(
	client *Client,
) *ContactReadRepository {
	return &ContactReadRepository{
		client: client,
	}
}

func (r *ContactReadRepository) prepareQuery(
	client *Client, limit *int, offset *int, o *ContactOrderInput, w *ContactWhereInput,
) (*ContactQuery, error) {
	var err error
	q := r.client.Contact.Query()
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
func (r *ContactReadRepository) GetWithClient(
	ctx context.Context, client *Client, w *ContactWhereInput, forUpdate bool,
) (*Contact, error) {
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
func (r *ContactReadRepository) ListWithClient(
	ctx context.Context, client *Client, limit *int, offset *int, o *ContactOrderInput, w *ContactWhereInput, forUpdate bool,
) ([]*Contact, error) {
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

func (r *ContactReadRepository) Count(ctx context.Context, w *ContactWhereInput) (int, error) {
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

func (r *ContactReadRepository) Get(ctx context.Context, w *ContactWhereInput) (*Contact, error) {
	return r.GetWithClient(ctx, r.client, w, false)
}

func (r *ContactReadRepository) List(
	ctx context.Context, limit *int, offset *int, o *ContactOrderInput, w *ContactWhereInput,
) ([]*Contact, error) {
	return r.ListWithClient(ctx, r.client, limit, offset, o, w, false)
}

type ContactSerializer struct {
	columns map[string]func(context.Context, *Contact) any
}

func NewContactSerializer(customColumns map[string]func(context.Context, *Contact) any, columns ...string) *ContactSerializer {
	columnsMap := map[string]func(context.Context, *Contact) any{}
	for _, col := range columns {
		switch col {

		case contact.FieldID:
			columnsMap[col] = func(ctx context.Context, c *Contact) any {
				return c.ID
			}

		case contact.FieldCreateTime:
			columnsMap[col] = func(ctx context.Context, c *Contact) any {
				return c.CreateTime
			}

		case contact.FieldUpdateTime:
			columnsMap[col] = func(ctx context.Context, c *Contact) any {
				return c.UpdateTime
			}

		case contact.FieldOwnerID:
			columnsMap[col] = func(ctx context.Context, c *Contact) any {
				return c.OwnerID
			}

		case contact.FieldAccountNumber:
			columnsMap[col] = func(ctx context.Context, c *Contact) any {
				return c.AccountNumber
			}

		case contact.FieldSuggestName:
			columnsMap[col] = func(ctx context.Context, c *Contact) any {
				return c.SuggestName
			}

		case contact.FieldBankName:
			columnsMap[col] = func(ctx context.Context, c *Contact) any {
				return c.BankName
			}

		default:
			panic(fmt.Sprintf("Unexpect column %s", col))
		}
	}

	for k, serializeFunc := range customColumns {
		columnsMap[k] = serializeFunc
	}

	return &ContactSerializer{
		columns: columnsMap,
	}
}

func (s *ContactSerializer) Serialize(ctx context.Context, c *Contact) map[string]any {
	result := make(map[string]any, len(s.columns))
	for col, serializeFunc := range s.columns {
		result[col] = serializeFunc(ctx, c)
	}
	return result
}

type ContactUpdateRepository struct {
	client   *Client
	isAtomic bool
}

func NewContactUpdateRepository(
	client *Client,
	isAtomic bool,
) *ContactUpdateRepository {
	return &ContactUpdateRepository{
		client:   client,
		isAtomic: isAtomic,
	}
}

// using in Tx
func (r *ContactUpdateRepository) UpdateWithClient(
	ctx context.Context, client *Client, instance *Contact, input *ContactUpdateInput,
) (*Contact, error) {
	newInstance, err := client.Contact.UpdateOne(instance).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	return newInstance, nil
}

func (r *ContactUpdateRepository) Update(
	ctx context.Context, instance *Contact, input *ContactUpdateInput,
) (*Contact, error) {
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

// ContactWhereInput represents a where input for filtering Contact queries.
type ContactWhereInput struct {
	Predicates []predicate.Contact  `json:"-"`
	Not        *ContactWhereInput   `json:"not,omitempty"`
	Or         []*ContactWhereInput `json:"or,omitempty"`
	And        []*ContactWhereInput `json:"and,omitempty"`

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

	// "owner_id" field predicates.
	OwnerID      *uuid.UUID  `json:"owner_id,omitempty" form:"owner_id" param:"owner_id" url:"owner_id"`
	OwnerIDNEQ   *uuid.UUID  `json:"owner_id_neq,omitempty" form:"owner_id_neq" param:"owner_id_neq" url:"owner_id_neq"`
	OwnerIDIn    []uuid.UUID `json:"owner_id_in,omitempty" form:"owner_id_in" param:"owner_id_in" url:"owner_id_in"`
	OwnerIDNotIn []uuid.UUID `json:"owner_id_not_in,omitempty" form:"owner_id_not_in" param:"owner_id_not_in" url:"owner_id_not_in"`

	// "account_number" field predicates.
	AccountNumber             *string  `json:"account_number,omitempty" form:"account_number" param:"account_number" url:"account_number"`
	AccountNumberNEQ          *string  `json:"account_number_neq,omitempty" form:"account_number_neq" param:"account_number_neq" url:"account_number_neq"`
	AccountNumberIn           []string `json:"account_number_in,omitempty" form:"account_number_in" param:"account_number_in" url:"account_number_in"`
	AccountNumberNotIn        []string `json:"account_number_not_in,omitempty" form:"account_number_not_in" param:"account_number_not_in" url:"account_number_not_in"`
	AccountNumberGT           *string  `json:"account_number_gt,omitempty" form:"account_number_gt" param:"account_number_gt" url:"account_number_gt"`
	AccountNumberGTE          *string  `json:"account_number_gte,omitempty" form:"account_number_gte" param:"account_number_gte" url:"account_number_gte"`
	AccountNumberLT           *string  `json:"account_number_lt,omitempty" form:"account_number_lt" param:"account_number_lt" url:"account_number_lt"`
	AccountNumberLTE          *string  `json:"account_number_lte,omitempty" form:"account_number_lte" param:"account_number_lte" url:"account_number_lte"`
	AccountNumberContains     *string  `json:"account_number_contains,omitempty" form:"account_number_contains" param:"account_number_contains" url:"account_number_contains"`
	AccountNumberHasPrefix    *string  `json:"account_number_has_prefix,omitempty" form:"account_number_has_prefix" param:"account_number_has_prefix" url:"account_number_has_prefix"`
	AccountNumberHasSuffix    *string  `json:"account_number_has_suffix,omitempty" form:"account_number_has_suffix" param:"account_number_has_suffix" url:"account_number_has_suffix"`
	AccountNumberEqualFold    *string  `json:"account_number_equal_fold,omitempty" form:"account_number_equal_fold" param:"account_number_equal_fold" url:"account_number_equal_fold"`
	AccountNumberContainsFold *string  `json:"account_number_contains_fold,omitempty" form:"account_number_contains_fold" param:"account_number_contains_fold" url:"account_number_contains_fold"`

	// "suggest_name" field predicates.
	SuggestName             *string  `json:"suggest_name,omitempty" form:"suggest_name" param:"suggest_name" url:"suggest_name"`
	SuggestNameNEQ          *string  `json:"suggest_name_neq,omitempty" form:"suggest_name_neq" param:"suggest_name_neq" url:"suggest_name_neq"`
	SuggestNameIn           []string `json:"suggest_name_in,omitempty" form:"suggest_name_in" param:"suggest_name_in" url:"suggest_name_in"`
	SuggestNameNotIn        []string `json:"suggest_name_not_in,omitempty" form:"suggest_name_not_in" param:"suggest_name_not_in" url:"suggest_name_not_in"`
	SuggestNameGT           *string  `json:"suggest_name_gt,omitempty" form:"suggest_name_gt" param:"suggest_name_gt" url:"suggest_name_gt"`
	SuggestNameGTE          *string  `json:"suggest_name_gte,omitempty" form:"suggest_name_gte" param:"suggest_name_gte" url:"suggest_name_gte"`
	SuggestNameLT           *string  `json:"suggest_name_lt,omitempty" form:"suggest_name_lt" param:"suggest_name_lt" url:"suggest_name_lt"`
	SuggestNameLTE          *string  `json:"suggest_name_lte,omitempty" form:"suggest_name_lte" param:"suggest_name_lte" url:"suggest_name_lte"`
	SuggestNameContains     *string  `json:"suggest_name_contains,omitempty" form:"suggest_name_contains" param:"suggest_name_contains" url:"suggest_name_contains"`
	SuggestNameHasPrefix    *string  `json:"suggest_name_has_prefix,omitempty" form:"suggest_name_has_prefix" param:"suggest_name_has_prefix" url:"suggest_name_has_prefix"`
	SuggestNameHasSuffix    *string  `json:"suggest_name_has_suffix,omitempty" form:"suggest_name_has_suffix" param:"suggest_name_has_suffix" url:"suggest_name_has_suffix"`
	SuggestNameEqualFold    *string  `json:"suggest_name_equal_fold,omitempty" form:"suggest_name_equal_fold" param:"suggest_name_equal_fold" url:"suggest_name_equal_fold"`
	SuggestNameContainsFold *string  `json:"suggest_name_contains_fold,omitempty" form:"suggest_name_contains_fold" param:"suggest_name_contains_fold" url:"suggest_name_contains_fold"`

	// "bank_name" field predicates.
	BankName             *string  `json:"bank_name,omitempty" form:"bank_name" param:"bank_name" url:"bank_name"`
	BankNameNEQ          *string  `json:"bank_name_neq,omitempty" form:"bank_name_neq" param:"bank_name_neq" url:"bank_name_neq"`
	BankNameIn           []string `json:"bank_name_in,omitempty" form:"bank_name_in" param:"bank_name_in" url:"bank_name_in"`
	BankNameNotIn        []string `json:"bank_name_not_in,omitempty" form:"bank_name_not_in" param:"bank_name_not_in" url:"bank_name_not_in"`
	BankNameGT           *string  `json:"bank_name_gt,omitempty" form:"bank_name_gt" param:"bank_name_gt" url:"bank_name_gt"`
	BankNameGTE          *string  `json:"bank_name_gte,omitempty" form:"bank_name_gte" param:"bank_name_gte" url:"bank_name_gte"`
	BankNameLT           *string  `json:"bank_name_lt,omitempty" form:"bank_name_lt" param:"bank_name_lt" url:"bank_name_lt"`
	BankNameLTE          *string  `json:"bank_name_lte,omitempty" form:"bank_name_lte" param:"bank_name_lte" url:"bank_name_lte"`
	BankNameContains     *string  `json:"bank_name_contains,omitempty" form:"bank_name_contains" param:"bank_name_contains" url:"bank_name_contains"`
	BankNameHasPrefix    *string  `json:"bank_name_has_prefix,omitempty" form:"bank_name_has_prefix" param:"bank_name_has_prefix" url:"bank_name_has_prefix"`
	BankNameHasSuffix    *string  `json:"bank_name_has_suffix,omitempty" form:"bank_name_has_suffix" param:"bank_name_has_suffix" url:"bank_name_has_suffix"`
	BankNameEqualFold    *string  `json:"bank_name_equal_fold,omitempty" form:"bank_name_equal_fold" param:"bank_name_equal_fold" url:"bank_name_equal_fold"`
	BankNameContainsFold *string  `json:"bank_name_contains_fold,omitempty" form:"bank_name_contains_fold" param:"bank_name_contains_fold" url:"bank_name_contains_fold"`

	// "owner" edge predicates.
	HasOwner     *bool                 `json:"has_owner,omitempty" form:"has_owner" param:"has_owner" url:"has_owner"`
	HasOwnerWith []*CustomerWhereInput `json:"has_owner_with,omitempty" form:"has_owner_with" param:"has_owner_with" url:"has_owner_with"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *ContactWhereInput) AddPredicates(predicates ...predicate.Contact) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the ContactWhereInput filter on the ContactQuery builder.
func (i *ContactWhereInput) Filter(q *ContactQuery) (*ContactQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyContactWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyContactWhereInput is returned in case the ContactWhereInput is empty.
var ErrEmptyContactWhereInput = errors.New("ent: empty predicate ContactWhereInput")

// P returns a predicate for filtering contacts.
// An error is returned if the input is empty or invalid.
func (i *ContactWhereInput) P() (predicate.Contact, error) {
	var predicates []predicate.Contact
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, contact.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Contact, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, contact.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Contact, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, contact.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, contact.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, contact.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, contact.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, contact.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, contact.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, contact.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, contact.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, contact.IDLTE(*i.IDLTE))
	}
	if i.CreateTime != nil {
		predicates = append(predicates, contact.CreateTimeEQ(*i.CreateTime))
	}
	if i.CreateTimeNEQ != nil {
		predicates = append(predicates, contact.CreateTimeNEQ(*i.CreateTimeNEQ))
	}
	if len(i.CreateTimeIn) > 0 {
		predicates = append(predicates, contact.CreateTimeIn(i.CreateTimeIn...))
	}
	if len(i.CreateTimeNotIn) > 0 {
		predicates = append(predicates, contact.CreateTimeNotIn(i.CreateTimeNotIn...))
	}
	if i.CreateTimeGT != nil {
		predicates = append(predicates, contact.CreateTimeGT(*i.CreateTimeGT))
	}
	if i.CreateTimeGTE != nil {
		predicates = append(predicates, contact.CreateTimeGTE(*i.CreateTimeGTE))
	}
	if i.CreateTimeLT != nil {
		predicates = append(predicates, contact.CreateTimeLT(*i.CreateTimeLT))
	}
	if i.CreateTimeLTE != nil {
		predicates = append(predicates, contact.CreateTimeLTE(*i.CreateTimeLTE))
	}
	if i.UpdateTime != nil {
		predicates = append(predicates, contact.UpdateTimeEQ(*i.UpdateTime))
	}
	if i.UpdateTimeNEQ != nil {
		predicates = append(predicates, contact.UpdateTimeNEQ(*i.UpdateTimeNEQ))
	}
	if len(i.UpdateTimeIn) > 0 {
		predicates = append(predicates, contact.UpdateTimeIn(i.UpdateTimeIn...))
	}
	if len(i.UpdateTimeNotIn) > 0 {
		predicates = append(predicates, contact.UpdateTimeNotIn(i.UpdateTimeNotIn...))
	}
	if i.UpdateTimeGT != nil {
		predicates = append(predicates, contact.UpdateTimeGT(*i.UpdateTimeGT))
	}
	if i.UpdateTimeGTE != nil {
		predicates = append(predicates, contact.UpdateTimeGTE(*i.UpdateTimeGTE))
	}
	if i.UpdateTimeLT != nil {
		predicates = append(predicates, contact.UpdateTimeLT(*i.UpdateTimeLT))
	}
	if i.UpdateTimeLTE != nil {
		predicates = append(predicates, contact.UpdateTimeLTE(*i.UpdateTimeLTE))
	}
	if i.OwnerID != nil {
		predicates = append(predicates, contact.OwnerIDEQ(*i.OwnerID))
	}
	if i.OwnerIDNEQ != nil {
		predicates = append(predicates, contact.OwnerIDNEQ(*i.OwnerIDNEQ))
	}
	if len(i.OwnerIDIn) > 0 {
		predicates = append(predicates, contact.OwnerIDIn(i.OwnerIDIn...))
	}
	if len(i.OwnerIDNotIn) > 0 {
		predicates = append(predicates, contact.OwnerIDNotIn(i.OwnerIDNotIn...))
	}
	if i.AccountNumber != nil {
		predicates = append(predicates, contact.AccountNumberEQ(*i.AccountNumber))
	}
	if i.AccountNumberNEQ != nil {
		predicates = append(predicates, contact.AccountNumberNEQ(*i.AccountNumberNEQ))
	}
	if len(i.AccountNumberIn) > 0 {
		predicates = append(predicates, contact.AccountNumberIn(i.AccountNumberIn...))
	}
	if len(i.AccountNumberNotIn) > 0 {
		predicates = append(predicates, contact.AccountNumberNotIn(i.AccountNumberNotIn...))
	}
	if i.AccountNumberGT != nil {
		predicates = append(predicates, contact.AccountNumberGT(*i.AccountNumberGT))
	}
	if i.AccountNumberGTE != nil {
		predicates = append(predicates, contact.AccountNumberGTE(*i.AccountNumberGTE))
	}
	if i.AccountNumberLT != nil {
		predicates = append(predicates, contact.AccountNumberLT(*i.AccountNumberLT))
	}
	if i.AccountNumberLTE != nil {
		predicates = append(predicates, contact.AccountNumberLTE(*i.AccountNumberLTE))
	}
	if i.AccountNumberContains != nil {
		predicates = append(predicates, contact.AccountNumberContains(*i.AccountNumberContains))
	}
	if i.AccountNumberHasPrefix != nil {
		predicates = append(predicates, contact.AccountNumberHasPrefix(*i.AccountNumberHasPrefix))
	}
	if i.AccountNumberHasSuffix != nil {
		predicates = append(predicates, contact.AccountNumberHasSuffix(*i.AccountNumberHasSuffix))
	}
	if i.AccountNumberEqualFold != nil {
		predicates = append(predicates, contact.AccountNumberEqualFold(*i.AccountNumberEqualFold))
	}
	if i.AccountNumberContainsFold != nil {
		predicates = append(predicates, contact.AccountNumberContainsFold(*i.AccountNumberContainsFold))
	}
	if i.SuggestName != nil {
		predicates = append(predicates, contact.SuggestNameEQ(*i.SuggestName))
	}
	if i.SuggestNameNEQ != nil {
		predicates = append(predicates, contact.SuggestNameNEQ(*i.SuggestNameNEQ))
	}
	if len(i.SuggestNameIn) > 0 {
		predicates = append(predicates, contact.SuggestNameIn(i.SuggestNameIn...))
	}
	if len(i.SuggestNameNotIn) > 0 {
		predicates = append(predicates, contact.SuggestNameNotIn(i.SuggestNameNotIn...))
	}
	if i.SuggestNameGT != nil {
		predicates = append(predicates, contact.SuggestNameGT(*i.SuggestNameGT))
	}
	if i.SuggestNameGTE != nil {
		predicates = append(predicates, contact.SuggestNameGTE(*i.SuggestNameGTE))
	}
	if i.SuggestNameLT != nil {
		predicates = append(predicates, contact.SuggestNameLT(*i.SuggestNameLT))
	}
	if i.SuggestNameLTE != nil {
		predicates = append(predicates, contact.SuggestNameLTE(*i.SuggestNameLTE))
	}
	if i.SuggestNameContains != nil {
		predicates = append(predicates, contact.SuggestNameContains(*i.SuggestNameContains))
	}
	if i.SuggestNameHasPrefix != nil {
		predicates = append(predicates, contact.SuggestNameHasPrefix(*i.SuggestNameHasPrefix))
	}
	if i.SuggestNameHasSuffix != nil {
		predicates = append(predicates, contact.SuggestNameHasSuffix(*i.SuggestNameHasSuffix))
	}
	if i.SuggestNameEqualFold != nil {
		predicates = append(predicates, contact.SuggestNameEqualFold(*i.SuggestNameEqualFold))
	}
	if i.SuggestNameContainsFold != nil {
		predicates = append(predicates, contact.SuggestNameContainsFold(*i.SuggestNameContainsFold))
	}
	if i.BankName != nil {
		predicates = append(predicates, contact.BankNameEQ(*i.BankName))
	}
	if i.BankNameNEQ != nil {
		predicates = append(predicates, contact.BankNameNEQ(*i.BankNameNEQ))
	}
	if len(i.BankNameIn) > 0 {
		predicates = append(predicates, contact.BankNameIn(i.BankNameIn...))
	}
	if len(i.BankNameNotIn) > 0 {
		predicates = append(predicates, contact.BankNameNotIn(i.BankNameNotIn...))
	}
	if i.BankNameGT != nil {
		predicates = append(predicates, contact.BankNameGT(*i.BankNameGT))
	}
	if i.BankNameGTE != nil {
		predicates = append(predicates, contact.BankNameGTE(*i.BankNameGTE))
	}
	if i.BankNameLT != nil {
		predicates = append(predicates, contact.BankNameLT(*i.BankNameLT))
	}
	if i.BankNameLTE != nil {
		predicates = append(predicates, contact.BankNameLTE(*i.BankNameLTE))
	}
	if i.BankNameContains != nil {
		predicates = append(predicates, contact.BankNameContains(*i.BankNameContains))
	}
	if i.BankNameHasPrefix != nil {
		predicates = append(predicates, contact.BankNameHasPrefix(*i.BankNameHasPrefix))
	}
	if i.BankNameHasSuffix != nil {
		predicates = append(predicates, contact.BankNameHasSuffix(*i.BankNameHasSuffix))
	}
	if i.BankNameEqualFold != nil {
		predicates = append(predicates, contact.BankNameEqualFold(*i.BankNameEqualFold))
	}
	if i.BankNameContainsFold != nil {
		predicates = append(predicates, contact.BankNameContainsFold(*i.BankNameContainsFold))
	}

	if i.HasOwner != nil {
		p := contact.HasOwner()
		if !*i.HasOwner {
			p = contact.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasOwnerWith) > 0 {
		with := make([]predicate.Customer, 0, len(i.HasOwnerWith))
		for _, w := range i.HasOwnerWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasOwnerWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, contact.HasOwnerWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyContactWhereInput
	case 1:
		return predicates[0], nil
	default:
		return contact.And(predicates...), nil
	}
}

// Contacts is a parsable slice of Contact.
type Contacts []*Contact

func (c Contacts) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
