package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Professor holds the schema definition for the Professor entity.
type Professor struct {
	ent.Schema
}

// Fields of the Professor.
func (Professor) Fields() []ent.Field {
	return []ent.Field{
		field.String("identity_card").NotEmpty().MaxLen(20).Unique(),
		field.Time("birth_date"),
		field.String("phone").NotEmpty().MaxLen(20),
		field.String("address").NotEmpty().MaxLen(255),
	}
}

// Edges of the Professor.
func (Professor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", Users.Type).Unique(),
		edge.From("boss", Professor.Type).Ref("subordinates").Unique(),
		edge.To("subordinates", Professor.Type),
		edge.From("subjects", Subject.Type).Ref("professor"),
		edge.From("careers", Careers.Type).Ref("leader").Unique(),
	}
}
