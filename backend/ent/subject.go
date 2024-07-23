// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"mocku/backend/ent/professor"
	"mocku/backend/ent/subject"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Subject is the model entity for the Subject schema.
type Subject struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreditUnits holds the value of the "credit_units" field.
	CreditUnits int `json:"credit_units,omitempty"`
	// Semester holds the value of the "semester" field.
	Semester int `json:"semester,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// PracticeHours holds the value of the "practice_hours" field.
	PracticeHours int `json:"practice_hours,omitempty"`
	// TheoryHours holds the value of the "theory_hours" field.
	TheoryHours int `json:"theory_hours,omitempty"`
	// LabHours holds the value of the "lab_hours" field.
	LabHours int `json:"lab_hours,omitempty"`
	// TotalHours holds the value of the "total_hours" field.
	TotalHours int `json:"total_hours,omitempty"`
	// ClassSchedule holds the value of the "class_schedule" field.
	ClassSchedule map[string][]string `json:"class_schedule,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubjectQuery when eager-loading is set.
	Edges             SubjectEdges `json:"edges"`
	subject_professor *int
	selectValues      sql.SelectValues
}

// SubjectEdges holds the relations/edges for other nodes in the graph.
type SubjectEdges struct {
	// Professor holds the value of the professor edge.
	Professor *Professor `json:"professor,omitempty"`
	// Career holds the value of the career edge.
	Career []*Careers `json:"career,omitempty"`
	// Notes holds the value of the notes edge.
	Notes []*Note `json:"notes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProfessorOrErr returns the Professor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubjectEdges) ProfessorOrErr() (*Professor, error) {
	if e.Professor != nil {
		return e.Professor, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: professor.Label}
	}
	return nil, &NotLoadedError{edge: "professor"}
}

// CareerOrErr returns the Career value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) CareerOrErr() ([]*Careers, error) {
	if e.loadedTypes[1] {
		return e.Career, nil
	}
	return nil, &NotLoadedError{edge: "career"}
}

// NotesOrErr returns the Notes value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) NotesOrErr() ([]*Note, error) {
	if e.loadedTypes[2] {
		return e.Notes, nil
	}
	return nil, &NotLoadedError{edge: "notes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subject) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subject.FieldClassSchedule:
			values[i] = new([]byte)
		case subject.FieldID, subject.FieldCreditUnits, subject.FieldSemester, subject.FieldPracticeHours, subject.FieldTheoryHours, subject.FieldLabHours, subject.FieldTotalHours:
			values[i] = new(sql.NullInt64)
		case subject.FieldName, subject.FieldDescription, subject.FieldCode:
			values[i] = new(sql.NullString)
		case subject.ForeignKeys[0]: // subject_professor
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subject fields.
func (s *Subject) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subject.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case subject.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case subject.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case subject.FieldCreditUnits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field credit_units", values[i])
			} else if value.Valid {
				s.CreditUnits = int(value.Int64)
			}
		case subject.FieldSemester:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field semester", values[i])
			} else if value.Valid {
				s.Semester = int(value.Int64)
			}
		case subject.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				s.Code = value.String
			}
		case subject.FieldPracticeHours:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field practice_hours", values[i])
			} else if value.Valid {
				s.PracticeHours = int(value.Int64)
			}
		case subject.FieldTheoryHours:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field theory_hours", values[i])
			} else if value.Valid {
				s.TheoryHours = int(value.Int64)
			}
		case subject.FieldLabHours:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lab_hours", values[i])
			} else if value.Valid {
				s.LabHours = int(value.Int64)
			}
		case subject.FieldTotalHours:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_hours", values[i])
			} else if value.Valid {
				s.TotalHours = int(value.Int64)
			}
		case subject.FieldClassSchedule:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field class_schedule", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.ClassSchedule); err != nil {
					return fmt.Errorf("unmarshal field class_schedule: %w", err)
				}
			}
		case subject.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field subject_professor", value)
			} else if value.Valid {
				s.subject_professor = new(int)
				*s.subject_professor = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Subject.
// This includes values selected through modifiers, order, etc.
func (s *Subject) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryProfessor queries the "professor" edge of the Subject entity.
func (s *Subject) QueryProfessor() *ProfessorQuery {
	return NewSubjectClient(s.config).QueryProfessor(s)
}

// QueryCareer queries the "career" edge of the Subject entity.
func (s *Subject) QueryCareer() *CareersQuery {
	return NewSubjectClient(s.config).QueryCareer(s)
}

// QueryNotes queries the "notes" edge of the Subject entity.
func (s *Subject) QueryNotes() *NoteQuery {
	return NewSubjectClient(s.config).QueryNotes(s)
}

// Update returns a builder for updating this Subject.
// Note that you need to call Subject.Unwrap() before calling this method if this Subject
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subject) Update() *SubjectUpdateOne {
	return NewSubjectClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Subject entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subject) Unwrap() *Subject {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subject is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subject) String() string {
	var builder strings.Builder
	builder.WriteString("Subject(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("credit_units=")
	builder.WriteString(fmt.Sprintf("%v", s.CreditUnits))
	builder.WriteString(", ")
	builder.WriteString("semester=")
	builder.WriteString(fmt.Sprintf("%v", s.Semester))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(s.Code)
	builder.WriteString(", ")
	builder.WriteString("practice_hours=")
	builder.WriteString(fmt.Sprintf("%v", s.PracticeHours))
	builder.WriteString(", ")
	builder.WriteString("theory_hours=")
	builder.WriteString(fmt.Sprintf("%v", s.TheoryHours))
	builder.WriteString(", ")
	builder.WriteString("lab_hours=")
	builder.WriteString(fmt.Sprintf("%v", s.LabHours))
	builder.WriteString(", ")
	builder.WriteString("total_hours=")
	builder.WriteString(fmt.Sprintf("%v", s.TotalHours))
	builder.WriteString(", ")
	builder.WriteString("class_schedule=")
	builder.WriteString(fmt.Sprintf("%v", s.ClassSchedule))
	builder.WriteByte(')')
	return builder.String()
}

// Subjects is a parsable slice of Subject.
type Subjects []*Subject
