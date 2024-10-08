// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mocku/backend/ent/module"
	"mocku/backend/ent/permission"
	"mocku/backend/ent/predicate"
	"mocku/backend/ent/role"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks    []Hook
	mutation *PermissionMutation
}

// Where appends a list predicates to the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PermissionUpdate) SetName(s string) *PermissionUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableName(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetDescription sets the "description" field.
func (pu *PermissionUpdate) SetDescription(s string) *PermissionUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableDescription(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// SetRead sets the "read" field.
func (pu *PermissionUpdate) SetRead(b bool) *PermissionUpdate {
	pu.mutation.SetRead(b)
	return pu
}

// SetNillableRead sets the "read" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableRead(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetRead(*b)
	}
	return pu
}

// SetCreate sets the "create" field.
func (pu *PermissionUpdate) SetCreate(b bool) *PermissionUpdate {
	pu.mutation.SetCreate(b)
	return pu
}

// SetNillableCreate sets the "create" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableCreate(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetCreate(*b)
	}
	return pu
}

// SetModify sets the "modify" field.
func (pu *PermissionUpdate) SetModify(b bool) *PermissionUpdate {
	pu.mutation.SetModify(b)
	return pu
}

// SetNillableModify sets the "modify" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableModify(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetModify(*b)
	}
	return pu
}

// SetDelete sets the "delete" field.
func (pu *PermissionUpdate) SetDelete(b bool) *PermissionUpdate {
	pu.mutation.SetDelete(b)
	return pu
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableDelete(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetDelete(*b)
	}
	return pu
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (pu *PermissionUpdate) AddRoleIDs(ids ...int) *PermissionUpdate {
	pu.mutation.AddRoleIDs(ids...)
	return pu
}

// AddRoles adds the "roles" edges to the Role entity.
func (pu *PermissionUpdate) AddRoles(r ...*Role) *PermissionUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddRoleIDs(ids...)
}

// AddModuleIDs adds the "module" edge to the Module entity by IDs.
func (pu *PermissionUpdate) AddModuleIDs(ids ...int) *PermissionUpdate {
	pu.mutation.AddModuleIDs(ids...)
	return pu
}

// AddModule adds the "module" edges to the Module entity.
func (pu *PermissionUpdate) AddModule(m ...*Module) *PermissionUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pu.AddModuleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (pu *PermissionUpdate) ClearRoles() *PermissionUpdate {
	pu.mutation.ClearRoles()
	return pu
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (pu *PermissionUpdate) RemoveRoleIDs(ids ...int) *PermissionUpdate {
	pu.mutation.RemoveRoleIDs(ids...)
	return pu
}

// RemoveRoles removes "roles" edges to Role entities.
func (pu *PermissionUpdate) RemoveRoles(r ...*Role) *PermissionUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveRoleIDs(ids...)
}

// ClearModule clears all "module" edges to the Module entity.
func (pu *PermissionUpdate) ClearModule() *PermissionUpdate {
	pu.mutation.ClearModule()
	return pu
}

// RemoveModuleIDs removes the "module" edge to Module entities by IDs.
func (pu *PermissionUpdate) RemoveModuleIDs(ids ...int) *PermissionUpdate {
	pu.mutation.RemoveModuleIDs(ids...)
	return pu
}

