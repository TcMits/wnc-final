package schema

import (
	"regexp"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Mixins of the Customer.
func (Customer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.Text("jwt_token_key").
			Optional().
			DefaultFunc(uuid.NewString).
			Sensitive().
			Annotations(
				entgql.Skip(entgql.SkipAll),
			),
		field.String("password").
			Optional().
			Sensitive().
			SchemaType(map[string]string{
				dialect.MySQL: "char(32)",
			}).
			Annotations(
				entgql.Skip(entgql.SkipAll),
			).
			MaxLen(255),
		field.String("username").
			Unique().
			MaxLen(128).
			Match(regexp.MustCompile("^[a-zA-Z0-9]{6,128}$")).
			Comment("Required. 128 characters or fewer. Letters, digits only."),
		field.String("first_name").
			Optional().
			Default("").
			MaxLen(128),
		field.String("last_name").
			Optional().
			Default("").
			MaxLen(128),
		field.String("phone_number").
			MaxLen(128).
			Unique(),
		field.String("email").
			Unique().
			Validate(func(s string) error {
				return validation.Validate(s, is.Email)
			}).
			MaxLen(255),
		field.Bool("is_active").
			Optional().
			Default(true),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("bank_accounts", BankAccount.Type).
			Ref("customer"),
		edge.
			From("contacts", Contact.Type).
			Ref("owner"),
	}
}

func (Customer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("create_time"),
	}
}
