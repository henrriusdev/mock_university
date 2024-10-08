// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mocku/backend/ent/cycle"
	"mocku/backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CycleUpdate is the builder for updating Cycle entities.
type CycleUpdate struct {
	config
	hooks    []Hook
	mutation *CycleMutation
}

// Where appends a list predicates to the CycleUpdate builder.
func (cu *CycleUpdate) Where(ps ...predicate.Cycle) *CycleUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CycleUpdate) SetName(s string) *CycleUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CycleUpdate) SetNillableName(s *string) *CycleUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetStartDate sets the "start_date" field.
func (cu *CycleUpdate) SetStartDate(t time.Time) *CycleUpdate {
	cu.mutation.SetStartDate(t)
	return cu
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (cu *CycleUpdate) SetNillableStartDate(t *time.Time) *CycleUpdate {
	if t != nil {
		cu.SetStartDate(*t)
	}
	return cu
}

// SetEndDate sets the "end_date" field.
func (cu *CycleUpdate) SetEndDate(t time.Time) *CycleUpdate {
	cu.mutation.SetEndDate(t)
	return cu
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (cu *CycleUpdate) SetNillableEndDate(t *time.Time) *CycleUpdate {
	if t != nil {
		cu.SetEndDate(*t)
	}
	return cu
}

// SetActive sets the "active" field.
func (cu *CycleUpdate) SetActive(b bool) *CycleUpdate {
	cu.mutation.SetActive(b)
	return cu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (cu *CycleUpdate) SetNillableActive(b *bool) *CycleUpdate {
	if b != nil {
		cu.SetActive(*b)
	}
	return cu
}

// Mutation returns the CycleMutation object of the builder.
func (cu *CycleUpdate) Mutation() *CycleMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CycleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CycleUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CycleUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CycleUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CycleUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := cycle.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Cycle.name": %w`, err)}
		}
	}
	return nil
}

func (cu *CycleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(cycle.Table, cycle.Columns, sqlgraph.NewFieldSpec(cycle.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(cycle.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.StartDate(); ok {
		_spec.SetField(cycle.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := cu.mutation.EndDate(); ok {
		_spec.SetField(cycle.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Active(); ok {
		_spec.SetField(cycle.FieldActive, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cycle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CycleUpdateOne is the builder for updating a single Cycle entity.
type CycleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CycleMutation
}

// SetName sets the "name" field.
func (cuo *CycleUpdateOne) SetName(s string) *CycleUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CycleUpdateOne) SetNillableName(s *string) *CycleUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetStartDate sets the "start_date" field.
func (cuo *CycleUpdateOne) SetStartDate(t time.Time) *CycleUpdateOne {
	cuo.mutation.SetStartDate(t)
	return cuo
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (cuo *CycleUpdateOne) SetNillableStartDate(t *time.Time) *CycleUpdateOne {
	if t != nil {
		cuo.SetStartDate(*t)
	}
	return cuo
}

// SetEndDate sets the "end_date" field.
func (cuo *CycleUpdateOne) SetEndDate(t time.Time) *CycleUpdateOne {
	cuo.mutation.SetEndDate(t)
	return cuo
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (cuo *CycleUpdateOne) SetNillableEndDate(t *time.Time) *CycleUpdateOne {
	if t != nil {
		cuo.SetEndDate(*t)
	}
	return cuo
}

// SetActive sets the "active" field.
func (cuo *CycleUpdateOne) SetActive(b bool) *CycleUpdateOne {
	cuo.mutation.SetActive(b)
	return cuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (cuo *CycleUpdateOne) SetNillableActive(b *bool) *CycleUpdateOne {
	if b != nil {
		cuo.SetActive(*b)
	}
	return cuo
}

// Mutation returns the CycleMutation object of the builder.
func (cuo *CycleUpdateOne) Mutation() *CycleMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CycleUpdate builder.
func (cuo *CycleUpdateOne) Where(ps ...predicate.Cycle) *CycleUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CycleUpdateOne) Select(field string, fields ...string) *CycleUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cycle entity.
func (cuo *CycleUpdateOne) Save(ctx context.Context) (*Cycle, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CycleUpdateOne) SaveX(ctx context.Context) *Cycle {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CycleUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CycleUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CycleUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := cycle.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Cycle.name": %w`, err)}
		}
	}
	return nil
}

func (cuo *CycleUpdateOne) sqlSave(ctx context.Context) (_node *Cycle, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(cycle.Table, cycle.Columns, sqlgraph.NewFieldSpec(cycle.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cycle.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cycle.FieldID)
		for _, f := range fields {
			if !cycle.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cycle.FieldID {
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
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(cycle.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.StartDate(); ok {
		_spec.SetField(cycle.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.EndDate(); ok {
		_spec.SetField(cycle.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Active(); ok {
		_spec.SetField(cycle.FieldActive, field.TypeBool, value)
	}
	_node = &Cycle{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cycle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
