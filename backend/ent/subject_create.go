// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mocku/backend/ent/careers"
	"mocku/backend/ent/note"
	"mocku/backend/ent/professor"
	"mocku/backend/ent/subject"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubjectCreate is the builder for creating a Subject entity.
type SubjectCreate struct {
	config
	mutation *SubjectMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *SubjectCreate) SetName(s string) *SubjectCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDescription sets the "description" field.
func (sc *SubjectCreate) SetDescription(s string) *SubjectCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetCreditUnits sets the "credit_units" field.
func (sc *SubjectCreate) SetCreditUnits(i int) *SubjectCreate {
	sc.mutation.SetCreditUnits(i)
	return sc
}

// SetSemester sets the "semester" field.
func (sc *SubjectCreate) SetSemester(i int) *SubjectCreate {
	sc.mutation.SetSemester(i)
	return sc
}

// SetCode sets the "code" field.
func (sc *SubjectCreate) SetCode(s string) *SubjectCreate {
	sc.mutation.SetCode(s)
	return sc
}

// SetPracticeHours sets the "practice_hours" field.
func (sc *SubjectCreate) SetPracticeHours(i int) *SubjectCreate {
	sc.mutation.SetPracticeHours(i)
	return sc
}

// SetTheoryHours sets the "theory_hours" field.
func (sc *SubjectCreate) SetTheoryHours(i int) *SubjectCreate {
	sc.mutation.SetTheoryHours(i)
	return sc
}

// SetLabHours sets the "lab_hours" field.
func (sc *SubjectCreate) SetLabHours(i int) *SubjectCreate {
	sc.mutation.SetLabHours(i)
	return sc
}

// SetTotalHours sets the "total_hours" field.
func (sc *SubjectCreate) SetTotalHours(i int) *SubjectCreate {
	sc.mutation.SetTotalHours(i)
	return sc
}

// SetClassSchedule sets the "class_schedule" field.
func (sc *SubjectCreate) SetClassSchedule(m map[string][]string) *SubjectCreate {
	sc.mutation.SetClassSchedule(m)
	return sc
}

// AddProfessorIDs adds the "professor" edge to the Professor entity by IDs.
func (sc *SubjectCreate) AddProfessorIDs(ids ...int) *SubjectCreate {
	sc.mutation.AddProfessorIDs(ids...)
	return sc
}

// AddProfessor adds the "professor" edges to the Professor entity.
func (sc *SubjectCreate) AddProfessor(p ...*Professor) *SubjectCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddProfessorIDs(ids...)
}

// AddCareerIDs adds the "career" edge to the Careers entity by IDs.
func (sc *SubjectCreate) AddCareerIDs(ids ...int) *SubjectCreate {
	sc.mutation.AddCareerIDs(ids...)
	return sc
}

// AddCareer adds the "career" edges to the Careers entity.
func (sc *SubjectCreate) AddCareer(c ...*Careers) *SubjectCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddCareerIDs(ids...)
}

// AddNoteIDs adds the "notes" edge to the Note entity by IDs.
func (sc *SubjectCreate) AddNoteIDs(ids ...int) *SubjectCreate {
	sc.mutation.AddNoteIDs(ids...)
	return sc
}

// AddNotes adds the "notes" edges to the Note entity.
func (sc *SubjectCreate) AddNotes(n ...*Note) *SubjectCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return sc.AddNoteIDs(ids...)
}

// Mutation returns the SubjectMutation object of the builder.
func (sc *SubjectCreate) Mutation() *SubjectMutation {
	return sc.mutation
}

