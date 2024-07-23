package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.String("reference").NotEmpty().MaxLen(20).Unique(),
		field.Time("date"),
		field.Float("amount").Positive(),
		field.String("description").NotEmpty().MaxLen(255),
		field.Int("fee_number").Positive(),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type).Unique(),
		edge.To("cycle", Cycle.Type).Unique(),
		edge.To("payment_method", PaymentMethod.Type).Unique(),
	}
}
