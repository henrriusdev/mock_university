package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Module holds the schema definition for the Module entity.
type Module struct {
	ent.Schema
}

// Fields of the Module.
func (Module) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MaxLen(100).Unique(),
	}
}

// Edges of the Module.
func (Module) Edges() []ent.Edge {
	return nil
}
