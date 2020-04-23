package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// Shorts holds the schema definition for the Shorts entity.
type Shorts struct {
	ent.Schema
}

// Fields of the Shorts.
func (Shorts) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").NotEmpty().Unique().Immutable(),
		field.UUID("value", uuid.UUID{}).Immutable(),
	}
}

// Edges of the Shorts.
func (Shorts) Edges() []ent.Edge {
	return nil
}
