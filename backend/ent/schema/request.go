package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Request holds the schema definition for the Request entity.
type Request struct {
	ent.Schema
}

// Fields of the Request.
func (Request) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").NotEmpty().MaxLen(100),                        // Tipo de petición
		field.String("status").Default("pending"),                          // Estado de la petición (pending, approved, rejected)
		field.String("title").NotEmpty().MaxLen(100),                       // Título de la petición
		field.Text("description").NotEmpty(),                               // Descripción detallada de la petición
		field.Time("created_at").Default(time.Now),                         // Fecha de creación
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now), // Fecha de última actualización
	}
}

// Edges of the Request.
func (Request) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("requester", Users.Type),
		edge.To("receiver", Users.Type),
	}
}
