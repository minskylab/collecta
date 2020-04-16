package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/rs/xid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", xid.ID{}),
		field.String("name").NotEmpty(),
		field.String("username").Optional(),
		field.Time("lastActivity").Default(func() time.Time {
			return time.Now()
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Account.Type),
		edge.To("contacts", Contact.Type),
		edge.To("surveys", Survey.Type),
		edge.From("domain", Domain.Type).Ref("users").Unique(),
	}
}
