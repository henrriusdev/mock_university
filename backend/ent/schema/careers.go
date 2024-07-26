package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Careers holds the schema definition for the Careers entity.
type Careers struct {
	ent.Schema
}

// Fields of the Careers.
func (Careers) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MaxLen(100),
		field.Text("description").NotEmpty(),
	}
}

// Edges of the Careers.
func (Careers) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("leader", Professor.Type).Unique(),
		edge.From("students", Student.Type).Ref("career"),
		edge.From("subjects", Subject.Type).Ref("career"),
	}
}
