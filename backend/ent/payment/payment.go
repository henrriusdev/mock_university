// Code generated by ent, DO NOT EDIT.

package payment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the payment type in the database.
	Label = "payment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldReference holds the string denoting the reference field in the database.
	FieldReference = "reference"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldFeeNumber holds the string denoting the fee_number field in the database.
	FieldFeeNumber = "fee_number"
	// EdgeStudent holds the string denoting the student edge name in mutations.
	EdgeStudent = "student"
	// EdgeCycle holds the string denoting the cycle edge name in mutations.
	EdgeCycle = "cycle"
	// EdgePaymentMethod holds the string denoting the payment_method edge name in mutations.
	EdgePaymentMethod = "payment_method"
	// Table holds the table name of the payment in the database.
	Table = "payments"
	// StudentTable is the table that holds the student relation/edge.
	StudentTable = "payments"
	// StudentInverseTable is the table name for the Student entity.
	// It exists in this package in order to avoid circular dependency with the "student" package.
	StudentInverseTable = "students"
	// StudentColumn is the table column denoting the student relation/edge.
	StudentColumn = "payment_student"
	// CycleTable is the table that holds the cycle relation/edge.
	CycleTable = "payments"
	// CycleInverseTable is the table name for the Cycle entity.
	// It exists in this package in order to avoid circular dependency with the "cycle" package.
	CycleInverseTable = "cycles"
	// CycleColumn is the table column denoting the cycle relation/edge.
	CycleColumn = "payment_cycle"
	// PaymentMethodTable is the table that holds the payment_method relation/edge.
	PaymentMethodTable = "payments"
	// PaymentMethodInverseTable is the table name for the PaymentMethod entity.
	// It exists in this package in order to avoid circular dependency with the "paymentmethod" package.
	PaymentMethodInverseTable = "payment_methods"
	// PaymentMethodColumn is the table column denoting the payment_method relation/edge.
	PaymentMethodColumn = "payment_payment_method"
)

// Columns holds all SQL columns for payment fields.
var Columns = []string{
	FieldID,
	FieldReference,
	FieldDate,
	FieldAmount,
	FieldDescription,
	FieldFeeNumber,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "payments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"payment_student",
	"payment_cycle",
	"payment_payment_method",
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
	// ReferenceValidator is a validator for the "reference" field. It is called by the builders before save.
	ReferenceValidator func(string) error
	// AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	AmountValidator func(float64) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// FeeNumberValidator is a validator for the "fee_number" field. It is called by the builders before save.
	FeeNumberValidator func(int) error
)

// OrderOption defines the ordering options for the Payment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByReference orders the results by the reference field.
func ByReference(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReference, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByFeeNumber orders the results by the fee_number field.
func ByFeeNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFeeNumber, opts...).ToFunc()
}

// ByStudentField orders the results by student field.
func ByStudentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStudentStep(), sql.OrderByField(field, opts...))
	}
}

// ByCycleField orders the results by cycle field.
func ByCycleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCycleStep(), sql.OrderByField(field, opts...))
	}
}

// ByPaymentMethodField orders the results by payment_method field.
func ByPaymentMethodField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPaymentMethodStep(), sql.OrderByField(field, opts...))
	}
}
func newStudentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StudentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, StudentTable, StudentColumn),
	)
}
func newCycleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CycleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CycleTable, CycleColumn),
	)
}
func newPaymentMethodStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PaymentMethodInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, PaymentMethodTable, PaymentMethodColumn),
	)
}
