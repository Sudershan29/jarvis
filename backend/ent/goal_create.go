// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/category"
	"backend/ent/goal"
	"backend/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GoalCreate is the builder for creating a Goal entity.
type GoalCreate struct {
	config
	mutation *GoalMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (gc *GoalCreate) SetName(s string) *GoalCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetDescription sets the "description" field.
func (gc *GoalCreate) SetDescription(s string) *GoalCreate {
	gc.mutation.SetDescription(s)
	return gc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gc *GoalCreate) SetNillableDescription(s *string) *GoalCreate {
	if s != nil {
		gc.SetDescription(*s)
	}
	return gc
}

// SetCreatedAt sets the "created_at" field.
func (gc *GoalCreate) SetCreatedAt(t time.Time) *GoalCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GoalCreate) SetNillableCreatedAt(t *time.Time) *GoalCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (gc *GoalCreate) AddCategoryIDs(ids ...int) *GoalCreate {
	gc.mutation.AddCategoryIDs(ids...)
	return gc
}

// AddCategories adds the "categories" edges to the Category entity.
func (gc *GoalCreate) AddCategories(c ...*Category) *GoalCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return gc.AddCategoryIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gc *GoalCreate) SetUserID(id int) *GoalCreate {
	gc.mutation.SetUserID(id)
	return gc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (gc *GoalCreate) SetNillableUserID(id *int) *GoalCreate {
	if id != nil {
		gc = gc.SetUserID(*id)
	}
	return gc
}

// SetUser sets the "user" edge to the User entity.
func (gc *GoalCreate) SetUser(u *User) *GoalCreate {
	return gc.SetUserID(u.ID)
}

// Mutation returns the GoalMutation object of the builder.
func (gc *GoalCreate) Mutation() *GoalMutation {
	return gc.mutation
}

// Save creates the Goal in the database.
func (gc *GoalCreate) Save(ctx context.Context) (*Goal, error) {
	gc.defaults()
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GoalCreate) SaveX(ctx context.Context) *Goal {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GoalCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GoalCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GoalCreate) defaults() {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := goal.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GoalCreate) check() error {
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Goal.name"`)}
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Goal.created_at"`)}
	}
	return nil
}

func (gc *GoalCreate) sqlSave(ctx context.Context) (*Goal, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GoalCreate) createSpec() (*Goal, *sqlgraph.CreateSpec) {
	var (
		_node = &Goal{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(goal.Table, sqlgraph.NewFieldSpec(goal.FieldID, field.TypeInt))
	)
	if value, ok := gc.mutation.Name(); ok {
		_spec.SetField(goal.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := gc.mutation.Description(); ok {
		_spec.SetField(goal.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.SetField(goal.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := gc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   goal.CategoriesTable,
			Columns: goal.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   goal.UserTable,
			Columns: []string{goal.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_goals = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GoalCreateBulk is the builder for creating many Goal entities in bulk.
type GoalCreateBulk struct {
	config
	builders []*GoalCreate
}

// Save creates the Goal entities in the database.
func (gcb *GoalCreateBulk) Save(ctx context.Context) ([]*Goal, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Goal, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GoalMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GoalCreateBulk) SaveX(ctx context.Context) []*Goal {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GoalCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GoalCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}