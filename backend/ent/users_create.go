// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mocku/backend/ent/activity"
	"mocku/backend/ent/blog"
	"mocku/backend/ent/notification"
	"mocku/backend/ent/professor"
	"mocku/backend/ent/request"
	"mocku/backend/ent/role"
	"mocku/backend/ent/student"
	"mocku/backend/ent/users"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UsersCreate is the builder for creating a Users entity.
type UsersCreate struct {
	config
	mutation *UsersMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (uc *UsersCreate) SetUsername(s string) *UsersCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetPassword sets the "password" field.
func (uc *UsersCreate) SetPassword(s string) *UsersCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetEmail sets the "email" field.
func (uc *UsersCreate) SetEmail(s string) *UsersCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetName sets the "name" field.
func (uc *UsersCreate) SetName(s string) *UsersCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetAvatar sets the "avatar" field.
func (uc *UsersCreate) SetAvatar(s string) *UsersCreate {
	uc.mutation.SetAvatar(s)
	return uc
}

// SetIsActive sets the "is_active" field.
func (uc *UsersCreate) SetIsActive(b bool) *UsersCreate {
	uc.mutation.SetIsActive(b)
	return uc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (uc *UsersCreate) SetNillableIsActive(b *bool) *UsersCreate {
	if b != nil {
		uc.SetIsActive(*b)
	}
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UsersCreate) SetCreatedAt(t time.Time) *UsersCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UsersCreate) SetNillableCreatedAt(t *time.Time) *UsersCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// AddRoleIDs adds the "role" edge to the Role entity by IDs.
func (uc *UsersCreate) AddRoleIDs(ids ...int) *UsersCreate {
	uc.mutation.AddRoleIDs(ids...)
	return uc
}

// AddRole adds the "role" edges to the Role entity.
func (uc *UsersCreate) AddRole(r ...*Role) *UsersCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddRoleIDs(ids...)
}

// AddRequestsMadeIDs adds the "requests_made" edge to the Request entity by IDs.
func (uc *UsersCreate) AddRequestsMadeIDs(ids ...int) *UsersCreate {
	uc.mutation.AddRequestsMadeIDs(ids...)
	return uc
}

// AddRequestsMade adds the "requests_made" edges to the Request entity.
func (uc *UsersCreate) AddRequestsMade(r ...*Request) *UsersCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddRequestsMadeIDs(ids...)
}

// AddRequestsReceivedIDs adds the "requests_received" edge to the Request entity by IDs.
func (uc *UsersCreate) AddRequestsReceivedIDs(ids ...int) *UsersCreate {
	uc.mutation.AddRequestsReceivedIDs(ids...)
	return uc
}

// AddRequestsReceived adds the "requests_received" edges to the Request entity.
func (uc *UsersCreate) AddRequestsReceived(r ...*Request) *UsersCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddRequestsReceivedIDs(ids...)
}

// AddBlogIDs adds the "blog" edge to the Blog entity by IDs.
func (uc *UsersCreate) AddBlogIDs(ids ...int) *UsersCreate {
	uc.mutation.AddBlogIDs(ids...)
	return uc
}

// AddBlog adds the "blog" edges to the Blog entity.
func (uc *UsersCreate) AddBlog(b ...*Blog) *UsersCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return uc.AddBlogIDs(ids...)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (uc *UsersCreate) AddNotificationIDs(ids ...int) *UsersCreate {
	uc.mutation.AddNotificationIDs(ids...)
	return uc
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (uc *UsersCreate) AddNotifications(n ...*Notification) *UsersCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return uc.AddNotificationIDs(ids...)
}

// AddActivityIDs adds the "activity" edge to the Activity entity by IDs.
func (uc *UsersCreate) AddActivityIDs(ids ...int) *UsersCreate {
	uc.mutation.AddActivityIDs(ids...)
	return uc
}

// AddActivity adds the "activity" edges to the Activity entity.
func (uc *UsersCreate) AddActivity(a ...*Activity) *UsersCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddActivityIDs(ids...)
}

// AddStudentIDs adds the "students" edge to the Student entity by IDs.
func (uc *UsersCreate) AddStudentIDs(ids ...int) *UsersCreate {
	uc.mutation.AddStudentIDs(ids...)
	return uc
}

// AddStudents adds the "students" edges to the Student entity.
func (uc *UsersCreate) AddStudents(s ...*Student) *UsersCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uc.AddStudentIDs(ids...)
}

// AddProfessorIDs adds the "professor" edge to the Professor entity by IDs.
func (uc *UsersCreate) AddProfessorIDs(ids ...int) *UsersCreate {
	uc.mutation.AddProfessorIDs(ids...)
	return uc
}

// AddProfessor adds the "professor" edges to the Professor entity.
func (uc *UsersCreate) AddProfessor(p ...*Professor) *UsersCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddProfessorIDs(ids...)
}

// Mutation returns the UsersMutation object of the builder.
func (uc *UsersCreate) Mutation() *UsersMutation {
	return uc.mutation
}

// Save creates the Users in the database.
func (uc *UsersCreate) Save(ctx context.Context) (*Users, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UsersCreate) SaveX(ctx context.Context) *Users {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UsersCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UsersCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UsersCreate) defaults() {
	if _, ok := uc.mutation.IsActive(); !ok {
		v := users.DefaultIsActive
		uc.mutation.SetIsActive(v)
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := users.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UsersCreate) check() error {
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Users.username"`)}
	}
	if v, ok := uc.mutation.Username(); ok {
		if err := users.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Users.username": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Users.password"`)}
	}
	if v, ok := uc.mutation.Password(); ok {
		if err := users.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Users.password": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Users.email"`)}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := users.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Users.email": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Users.name"`)}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := users.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Users.name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New(`ent: missing required field "Users.avatar"`)}
	}
	if v, ok := uc.mutation.Avatar(); ok {
		if err := users.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "Users.avatar": %w`, err)}
		}
	}
	if _, ok := uc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "Users.is_active"`)}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Users.created_at"`)}
	}
	return nil
}

func (uc *UsersCreate) sqlSave(ctx context.Context) (*Users, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UsersCreate) createSpec() (*Users, *sqlgraph.CreateSpec) {
	var (
		_node = &Users{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(users.Table, sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt))
	)
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(users.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(users.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(users.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(users.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.SetField(users.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := uc.mutation.IsActive(); ok {
		_spec.SetField(users.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(users.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := uc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   users.RoleTable,
			Columns: users.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.RequestsMadeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   users.RequestsMadeTable,
			Columns: users.RequestsMadePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.RequestsReceivedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   users.RequestsReceivedTable,
			Columns: users.RequestsReceivedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(request.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.BlogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   users.BlogTable,
			Columns: []string{users.BlogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blog.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   users.NotificationsTable,
			Columns: users.NotificationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ActivityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   users.ActivityTable,
			Columns: users.ActivityPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activity.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   users.StudentsTable,
			Columns: []string{users.StudentsColumn},
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
	if nodes := uc.mutation.ProfessorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   users.ProfessorTable,
			Columns: []string{users.ProfessorColumn},
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
	return _node, _spec
}

// UsersCreateBulk is the builder for creating many Users entities in bulk.
type UsersCreateBulk struct {
	config
	err      error
	builders []*UsersCreate
}

// Save creates the Users entities in the database.
func (ucb *UsersCreateBulk) Save(ctx context.Context) ([]*Users, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*Users, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UsersMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UsersCreateBulk) SaveX(ctx context.Context) []*Users {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UsersCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UsersCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
