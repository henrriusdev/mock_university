package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Subject holds the schema definition for the Subject entity.
type Subject struct {
	ent.Schema
}

// Fields of the Subject.
func (Subject) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MaxLen(100),
		field.String("description").NotEmpty().MaxLen(255),
		field.Int("credit_units").Positive(),
		field.Int("semester").Positive(),
		field.String("code").NotEmpty().MaxLen(20).Unique(),
		field.Int("practice_hours").Positive(),
		field.Int("theory_hours").Positive(),
		field.Int("lab_hours").Positive(),
		field.Int("total_hours").Positive(),
		field.JSON("class_schedule", map[string][]string{}).Optional(),
	}
}

// Edges of the Subject.
func (Subject) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("professor", Professor.Type),
	}
}
