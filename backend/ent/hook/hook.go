// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	"mocku/backend/ent"
)

// The ActivityFunc type is an adapter to allow the use of ordinary
// function as Activity mutator.
type ActivityFunc func(context.Context, *ent.ActivityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ActivityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ActivityMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ActivityMutation", m)
}

// The BlogFunc type is an adapter to allow the use of ordinary
// function as Blog mutator.
type BlogFunc func(context.Context, *ent.BlogMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BlogFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BlogMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BlogMutation", m)
}

// The CareersFunc type is an adapter to allow the use of ordinary
// function as Careers mutator.
type CareersFunc func(context.Context, *ent.CareersMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CareersFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CareersMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CareersMutation", m)
}

// The ConfigurationFunc type is an adapter to allow the use of ordinary
// function as Configuration mutator.
type ConfigurationFunc func(context.Context, *ent.ConfigurationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ConfigurationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ConfigurationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ConfigurationMutation", m)
}

// The CycleFunc type is an adapter to allow the use of ordinary
// function as Cycle mutator.
type CycleFunc func(context.Context, *ent.CycleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CycleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CycleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CycleMutation", m)
}

// The ModuleFunc type is an adapter to allow the use of ordinary
// function as Module mutator.
type ModuleFunc func(context.Context, *ent.ModuleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ModuleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ModuleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ModuleMutation", m)
}

// The NoteFunc type is an adapter to allow the use of ordinary
// function as Note mutator.
type NoteFunc func(context.Context, *ent.NoteMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NoteFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.NoteMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NoteMutation", m)
}

// The NotificationFunc type is an adapter to allow the use of ordinary
// function as Notification mutator.
type NotificationFunc func(context.Context, *ent.NotificationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NotificationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.NotificationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NotificationMutation", m)
}

// The PaymentFunc type is an adapter to allow the use of ordinary
// function as Payment mutator.
type PaymentFunc func(context.Context, *ent.PaymentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PaymentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PaymentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PaymentMutation", m)
}

// The PaymentMethodFunc type is an adapter to allow the use of ordinary
// function as PaymentMethod mutator.
type PaymentMethodFunc func(context.Context, *ent.PaymentMethodMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PaymentMethodFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PaymentMethodMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PaymentMethodMutation", m)
}

// The PermissionFunc type is an adapter to allow the use of ordinary
// function as Permission mutator.
type PermissionFunc func(context.Context, *ent.PermissionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PermissionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PermissionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PermissionMutation", m)
}

// The ProfessorFunc type is an adapter to allow the use of ordinary
// function as Professor mutator.
type ProfessorFunc func(context.Context, *ent.ProfessorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProfessorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProfessorMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProfessorMutation", m)
}

// The RequestFunc type is an adapter to allow the use of ordinary
// function as Request mutator.
type RequestFunc func(context.Context, *ent.RequestMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RequestFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RequestMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RequestMutation", m)
}

// The RoleFunc type is an adapter to allow the use of ordinary
// function as Role mutator.
type RoleFunc func(context.Context, *ent.RoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RoleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoleMutation", m)
}

// The StudentFunc type is an adapter to allow the use of ordinary
// function as Student mutator.
type StudentFunc func(context.Context, *ent.StudentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StudentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.StudentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StudentMutation", m)
}

// The SubjectFunc type is an adapter to allow the use of ordinary
// function as Subject mutator.
type SubjectFunc func(context.Context, *ent.SubjectMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SubjectFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SubjectMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SubjectMutation", m)
}

// The UsersFunc type is an adapter to allow the use of ordinary
// function as Users mutator.
type UsersFunc func(context.Context, *ent.UsersMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UsersFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UsersMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UsersMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