// Save creates the Subject in the database.
func (sc *SubjectCreate) Save(ctx context.Context) (*Subject, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubjectCreate) SaveX(ctx context.Context) *Subject {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubjectCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubjectCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubjectCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Subject.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := subject.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Subject.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Subject.description"`)}
	}
	if v, ok := sc.mutation.Description(); ok {
		if err := subject.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Subject.description": %w`, err)}
		}
	}
	if _, ok := sc.mutation.CreditUnits(); !ok {
		return &ValidationError{Name: "credit_units", err: errors.New(`ent: missing required field "Subject.credit_units"`)}
	}
	if v, ok := sc.mutation.CreditUnits(); ok {
		if err := subject.CreditUnitsValidator(v); err != nil {
			return &ValidationError{Name: "credit_units", err: fmt.Errorf(`ent: validator failed for field "Subject.credit_units": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Semester(); !ok {
		return &ValidationError{Name: "semester", err: errors.New(`ent: missing required field "Subject.semester"`)}
	}
	if v, ok := sc.mutation.Semester(); ok {
		if err := subject.SemesterValidator(v); err != nil {
			return &ValidationError{Name: "semester", err: fmt.Errorf(`ent: validator failed for field "Subject.semester": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Subject.code"`)}
	}
	if v, ok := sc.mutation.Code(); ok {
		if err := subject.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Subject.code": %w`, err)}
		}
	}
	if _, ok := sc.mutation.PracticeHours(); !ok {
		return &ValidationError{Name: "practice_hours", err: errors.New(`ent: missing required field "Subject.practice_hours"`)}
	}
	if v, ok := sc.mutation.PracticeHours(); ok {
		if err := subject.PracticeHoursValidator(v); err != nil {
			return &ValidationError{Name: "practice_hours", err: fmt.Errorf(`ent: validator failed for field "Subject.practice_hours": %w`, err)}
		}
	}
	if _, ok := sc.mutation.TheoryHours(); !ok {
		return &ValidationError{Name: "theory_hours", err: errors.New(`ent: missing required field "Subject.theory_hours"`)}
	}
	if v, ok := sc.mutation.TheoryHours(); ok {
		if err := subject.TheoryHoursValidator(v); err != nil {
			return &ValidationError{Name: "theory_hours", err: fmt.Errorf(`ent: validator failed for field "Subject.theory_hours": %w`, err)}
		}
	}
	if _, ok := sc.mutation.LabHours(); !ok {
		return &ValidationError{Name: "lab_hours", err: errors.New(`ent: missing required field "Subject.lab_hours"`)}
	}
	if v, ok := sc.mutation.LabHours(); ok {
		if err := subject.LabHoursValidator(v); err != nil {
			return &ValidationError{Name: "lab_hours", err: fmt.Errorf(`ent: validator failed for field "Subject.lab_hours": %w`, err)}
		}
	}
	if _, ok := sc.mutation.TotalHours(); !ok {
		return &ValidationError{Name: "total_hours", err: errors.New(`ent: missing required field "Subject.total_hours"`)}
	}
	if v, ok := sc.mutation.TotalHours(); ok {
		if err := subject.TotalHoursValidator(v); err != nil {
			return &ValidationError{Name: "total_hours", err: fmt.Errorf(`ent: validator failed for field "Subject.total_hours": %w`, err)}
		}
	}
	return nil
}

func (sc *SubjectCreate) sqlSave(ctx context.Context) (*Subject, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubjectCreate) createSpec() (*Subject, *sqlgraph.CreateSpec) {
	var (
		_node = &Subject{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(subject.Table, sqlgraph.NewFieldSpec(subject.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(subject.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(subject.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := sc.mutation.CreditUnits(); ok {
		_spec.SetField(subject.FieldCreditUnits, field.TypeInt, value)
		_node.CreditUnits = value
	}
	if value, ok := sc.mutation.Semester(); ok {
		_spec.SetField(subject.FieldSemester, field.TypeInt, value)
		_node.Semester = value
	}
	if value, ok := sc.mutation.Code(); ok {
		_spec.SetField(subject.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := sc.mutation.PracticeHours(); ok {
		_spec.SetField(subject.FieldPracticeHours, field.TypeInt, value)
		_node.PracticeHours = value
	}
	if value, ok := sc.mutation.TheoryHours(); ok {
		_spec.SetField(subject.FieldTheoryHours, field.TypeInt, value)
		_node.TheoryHours = value
	}
	if value, ok := sc.mutation.LabHours(); ok {
		_spec.SetField(subject.FieldLabHours, field.TypeInt, value)
		_node.LabHours = value
	}
	if value, ok := sc.mutation.TotalHours(); ok {
		_spec.SetField(subject.FieldTotalHours, field.TypeInt, value)
		_node.TotalHours = value
	}
	if value, ok := sc.mutation.ClassSchedule(); ok {
		_spec.SetField(subject.FieldClassSchedule, field.TypeJSON, value)
		_node.ClassSchedule = value
	}
	if nodes := sc.mutation.ProfessorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   subject.ProfessorTable,
			Columns: subject.ProfessorPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(professor.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.CareerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.CareerTable,
			Columns: []string{subject.CareerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(careers.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.NotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subject.NotesTable,
			Columns: subject.NotesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubjectCreateBulk is the builder for creating many Subject entities in bulk.
type SubjectCreateBulk struct {
	config
	err      error
	builders []*SubjectCreate
}

// Save creates the Subject entities in the database.
func (scb *SubjectCreateBulk) Save(ctx context.Context) ([]*Subject, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subject, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubjectMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubjectCreateBulk) SaveX(ctx context.Context) []*Subject {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubjectCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubjectCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
