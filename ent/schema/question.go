package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// Question holds the schema definition for the Question entity.
type Question struct {
	ent.Schema
}

// Fields of the Question.
func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("hash"),
		field.String("title").NotEmpty(),
		field.String("description"),
		field.JSON("metadata", map[string]interface{}{}).Optional(),
		field.String("validator").Optional().Immutable(),
		field.Bool("anonymous").Default(false),
	}
}

// Edges of the Question.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("answers", Answer.Type),
		edge.To("input", Input.Type).Unique(),
		edge.From("flow", Flow.Type).Unique().Ref("questions"),
	}
}
