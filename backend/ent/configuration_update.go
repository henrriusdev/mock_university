// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mocku/backend/ent/configuration"
	"mocku/backend/ent/cycle"
	"mocku/backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// ConfigurationUpdate is the builder for updating Configuration entities.
type ConfigurationUpdate struct {
	config
	hooks    []Hook
	mutation *ConfigurationMutation
}

// Where appends a list predicates to the ConfigurationUpdate builder.
func (cu *ConfigurationUpdate) Where(ps ...predicate.Configuration) *ConfigurationUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetStartRegistrationSubjects sets the "start_registration_subjects" field.
func (cu *ConfigurationUpdate) SetStartRegistrationSubjects(t time.Time) *ConfigurationUpdate {
	cu.mutation.SetStartRegistrationSubjects(t)
	return cu
}

// SetNillableStartRegistrationSubjects sets the "start_registration_subjects" field if the given value is not nil.
func (cu *ConfigurationUpdate) SetNillableStartRegistrationSubjects(t *time.Time) *ConfigurationUpdate {
	if t != nil {
		cu.SetStartRegistrationSubjects(*t)
	}
	return cu
}

// SetEndRegistrationSubjects sets the "end_registration_subjects" field.
func (cu *ConfigurationUpdate) SetEndRegistrationSubjects(t time.Time) *ConfigurationUpdate {
	cu.mutation.SetEndRegistrationSubjects(t)
	return cu
}

// SetNillableEndRegistrationSubjects sets the "end_registration_subjects" field if the given value is not nil.
func (cu *ConfigurationUpdate) SetNillableEndRegistrationSubjects(t *time.Time) *ConfigurationUpdate {
	if t != nil {
		cu.SetEndRegistrationSubjects(*t)
	}
	return cu
}

// SetBlockNotPayInscription sets the "block_not_pay_inscription" field.
func (cu *ConfigurationUpdate) SetBlockNotPayInscription(b bool) *ConfigurationUpdate {
	cu.mutation.SetBlockNotPayInscription(b)
	return cu
}

// SetNillableBlockNotPayInscription sets the "block_not_pay_inscription" field if the given value is not nil.
func (cu *ConfigurationUpdate) SetNillableBlockNotPayInscription(b *bool) *ConfigurationUpdate {
	if b != nil {
		cu.SetBlockNotPayInscription(*b)
	}
	return cu
}

// SetFeeDates sets the "fee_dates" field.
func (cu *ConfigurationUpdate) SetFeeDates(t []time.Time) *ConfigurationUpdate {
	cu.mutation.SetFeeDates(t)
	return cu
}

// AppendFeeDates appends t to the "fee_dates" field.
func (cu *ConfigurationUpdate) AppendFeeDates(t []time.Time) *ConfigurationUpdate {
	cu.mutation.AppendFeeDates(t)
	return cu
}

// ClearFeeDates clears the value of the "fee_dates" field.
func (cu *ConfigurationUpdate) ClearFeeDates() *ConfigurationUpdate {
	cu.mutation.ClearFeeDates()
	return cu
}

// SetNumberFees sets the "number_fees" field.
func (cu *ConfigurationUpdate) SetNumberFees(i int) *ConfigurationUpdate {
	cu.mutation.ResetNumberFees()
	cu.mutation.SetNumberFees(i)
	return cu
}

// SetNillableNumberFees sets the "number_fees" field if the given value is not nil.
func (cu *ConfigurationUpdate) SetNillableNumberFees(i *int) *ConfigurationUpdate {
	if i != nil {
		cu.SetNumberFees(*i)
	}
	return cu
}

// AddNumberFees adds i to the "number_fees" field.
func (cu *ConfigurationUpdate) AddNumberFees(i int) *ConfigurationUpdate {
	cu.mutation.AddNumberFees(i)
	return cu
}

// SetNumberNotes sets the "number_notes" field.
func (cu *ConfigurationUpdate) SetNumberNotes(i int) *ConfigurationUpdate {
	cu.mutation.ResetNumberNotes()
	cu.mutation.SetNumberNotes(i)
	return cu
}

// SetNillableNumberNotes sets the "number_notes" field if the given value is not nil.
func (cu *ConfigurationUpdate) SetNillableNumberNotes(i *int) *ConfigurationUpdate {
	if i != nil {
		cu.SetNumberNotes(*i)
	}
	return cu
}

