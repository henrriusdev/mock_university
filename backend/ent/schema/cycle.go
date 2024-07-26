package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Cycle holds the schema definition for the Cycle entity.
type Cycle struct {
	ent.Schema
}

// Fields of the Cycle.
func (Cycle) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MaxLen(100),
		field.Time("start_date").Default(time.Now),
		field.Time("end_date").Default(time.Now),
		field.Bool("active").Default(false),
	}
}

// Edges of the Cycle.
func (Cycle) Edges() []ent.Edge {
	return nil
}
