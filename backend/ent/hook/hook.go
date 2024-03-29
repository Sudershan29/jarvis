// Code generated by ent, DO NOT EDIT.

package hook

import (
	"backend/ent"
	"context"
	"fmt"
)

// The CalendarFunc type is an adapter to allow the use of ordinary
// function as Calendar mutator.
type CalendarFunc func(context.Context, *ent.CalendarMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CalendarFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CalendarMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CalendarMutation", m)
}

// The CategoryFunc type is an adapter to allow the use of ordinary
// function as Category mutator.
type CategoryFunc func(context.Context, *ent.CategoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CategoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CategoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CategoryMutation", m)
}

// The GoalFunc type is an adapter to allow the use of ordinary
// function as Goal mutator.
type GoalFunc func(context.Context, *ent.GoalMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoalFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GoalMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoalMutation", m)
}

// The HobbyFunc type is an adapter to allow the use of ordinary
// function as Hobby mutator.
type HobbyFunc func(context.Context, *ent.HobbyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HobbyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.HobbyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HobbyMutation", m)
}

// The MeetingFunc type is an adapter to allow the use of ordinary
// function as Meeting mutator.
type MeetingFunc func(context.Context, *ent.MeetingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MeetingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MeetingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MeetingMutation", m)
}

// The PreferenceFunc type is an adapter to allow the use of ordinary
// function as Preference mutator.
type PreferenceFunc func(context.Context, *ent.PreferenceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PreferenceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PreferenceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PreferenceMutation", m)
}

// The ProposalFunc type is an adapter to allow the use of ordinary
// function as Proposal mutator.
type ProposalFunc func(context.Context, *ent.ProposalMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProposalFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProposalMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProposalMutation", m)
}

// The SkillFunc type is an adapter to allow the use of ordinary
// function as Skill mutator.
type SkillFunc func(context.Context, *ent.SkillMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SkillFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SkillMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SkillMutation", m)
}

// The TaskFunc type is an adapter to allow the use of ordinary
// function as Task mutator.
type TaskFunc func(context.Context, *ent.TaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TaskMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskMutation", m)
}

// The TimePreferenceFunc type is an adapter to allow the use of ordinary
// function as TimePreference mutator.
type TimePreferenceFunc func(context.Context, *ent.TimePreferenceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TimePreferenceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TimePreferenceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TimePreferenceMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
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
