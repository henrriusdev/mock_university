package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().MaxLen(100),
		field.Text("message").NotEmpty(),
		field.String("status").Default("unread"), // unread, read
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("recipient", Users.Type).Unique(),
	}
}
