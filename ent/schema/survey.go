package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/collecta/uuid"
)

// Survey holds the schema definition for the Survey entity.
type Survey struct {
	ent.Schema
}

// Fields of the Survey.
func (Survey) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.Strings("tags"),
		field.Time("lastInteraction"),
		field.Time("dueDate").Immutable().Default(func() time.Time {
			return time.Now().Add(3 * 24 * time.Hour) // 3 days
		}).Immutable(),
		field.String("title").NotEmpty(),
		field.String("description").Optional(),
		// field.JSON("metadata", map[string]interface{}{}).Optional(),
		field.Bool("done").Default(false).Optional(),
		field.Bool("isPublic").Default(false).Optional(),
	}
}

// Edges of the Survey.
func (Survey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("flow", Flow.Type).Unique().Required(),
		edge.From("for", Person.Type).Ref("surveys").Unique().Required(),
		edge.From("owner", Domain.Type).Ref("surveys").Unique(),
	}
}
