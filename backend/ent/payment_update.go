// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mocku/backend/ent/cycle"
	"mocku/backend/ent/payment"
	"mocku/backend/ent/paymentmethod"
	"mocku/backend/ent/predicate"
	"mocku/backend/ent/student"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PaymentUpdate is the builder for updating Payment entities.
type PaymentUpdate struct {
	config
	hooks    []Hook
	mutation *PaymentMutation
}

// Where appends a list predicates to the PaymentUpdate builder.
func (pu *PaymentUpdate) Where(ps ...predicate.Payment) *PaymentUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetReference sets the "reference" field.
func (pu *PaymentUpdate) SetReference(s string) *PaymentUpdate {
	pu.mutation.SetReference(s)
	return pu
}

// SetNillableReference sets the "reference" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableReference(s *string) *PaymentUpdate {
	if s != nil {
		pu.SetReference(*s)
	}
	return pu
}

// SetDate sets the "date" field.
func (pu *PaymentUpdate) SetDate(t time.Time) *PaymentUpdate {
	pu.mutation.SetDate(t)
	return pu
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableDate(t *time.Time) *PaymentUpdate {
	if t != nil {
		pu.SetDate(*t)
	}
	return pu
}

// SetAmount sets the "amount" field.
func (pu *PaymentUpdate) SetAmount(f float64) *PaymentUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(f)
	return pu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableAmount(f *float64) *PaymentUpdate {
	if f != nil {
		pu.SetAmount(*f)
	}
	return pu
}

// AddAmount adds f to the "amount" field.
func (pu *PaymentUpdate) AddAmount(f float64) *PaymentUpdate {
	pu.mutation.AddAmount(f)
	return pu
}

// SetDescription sets the "description" field.
func (pu *PaymentUpdate) SetDescription(s string) *PaymentUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableDescription(s *string) *PaymentUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// SetFeeNumber sets the "fee_number" field.
func (pu *PaymentUpdate) SetFeeNumber(i int) *PaymentUpdate {
	pu.mutation.ResetFeeNumber()
	pu.mutation.SetFeeNumber(i)
	return pu
}

// SetNillableFeeNumber sets the "fee_number" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableFeeNumber(i *int) *PaymentUpdate {
	if i != nil {
		pu.SetFeeNumber(*i)
	}
	return pu
}

