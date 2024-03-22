// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/predicate"
	"backend/ent/timepreference"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TimePreferenceDelete is the builder for deleting a TimePreference entity.
type TimePreferenceDelete struct {
	config
	hooks    []Hook
	mutation *TimePreferenceMutation
}

// Where appends a list predicates to the TimePreferenceDelete builder.
func (tpd *TimePreferenceDelete) Where(ps ...predicate.TimePreference) *TimePreferenceDelete {
	tpd.mutation.Where(ps...)
	return tpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tpd *TimePreferenceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tpd.sqlExec, tpd.mutation, tpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tpd *TimePreferenceDelete) ExecX(ctx context.Context) int {
	n, err := tpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tpd *TimePreferenceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(timepreference.Table, sqlgraph.NewFieldSpec(timepreference.FieldID, field.TypeInt))
	if ps := tpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tpd.mutation.done = true
	return affected, err
}

// TimePreferenceDeleteOne is the builder for deleting a single TimePreference entity.
type TimePreferenceDeleteOne struct {
	tpd *TimePreferenceDelete
}

// Where appends a list predicates to the TimePreferenceDelete builder.
func (tpdo *TimePreferenceDeleteOne) Where(ps ...predicate.TimePreference) *TimePreferenceDeleteOne {
	tpdo.tpd.mutation.Where(ps...)
	return tpdo
}

// Exec executes the deletion query.
func (tpdo *TimePreferenceDeleteOne) Exec(ctx context.Context) error {
	n, err := tpdo.tpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{timepreference.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tpdo *TimePreferenceDeleteOne) ExecX(ctx context.Context) {
	if err := tpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
