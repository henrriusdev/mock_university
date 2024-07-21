package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct {
	ent.Schema
}

// Fields of the Blog.
func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.Text("text").NotEmpty(),
		field.Time("created_at").Immutable(),
		field.Time("updated_at").Optional(),
		field.Bool("published").Default(false),
		field.Strings("tags").Optional(),
		field.Strings("categories").Optional(),
		field.Int("view_count").Default(0),
		field.Strings("images").Optional(),
		field.String("slug").Unique(),
	}
}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", Users.Type).Unique(),
	}
}
