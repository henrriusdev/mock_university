package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty().MaxLen(30).Unique(),
		field.String("password").Sensitive().NotEmpty().MaxLen(30),
		field.String("email").NotEmpty().MaxLen(100).Unique(),
		field.String("name").NotEmpty().MaxLen(100),
		field.String("avatar").MaxLen(255),
		field.Bool("is_active").Default(true),
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}
