package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Configuration holds the schema definition for the Configuration entity.
type Configuration struct {
	ent.Schema
}

// Fields of the Configuration.
func (Configuration) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_registration_subjects"),
		field.Time("end_registration_subjects"),
		field.Bool("block_not_pay_inscription").Default(false),
		field.JSON("fee_dates", []time.Time{}).Optional(),       // Arreglo de fechas para las cuotas
		field.Int("number_fees").Default(0),                     // Cantidad de cuotas durante el ciclo
		field.Int("number_notes").Default(3),                    // Cantidad de notas a ingresar
		field.JSON("notes_Percentages", []float64{}).Optional(), // Porcentajes de las notas
	}
}

// Edges of the Configuration.
func (Configuration) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cycle", Cycle.Type).Unique(),
	}
}
