package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Note holds the schema definition for the Note entity.
type Note struct {
	ent.Schema
}

// Fields of the Note.
func (Note) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("notes", []float64{}).Optional(),
		field.Float32("average").Optional(),
	}
}

// Edges of the Note.
func (Note) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type),
		edge.To("professor", Professor.Type),
		edge.To("subject", Subject.Type),
		edge.To("cycle", Cycle.Type),
	}
}
