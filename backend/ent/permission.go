// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mocku/backend/ent/permission"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Permission is the model entity for the Permission schema.
type Permission struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Read holds the value of the "read" field.
	Read bool `json:"read,omitempty"`
	// Create holds the value of the "create" field.
	Create bool `json:"create,omitempty"`
	// Modify holds the value of the "modify" field.
	Modify bool `json:"modify,omitempty"`
	// Delete holds the value of the "delete" field.
	Delete bool `json:"delete,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PermissionQuery when eager-loading is set.
	Edges        PermissionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PermissionEdges holds the relations/edges for other nodes in the graph.
type PermissionEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// Module holds the value of the module edge.
	Module []*Module `json:"module,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// ModuleOrErr returns the Module value or an error if the edge
// was not loaded in eager-loading.
func (e PermissionEdges) ModuleOrErr() ([]*Module, error) {
	if e.loadedTypes[1] {
		return e.Module, nil
	}
	return nil, &NotLoadedError{edge: "module"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Permission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case permission.FieldRead, permission.FieldCreate, permission.FieldModify, permission.FieldDelete:
			values[i] = new(sql.NullBool)
		case permission.FieldID:
			values[i] = new(sql.NullInt64)
		case permission.FieldName, permission.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Permission fields.
func (pe *Permission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case permission.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case permission.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pe.Description = value.String
			}
		case permission.FieldRead:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field read", values[i])
			} else if value.Valid {
				pe.Read = value.Bool
			}
		case permission.FieldCreate:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field create", values[i])
			} else if value.Valid {
				pe.Create = value.Bool
			}
		case permission.FieldModify:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field modify", values[i])
			} else if value.Valid {
				pe.Modify = value.Bool
			}
		case permission.FieldDelete:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				pe.Delete = value.Bool
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Permission.
// This includes values selected through modifiers, order, etc.
func (pe *Permission) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// QueryRoles queries the "roles" edge of the Permission entity.
func (pe *Permission) QueryRoles() *RoleQuery {
	return NewPermissionClient(pe.config).QueryRoles(pe)
}

// QueryModule queries the "module" edge of the Permission entity.
func (pe *Permission) QueryModule() *ModuleQuery {
	return NewPermissionClient(pe.config).QueryModule(pe)
}

// Update returns a builder for updating this Permission.
// Note that you need to call Permission.Unwrap() before calling this method if this Permission
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Permission) Update() *PermissionUpdateOne {
	return NewPermissionClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Permission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Permission) Unwrap() *Permission {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Permission is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Permission) String() string {
	var builder strings.Builder
	builder.WriteString("Permission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pe.Description)
	builder.WriteString(", ")
	builder.WriteString("read=")
	builder.WriteString(fmt.Sprintf("%v", pe.Read))
	builder.WriteString(", ")
	builder.WriteString("create=")
	builder.WriteString(fmt.Sprintf("%v", pe.Create))
	builder.WriteString(", ")
	builder.WriteString("modify=")
	builder.WriteString(fmt.Sprintf("%v", pe.Modify))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", pe.Delete))
	builder.WriteByte(')')
	return builder.String()
}

// Permissions is a parsable slice of Permission.
type Permissions []*Permission
