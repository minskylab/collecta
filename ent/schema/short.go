package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// Short holds the schema definition for the Short entity.
type Short struct {
	ent.Schema
}

// Fields of the Short.
func (Short) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").NotEmpty().Unique().Immutable(),
		field.UUID("value", uuid.UUID{}).Immutable(),
	}
}

// Edges of the Short.
func (Short) Edges() []ent.Edge {
	return nil
}
