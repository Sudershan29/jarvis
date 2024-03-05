// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/skill"
	"backend/ent/task"
	"backend/ent/timepreference"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TimePreferenceCreate is the builder for creating a TimePreference entity.
type TimePreferenceCreate struct {
	config
	mutation *TimePreferenceMutation
	hooks    []Hook
}

// SetDay sets the "day" field.
func (tpc *TimePreferenceCreate) SetDay(s string) *TimePreferenceCreate {
	tpc.mutation.SetDay(s)
	return tpc
}

// AddSkillIDs adds the "skills" edge to the Skill entity by IDs.
func (tpc *TimePreferenceCreate) AddSkillIDs(ids ...int) *TimePreferenceCreate {
	tpc.mutation.AddSkillIDs(ids...)
	return tpc
}

// AddSkills adds the "skills" edges to the Skill entity.
func (tpc *TimePreferenceCreate) AddSkills(s ...*Skill) *TimePreferenceCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tpc.AddSkillIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (tpc *TimePreferenceCreate) AddTaskIDs(ids ...int) *TimePreferenceCreate {
	tpc.mutation.AddTaskIDs(ids...)
	return tpc
}

// AddTasks adds the "tasks" edges to the Task entity.
func (tpc *TimePreferenceCreate) AddTasks(t ...*Task) *TimePreferenceCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tpc.AddTaskIDs(ids...)
}

// Mutation returns the TimePreferenceMutation object of the builder.
func (tpc *TimePreferenceCreate) Mutation() *TimePreferenceMutation {
	return tpc.mutation
}

// Save creates the TimePreference in the database.
func (tpc *TimePreferenceCreate) Save(ctx context.Context) (*TimePreference, error) {
	return withHooks(ctx, tpc.sqlSave, tpc.mutation, tpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tpc *TimePreferenceCreate) SaveX(ctx context.Context) *TimePreference {
	v, err := tpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tpc *TimePreferenceCreate) Exec(ctx context.Context) error {
	_, err := tpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpc *TimePreferenceCreate) ExecX(ctx context.Context) {
	if err := tpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tpc *TimePreferenceCreate) check() error {
	if _, ok := tpc.mutation.Day(); !ok {
		return &ValidationError{Name: "day", err: errors.New(`ent: missing required field "TimePreference.day"`)}
	}
	return nil
}

func (tpc *TimePreferenceCreate) sqlSave(ctx context.Context) (*TimePreference, error) {
	if err := tpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tpc.mutation.id = &_node.ID
	tpc.mutation.done = true
	return _node, nil
}

func (tpc *TimePreferenceCreate) createSpec() (*TimePreference, *sqlgraph.CreateSpec) {
	var (
		_node = &TimePreference{config: tpc.config}
		_spec = sqlgraph.NewCreateSpec(timepreference.Table, sqlgraph.NewFieldSpec(timepreference.FieldID, field.TypeInt))
	)
	if value, ok := tpc.mutation.Day(); ok {
		_spec.SetField(timepreference.FieldDay, field.TypeString, value)
		_node.Day = value
	}
	if nodes := tpc.mutation.SkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timepreference.SkillsTable,
			Columns: timepreference.SkillsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tpc.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   timepreference.TasksTable,
			Columns: timepreference.TasksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TimePreferenceCreateBulk is the builder for creating many TimePreference entities in bulk.
type TimePreferenceCreateBulk struct {
	config
	builders []*TimePreferenceCreate
}

// Save creates the TimePreference entities in the database.
func (tpcb *TimePreferenceCreateBulk) Save(ctx context.Context) ([]*TimePreference, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tpcb.builders))
	nodes := make([]*TimePreference, len(tpcb.builders))
	mutators := make([]Mutator, len(tpcb.builders))
	for i := range tpcb.builders {
		func(i int, root context.Context) {
			builder := tpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TimePreferenceMutation)
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
					_, err = mutators[i+1].Mutate(root, tpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tpcb *TimePreferenceCreateBulk) SaveX(ctx context.Context) []*TimePreference {
	v, err := tpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tpcb *TimePreferenceCreateBulk) Exec(ctx context.Context) error {
	_, err := tpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpcb *TimePreferenceCreateBulk) ExecX(ctx context.Context) {
	if err := tpcb.Exec(ctx); err != nil {
		panic(err)
	}
}