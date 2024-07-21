package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permissions.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MaxLen(100),
		field.Text("description").NotEmpty(),
		field.Bool("read").Default(false),
		field.Bool("create").Default(false),
		field.Bool("update").Default(false),
		field.Bool("delete").Default(false),
	}
}

// Edges of the Permissions.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type),
		edge.To("module", Module.Type),
	}
}
