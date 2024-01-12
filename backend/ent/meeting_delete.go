// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/meeting"
	"backend/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MeetingDelete is the builder for deleting a Meeting entity.
type MeetingDelete struct {
	config
	hooks    []Hook
	mutation *MeetingMutation
}

// Where appends a list predicates to the MeetingDelete builder.
func (md *MeetingDelete) Where(ps ...predicate.Meeting) *MeetingDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MeetingDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, md.sqlExec, md.mutation, md.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MeetingDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MeetingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(meeting.Table, sqlgraph.NewFieldSpec(meeting.FieldID, field.TypeInt))
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, md.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	md.mutation.done = true
	return affected, err
}

// MeetingDeleteOne is the builder for deleting a single Meeting entity.
type MeetingDeleteOne struct {
	md *MeetingDelete
}

// Where appends a list predicates to the MeetingDelete builder.
func (mdo *MeetingDeleteOne) Where(ps ...predicate.Meeting) *MeetingDeleteOne {
	mdo.md.mutation.Where(ps...)
	return mdo
}

// Exec executes the deletion query.
func (mdo *MeetingDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{meeting.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MeetingDeleteOne) ExecX(ctx context.Context) {
	if err := mdo.Exec(ctx); err != nil {
		panic(err)
	}
}