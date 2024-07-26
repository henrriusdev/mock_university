// Code generated by ent, DO NOT EDIT.

package subject

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the subject type in the database.
	Label = "subject"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCreditUnits holds the string denoting the credit_units field in the database.
	FieldCreditUnits = "credit_units"
	// FieldSemester holds the string denoting the semester field in the database.
	FieldSemester = "semester"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldPracticeHours holds the string denoting the practice_hours field in the database.
	FieldPracticeHours = "practice_hours"
	// FieldTheoryHours holds the string denoting the theory_hours field in the database.
	FieldTheoryHours = "theory_hours"
	// FieldLabHours holds the string denoting the lab_hours field in the database.
	FieldLabHours = "lab_hours"
	// FieldTotalHours holds the string denoting the total_hours field in the database.
	FieldTotalHours = "total_hours"
	// FieldClassSchedule holds the string denoting the class_schedule field in the database.
	FieldClassSchedule = "class_schedule"
	// EdgeProfessor holds the string denoting the professor edge name in mutations.
	EdgeProfessor = "professor"
	// EdgeCareer holds the string denoting the career edge name in mutations.
	EdgeCareer = "career"
	// EdgeNotes holds the string denoting the notes edge name in mutations.
	EdgeNotes = "notes"
	// Table holds the table name of the subject in the database.
	Table = "subjects"
	// ProfessorTable is the table that holds the professor relation/edge.
	ProfessorTable = "subjects"
	// ProfessorInverseTable is the table name for the Professor entity.
	// It exists in this package in order to avoid circular dependency with the "professor" package.
	ProfessorInverseTable = "professors"
	// ProfessorColumn is the table column denoting the professor relation/edge.
	ProfessorColumn = "subject_professor"
	// CareerTable is the table that holds the career relation/edge. The primary key declared below.
	CareerTable = "subject_career"
	// CareerInverseTable is the table name for the Careers entity.
	// It exists in this package in order to avoid circular dependency with the "careers" package.
	CareerInverseTable = "careers"
	// NotesTable is the table that holds the notes relation/edge.
	NotesTable = "notes"
	// NotesInverseTable is the table name for the Note entity.
	// It exists in this package in order to avoid circular dependency with the "note" package.
	NotesInverseTable = "notes"
	// NotesColumn is the table column denoting the notes relation/edge.
	NotesColumn = "note_subject"
)

// Columns holds all SQL columns for subject fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldCreditUnits,
	FieldSemester,
	FieldCode,
	FieldPracticeHours,
	FieldTheoryHours,
	FieldLabHours,
	FieldTotalHours,
	FieldClassSchedule,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "subjects"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"subject_professor",
}

var (
	// CareerPrimaryKey and CareerColumn2 are the table columns denoting the
	// primary key for the career relation (M2M).
	CareerPrimaryKey = []string{"subject_id", "careers_id"}
)

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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// CreditUnitsValidator is a validator for the "credit_units" field. It is called by the builders before save.
	CreditUnitsValidator func(int) error
	// SemesterValidator is a validator for the "semester" field. It is called by the builders before save.
	SemesterValidator func(int) error
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// PracticeHoursValidator is a validator for the "practice_hours" field. It is called by the builders before save.
	PracticeHoursValidator func(int) error
	// TheoryHoursValidator is a validator for the "theory_hours" field. It is called by the builders before save.
	TheoryHoursValidator func(int) error
	// LabHoursValidator is a validator for the "lab_hours" field. It is called by the builders before save.
	LabHoursValidator func(int) error
	// TotalHoursValidator is a validator for the "total_hours" field. It is called by the builders before save.
	TotalHoursValidator func(int) error
)

// OrderOption defines the ordering options for the Subject queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCreditUnits orders the results by the credit_units field.
func ByCreditUnits(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreditUnits, opts...).ToFunc()
}

// BySemester orders the results by the semester field.
func BySemester(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSemester, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByPracticeHours orders the results by the practice_hours field.
func ByPracticeHours(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPracticeHours, opts...).ToFunc()
}

// ByTheoryHours orders the results by the theory_hours field.
func ByTheoryHours(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTheoryHours, opts...).ToFunc()
}

// ByLabHours orders the results by the lab_hours field.
func ByLabHours(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLabHours, opts...).ToFunc()
}

// ByTotalHours orders the results by the total_hours field.
func ByTotalHours(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalHours, opts...).ToFunc()
}

// ByProfessorField orders the results by professor field.
func ByProfessorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProfessorStep(), sql.OrderByField(field, opts...))
	}
}

// ByCareerCount orders the results by career count.
func ByCareerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCareerStep(), opts...)
	}
}

// ByCareer orders the results by career terms.
func ByCareer(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCareerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNotesCount orders the results by notes count.
func ByNotesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotesStep(), opts...)
	}
}

// ByNotes orders the results by notes terms.
func ByNotes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProfessorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProfessorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ProfessorTable, ProfessorColumn),
	)
}
func newCareerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CareerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, CareerTable, CareerPrimaryKey...),
	)
}
func newNotesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, NotesTable, NotesColumn),
	)
}
