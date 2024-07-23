// Code generated by ent, DO NOT EDIT.

package activity

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the activity type in the database.
	Label = "activity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the activity in the database.
	Table = "activities"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "activities"
	// UserInverseTable is the table name for the Users entity.
	// It exists in this package in order to avoid circular dependency with the "users" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "activity_user"
)

// Columns holds all SQL columns for activity fields.
var Columns = []string{
	FieldID,
	FieldAction,
	FieldDescription,
	FieldTimestamp,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "activities"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"activity_user",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// ActionValidator is a validator for the "action" field. It is called by the builders before save.
	ActionValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultTimestamp holds the default value on creation for the "timestamp" field.
	DefaultTimestamp func() time.Time
)

// OrderOption defines the ordering options for the Activity queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAction orders the results by the action field.
func ByAction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAction, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByTimestamp orders the results by the timestamp field.
func ByTimestamp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimestamp, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