// AddFeeNumber adds i to the "fee_number" field.
func (pu *PaymentUpdate) AddFeeNumber(i int) *PaymentUpdate {
	pu.mutation.AddFeeNumber(i)
	return pu
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (pu *PaymentUpdate) SetStudentID(id int) *PaymentUpdate {
	pu.mutation.SetStudentID(id)
	return pu
}

// SetNillableStudentID sets the "student" edge to the Student entity by ID if the given value is not nil.
func (pu *PaymentUpdate) SetNillableStudentID(id *int) *PaymentUpdate {
	if id != nil {
		pu = pu.SetStudentID(*id)
	}
	return pu
}

// SetStudent sets the "student" edge to the Student entity.
func (pu *PaymentUpdate) SetStudent(s *Student) *PaymentUpdate {
	return pu.SetStudentID(s.ID)
}

// SetCycleID sets the "cycle" edge to the Cycle entity by ID.
func (pu *PaymentUpdate) SetCycleID(id int) *PaymentUpdate {
	pu.mutation.SetCycleID(id)
	return pu
}

// SetNillableCycleID sets the "cycle" edge to the Cycle entity by ID if the given value is not nil.
func (pu *PaymentUpdate) SetNillableCycleID(id *int) *PaymentUpdate {
	if id != nil {
		pu = pu.SetCycleID(*id)
	}
	return pu
}

// SetCycle sets the "cycle" edge to the Cycle entity.
func (pu *PaymentUpdate) SetCycle(c *Cycle) *PaymentUpdate {
	return pu.SetCycleID(c.ID)
}

// SetPaymentMethodID sets the "payment_method" edge to the PaymentMethod entity by ID.
func (pu *PaymentUpdate) SetPaymentMethodID(id int) *PaymentUpdate {
	pu.mutation.SetPaymentMethodID(id)
	return pu
}

// SetNillablePaymentMethodID sets the "payment_method" edge to the PaymentMethod entity by ID if the given value is not nil.
func (pu *PaymentUpdate) SetNillablePaymentMethodID(id *int) *PaymentUpdate {
	if id != nil {
		pu = pu.SetPaymentMethodID(*id)
	}
	return pu
}

// SetPaymentMethod sets the "payment_method" edge to the PaymentMethod entity.
func (pu *PaymentUpdate) SetPaymentMethod(p *PaymentMethod) *PaymentUpdate {
	return pu.SetPaymentMethodID(p.ID)
}

// Mutation returns the PaymentMutation object of the builder.
func (pu *PaymentUpdate) Mutation() *PaymentMutation {
	return pu.mutation
}

// ClearStudent clears the "student" edge to the Student entity.
func (pu *PaymentUpdate) ClearStudent() *PaymentUpdate {
	pu.mutation.ClearStudent()
	return pu
}

// ClearCycle clears the "cycle" edge to the Cycle entity.
func (pu *PaymentUpdate) ClearCycle() *PaymentUpdate {
	pu.mutation.ClearCycle()
	return pu
}

// ClearPaymentMethod clears the "payment_method" edge to the PaymentMethod entity.
func (pu *PaymentUpdate) ClearPaymentMethod() *PaymentUpdate {
	pu.mutation.ClearPaymentMethod()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PaymentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaymentUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaymentUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaymentUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PaymentUpdate) check() error {
	if v, ok := pu.mutation.Reference(); ok {
		if err := payment.ReferenceValidator(v); err != nil {
			return &ValidationError{Name: "reference", err: fmt.Errorf(`ent: validator failed for field "Payment.reference": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Amount(); ok {
		if err := payment.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Payment.amount": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Description(); ok {
		if err := payment.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Payment.description": %w`, err)}
		}
	}
	if v, ok := pu.mutation.FeeNumber(); ok {
		if err := payment.FeeNumberValidator(v); err != nil {
			return &ValidationError{Name: "fee_number", err: fmt.Errorf(`ent: validator failed for field "Payment.fee_number": %w`, err)}
		}
	}
	return nil
}

func (pu *PaymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Reference(); ok {
		_spec.SetField(payment.FieldReference, field.TypeString, value)
	}
	if value, ok := pu.mutation.Date(); ok {
		_spec.SetField(payment.FieldDate, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		_spec.AddField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(payment.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.FeeNumber(); ok {
		_spec.SetField(payment.FieldFeeNumber, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedFeeNumber(); ok {
		_spec.AddField(payment.FieldFeeNumber, field.TypeInt, value)
	}
	if pu.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payment.StudentTable,
			Columns: []string{payment.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payment.StudentTable,
			Columns: []string{payment.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.CycleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payment.CycleTable,
			Columns: []string{payment.CycleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cycle.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CycleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payment.CycleTable,
			Columns: []string{payment.CycleColumn},
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
	if pu.mutation.PaymentMethodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payment.PaymentMethodTable,
			Columns: []string{payment.PaymentMethodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentmethod.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PaymentMethodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payment.PaymentMethodTable,
			Columns: []string{payment.PaymentMethodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentmethod.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PaymentUpdateOne is the builder for updating a single Payment entity.
type PaymentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaymentMutation
}

// SetReference sets the "reference" field.
func (puo *PaymentUpdateOne) SetReference(s string) *PaymentUpdateOne {
	puo.mutation.SetReference(s)
	return puo
}

// SetNillableReference sets the "reference" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableReference(s *string) *PaymentUpdateOne {
	if s != nil {
		puo.SetReference(*s)
	}
	return puo
}

// SetDate sets the "date" field.
func (puo *PaymentUpdateOne) SetDate(t time.Time) *PaymentUpdateOne {
	puo.mutation.SetDate(t)
	return puo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableDate(t *time.Time) *PaymentUpdateOne {
	if t != nil {
		puo.SetDate(*t)
	}
	return puo
}

// SetAmount sets the "amount" field.
func (puo *PaymentUpdateOne) SetAmount(f float64) *PaymentUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(f)
	return puo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableAmount(f *float64) *PaymentUpdateOne {
	if f != nil {
		puo.SetAmount(*f)
	}
	return puo
}

// AddAmount adds f to the "amount" field.
func (puo *PaymentUpdateOne) AddAmount(f float64) *PaymentUpdateOne {
	puo.mutation.AddAmount(f)
	return puo
}

// SetDescription sets the "description" field.
func (puo *PaymentUpdateOne) SetDescription(s string) *PaymentUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableDescription(s *string) *PaymentUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// SetFeeNumber sets the "fee_number" field.
func (puo *PaymentUpdateOne) SetFeeNumber(i int) *PaymentUpdateOne {
	puo.mutation.ResetFeeNumber()
	puo.mutation.SetFeeNumber(i)
	return puo
}

// SetNillableFeeNumber sets the "fee_number" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableFeeNumber(i *int) *PaymentUpdateOne {
	if i != nil {
		puo.SetFeeNumber(*i)
	}
	return puo
}

// AddFeeNumber adds i to the "fee_number" field.
func (puo *PaymentUpdateOne) AddFeeNumber(i int) *PaymentUpdateOne {
	puo.mutation.AddFeeNumber(i)
	return puo
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (puo *PaymentUpdateOne) SetStudentID(id int) *PaymentUpdateOne {
	puo.mutation.SetStudentID(id)
	return puo
}

// SetNillableStudentID sets the "student" edge to the Student entity by ID if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableStudentID(id *int) *PaymentUpdateOne {
	if id != nil {
		puo = puo.SetStudentID(*id)
	}
	return puo
}

// SetStudent sets the "student" edge to the Student entity.
func (puo *PaymentUpdateOne) SetStudent(s *Student) *PaymentUpdateOne {
	return puo.SetStudentID(s.ID)
}

// SetCycleID sets the "cycle" edge to the Cycle entity by ID.
func (puo *PaymentUpdateOne) SetCycleID(id int) *PaymentUpdateOne {
	puo.mutation.SetCycleID(id)
	return puo
}

// SetNillableCycleID sets the "cycle" edge to the Cycle entity by ID if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableCycleID(id *int) *PaymentUpdateOne {
	if id != nil {
		puo = puo.SetCycleID(*id)
	}
	return puo
}

// SetCycle sets the "cycle" edge to the Cycle entity.
func (puo *PaymentUpdateOne) SetCycle(c *Cycle) *PaymentUpdateOne {
	return puo.SetCycleID(c.ID)
}

// SetPaymentMethodID sets the "payment_method" edge to the PaymentMethod entity by ID.
func (puo *PaymentUpdateOne) SetPaymentMethodID(id int) *PaymentUpdateOne {
	puo.mutation.SetPaymentMethodID(id)
	return puo
}

// SetNillablePaymentMethodID sets the "payment_method" edge to the PaymentMethod entity by ID if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillablePaymentMethodID(id *int) *PaymentUpdateOne {
	if id != nil {
		puo = puo.SetPaymentMethodID(*id)
	}
	return puo
}

// SetPaymentMethod sets the "payment_method" edge to the PaymentMethod entity.
func (puo *PaymentUpdateOne) SetPaymentMethod(p *PaymentMethod) *PaymentUpdateOne {
	return puo.SetPaymentMethodID(p.ID)
}

// Mutation returns the PaymentMutation object of the builder.
func (puo *PaymentUpdateOne) Mutation() *PaymentMutation {
	return puo.mutation
}

// ClearStudent clears the "student" edge to the Student entity.
func (puo *PaymentUpdateOne) ClearStudent() *PaymentUpdateOne {
	puo.mutation.ClearStudent()
	return puo
}

// ClearCycle clears the "cycle" edge to the Cycle entity.
func (puo *PaymentUpdateOne) ClearCycle() *PaymentUpdateOne {
	puo.mutation.ClearCycle()
	return puo
}

// ClearPaymentMethod clears the "payment_method" edge to the PaymentMethod entity.
func (puo *PaymentUpdateOne) ClearPaymentMethod() *PaymentUpdateOne {
	puo.mutation.ClearPaymentMethod()
	return puo
}

// Where appends a list predicates to the PaymentUpdate builder.
func (puo *PaymentUpdateOne) Where(ps ...predicate.Payment) *PaymentUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PaymentUpdateOne) Select(field string, fields ...string) *PaymentUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Payment entity.
func (puo *PaymentUpdateOne) Save(ctx context.Context) (*Payment, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaymentUpdateOne) SaveX(ctx context.Context) *Payment {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PaymentUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaymentUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PaymentUpdateOne) check() error {
	if v, ok := puo.mutation.Reference(); ok {
		if err := payment.ReferenceValidator(v); err != nil {
			return &ValidationError{Name: "reference", err: fmt.Errorf(`ent: validator failed for field "Payment.reference": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Amount(); ok {
		if err := payment.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Payment.amount": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Description(); ok {
		if err := payment.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Payment.description": %w`, err)}
		}
	}
	if v, ok := puo.mutation.FeeNumber(); ok {
		if err := payment.FeeNumberValidator(v); err != nil {
			return &ValidationError{Name: "fee_number", err: fmt.Errorf(`ent: validator failed for field "Payment.fee_number": %w`, err)}
		}
	}
	return nil
}

func (puo *PaymentUpdateOne) sqlSave(ctx context.Context) (_node *Payment, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Payment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payment.FieldID)
		for _, f := range fields {
			if !payment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != payment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Reference(); ok {
		_spec.SetField(payment.FieldReference, field.TypeString, value)
	}
	if value, ok := puo.mutation.Date(); ok {
		_spec.SetField(payment.FieldDate, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		_spec.AddField(payment.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(payment.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.FeeNumber(); ok {
		_spec.SetField(payment.FieldFeeNumber, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedFeeNumber(); ok {
		_spec.AddField(payment.FieldFeeNumber, field.TypeInt, value)
	}
	if puo.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payment.StudentTable,
			Columns: []string{payment.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payment.StudentTable,
			Columns: []string{payment.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.CycleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payment.CycleTable,
			Columns: []string{payment.CycleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cycle.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CycleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payment.CycleTable,
			Columns: []string{payment.CycleColumn},
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
	if puo.mutation.PaymentMethodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payment.PaymentMethodTable,
			Columns: []string{payment.PaymentMethodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentmethod.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PaymentMethodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   payment.PaymentMethodTable,
			Columns: []string{payment.PaymentMethodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(paymentmethod.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Payment{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
