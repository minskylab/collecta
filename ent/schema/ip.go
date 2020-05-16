package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// IP holds the schema definition for the IP entity.
type IP struct {
	ent.Schema
}

// Fields of the IP.
func (IP) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip"),
	}
}

// Edges of the IP.
func (IP) Edges() []ent.Edge {
	return nil
}
