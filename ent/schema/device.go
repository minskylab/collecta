package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.String("device"),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return nil
}
