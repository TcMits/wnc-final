package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Mixins of the Transaction.
func (Transaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.
			UUID("source_transaction_id", uuid.UUID{}).
			Optional().
			Nillable(),
		field.
			Enum("status").
			Default("draft").
			Values("draft", "verified", "success"),
		field.
			String("receiver_bank_account_number").
			MaxLen(255),
		field.
			String("receiver_bank_name").
			MaxLen(255),
		field.
			String("receiver_name").
			MaxLen(256),
		field.UUID("receiver_id", uuid.UUID{}).
			Optional().
			Nillable(),
		field.
			String("sender_bank_account_number").
			MaxLen(255),
		field.
			String("sender_bank_name").
			MaxLen(255),
		field.
			String("sender_name").
			MaxLen(256),
		field.UUID("sender_id", uuid.UUID{}).
			Nillable(),
		field.Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(20,2)",
				dialect.Postgres: "numeric",
			}),
		field.
			Enum("transaction_type").
			Values("internal", "external"),
		field.
			Text("description").
			Optional().
			Default(""),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("fee_transaction", Transaction.Type).
			Unique().
			From("source_transaction").
			Field("source_transaction_id").
			Annotations(entsql.Annotation{
				OnDelete: entsql.SetNull,
			}).
			Unique(),
		edge.
			To("receiver", BankAccount.Type).
			Unique().
			Field("receiver_id").
			Annotations(entsql.Annotation{
				OnDelete: entsql.SetNull,
			}),
		edge.
			To("sender", BankAccount.Type).
			Required().
			Unique().
			Field("sender_id").
			Annotations(entsql.Annotation{
				OnDelete: entsql.SetNull,
			}),
		edge.
			To("debt", Debt.Type).
			Unique(),
	}
}

func (Transaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("source_transaction_id"),
		index.Fields("receiver_id"),
		index.Fields("sender_id"),
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}
