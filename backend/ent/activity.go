// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mocku/backend/ent/activity"
	"mocku/backend/ent/users"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Activity is the model entity for the Activity schema.
type Activity struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Action holds the value of the "action" field.
	Action string `json:"action,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ActivityQuery when eager-loading is set.
	Edges         ActivityEdges `json:"edges"`
	activity_user *int
	selectValues  sql.SelectValues
}

// ActivityEdges holds the relations/edges for other nodes in the graph.
type ActivityEdges struct {
	// User holds the value of the user edge.
	User *Users `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ActivityEdges) UserOrErr() (*Users, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: users.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Activity) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case activity.FieldID:
			values[i] = new(sql.NullInt64)
		case activity.FieldAction, activity.FieldDescription:
			values[i] = new(sql.NullString)
		case activity.FieldTimestamp:
			values[i] = new(sql.NullTime)
		case activity.ForeignKeys[0]: // activity_user
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Activity fields.
func (a *Activity) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case activity.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case activity.FieldAction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field action", values[i])
			} else if value.Valid {
				a.Action = value.String
			}
		case activity.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case activity.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				a.Timestamp = value.Time
			}
		case activity.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field activity_user", value)
			} else if value.Valid {
				a.activity_user = new(int)
				*a.activity_user = int(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Activity.
// This includes values selected through modifiers, order, etc.
func (a *Activity) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Activity entity.
func (a *Activity) QueryUser() *UsersQuery {
	return NewActivityClient(a.config).QueryUser(a)
}

// Update returns a builder for updating this Activity.
// Note that you need to call Activity.Unwrap() before calling this method if this Activity
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Activity) Update() *ActivityUpdateOne {
	return NewActivityClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Activity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Activity) Unwrap() *Activity {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Activity is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Activity) String() string {
	var builder strings.Builder
	builder.WriteString("Activity(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("action=")
	builder.WriteString(a.Action)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(a.Description)
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(a.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Activities is a parsable slice of Activity.
type Activities []*Activity