// RemoveModule removes "module" edges to Module entities.
func (pu *PermissionUpdate) RemoveModule(m ...*Module) *PermissionUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pu.RemoveModuleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PermissionUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := permission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Permission.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Description(); ok {
		if err := permission.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Permission.description": %w`, err)}
		}
	}
	return nil
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(permission.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.Read(); ok {
		_spec.SetField(permission.FieldRead, field.TypeBool, value)
	}
	if value, ok := pu.mutation.Create(); ok {
		_spec.SetField(permission.FieldCreate, field.TypeBool, value)
	}
	if value, ok := pu.mutation.Modify(); ok {
		_spec.SetField(permission.FieldModify, field.TypeBool, value)
	}
	if value, ok := pu.mutation.Delete(); ok {
		_spec.SetField(permission.FieldDelete, field.TypeBool, value)
	}
	if pu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !pu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.ModuleTable,
			Columns: []string{permission.ModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(module.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedModuleIDs(); len(nodes) > 0 && !pu.mutation.ModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.ModuleTable,
			Columns: []string{permission.ModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(module.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ModuleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.ModuleTable,
			Columns: []string{permission.ModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(module.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PermissionMutation
}

// SetName sets the "name" field.
func (puo *PermissionUpdateOne) SetName(s string) *PermissionUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableName(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetDescription sets the "description" field.
func (puo *PermissionUpdateOne) SetDescription(s string) *PermissionUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableDescription(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// SetRead sets the "read" field.
func (puo *PermissionUpdateOne) SetRead(b bool) *PermissionUpdateOne {
	puo.mutation.SetRead(b)
	return puo
}

// SetNillableRead sets the "read" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableRead(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetRead(*b)
	}
	return puo
}

// SetCreate sets the "create" field.
func (puo *PermissionUpdateOne) SetCreate(b bool) *PermissionUpdateOne {
	puo.mutation.SetCreate(b)
	return puo
}

// SetNillableCreate sets the "create" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableCreate(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetCreate(*b)
	}
	return puo
}

// SetModify sets the "modify" field.
func (puo *PermissionUpdateOne) SetModify(b bool) *PermissionUpdateOne {
	puo.mutation.SetModify(b)
	return puo
}

// SetNillableModify sets the "modify" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableModify(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetModify(*b)
	}
	return puo
}

// SetDelete sets the "delete" field.
func (puo *PermissionUpdateOne) SetDelete(b bool) *PermissionUpdateOne {
	puo.mutation.SetDelete(b)
	return puo
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableDelete(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetDelete(*b)
	}
	return puo
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (puo *PermissionUpdateOne) AddRoleIDs(ids ...int) *PermissionUpdateOne {
	puo.mutation.AddRoleIDs(ids...)
	return puo
}

// AddRoles adds the "roles" edges to the Role entity.
func (puo *PermissionUpdateOne) AddRoles(r ...*Role) *PermissionUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddRoleIDs(ids...)
}

// AddModuleIDs adds the "module" edge to the Module entity by IDs.
func (puo *PermissionUpdateOne) AddModuleIDs(ids ...int) *PermissionUpdateOne {
	puo.mutation.AddModuleIDs(ids...)
	return puo
}

// AddModule adds the "module" edges to the Module entity.
func (puo *PermissionUpdateOne) AddModule(m ...*Module) *PermissionUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return puo.AddModuleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (puo *PermissionUpdateOne) ClearRoles() *PermissionUpdateOne {
	puo.mutation.ClearRoles()
	return puo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (puo *PermissionUpdateOne) RemoveRoleIDs(ids ...int) *PermissionUpdateOne {
	puo.mutation.RemoveRoleIDs(ids...)
	return puo
}

// RemoveRoles removes "roles" edges to Role entities.
func (puo *PermissionUpdateOne) RemoveRoles(r ...*Role) *PermissionUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveRoleIDs(ids...)
}

// ClearModule clears all "module" edges to the Module entity.
func (puo *PermissionUpdateOne) ClearModule() *PermissionUpdateOne {
	puo.mutation.ClearModule()
	return puo
}

// RemoveModuleIDs removes the "module" edge to Module entities by IDs.
func (puo *PermissionUpdateOne) RemoveModuleIDs(ids ...int) *PermissionUpdateOne {
	puo.mutation.RemoveModuleIDs(ids...)
	return puo
}

// RemoveModule removes "module" edges to Module entities.
func (puo *PermissionUpdateOne) RemoveModule(m ...*Module) *PermissionUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return puo.RemoveModuleIDs(ids...)
}

// Where appends a list predicates to the PermissionUpdate builder.
func (puo *PermissionUpdateOne) Where(ps ...predicate.Permission) *PermissionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PermissionUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := permission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Permission.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Description(); ok {
		if err := permission.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Permission.description": %w`, err)}
		}
	}
	return nil
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Permission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
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
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(permission.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.Read(); ok {
		_spec.SetField(permission.FieldRead, field.TypeBool, value)
	}
	if value, ok := puo.mutation.Create(); ok {
		_spec.SetField(permission.FieldCreate, field.TypeBool, value)
	}
	if value, ok := puo.mutation.Modify(); ok {
		_spec.SetField(permission.FieldModify, field.TypeBool, value)
	}
	if value, ok := puo.mutation.Delete(); ok {
		_spec.SetField(permission.FieldDelete, field.TypeBool, value)
	}
	if puo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !puo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.ModuleTable,
			Columns: []string{permission.ModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(module.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedModuleIDs(); len(nodes) > 0 && !puo.mutation.ModuleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.ModuleTable,
			Columns: []string{permission.ModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(module.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ModuleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.ModuleTable,
			Columns: []string{permission.ModuleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(module.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
