package schema

import "github.com/facebookincubator/ent"

// IP holds the schema definition for the IP entity.
type IP struct {
	ent.Schema
}

// Fields of the IP.
func (IP) Fields() []ent.Field {
	return nil
}

// Edges of the IP.
func (IP) Edges() []ent.Edge {
	return nil
}
