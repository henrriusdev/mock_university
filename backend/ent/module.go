// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mocku/backend/ent/module"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Module is the model entity for the Module schema.
type Module struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name              string `json:"name,omitempty"`
	permission_module *int
	selectValues      sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Module) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case module.FieldID:
			values[i] = new(sql.NullInt64)
		case module.FieldName:
			values[i] = new(sql.NullString)
		case module.ForeignKeys[0]: // permission_module
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Module fields.
func (m *Module) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case module.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case module.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case module.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field permission_module", value)
			} else if value.Valid {
				m.permission_module = new(int)
				*m.permission_module = int(value.Int64)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Module.
// This includes values selected through modifiers, order, etc.
func (m *Module) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// Update returns a builder for updating this Module.
// Note that you need to call Module.Unwrap() before calling this method if this Module
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Module) Update() *ModuleUpdateOne {
	return NewModuleClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Module entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Module) Unwrap() *Module {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Module is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Module) String() string {
	var builder strings.Builder
	builder.WriteString("Module(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Modules is a parsable slice of Module.
type Modules []*Module
