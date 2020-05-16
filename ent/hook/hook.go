// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/minskylab/collecta/ent"
)

// The AccountFunc type is an adapter to allow the use of ordinary
// function as Account mutator.
type AccountFunc func(context.Context, *ent.AccountMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AccountFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AccountMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AccountMutation", m)
	}
	return f(ctx, mv)
}

// The AnswerFunc type is an adapter to allow the use of ordinary
// function as Answer mutator.
type AnswerFunc func(context.Context, *ent.AnswerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AnswerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AnswerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AnswerMutation", m)
	}
	return f(ctx, mv)
}

// The ContactFunc type is an adapter to allow the use of ordinary
// function as Contact mutator.
type ContactFunc func(context.Context, *ent.ContactMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ContactFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ContactMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ContactMutation", m)
	}
	return f(ctx, mv)
}

// The DeviceFunc type is an adapter to allow the use of ordinary
// function as Device mutator.
type DeviceFunc func(context.Context, *ent.DeviceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DeviceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DeviceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DeviceMutation", m)
	}
	return f(ctx, mv)
}

// The DomainFunc type is an adapter to allow the use of ordinary
// function as Domain mutator.
type DomainFunc func(context.Context, *ent.DomainMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DomainFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DomainMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DomainMutation", m)
	}
	return f(ctx, mv)
}

// The FlowFunc type is an adapter to allow the use of ordinary
// function as Flow mutator.
type FlowFunc func(context.Context, *ent.FlowMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FlowFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FlowMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FlowMutation", m)
	}
	return f(ctx, mv)
}

// The IPFunc type is an adapter to allow the use of ordinary
// function as IP mutator.
type IPFunc func(context.Context, *ent.IPMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f IPFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.IPMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.IPMutation", m)
	}
	return f(ctx, mv)
}

// The InputFunc type is an adapter to allow the use of ordinary
// function as Input mutator.
type InputFunc func(context.Context, *ent.InputMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f InputFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.InputMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.InputMutation", m)
	}
	return f(ctx, mv)
}

// The PersonFunc type is an adapter to allow the use of ordinary
// function as Person mutator.
type PersonFunc func(context.Context, *ent.PersonMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PersonFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PersonMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PersonMutation", m)
	}
	return f(ctx, mv)
}

// The QuestionFunc type is an adapter to allow the use of ordinary
// function as Question mutator.
type QuestionFunc func(context.Context, *ent.QuestionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f QuestionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.QuestionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.QuestionMutation", m)
	}
	return f(ctx, mv)
}

// The ShortFunc type is an adapter to allow the use of ordinary
// function as Short mutator.
type ShortFunc func(context.Context, *ent.ShortMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ShortFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ShortMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ShortMutation", m)
	}
	return f(ctx, mv)
}

// The SurveyFunc type is an adapter to allow the use of ordinary
// function as Survey mutator.
type SurveyFunc func(context.Context, *ent.SurveyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SurveyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SurveyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SurveyMutation", m)
	}
	return f(ctx, mv)
}

// On executes the given hook only of the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Op().Is(op) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
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
//
func Reject(op ent.Op) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Op().Is(op) {
				return nil, fmt.Errorf("%s operation is not allowed", m.Op())
			}
			return next.Mutate(ctx, m)
		})
	}
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
