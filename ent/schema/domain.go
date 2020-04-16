package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/rs/xid"
)

// Domain holds the schema definition for the Domain entity.
type Domain struct {
	ent.Schema
}

// Fields of the Domain.
func (Domain) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", xid.ID{}),
		field.Strings("tags"),
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty().Unique(),
		field.String("domain").Unique(),
	}
}

// Edges of the Domain.
func (Domain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("surveys", Survey.Type),
		edge.To("users", User.Type),
	}
}
