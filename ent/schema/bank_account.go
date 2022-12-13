package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/TcMits/wnc-final/pkg/tool"
	"github.com/google/uuid"
)

// BankAccount holds the schema definition for the BankAccount entity.
type BankAccount struct {
	ent.Schema
}

// Mixins of the BankAccount.
func (BankAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the BankAccount.
func (BankAccount) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("customer_id", uuid.UUID{}),
		field.Float("cash_in").
			Positive().
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(20,2)",
				dialect.Postgres: "numeric",
			}),
		field.Float("cash_out").
			Positive().
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(20,2)",
				dialect.Postgres: "numeric",
			}),
		field.String("account_number").
			DefaultFunc(tool.GetTimeStampID).
			Unique().
			MaxLen(255).
			Immutable(),
		field.Bool("is_for_payment").
			Default(false),
	}
}

// Edges of the BankAccount.
func (BankAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("customer", Customer.Type).
			Unique().
			Required().
			Field("customer_id").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
		edge.
			From("sent_transaction", Transaction.Type).
			Ref("sender"),
		edge.
			From("received_transaction", Transaction.Type).
			Ref("receiver"),
		edge.
			From("owned_debts", Debt.Type).
			Ref("owner"),
		edge.
			From("received_debts", Debt.Type).
			Ref("receiver"),
	}
}

func (BankAccount) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account_number"),
		index.Fields("customer_id"),
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}
