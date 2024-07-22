// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"mocku/backend/ent/note"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Note is the model entity for the Note schema.
type Note struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Notes holds the value of the "notes" field.
	Notes []float64 `json:"notes,omitempty"`
	// Average holds the value of the "average" field.
	Average float32 `json:"average,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NoteQuery when eager-loading is set.
	Edges        NoteEdges `json:"edges"`
	selectValues sql.SelectValues
}

// NoteEdges holds the relations/edges for other nodes in the graph.
type NoteEdges struct {
	// Student holds the value of the student edge.
	Student []*Student `json:"student,omitempty"`
	// Subject holds the value of the subject edge.
	Subject []*Subject `json:"subject,omitempty"`
	// Cycle holds the value of the cycle edge.
	Cycle []*Cycle `json:"cycle,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// StudentOrErr returns the Student value or an error if the edge
// was not loaded in eager-loading.
func (e NoteEdges) StudentOrErr() ([]*Student, error) {
	if e.loadedTypes[0] {
		return e.Student, nil
	}
	return nil, &NotLoadedError{edge: "student"}
}

// SubjectOrErr returns the Subject value or an error if the edge
// was not loaded in eager-loading.
func (e NoteEdges) SubjectOrErr() ([]*Subject, error) {
	if e.loadedTypes[1] {
		return e.Subject, nil
	}
	return nil, &NotLoadedError{edge: "subject"}
}

// CycleOrErr returns the Cycle value or an error if the edge
// was not loaded in eager-loading.
func (e NoteEdges) CycleOrErr() ([]*Cycle, error) {
	if e.loadedTypes[2] {
		return e.Cycle, nil
	}
	return nil, &NotLoadedError{edge: "cycle"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Note) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case note.FieldNotes:
			values[i] = new([]byte)
		case note.FieldAverage:
			values[i] = new(sql.NullFloat64)
		case note.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Note fields.
func (n *Note) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case note.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case note.FieldNotes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field notes", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &n.Notes); err != nil {
					return fmt.Errorf("unmarshal field notes: %w", err)
				}
			}
		case note.FieldAverage:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field average", values[i])
			} else if value.Valid {
				n.Average = float32(value.Float64)
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Note.
// This includes values selected through modifiers, order, etc.
func (n *Note) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryStudent queries the "student" edge of the Note entity.
func (n *Note) QueryStudent() *StudentQuery {
	return NewNoteClient(n.config).QueryStudent(n)
}

// QuerySubject queries the "subject" edge of the Note entity.
func (n *Note) QuerySubject() *SubjectQuery {
	return NewNoteClient(n.config).QuerySubject(n)
}

// QueryCycle queries the "cycle" edge of the Note entity.
func (n *Note) QueryCycle() *CycleQuery {
	return NewNoteClient(n.config).QueryCycle(n)
}

// Update returns a builder for updating this Note.
// Note that you need to call Note.Unwrap() before calling this method if this Note
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Note) Update() *NoteUpdateOne {
	return NewNoteClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Note entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Note) Unwrap() *Note {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Note is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Note) String() string {
	var builder strings.Builder
	builder.WriteString("Note(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("notes=")
	builder.WriteString(fmt.Sprintf("%v", n.Notes))
	builder.WriteString(", ")
	builder.WriteString("average=")
	builder.WriteString(fmt.Sprintf("%v", n.Average))
	builder.WriteByte(')')
	return builder.String()
}

// Notes is a parsable slice of Note.
type Notes []*Note
