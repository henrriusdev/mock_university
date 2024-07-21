package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PaymentMethod holds the schema definition for the PaymentMethod entity.
type PaymentMethod struct {
	ent.Schema
}

// Fields of the PaymentMethod.
func (PaymentMethod) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MaxLen(100),
	}
}

// Edges of the PaymentMethod.
func (PaymentMethod) Edges() []ent.Edge {
	return nil
}
