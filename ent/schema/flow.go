package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// Flow holds the schema definition for the Flow entity.
type Flow struct {
	ent.Schema
}

// Fields of the Flow.
func (Flow) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.UUID("state", uuid.UUID{}),
		field.String("stateTable").NotEmpty(),
		field.UUID("initialState", uuid.UUID{}),
		field.UUID("terminationState", uuid.UUID{}),
		field.UUID("pastState", uuid.UUID{}).Optional(),
		field.Strings("inputs").Immutable().Optional(),
	}
}

// Edges of the Flow.
func (Flow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("survey", Survey.Type).Ref("flow").Unique(),
		edge.To("questions", Question.Type).Required(),
	}
}
