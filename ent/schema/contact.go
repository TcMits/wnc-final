package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Contact holds the schema definition for the Contact entity.
type Contact struct {
	ent.Schema
}

// Mixins of the Contact.
func (Contact) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Contact.
func (Contact) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("owner_id", uuid.UUID{}),
		field.
			String("account_number").
			MaxLen(255),
		field.
			String("suggest_name").
			MaxLen(255),
		field.Text("bank_name"),
	}
}

// Edges of the Contact.
func (Contact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("owner", Customer.Type).
			Unique().
			Required().
			Field("owner_id").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

func (Contact) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("suggest_name"),
		index.Fields("create_time"),
		index.Fields("owner_id"),
		index.Fields("account_number", "bank_name", "owner_id").Unique(),
	}
}
