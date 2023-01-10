package ent

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/TcMits/wnc-final/ent/partner"
	"github.com/TcMits/wnc-final/ent/predicate"
	"github.com/google/uuid"
)

// Partner is the model entity for the Partner schema.
type Partner struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// APIKey holds the value of the "api_key" field.
	APIKey string `json:"api_key,omitempty"`
	// SecretKey holds the value of the "secret_key" field.
	SecretKey string `json:"secret_key,omitempty"`
	// PublicKey holds the value of the "public_key" field.
	PublicKey string `json:"public_key,omitempty"`
	// PrivateKey holds the value of the "private_key" field.
	PrivateKey string `json:"private_key,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Partner) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case partner.FieldIsActive:
			values[i] = new(sql.NullBool)
		case partner.FieldAPIKey, partner.FieldSecretKey, partner.FieldPublicKey, partner.FieldPrivateKey, partner.FieldName:
			values[i] = new(sql.NullString)
		case partner.FieldCreateTime, partner.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case partner.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Partner", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Partner fields.
func (pa *Partner) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case partner.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pa.ID = *value
			}
		case partner.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pa.CreateTime = value.Time
			}
		case partner.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pa.UpdateTime = value.Time
			}
		case partner.FieldAPIKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field api_key", values[i])
			} else if value.Valid {
				pa.APIKey = value.String
			}
		case partner.FieldSecretKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret_key", values[i])
			} else if value.Valid {
				pa.SecretKey = value.String
			}
		case partner.FieldPublicKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_key", values[i])
			} else if value.Valid {
				pa.PublicKey = value.String
			}
		case partner.FieldPrivateKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field private_key", values[i])
			} else if value.Valid {
				pa.PrivateKey = value.String
			}
		case partner.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pa.Name = value.String
			}
		case partner.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				pa.IsActive = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Partner.
// Note that you need to call Partner.Unwrap() before calling this method if this Partner
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Partner) Update() *PartnerUpdateOne {
	return (&PartnerClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the Partner entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Partner) Unwrap() *Partner {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Partner is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Partner) String() string {
	var builder strings.Builder
	builder.WriteString("Partner(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("create_time=")
	builder.WriteString(pa.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(pa.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("api_key=")
	builder.WriteString(pa.APIKey)
	builder.WriteString(", ")
	builder.WriteString("secret_key=")
	builder.WriteString(pa.SecretKey)
	builder.WriteString(", ")
	builder.WriteString("public_key=")
	builder.WriteString(pa.PublicKey)
	builder.WriteString(", ")
	builder.WriteString("private_key=")
	builder.WriteString(pa.PrivateKey)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pa.Name)
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", pa.IsActive))
	builder.WriteByte(')')
	return builder.String()
}

type PartnerCreateRepository struct {
	client   *Client
	isAtomic bool
}

func NewPartnerCreateRepository(
	client *Client,
	isAtomic bool,
) *PartnerCreateRepository {
	return &PartnerCreateRepository{
		client:   client,
		isAtomic: isAtomic,
	}
}

// using in Tx
func (r *PartnerCreateRepository) CreateWithClient(
	ctx context.Context, client *Client, input *PartnerCreateInput,
) (*Partner, error) {
	instance, err := client.Partner.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (r *PartnerCreateRepository) Create(
	ctx context.Context, input *PartnerCreateInput,
) (*Partner, error) {
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

type PartnerDeleteRepository struct {
	client   *Client
	isAtomic bool
}

func NewPartnerDeleteRepository(
	client *Client,
	isAtomic bool,
) *PartnerDeleteRepository {
	return &PartnerDeleteRepository{
		client:   client,
		isAtomic: isAtomic,
	}
}

// using in Tx
func (r *PartnerDeleteRepository) DeleteWithClient(
	ctx context.Context, client *Client, instance *Partner,
) error {
	err := client.Partner.DeleteOne(instance).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *PartnerDeleteRepository) Delete(
	ctx context.Context, instance *Partner,
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

// PartnerCreateInput represents a mutation input for creating partners.
type PartnerCreateInput struct {
	CreateTime *time.Time `json:"create_time,omitempty" form:"create_time"`
	UpdateTime *time.Time `json:"update_time,omitempty" form:"update_time"`
	APIKey     string     `json:"api_key,omitempty" form:"api_key"`
	SecretKey  string     `json:"secret_key,omitempty" form:"secret_key"`
	PrivateKey string     `json:"private_key,omitempty" form:"private_key"`
	PublicKey  string     `json:"public_key,omitempty" form:"public_key"`
	Name       *string    `json:"name,omitempty" form:"name"`
	IsActive   *bool      `json:"is_active,omitempty" form:"is_active"`
}

// Mutate applies the PartnerCreateInput on the PartnerCreate builder.
func (i *PartnerCreateInput) Mutate(m *PartnerMutation) {
	if v := i.CreateTime; v != nil {
		m.SetCreateTime(*v)
	}
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	m.SetAPIKey(i.APIKey)
	m.SetSecretKey(i.SecretKey)
	m.SetPublicKey(i.PublicKey)
	m.SetPrivateKey(i.PrivateKey)
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
}

// SetInput applies the change-set in the PartnerCreateInput on the create builder.
func (c *PartnerCreate) SetInput(i *PartnerCreateInput) *PartnerCreate {
	i.Mutate(c.Mutation())
	return c
}

// PartnerUpdateInput represents a mutation input for updating partners.
type PartnerUpdateInput struct {
	ID            uuid.UUID
	UpdateTime    *time.Time `json:"update_time,omitempty" form:"update_time"`
	APIKey        *string    `json:"api_key,omitempty" form:"api_key"`
	SecretKey     *string    `json:"secret_key,omitempty" form:"secret_key"`
	PublicKey     *string    `json:"public_key,omitempty" form:"public_key"`
	PrivateKey    *string    `json:"private_key,omitempty" form:"private_key"`
	Name          *string    `json:"name,omitempty" form:"name"`
	ClearName     bool
	IsActive      *bool `json:"is_active,omitempty" form:"is_active"`
	ClearIsActive bool
}

// Mutate applies the PartnerUpdateInput on the PartnerMutation.
func (i *PartnerUpdateInput) Mutate(m *PartnerMutation) {
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	if v := i.APIKey; v != nil {
		m.SetAPIKey(*v)
	}
	if v := i.SecretKey; v != nil {
		m.SetSecretKey(*v)
	}
	if v := i.PublicKey; v != nil {
		m.SetPublicKey(*v)
	}
	if v := i.PrivateKey; v != nil {
		m.SetPrivateKey(*v)
	}
	if i.ClearName {
		m.ClearName()
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearIsActive {
		m.ClearIsActive()
	}
	if v := i.IsActive; v != nil {
		m.SetIsActive(*v)
	}
}

// SetInput applies the change-set in the PartnerUpdateInput on the update builder.
func (u *PartnerUpdate) SetInput(i *PartnerUpdateInput) *PartnerUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the PartnerUpdateInput on the update-one builder.
func (u *PartnerUpdateOne) SetInput(i *PartnerUpdateInput) *PartnerUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

type PartnerReadRepository struct {
	client *Client
}

func NewPartnerReadRepository(
	client *Client,
) *PartnerReadRepository {
	return &PartnerReadRepository{
		client: client,
	}
}

func (r *PartnerReadRepository) prepareQuery(
	client *Client, limit *int, offset *int, o *PartnerOrderInput, w *PartnerWhereInput,
) (*PartnerQuery, error) {
	var err error
	q := r.client.Partner.Query()
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
func (r *PartnerReadRepository) GetWithClient(
	ctx context.Context, client *Client, w *PartnerWhereInput, forUpdate bool,
) (*Partner, error) {
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
func (r *PartnerReadRepository) ListWithClient(
	ctx context.Context, client *Client, limit *int, offset *int, o *PartnerOrderInput, w *PartnerWhereInput, forUpdate bool,
) ([]*Partner, error) {
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

func (r *PartnerReadRepository) Count(ctx context.Context, w *PartnerWhereInput) (int, error) {
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

func (r *PartnerReadRepository) Get(ctx context.Context, w *PartnerWhereInput) (*Partner, error) {
	return r.GetWithClient(ctx, r.client, w, false)
}

func (r *PartnerReadRepository) List(
	ctx context.Context, limit *int, offset *int, o *PartnerOrderInput, w *PartnerWhereInput,
) ([]*Partner, error) {
	return r.ListWithClient(ctx, r.client, limit, offset, o, w, false)
}

type PartnerSerializer struct {
	columns map[string]func(context.Context, *Partner) any
}

func NewPartnerSerializer(customColumns map[string]func(context.Context, *Partner) any, columns ...string) *PartnerSerializer {
	columnsMap := map[string]func(context.Context, *Partner) any{}
	for _, col := range columns {
		switch col {

		case partner.FieldID:
			columnsMap[col] = func(ctx context.Context, pa *Partner) any {
				return pa.ID
			}

		case partner.FieldCreateTime:
			columnsMap[col] = func(ctx context.Context, pa *Partner) any {
				return pa.CreateTime
			}

		case partner.FieldUpdateTime:
			columnsMap[col] = func(ctx context.Context, pa *Partner) any {
				return pa.UpdateTime
			}

		case partner.FieldAPIKey:
			columnsMap[col] = func(ctx context.Context, pa *Partner) any {
				return pa.APIKey
			}

		case partner.FieldSecretKey:
			columnsMap[col] = func(ctx context.Context, pa *Partner) any {
				return pa.SecretKey
			}

		case partner.FieldPublicKey:
			columnsMap[col] = func(ctx context.Context, pa *Partner) any {
				return pa.PublicKey
			}

		case partner.FieldPrivateKey:
			columnsMap[col] = func(ctx context.Context, pa *Partner) any {
				return pa.PrivateKey
			}

		case partner.FieldName:
			columnsMap[col] = func(ctx context.Context, pa *Partner) any {
				return pa.Name
			}

		case partner.FieldIsActive:
			columnsMap[col] = func(ctx context.Context, pa *Partner) any {
				return pa.IsActive
			}

		default:
			panic(fmt.Sprintf("Unexpect column %s", col))
		}
	}

	for k, serializeFunc := range customColumns {
		columnsMap[k] = serializeFunc
	}

	return &PartnerSerializer{
		columns: columnsMap,
	}
}

func (s *PartnerSerializer) Serialize(ctx context.Context, pa *Partner) map[string]any {
	result := make(map[string]any, len(s.columns))
	for col, serializeFunc := range s.columns {
		result[col] = serializeFunc(ctx, pa)
	}
	return result
}

type PartnerUpdateRepository struct {
	client   *Client
	isAtomic bool
}

func NewPartnerUpdateRepository(
	client *Client,
	isAtomic bool,
) *PartnerUpdateRepository {
	return &PartnerUpdateRepository{
		client:   client,
		isAtomic: isAtomic,
	}
}

// using in Tx
func (r *PartnerUpdateRepository) UpdateWithClient(
	ctx context.Context, client *Client, instance *Partner, input *PartnerUpdateInput,
) (*Partner, error) {
	newInstance, err := client.Partner.UpdateOne(instance).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	return newInstance, nil
}

func (r *PartnerUpdateRepository) Update(
	ctx context.Context, instance *Partner, input *PartnerUpdateInput,
) (*Partner, error) {
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

// PartnerWhereInput represents a where input for filtering Partner queries.
type PartnerWhereInput struct {
	Predicates []predicate.Partner  `json:"-"`
	Not        *PartnerWhereInput   `json:"not,omitempty"`
	Or         []*PartnerWhereInput `json:"or,omitempty"`
	And        []*PartnerWhereInput `json:"and,omitempty"`

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

	// "api_key" field predicates.
	APIKey             *string  `json:"api_key,omitempty" form:"api_key" param:"api_key" url:"api_key"`
	APIKeyNEQ          *string  `json:"api_key_neq,omitempty" form:"api_key_neq" param:"api_key_neq" url:"api_key_neq"`
	APIKeyIn           []string `json:"api_key_in,omitempty" form:"api_key_in" param:"api_key_in" url:"api_key_in"`
	APIKeyNotIn        []string `json:"api_key_not_in,omitempty" form:"api_key_not_in" param:"api_key_not_in" url:"api_key_not_in"`
	APIKeyGT           *string  `json:"api_key_gt,omitempty" form:"api_key_gt" param:"api_key_gt" url:"api_key_gt"`
	APIKeyGTE          *string  `json:"api_key_gte,omitempty" form:"api_key_gte" param:"api_key_gte" url:"api_key_gte"`
	APIKeyLT           *string  `json:"api_key_lt,omitempty" form:"api_key_lt" param:"api_key_lt" url:"api_key_lt"`
	APIKeyLTE          *string  `json:"api_key_lte,omitempty" form:"api_key_lte" param:"api_key_lte" url:"api_key_lte"`
	APIKeyContains     *string  `json:"api_key_contains,omitempty" form:"api_key_contains" param:"api_key_contains" url:"api_key_contains"`
	APIKeyHasPrefix    *string  `json:"api_key_has_prefix,omitempty" form:"api_key_has_prefix" param:"api_key_has_prefix" url:"api_key_has_prefix"`
	APIKeyHasSuffix    *string  `json:"api_key_has_suffix,omitempty" form:"api_key_has_suffix" param:"api_key_has_suffix" url:"api_key_has_suffix"`
	APIKeyEqualFold    *string  `json:"api_key_equal_fold,omitempty" form:"api_key_equal_fold" param:"api_key_equal_fold" url:"api_key_equal_fold"`
	APIKeyContainsFold *string  `json:"api_key_contains_fold,omitempty" form:"api_key_contains_fold" param:"api_key_contains_fold" url:"api_key_contains_fold"`

	// "secret_key" field predicates.
	SecretKey             *string  `json:"secret_key,omitempty" form:"secret_key" param:"secret_key" url:"secret_key"`
	SecretKeyNEQ          *string  `json:"secret_key_neq,omitempty" form:"secret_key_neq" param:"secret_key_neq" url:"secret_key_neq"`
	SecretKeyIn           []string `json:"secret_key_in,omitempty" form:"secret_key_in" param:"secret_key_in" url:"secret_key_in"`
	SecretKeyNotIn        []string `json:"secret_key_not_in,omitempty" form:"secret_key_not_in" param:"secret_key_not_in" url:"secret_key_not_in"`
	SecretKeyGT           *string  `json:"secret_key_gt,omitempty" form:"secret_key_gt" param:"secret_key_gt" url:"secret_key_gt"`
	SecretKeyGTE          *string  `json:"secret_key_gte,omitempty" form:"secret_key_gte" param:"secret_key_gte" url:"secret_key_gte"`
	SecretKeyLT           *string  `json:"secret_key_lt,omitempty" form:"secret_key_lt" param:"secret_key_lt" url:"secret_key_lt"`
	SecretKeyLTE          *string  `json:"secret_key_lte,omitempty" form:"secret_key_lte" param:"secret_key_lte" url:"secret_key_lte"`
	SecretKeyContains     *string  `json:"secret_key_contains,omitempty" form:"secret_key_contains" param:"secret_key_contains" url:"secret_key_contains"`
	SecretKeyHasPrefix    *string  `json:"secret_key_has_prefix,omitempty" form:"secret_key_has_prefix" param:"secret_key_has_prefix" url:"secret_key_has_prefix"`
	SecretKeyHasSuffix    *string  `json:"secret_key_has_suffix,omitempty" form:"secret_key_has_suffix" param:"secret_key_has_suffix" url:"secret_key_has_suffix"`
	SecretKeyEqualFold    *string  `json:"secret_key_equal_fold,omitempty" form:"secret_key_equal_fold" param:"secret_key_equal_fold" url:"secret_key_equal_fold"`
	SecretKeyContainsFold *string  `json:"secret_key_contains_fold,omitempty" form:"secret_key_contains_fold" param:"secret_key_contains_fold" url:"secret_key_contains_fold"`

	// "public_key" field predicates.
	PublicKey             *string  `json:"public_key,omitempty" form:"public_key" param:"public_key" url:"public_key"`
	PublicKeyNEQ          *string  `json:"public_key_neq,omitempty" form:"public_key_neq" param:"public_key_neq" url:"public_key_neq"`
	PublicKeyIn           []string `json:"public_key_in,omitempty" form:"public_key_in" param:"public_key_in" url:"public_key_in"`
	PublicKeyNotIn        []string `json:"public_key_not_in,omitempty" form:"public_key_not_in" param:"public_key_not_in" url:"public_key_not_in"`
	PublicKeyGT           *string  `json:"public_key_gt,omitempty" form:"public_key_gt" param:"public_key_gt" url:"public_key_gt"`
	PublicKeyGTE          *string  `json:"public_key_gte,omitempty" form:"public_key_gte" param:"public_key_gte" url:"public_key_gte"`
	PublicKeyLT           *string  `json:"public_key_lt,omitempty" form:"public_key_lt" param:"public_key_lt" url:"public_key_lt"`
	PublicKeyLTE          *string  `json:"public_key_lte,omitempty" form:"public_key_lte" param:"public_key_lte" url:"public_key_lte"`
	PublicKeyContains     *string  `json:"public_key_contains,omitempty" form:"public_key_contains" param:"public_key_contains" url:"public_key_contains"`
	PublicKeyHasPrefix    *string  `json:"public_key_has_prefix,omitempty" form:"public_key_has_prefix" param:"public_key_has_prefix" url:"public_key_has_prefix"`
	PublicKeyHasSuffix    *string  `json:"public_key_has_suffix,omitempty" form:"public_key_has_suffix" param:"public_key_has_suffix" url:"public_key_has_suffix"`
	PublicKeyEqualFold    *string  `json:"public_key_equal_fold,omitempty" form:"public_key_equal_fold" param:"public_key_equal_fold" url:"public_key_equal_fold"`
	PublicKeyContainsFold *string  `json:"public_key_contains_fold,omitempty" form:"public_key_contains_fold" param:"public_key_contains_fold" url:"public_key_contains_fold"`

	// "private_key" field predicates.
	PrivateKey             *string  `json:"private_key,omitempty" form:"private_key" param:"private_key" url:"private_key"`
	PrivateKeyNEQ          *string  `json:"private_key_neq,omitempty" form:"private_key_neq" param:"private_key_neq" url:"private_key_neq"`
	PrivateKeyIn           []string `json:"private_key_in,omitempty" form:"private_key_in" param:"private_key_in" url:"private_key_in"`
	PrivateKeyNotIn        []string `json:"private_key_not_in,omitempty" form:"private_key_not_in" param:"private_key_not_in" url:"private_key_not_in"`
	PrivateKeyGT           *string  `json:"private_key_gt,omitempty" form:"private_key_gt" param:"private_key_gt" url:"private_key_gt"`
	PrivateKeyGTE          *string  `json:"private_key_gte,omitempty" form:"private_key_gte" param:"private_key_gte" url:"private_key_gte"`
	PrivateKeyLT           *string  `json:"private_key_lt,omitempty" form:"private_key_lt" param:"private_key_lt" url:"private_key_lt"`
	PrivateKeyLTE          *string  `json:"private_key_lte,omitempty" form:"private_key_lte" param:"private_key_lte" url:"private_key_lte"`
	PrivateKeyContains     *string  `json:"private_key_contains,omitempty" form:"private_key_contains" param:"private_key_contains" url:"private_key_contains"`
	PrivateKeyHasPrefix    *string  `json:"private_key_has_prefix,omitempty" form:"private_key_has_prefix" param:"private_key_has_prefix" url:"private_key_has_prefix"`
	PrivateKeyHasSuffix    *string  `json:"private_key_has_suffix,omitempty" form:"private_key_has_suffix" param:"private_key_has_suffix" url:"private_key_has_suffix"`
	PrivateKeyEqualFold    *string  `json:"private_key_equal_fold,omitempty" form:"private_key_equal_fold" param:"private_key_equal_fold" url:"private_key_equal_fold"`
	PrivateKeyContainsFold *string  `json:"private_key_contains_fold,omitempty" form:"private_key_contains_fold" param:"private_key_contains_fold" url:"private_key_contains_fold"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty" form:"name" param:"name" url:"name"`
	NameNEQ          *string  `json:"name_neq,omitempty" form:"name_neq" param:"name_neq" url:"name_neq"`
	NameIn           []string `json:"name_in,omitempty" form:"name_in" param:"name_in" url:"name_in"`
	NameNotIn        []string `json:"name_not_in,omitempty" form:"name_not_in" param:"name_not_in" url:"name_not_in"`
	NameGT           *string  `json:"name_gt,omitempty" form:"name_gt" param:"name_gt" url:"name_gt"`
	NameGTE          *string  `json:"name_gte,omitempty" form:"name_gte" param:"name_gte" url:"name_gte"`
	NameLT           *string  `json:"name_lt,omitempty" form:"name_lt" param:"name_lt" url:"name_lt"`
	NameLTE          *string  `json:"name_lte,omitempty" form:"name_lte" param:"name_lte" url:"name_lte"`
	NameContains     *string  `json:"name_contains,omitempty" form:"name_contains" param:"name_contains" url:"name_contains"`
	NameHasPrefix    *string  `json:"name_has_prefix,omitempty" form:"name_has_prefix" param:"name_has_prefix" url:"name_has_prefix"`
	NameHasSuffix    *string  `json:"name_has_suffix,omitempty" form:"name_has_suffix" param:"name_has_suffix" url:"name_has_suffix"`
	NameIsNil        bool     `json:"name_is_nil,omitempty" form:"name_is_nil" param:"name_is_nil" url:"name_is_nil"`
	NameNotNil       bool     `json:"name_not_nil,omitempty" form:"name_not_nil" param:"name_not_nil" url:"name_not_nil"`
	NameEqualFold    *string  `json:"name_equal_fold,omitempty" form:"name_equal_fold" param:"name_equal_fold" url:"name_equal_fold"`
	NameContainsFold *string  `json:"name_contains_fold,omitempty" form:"name_contains_fold" param:"name_contains_fold" url:"name_contains_fold"`

	// "is_active" field predicates.
	IsActive       *bool `json:"is_active,omitempty" form:"is_active" param:"is_active" url:"is_active"`
	IsActiveNEQ    *bool `json:"is_active_neq,omitempty" form:"is_active_neq" param:"is_active_neq" url:"is_active_neq"`
	IsActiveIsNil  bool  `json:"is_active_is_nil,omitempty" form:"is_active_is_nil" param:"is_active_is_nil" url:"is_active_is_nil"`
	IsActiveNotNil bool  `json:"is_active_not_nil,omitempty" form:"is_active_not_nil" param:"is_active_not_nil" url:"is_active_not_nil"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *PartnerWhereInput) AddPredicates(predicates ...predicate.Partner) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the PartnerWhereInput filter on the PartnerQuery builder.
func (i *PartnerWhereInput) Filter(q *PartnerQuery) (*PartnerQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyPartnerWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyPartnerWhereInput is returned in case the PartnerWhereInput is empty.
var ErrEmptyPartnerWhereInput = errors.New("ent: empty predicate PartnerWhereInput")

// P returns a predicate for filtering partners.
// An error is returned if the input is empty or invalid.
func (i *PartnerWhereInput) P() (predicate.Partner, error) {
	var predicates []predicate.Partner
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, partner.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Partner, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, partner.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Partner, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, partner.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, partner.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, partner.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, partner.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, partner.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, partner.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, partner.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, partner.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, partner.IDLTE(*i.IDLTE))
	}
	if i.CreateTime != nil {
		predicates = append(predicates, partner.CreateTimeEQ(*i.CreateTime))
	}
	if i.CreateTimeNEQ != nil {
		predicates = append(predicates, partner.CreateTimeNEQ(*i.CreateTimeNEQ))
	}
	if len(i.CreateTimeIn) > 0 {
		predicates = append(predicates, partner.CreateTimeIn(i.CreateTimeIn...))
	}
	if len(i.CreateTimeNotIn) > 0 {
		predicates = append(predicates, partner.CreateTimeNotIn(i.CreateTimeNotIn...))
	}
	if i.CreateTimeGT != nil {
		predicates = append(predicates, partner.CreateTimeGT(*i.CreateTimeGT))
	}
	if i.CreateTimeGTE != nil {
		predicates = append(predicates, partner.CreateTimeGTE(*i.CreateTimeGTE))
	}
	if i.CreateTimeLT != nil {
		predicates = append(predicates, partner.CreateTimeLT(*i.CreateTimeLT))
	}
	if i.CreateTimeLTE != nil {
		predicates = append(predicates, partner.CreateTimeLTE(*i.CreateTimeLTE))
	}
	if i.UpdateTime != nil {
		predicates = append(predicates, partner.UpdateTimeEQ(*i.UpdateTime))
	}
	if i.UpdateTimeNEQ != nil {
		predicates = append(predicates, partner.UpdateTimeNEQ(*i.UpdateTimeNEQ))
	}
	if len(i.UpdateTimeIn) > 0 {
		predicates = append(predicates, partner.UpdateTimeIn(i.UpdateTimeIn...))
	}
	if len(i.UpdateTimeNotIn) > 0 {
		predicates = append(predicates, partner.UpdateTimeNotIn(i.UpdateTimeNotIn...))
	}
	if i.UpdateTimeGT != nil {
		predicates = append(predicates, partner.UpdateTimeGT(*i.UpdateTimeGT))
	}
	if i.UpdateTimeGTE != nil {
		predicates = append(predicates, partner.UpdateTimeGTE(*i.UpdateTimeGTE))
	}
	if i.UpdateTimeLT != nil {
		predicates = append(predicates, partner.UpdateTimeLT(*i.UpdateTimeLT))
	}
	if i.UpdateTimeLTE != nil {
		predicates = append(predicates, partner.UpdateTimeLTE(*i.UpdateTimeLTE))
	}
	if i.APIKey != nil {
		predicates = append(predicates, partner.APIKeyEQ(*i.APIKey))
	}
	if i.APIKeyNEQ != nil {
		predicates = append(predicates, partner.APIKeyNEQ(*i.APIKeyNEQ))
	}
	if len(i.APIKeyIn) > 0 {
		predicates = append(predicates, partner.APIKeyIn(i.APIKeyIn...))
	}
	if len(i.APIKeyNotIn) > 0 {
		predicates = append(predicates, partner.APIKeyNotIn(i.APIKeyNotIn...))
	}
	if i.APIKeyGT != nil {
		predicates = append(predicates, partner.APIKeyGT(*i.APIKeyGT))
	}
	if i.APIKeyGTE != nil {
		predicates = append(predicates, partner.APIKeyGTE(*i.APIKeyGTE))
	}
	if i.APIKeyLT != nil {
		predicates = append(predicates, partner.APIKeyLT(*i.APIKeyLT))
	}
	if i.APIKeyLTE != nil {
		predicates = append(predicates, partner.APIKeyLTE(*i.APIKeyLTE))
	}
	if i.APIKeyContains != nil {
		predicates = append(predicates, partner.APIKeyContains(*i.APIKeyContains))
	}
	if i.APIKeyHasPrefix != nil {
		predicates = append(predicates, partner.APIKeyHasPrefix(*i.APIKeyHasPrefix))
	}
	if i.APIKeyHasSuffix != nil {
		predicates = append(predicates, partner.APIKeyHasSuffix(*i.APIKeyHasSuffix))
	}
	if i.APIKeyEqualFold != nil {
		predicates = append(predicates, partner.APIKeyEqualFold(*i.APIKeyEqualFold))
	}
	if i.APIKeyContainsFold != nil {
		predicates = append(predicates, partner.APIKeyContainsFold(*i.APIKeyContainsFold))
	}
	if i.SecretKey != nil {
		predicates = append(predicates, partner.SecretKeyEQ(*i.SecretKey))
	}
	if i.SecretKeyNEQ != nil {
		predicates = append(predicates, partner.SecretKeyNEQ(*i.SecretKeyNEQ))
	}
	if len(i.SecretKeyIn) > 0 {
		predicates = append(predicates, partner.SecretKeyIn(i.SecretKeyIn...))
	}
	if len(i.SecretKeyNotIn) > 0 {
		predicates = append(predicates, partner.SecretKeyNotIn(i.SecretKeyNotIn...))
	}
	if i.SecretKeyGT != nil {
		predicates = append(predicates, partner.SecretKeyGT(*i.SecretKeyGT))
	}
	if i.SecretKeyGTE != nil {
		predicates = append(predicates, partner.SecretKeyGTE(*i.SecretKeyGTE))
	}
	if i.SecretKeyLT != nil {
		predicates = append(predicates, partner.SecretKeyLT(*i.SecretKeyLT))
	}
	if i.SecretKeyLTE != nil {
		predicates = append(predicates, partner.SecretKeyLTE(*i.SecretKeyLTE))
	}
	if i.SecretKeyContains != nil {
		predicates = append(predicates, partner.SecretKeyContains(*i.SecretKeyContains))
	}
	if i.SecretKeyHasPrefix != nil {
		predicates = append(predicates, partner.SecretKeyHasPrefix(*i.SecretKeyHasPrefix))
	}
	if i.SecretKeyHasSuffix != nil {
		predicates = append(predicates, partner.SecretKeyHasSuffix(*i.SecretKeyHasSuffix))
	}
	if i.SecretKeyEqualFold != nil {
		predicates = append(predicates, partner.SecretKeyEqualFold(*i.SecretKeyEqualFold))
	}
	if i.SecretKeyContainsFold != nil {
		predicates = append(predicates, partner.SecretKeyContainsFold(*i.SecretKeyContainsFold))
	}
	if i.PublicKey != nil {
		predicates = append(predicates, partner.PublicKeyEQ(*i.PublicKey))
	}
	if i.PublicKeyNEQ != nil {
		predicates = append(predicates, partner.PublicKeyNEQ(*i.PublicKeyNEQ))
	}
	if len(i.PublicKeyIn) > 0 {
		predicates = append(predicates, partner.PublicKeyIn(i.PublicKeyIn...))
	}
	if len(i.PublicKeyNotIn) > 0 {
		predicates = append(predicates, partner.PublicKeyNotIn(i.PublicKeyNotIn...))
	}
	if i.PublicKeyGT != nil {
		predicates = append(predicates, partner.PublicKeyGT(*i.PublicKeyGT))
	}
	if i.PublicKeyGTE != nil {
		predicates = append(predicates, partner.PublicKeyGTE(*i.PublicKeyGTE))
	}
	if i.PublicKeyLT != nil {
		predicates = append(predicates, partner.PublicKeyLT(*i.PublicKeyLT))
	}
	if i.PublicKeyLTE != nil {
		predicates = append(predicates, partner.PublicKeyLTE(*i.PublicKeyLTE))
	}
	if i.PublicKeyContains != nil {
		predicates = append(predicates, partner.PublicKeyContains(*i.PublicKeyContains))
	}
	if i.PublicKeyHasPrefix != nil {
		predicates = append(predicates, partner.PublicKeyHasPrefix(*i.PublicKeyHasPrefix))
	}
	if i.PublicKeyHasSuffix != nil {
		predicates = append(predicates, partner.PublicKeyHasSuffix(*i.PublicKeyHasSuffix))
	}
	if i.PublicKeyEqualFold != nil {
		predicates = append(predicates, partner.PublicKeyEqualFold(*i.PublicKeyEqualFold))
	}
	if i.PublicKeyContainsFold != nil {
		predicates = append(predicates, partner.PublicKeyContainsFold(*i.PublicKeyContainsFold))
	}
	if i.PrivateKey != nil {
		predicates = append(predicates, partner.PrivateKeyEQ(*i.PrivateKey))
	}
	if i.PrivateKeyNEQ != nil {
		predicates = append(predicates, partner.PrivateKeyNEQ(*i.PrivateKeyNEQ))
	}
	if len(i.PrivateKeyIn) > 0 {
		predicates = append(predicates, partner.PrivateKeyIn(i.PrivateKeyIn...))
	}
	if len(i.PrivateKeyNotIn) > 0 {
		predicates = append(predicates, partner.PrivateKeyNotIn(i.PrivateKeyNotIn...))
	}
	if i.PrivateKeyGT != nil {
		predicates = append(predicates, partner.PrivateKeyGT(*i.PrivateKeyGT))
	}
	if i.PrivateKeyGTE != nil {
		predicates = append(predicates, partner.PrivateKeyGTE(*i.PrivateKeyGTE))
	}
	if i.PrivateKeyLT != nil {
		predicates = append(predicates, partner.PrivateKeyLT(*i.PrivateKeyLT))
	}
	if i.PrivateKeyLTE != nil {
		predicates = append(predicates, partner.PrivateKeyLTE(*i.PrivateKeyLTE))
	}
	if i.PrivateKeyContains != nil {
		predicates = append(predicates, partner.PrivateKeyContains(*i.PrivateKeyContains))
	}
	if i.PrivateKeyHasPrefix != nil {
		predicates = append(predicates, partner.PrivateKeyHasPrefix(*i.PrivateKeyHasPrefix))
	}
	if i.PrivateKeyHasSuffix != nil {
		predicates = append(predicates, partner.PrivateKeyHasSuffix(*i.PrivateKeyHasSuffix))
	}
	if i.PrivateKeyEqualFold != nil {
		predicates = append(predicates, partner.PrivateKeyEqualFold(*i.PrivateKeyEqualFold))
	}
	if i.PrivateKeyContainsFold != nil {
		predicates = append(predicates, partner.PrivateKeyContainsFold(*i.PrivateKeyContainsFold))
	}
	if i.Name != nil {
		predicates = append(predicates, partner.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, partner.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, partner.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, partner.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, partner.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, partner.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, partner.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, partner.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, partner.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, partner.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, partner.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameIsNil {
		predicates = append(predicates, partner.NameIsNil())
	}
	if i.NameNotNil {
		predicates = append(predicates, partner.NameNotNil())
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, partner.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, partner.NameContainsFold(*i.NameContainsFold))
	}
	if i.IsActive != nil {
		predicates = append(predicates, partner.IsActiveEQ(*i.IsActive))
	}
	if i.IsActiveNEQ != nil {
		predicates = append(predicates, partner.IsActiveNEQ(*i.IsActiveNEQ))
	}
	if i.IsActiveIsNil {
		predicates = append(predicates, partner.IsActiveIsNil())
	}
	if i.IsActiveNotNil {
		predicates = append(predicates, partner.IsActiveNotNil())
	}

	switch len(predicates) {
	case 0:
		return nil, ErrEmptyPartnerWhereInput
	case 1:
		return predicates[0], nil
	default:
		return partner.And(predicates...), nil
	}
}

// Partners is a parsable slice of Partner.
type Partners []*Partner

func (pa Partners) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
