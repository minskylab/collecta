package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/rs/xid"
)

// Flow holds the schema definition for the Flow entity.
type Flow struct {
	ent.Schema
}

// Fields of the Flow.
func (Flow) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", xid.ID{}),
		field.UUID("state", xid.ID{}),
		field.String("stateTable").NotEmpty(),
		field.Strings("inputs").Immutable().Optional(),
	}
}

// Edges of the Flow.
func (Flow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("questions", Question.Type).Required(),
	}
}