// AddNumberNotes adds i to the "number_notes" field.
func (cu *ConfigurationUpdate) AddNumberNotes(i int) *ConfigurationUpdate {
	cu.mutation.AddNumberNotes(i)
	return cu
}

// SetNotesPercentages sets the "notes_Percentages" field.
func (cu *ConfigurationUpdate) SetNotesPercentages(f []float64) *ConfigurationUpdate {
	cu.mutation.SetNotesPercentages(f)
	return cu
}

// AppendNotesPercentages appends f to the "notes_Percentages" field.
func (cu *ConfigurationUpdate) AppendNotesPercentages(f []float64) *ConfigurationUpdate {
	cu.mutation.AppendNotesPercentages(f)
	return cu
}

// ClearNotesPercentages clears the value of the "notes_Percentages" field.
func (cu *ConfigurationUpdate) ClearNotesPercentages() *ConfigurationUpdate {
	cu.mutation.ClearNotesPercentages()
	return cu
}

// SetCycleID sets the "cycle" edge to the Cycle entity by ID.
func (cu *ConfigurationUpdate) SetCycleID(id int) *ConfigurationUpdate {
	cu.mutation.SetCycleID(id)
	return cu
}

// SetNillableCycleID sets the "cycle" edge to the Cycle entity by ID if the given value is not nil.
func (cu *ConfigurationUpdate) SetNillableCycleID(id *int) *ConfigurationUpdate {
	if id != nil {
		cu = cu.SetCycleID(*id)
	}
	return cu
}

// SetCycle sets the "cycle" edge to the Cycle entity.
func (cu *ConfigurationUpdate) SetCycle(c *Cycle) *ConfigurationUpdate {
	return cu.SetCycleID(c.ID)
}

// Mutation returns the ConfigurationMutation object of the builder.
func (cu *ConfigurationUpdate) Mutation() *ConfigurationMutation {
	return cu.mutation
}

