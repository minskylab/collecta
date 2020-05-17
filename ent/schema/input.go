package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// Input holds the schema definition for the Input entity.
type Input struct {
	ent.Schema
}

// Fields of the Input.
func (Input) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.Enum("kind").Values("Text", "Options", "Satisfaction", "Boolean").Immutable(),
		field.Bool("multiple").Default(false).Optional(),
		field.Strings("defaults").Optional(),
		field.JSON("options", map[string]interface{}{}).Optional(),
	}
}

// Edges of the Input.
func (Input) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", Question.Type).Ref("input").Unique().Required(),
	}
}
