package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("name").NotEmpty(),
		field.Time("lastActivity").Default(time.Now),
		field.String("username").Optional(),
		field.String("picture").Optional(),
		field.Strings("roles").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Account.Type),
		edge.To("contacts", Contact.Type),
		edge.To("surveys", Survey.Type),
		edge.From("domains", Domain.Type).Ref("users"),
		edge.From("adminOf", Domain.Type).Ref("admins"),
	}
}