// ClearCycle clears the "cycle" edge to the Cycle entity.
func (cu *ConfigurationUpdate) ClearCycle() *ConfigurationUpdate {
	cu.mutation.ClearCycle()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConfigurationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConfigurationUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConfigurationUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConfigurationUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ConfigurationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(configuration.Table, configuration.Columns, sqlgraph.NewFieldSpec(configuration.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.StartRegistrationSubjects(); ok {
		_spec.SetField(configuration.FieldStartRegistrationSubjects, field.TypeTime, value)
	}
	if value, ok := cu.mutation.EndRegistrationSubjects(); ok {
		_spec.SetField(configuration.FieldEndRegistrationSubjects, field.TypeTime, value)
	}
	if value, ok := cu.mutation.BlockNotPayInscription(); ok {
		_spec.SetField(configuration.FieldBlockNotPayInscription, field.TypeBool, value)
	}
	if value, ok := cu.mutation.FeeDates(); ok {
		_spec.SetField(configuration.FieldFeeDates, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedFeeDates(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, configuration.FieldFeeDates, value)
		})
	}
	if cu.mutation.FeeDatesCleared() {
		_spec.ClearField(configuration.FieldFeeDates, field.TypeJSON)
	}
	if value, ok := cu.mutation.NumberFees(); ok {
		_spec.SetField(configuration.FieldNumberFees, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedNumberFees(); ok {
		_spec.AddField(configuration.FieldNumberFees, field.TypeInt, value)
	}
	if value, ok := cu.mutation.NumberNotes(); ok {
		_spec.SetField(configuration.FieldNumberNotes, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedNumberNotes(); ok {
		_spec.AddField(configuration.FieldNumberNotes, field.TypeInt, value)
	}
	if value, ok := cu.mutation.NotesPercentages(); ok {
		_spec.SetField(configuration.FieldNotesPercentages, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedNotesPercentages(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, configuration.FieldNotesPercentages, value)
		})
	}
	if cu.mutation.NotesPercentagesCleared() {
		_spec.ClearField(configuration.FieldNotesPercentages, field.TypeJSON)
	}
	if cu.mutation.CycleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   configuration.CycleTable,
			Columns: []string{configuration.CycleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cycle.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CycleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   configuration.CycleTable,
			Columns: []string{configuration.CycleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cycle.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{configuration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ConfigurationUpdateOne is the builder for updating a single Configuration entity.
type ConfigurationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConfigurationMutation
}

// SetStartRegistrationSubjects sets the "start_registration_subjects" field.
func (cuo *ConfigurationUpdateOne) SetStartRegistrationSubjects(t time.Time) *ConfigurationUpdateOne {
	cuo.mutation.SetStartRegistrationSubjects(t)
	return cuo
}

// SetNillableStartRegistrationSubjects sets the "start_registration_subjects" field if the given value is not nil.
func (cuo *ConfigurationUpdateOne) SetNillableStartRegistrationSubjects(t *time.Time) *ConfigurationUpdateOne {
	if t != nil {
		cuo.SetStartRegistrationSubjects(*t)
	}
	return cuo
}

// SetEndRegistrationSubjects sets the "end_registration_subjects" field.
func (cuo *ConfigurationUpdateOne) SetEndRegistrationSubjects(t time.Time) *ConfigurationUpdateOne {
	cuo.mutation.SetEndRegistrationSubjects(t)
	return cuo
}

// SetNillableEndRegistrationSubjects sets the "end_registration_subjects" field if the given value is not nil.
func (cuo *ConfigurationUpdateOne) SetNillableEndRegistrationSubjects(t *time.Time) *ConfigurationUpdateOne {
	if t != nil {
		cuo.SetEndRegistrationSubjects(*t)
	}
	return cuo
}

// SetBlockNotPayInscription sets the "block_not_pay_inscription" field.
func (cuo *ConfigurationUpdateOne) SetBlockNotPayInscription(b bool) *ConfigurationUpdateOne {
	cuo.mutation.SetBlockNotPayInscription(b)
	return cuo
}

// SetNillableBlockNotPayInscription sets the "block_not_pay_inscription" field if the given value is not nil.
func (cuo *ConfigurationUpdateOne) SetNillableBlockNotPayInscription(b *bool) *ConfigurationUpdateOne {
	if b != nil {
		cuo.SetBlockNotPayInscription(*b)
	}
	return cuo
}

// SetFeeDates sets the "fee_dates" field.
func (cuo *ConfigurationUpdateOne) SetFeeDates(t []time.Time) *ConfigurationUpdateOne {
	cuo.mutation.SetFeeDates(t)
	return cuo
}

// AppendFeeDates appends t to the "fee_dates" field.
func (cuo *ConfigurationUpdateOne) AppendFeeDates(t []time.Time) *ConfigurationUpdateOne {
	cuo.mutation.AppendFeeDates(t)
	return cuo
}

// ClearFeeDates clears the value of the "fee_dates" field.
func (cuo *ConfigurationUpdateOne) ClearFeeDates() *ConfigurationUpdateOne {
	cuo.mutation.ClearFeeDates()
	return cuo
}

// SetNumberFees sets the "number_fees" field.
func (cuo *ConfigurationUpdateOne) SetNumberFees(i int) *ConfigurationUpdateOne {
	cuo.mutation.ResetNumberFees()
	cuo.mutation.SetNumberFees(i)
	return cuo
}

// SetNillableNumberFees sets the "number_fees" field if the given value is not nil.
func (cuo *ConfigurationUpdateOne) SetNillableNumberFees(i *int) *ConfigurationUpdateOne {
	if i != nil {
		cuo.SetNumberFees(*i)
	}
	return cuo
}

// AddNumberFees adds i to the "number_fees" field.
func (cuo *ConfigurationUpdateOne) AddNumberFees(i int) *ConfigurationUpdateOne {
	cuo.mutation.AddNumberFees(i)
	return cuo
}

// SetNumberNotes sets the "number_notes" field.
func (cuo *ConfigurationUpdateOne) SetNumberNotes(i int) *ConfigurationUpdateOne {
	cuo.mutation.ResetNumberNotes()
	cuo.mutation.SetNumberNotes(i)
	return cuo
}

// SetNillableNumberNotes sets the "number_notes" field if the given value is not nil.
func (cuo *ConfigurationUpdateOne) SetNillableNumberNotes(i *int) *ConfigurationUpdateOne {
	if i != nil {
		cuo.SetNumberNotes(*i)
	}
	return cuo
}

// AddNumberNotes adds i to the "number_notes" field.
func (cuo *ConfigurationUpdateOne) AddNumberNotes(i int) *ConfigurationUpdateOne {
	cuo.mutation.AddNumberNotes(i)
	return cuo
}

// SetNotesPercentages sets the "notes_Percentages" field.
func (cuo *ConfigurationUpdateOne) SetNotesPercentages(f []float64) *ConfigurationUpdateOne {
	cuo.mutation.SetNotesPercentages(f)
	return cuo
}

// AppendNotesPercentages appends f to the "notes_Percentages" field.
func (cuo *ConfigurationUpdateOne) AppendNotesPercentages(f []float64) *ConfigurationUpdateOne {
	cuo.mutation.AppendNotesPercentages(f)
	return cuo
}

// ClearNotesPercentages clears the value of the "notes_Percentages" field.
func (cuo *ConfigurationUpdateOne) ClearNotesPercentages() *ConfigurationUpdateOne {
	cuo.mutation.ClearNotesPercentages()
	return cuo
}

// SetCycleID sets the "cycle" edge to the Cycle entity by ID.
func (cuo *ConfigurationUpdateOne) SetCycleID(id int) *ConfigurationUpdateOne {
	cuo.mutation.SetCycleID(id)
	return cuo
}

// SetNillableCycleID sets the "cycle" edge to the Cycle entity by ID if the given value is not nil.
func (cuo *ConfigurationUpdateOne) SetNillableCycleID(id *int) *ConfigurationUpdateOne {
	if id != nil {
		cuo = cuo.SetCycleID(*id)
	}
	return cuo
}

// SetCycle sets the "cycle" edge to the Cycle entity.
func (cuo *ConfigurationUpdateOne) SetCycle(c *Cycle) *ConfigurationUpdateOne {
	return cuo.SetCycleID(c.ID)
}

// Mutation returns the ConfigurationMutation object of the builder.
func (cuo *ConfigurationUpdateOne) Mutation() *ConfigurationMutation {
	return cuo.mutation
}

// ClearCycle clears the "cycle" edge to the Cycle entity.
func (cuo *ConfigurationUpdateOne) ClearCycle() *ConfigurationUpdateOne {
	cuo.mutation.ClearCycle()
	return cuo
}

// Where appends a list predicates to the ConfigurationUpdate builder.
func (cuo *ConfigurationUpdateOne) Where(ps ...predicate.Configuration) *ConfigurationUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConfigurationUpdateOne) Select(field string, fields ...string) *ConfigurationUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Configuration entity.
func (cuo *ConfigurationUpdateOne) Save(ctx context.Context) (*Configuration, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConfigurationUpdateOne) SaveX(ctx context.Context) *Configuration {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConfigurationUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConfigurationUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ConfigurationUpdateOne) sqlSave(ctx context.Context) (_node *Configuration, err error) {
	_spec := sqlgraph.NewUpdateSpec(configuration.Table, configuration.Columns, sqlgraph.NewFieldSpec(configuration.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Configuration.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, configuration.FieldID)
		for _, f := range fields {
			if !configuration.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != configuration.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.StartRegistrationSubjects(); ok {
		_spec.SetField(configuration.FieldStartRegistrationSubjects, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.EndRegistrationSubjects(); ok {
		_spec.SetField(configuration.FieldEndRegistrationSubjects, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.BlockNotPayInscription(); ok {
		_spec.SetField(configuration.FieldBlockNotPayInscription, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.FeeDates(); ok {
		_spec.SetField(configuration.FieldFeeDates, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedFeeDates(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, configuration.FieldFeeDates, value)
		})
	}
	if cuo.mutation.FeeDatesCleared() {
		_spec.ClearField(configuration.FieldFeeDates, field.TypeJSON)
	}
	if value, ok := cuo.mutation.NumberFees(); ok {
		_spec.SetField(configuration.FieldNumberFees, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedNumberFees(); ok {
		_spec.AddField(configuration.FieldNumberFees, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.NumberNotes(); ok {
		_spec.SetField(configuration.FieldNumberNotes, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedNumberNotes(); ok {
		_spec.AddField(configuration.FieldNumberNotes, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.NotesPercentages(); ok {
		_spec.SetField(configuration.FieldNotesPercentages, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedNotesPercentages(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, configuration.FieldNotesPercentages, value)
		})
	}
	if cuo.mutation.NotesPercentagesCleared() {
		_spec.ClearField(configuration.FieldNotesPercentages, field.TypeJSON)
	}
	if cuo.mutation.CycleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   configuration.CycleTable,
			Columns: []string{configuration.CycleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cycle.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CycleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   configuration.CycleTable,
			Columns: []string{configuration.CycleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cycle.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Configuration{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{configuration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
