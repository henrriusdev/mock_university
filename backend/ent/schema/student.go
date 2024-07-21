package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Students.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("identity_card").NotEmpty().MaxLen(20).Unique(),
		field.Time("birth_date"),
		field.String("phone").NotEmpty().MaxLen(20),
		field.String("address").NotEmpty().MaxLen(255),
		field.Int("number").NonNegative(),
		field.String("district").NotEmpty().MaxLen(100),
		field.String("city").NotEmpty().MaxLen(100),
		field.Int("postal_code").NonNegative(),
		field.Int("credit_units_accumulated").NonNegative(),
	}
}

// Edges of the Students.
func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", Users.Type).Unique(),
	}
}
