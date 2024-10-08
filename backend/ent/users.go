// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mocku/backend/ent/role"
	"mocku/backend/ent/users"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Users is the model entity for the Users schema.
type Users struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UsersQuery when eager-loading is set.
	Edges        UsersEdges `json:"edges"`
	users_role   *int
	selectValues sql.SelectValues
}

// UsersEdges holds the relations/edges for other nodes in the graph.
type UsersEdges struct {
	// Role holds the value of the role edge.
	Role *Role `json:"role,omitempty"`
	// RequestsMade holds the value of the requests_made edge.
	RequestsMade []*Request `json:"requests_made,omitempty"`
	// RequestsReceived holds the value of the requests_received edge.
	RequestsReceived []*Request `json:"requests_received,omitempty"`
	// Blog holds the value of the blog edge.
	Blog []*Blog `json:"blog,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// Activity holds the value of the activity edge.
	Activity []*Activity `json:"activity,omitempty"`
	// Students holds the value of the students edge.
	Students []*Student `json:"students,omitempty"`
	// Professor holds the value of the professor edge.
	Professor []*Professor `json:"professor,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UsersEdges) RoleOrErr() (*Role, error) {
	if e.Role != nil {
		return e.Role, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: role.Label}
	}
	return nil, &NotLoadedError{edge: "role"}
}

// RequestsMadeOrErr returns the RequestsMade value or an error if the edge
// was not loaded in eager-loading.
func (e UsersEdges) RequestsMadeOrErr() ([]*Request, error) {
	if e.loadedTypes[1] {
		return e.RequestsMade, nil
	}
	return nil, &NotLoadedError{edge: "requests_made"}
}

// RequestsReceivedOrErr returns the RequestsReceived value or an error if the edge
// was not loaded in eager-loading.
func (e UsersEdges) RequestsReceivedOrErr() ([]*Request, error) {
	if e.loadedTypes[2] {
		return e.RequestsReceived, nil
	}
	return nil, &NotLoadedError{edge: "requests_received"}
}

// BlogOrErr returns the Blog value or an error if the edge
// was not loaded in eager-loading.
func (e UsersEdges) BlogOrErr() ([]*Blog, error) {
	if e.loadedTypes[3] {
		return e.Blog, nil
	}
	return nil, &NotLoadedError{edge: "blog"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e UsersEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[4] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// ActivityOrErr returns the Activity value or an error if the edge
// was not loaded in eager-loading.
func (e UsersEdges) ActivityOrErr() ([]*Activity, error) {
	if e.loadedTypes[5] {
		return e.Activity, nil
	}
	return nil, &NotLoadedError{edge: "activity"}
}

// StudentsOrErr returns the Students value or an error if the edge
// was not loaded in eager-loading.
func (e UsersEdges) StudentsOrErr() ([]*Student, error) {
	if e.loadedTypes[6] {
		return e.Students, nil
	}
	return nil, &NotLoadedError{edge: "students"}
}

// ProfessorOrErr returns the Professor value or an error if the edge
// was not loaded in eager-loading.
func (e UsersEdges) ProfessorOrErr() ([]*Professor, error) {
	if e.loadedTypes[7] {
		return e.Professor, nil
	}
	return nil, &NotLoadedError{edge: "professor"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Users) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case users.FieldIsActive:
			values[i] = new(sql.NullBool)
		case users.FieldID:
			values[i] = new(sql.NullInt64)
		case users.FieldUsername, users.FieldPassword, users.FieldEmail, users.FieldName, users.FieldAvatar:
			values[i] = new(sql.NullString)
		case users.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case users.ForeignKeys[0]: // users_role
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Users fields.
func (u *Users) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case users.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case users.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case users.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case users.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case users.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case users.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case users.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				u.IsActive = value.Bool
			}
		case users.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case users.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field users_role", value)
			} else if value.Valid {
				u.users_role = new(int)
				*u.users_role = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Users.
// This includes values selected through modifiers, order, etc.
func (u *Users) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryRole queries the "role" edge of the Users entity.
func (u *Users) QueryRole() *RoleQuery {
	return NewUsersClient(u.config).QueryRole(u)
}

// QueryRequestsMade queries the "requests_made" edge of the Users entity.
func (u *Users) QueryRequestsMade() *RequestQuery {
	return NewUsersClient(u.config).QueryRequestsMade(u)
}

// QueryRequestsReceived queries the "requests_received" edge of the Users entity.
func (u *Users) QueryRequestsReceived() *RequestQuery {
	return NewUsersClient(u.config).QueryRequestsReceived(u)
}

// QueryBlog queries the "blog" edge of the Users entity.
func (u *Users) QueryBlog() *BlogQuery {
	return NewUsersClient(u.config).QueryBlog(u)
}

// QueryNotifications queries the "notifications" edge of the Users entity.
func (u *Users) QueryNotifications() *NotificationQuery {
	return NewUsersClient(u.config).QueryNotifications(u)
}

// QueryActivity queries the "activity" edge of the Users entity.
func (u *Users) QueryActivity() *ActivityQuery {
	return NewUsersClient(u.config).QueryActivity(u)
}

// QueryStudents queries the "students" edge of the Users entity.
func (u *Users) QueryStudents() *StudentQuery {
	return NewUsersClient(u.config).QueryStudents(u)
}

// QueryProfessor queries the "professor" edge of the Users entity.
func (u *Users) QueryProfessor() *ProfessorQuery {
	return NewUsersClient(u.config).QueryProfessor(u)
}

// Update returns a builder for updating this Users.
// Note that you need to call Users.Unwrap() before calling this method if this Users
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Users) Update() *UsersUpdateOne {
	return NewUsersClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the Users entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Users) Unwrap() *Users {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Users is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Users) String() string {
	var builder strings.Builder
	builder.WriteString("Users(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", u.IsActive))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UsersSlice is a parsable slice of Users.
type UsersSlice []*Users
