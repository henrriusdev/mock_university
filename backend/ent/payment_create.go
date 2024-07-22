// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mocku/backend/ent/cycle"
	"mocku/backend/ent/payment"
	"mocku/backend/ent/paymentmethod"
	"mocku/backend/ent/student"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PaymentCreate is the builder for creating a Payment entity.
type PaymentCreate struct {
	config
	mutation *PaymentMutation
	hooks    []Hook
}

// SetReference sets the "reference" field.
func (pc *PaymentCreate) SetReference(s string) *PaymentCreate {
	pc.mutation.SetReference(s)
	return pc
}

// SetDate sets the "date" field.
func (pc *PaymentCreate) SetDate(t time.Time) *PaymentCreate {
	pc.mutation.SetDate(t)
	return pc
}

// SetAmount sets the "amount" field.
func (pc *PaymentCreate) SetAmount(f float64) *PaymentCreate {
	pc.mutation.SetAmount(f)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PaymentCreate) SetDescription(s string) *PaymentCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetFeeNumber sets the "fee_number" field.
func (pc *PaymentCreate) SetFeeNumber(i int) *PaymentCreate {
	pc.mutation.SetFeeNumber(i)
	return pc
}

// AddStudentIDs adds the "student" edge to the Student entity by IDs.
func (pc *PaymentCreate) AddStudentIDs(ids ...int) *PaymentCreate {
	pc.mutation.AddStudentIDs(ids...)
	return pc
}

// AddStudent adds the "student" edges to the Student entity.
func (pc *PaymentCreate) AddStudent(s ...*Student) *PaymentCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddStudentIDs(ids...)
}

// AddCycleIDs adds the "cycle" edge to the Cycle entity by IDs.
func (pc *PaymentCreate) AddCycleIDs(ids ...int) *PaymentCreate {
	pc.mutation.AddCycleIDs(ids...)
	return pc
}

// AddCycle adds the "cycle" edges to the Cycle entity.
func (pc *PaymentCreate) AddCycle(c ...*Cycle) *PaymentCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCycleIDs(ids...)
}

// AddPaymentMethodIDs adds the "payment_method" edge to the PaymentMethod entity by IDs.
func (pc *PaymentCreate) AddPaymentMethodIDs(ids ...int) *PaymentCreate {
	pc.mutation.AddPaymentMethodIDs(ids...)
	return pc
}

// AddPaymentMethod adds the "payment_method" edges to the PaymentMethod entity.
func (pc *PaymentCreate) AddPaymentMethod(p ...*PaymentMethod) *PaymentCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPaymentMethodIDs(ids...)
}

// Mutation returns the PaymentMutation object of the builder.
func (pc *PaymentCreate) Mutation() *PaymentMutation {
	return pc.mutation
}

// Save creates the Payment in the database.
func (pc *PaymentCreate) Save(ctx context.Context) (*Payment, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PaymentCreate) SaveX(ctx context.Context) *Payment {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PaymentCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PaymentCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PaymentCreate) check() error {
	if _, ok := pc.mutation.Reference(); !ok {
		return &ValidationError{Name: "reference", err: errors.New(`ent: missing required field "Payment.reference"`)}
	}
	if v, ok := pc.mutation.Reference(); ok {
		if err := payment.ReferenceValidator(v); err != nil {
			return &ValidationError{Name: "reference", err: fmt.Errorf(`ent: validator failed for field "Payment.reference": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "Payment.date"`)}
	}
	if _, ok := pc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Payment.amount"`)}
	}
	if v, ok := pc.mutation.Amount(); ok {
		if err := payment.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Payment.amount": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Payment.description"`)}
	}
	if v, ok := pc.mutation.Description(); ok {
		if err := payment.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Payment.description": %w`, err)}
		}
	}
	if _, ok := pc.mutation.FeeNumber(); !ok {
		return &ValidationError{Name: "fee_number", err: errors.New(`ent: missing required field "Payment.fee_number"`)}
	}
	if v, ok := pc.mutation.FeeNumber(); ok {
		if err := payment.FeeNumberValidator(v); err != nil {
			return &ValidationError{Name: "fee_number", err: fmt.Errorf(`ent: validator failed for field "Payment.fee_number": %w`, err)}
		}
	}
	return nil
}

func (pc *PaymentCreate) sqlSave(ctx context.Context) (*Payment, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PaymentCreate) createSpec() (*Payment, *sqlgraph.CreateSpec) {
	var (
		_node = &Payment{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(payment.Table, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Reference(); ok {
		_spec.SetField(payment.FieldReference, field.TypeString, value)
		_node.Reference = value
	}
	if value, ok := pc.mutation.Date(); ok {
		_spec.SetField(payment.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	if value, ok := pc.mutation.Amount(); ok {
		_spec.SetField(payment.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(payment.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.FeeNumber(); ok {
		_spec.SetField(payment.FieldFeeNumber, field.TypeInt, value)
		_node.FeeNumber = value
	}
	if nodes := pc.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   payment.StudentTable,
			Columns: payment.StudentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CycleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PaymentMethodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PaymentCreateBulk is the builder for creating many Payment entities in bulk.
type PaymentCreateBulk struct {
	config
	err      error
	builders []*PaymentCreate
}

// Save creates the Payment entities in the database.
func (pcb *PaymentCreateBulk) Save(ctx context.Context) ([]*Payment, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Payment, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PaymentMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PaymentCreateBulk) SaveX(ctx context.Context) []*Payment {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PaymentCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PaymentCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
