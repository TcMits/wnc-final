package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Partner holds the schema definition for the Partner entity.
type Partner struct {
	ent.Schema
}

// Mixins of the Partner.
func (Partner) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Partner.
func (Partner) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("api_key").
			Unique(),
		field.String("secret_key").
			Unique(),
		field.String("public_key").
			Unique(),
		field.String("private_key").
			Unique(),
		field.String("name").
			Optional().
			Default("").
			MaxLen(128),
		field.Bool("is_active").
			Optional().
			Default(true),
	}
}

// Edges of the Partner.
func (Partner) Edges() []ent.Edge {
	return nil
}

func (Partner) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("create_time"),
	}
}
