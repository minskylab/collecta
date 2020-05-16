package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// Contact holds the schema definition for the Contact entity.
type Contact struct {
	ent.Schema
}

// Fields of the Contact.
func (Contact) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("name"),
		field.String("value").NotEmpty(),
		field.Enum("kind").Values("Email", "Phone").Default("Email"),
		field.Bool("principal"),
		field.Bool("validated"),
		field.Bool("fromAccount").Default(false),
	}
}

// Edges of the Contact.
func (Contact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Person.Type).Ref("contacts").Unique().Required(),
	}
}
