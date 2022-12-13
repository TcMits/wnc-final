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

// Debt holds the schema definition for the Debt entity.
type Debt struct {
	ent.Schema
}

// Mixins of the Debt.
func (Debt) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Debt.
func (Debt) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.
			String("owner_bank_account_number").
			MaxLen(255),
		field.
			String("owner_bank_name").
			MaxLen(255),
		field.
			String("owner_name").
			MaxLen(256),
		field.
			UUID("owner_id", uuid.UUID{}).
			Nillable(),
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
			Nillable(),
		field.UUID("transaction_id", uuid.UUID{}).
			Optional().
			Nillable(),
		field.
			Enum("status").
			Default("pending").
			Values("pending", "cancelled", "fulfilled"),
		field.
			Text("description").
			Optional().
			Default(""),
		field.Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(20,2)",
				dialect.Postgres: "numeric",
			}),
	}
}

// Edges of the Debt.
func (Debt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("owner", BankAccount.Type).
			Required().
			Unique().
			Field("owner_id").
			Annotations(entsql.Annotation{
				OnDelete: entsql.SetNull,
			}),
		edge.
			To("receiver", BankAccount.Type).
			Required().
			Unique().
			Field("receiver_id").
			Annotations(entsql.Annotation{
				OnDelete: entsql.SetNull,
			}),
		edge.
			From("transaction", Transaction.Type).
			Field("transaction_id").
			Ref("debt").
			Unique().
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

func (Debt) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id"),
		index.Fields("receiver_id"),
		index.Fields("transaction_id"),
		index.Fields("create_time"),
		index.Fields("update_time"),
	}
}
