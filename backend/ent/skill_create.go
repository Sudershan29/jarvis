// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/category"
	"backend/ent/proposal"
	"backend/ent/skill"
	"backend/ent/timepreference"
	"backend/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SkillCreate is the builder for creating a Skill entity.
type SkillCreate struct {
	config
	mutation *SkillMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *SkillCreate) SetName(s string) *SkillCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetLevel sets the "level" field.
func (sc *SkillCreate) SetLevel(s string) *SkillCreate {
	sc.mutation.SetLevel(s)
	return sc
}

// SetProgress sets the "progress" field.
func (sc *SkillCreate) SetProgress(i int) *SkillCreate {
	sc.mutation.SetProgress(i)
	return sc
}

// SetNillableProgress sets the "progress" field if the given value is not nil.
func (sc *SkillCreate) SetNillableProgress(i *int) *SkillCreate {
	if i != nil {
		sc.SetProgress(*i)
	}
	return sc
}

// SetDuration sets the "duration" field.
func (sc *SkillCreate) SetDuration(i int) *SkillCreate {
	sc.mutation.SetDuration(i)
	return sc
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (sc *SkillCreate) SetNillableDuration(i *int) *SkillCreate {
	if i != nil {
		sc.SetDuration(*i)
	}
	return sc
}

// SetDurationAchieved sets the "duration_achieved" field.
func (sc *SkillCreate) SetDurationAchieved(i int) *SkillCreate {
	sc.mutation.SetDurationAchieved(i)
	return sc
}

// SetNillableDurationAchieved sets the "duration_achieved" field if the given value is not nil.
func (sc *SkillCreate) SetNillableDurationAchieved(i *int) *SkillCreate {
	if i != nil {
		sc.SetDurationAchieved(*i)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SkillCreate) SetCreatedAt(t time.Time) *SkillCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SkillCreate) SetNillableCreatedAt(t *time.Time) *SkillCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (sc *SkillCreate) AddCategoryIDs(ids ...int) *SkillCreate {
	sc.mutation.AddCategoryIDs(ids...)
	return sc
}

// AddCategories adds the "categories" edges to the Category entity.
func (sc *SkillCreate) AddCategories(c ...*Category) *SkillCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddCategoryIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (sc *SkillCreate) SetUserID(id int) *SkillCreate {
	sc.mutation.SetUserID(id)
	return sc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (sc *SkillCreate) SetNillableUserID(id *int) *SkillCreate {
	if id != nil {
		sc = sc.SetUserID(*id)
	}
	return sc
}

// SetUser sets the "user" edge to the User entity.
func (sc *SkillCreate) SetUser(u *User) *SkillCreate {
	return sc.SetUserID(u.ID)
}

// AddTimePreferenceIDs adds the "time_preferences" edge to the TimePreference entity by IDs.
func (sc *SkillCreate) AddTimePreferenceIDs(ids ...int) *SkillCreate {
	sc.mutation.AddTimePreferenceIDs(ids...)
	return sc
}

// AddTimePreferences adds the "time_preferences" edges to the TimePreference entity.
func (sc *SkillCreate) AddTimePreferences(t ...*TimePreference) *SkillCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddTimePreferenceIDs(ids...)
}

// AddProposalIDs adds the "proposals" edge to the Proposal entity by IDs.
func (sc *SkillCreate) AddProposalIDs(ids ...int) *SkillCreate {
	sc.mutation.AddProposalIDs(ids...)
	return sc
}

// AddProposals adds the "proposals" edges to the Proposal entity.
func (sc *SkillCreate) AddProposals(p ...*Proposal) *SkillCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddProposalIDs(ids...)
}

// Mutation returns the SkillMutation object of the builder.
func (sc *SkillCreate) Mutation() *SkillMutation {
	return sc.mutation
}

// Save creates the Skill in the database.
func (sc *SkillCreate) Save(ctx context.Context) (*Skill, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SkillCreate) SaveX(ctx context.Context) *Skill {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SkillCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SkillCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SkillCreate) defaults() {
	if _, ok := sc.mutation.Progress(); !ok {
		v := skill.DefaultProgress
		sc.mutation.SetProgress(v)
	}
	if _, ok := sc.mutation.Duration(); !ok {
		v := skill.DefaultDuration
		sc.mutation.SetDuration(v)
	}
	if _, ok := sc.mutation.DurationAchieved(); !ok {
		v := skill.DefaultDurationAchieved
		sc.mutation.SetDurationAchieved(v)
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := skill.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SkillCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Skill.name"`)}
	}
	if _, ok := sc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "Skill.level"`)}
	}
	if _, ok := sc.mutation.Progress(); !ok {
		return &ValidationError{Name: "progress", err: errors.New(`ent: missing required field "Skill.progress"`)}
	}
	if _, ok := sc.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New(`ent: missing required field "Skill.duration"`)}
	}
	if _, ok := sc.mutation.DurationAchieved(); !ok {
		return &ValidationError{Name: "duration_achieved", err: errors.New(`ent: missing required field "Skill.duration_achieved"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Skill.created_at"`)}
	}
	return nil
}

func (sc *SkillCreate) sqlSave(ctx context.Context) (*Skill, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SkillCreate) createSpec() (*Skill, *sqlgraph.CreateSpec) {
	var (
		_node = &Skill{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(skill.Table, sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(skill.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Level(); ok {
		_spec.SetField(skill.FieldLevel, field.TypeString, value)
		_node.Level = value
	}
	if value, ok := sc.mutation.Progress(); ok {
		_spec.SetField(skill.FieldProgress, field.TypeInt, value)
		_node.Progress = value
	}
	if value, ok := sc.mutation.Duration(); ok {
		_spec.SetField(skill.FieldDuration, field.TypeInt, value)
		_node.Duration = value
	}
	if value, ok := sc.mutation.DurationAchieved(); ok {
		_spec.SetField(skill.FieldDurationAchieved, field.TypeInt, value)
		_node.DurationAchieved = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(skill.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := sc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   skill.CategoriesTable,
			Columns: skill.CategoriesPrimaryKey,
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
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   skill.UserTable,
			Columns: []string{skill.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_skills = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.TimePreferencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   skill.TimePreferencesTable,
			Columns: skill.TimePreferencesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timepreference.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ProposalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   skill.ProposalsTable,
			Columns: []string{skill.ProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proposal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SkillCreateBulk is the builder for creating many Skill entities in bulk.
type SkillCreateBulk struct {
	config
	builders []*SkillCreate
}

// Save creates the Skill entities in the database.
func (scb *SkillCreateBulk) Save(ctx context.Context) ([]*Skill, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Skill, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SkillMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SkillCreateBulk) SaveX(ctx context.Context) []*Skill {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SkillCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SkillCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
