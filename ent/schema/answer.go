package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// Answer holds the schema definition for the Answer entity.
type Answer struct {
	ent.Schema
}

// Fields of the Answer.
func (Answer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.Time("at").Immutable().Default(time.Now),
		field.Strings("responses").Immutable(),
		field.String("validator").Optional().Immutable(),
		field.Bool("valid").Optional(),
	}
}

// Edges of the Answer.
func (Answer) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.From("user", User.Type).Ref("answers").Required().Unique(),
		edge.From("question", Question.Type).Ref("answers").Required().Unique(),
	}
}
