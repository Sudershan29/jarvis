// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/proposal"
	"backend/ent/skill"
	"backend/ent/task"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProposalCreate is the builder for creating a Proposal entity.
type ProposalCreate struct {
	config
	mutation *ProposalMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *ProposalCreate) SetName(s string) *ProposalCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetAllocatedDuration sets the "allocated_duration" field.
func (pc *ProposalCreate) SetAllocatedDuration(i int) *ProposalCreate {
	pc.mutation.SetAllocatedDuration(i)
	return pc
}

// SetAchievedDuration sets the "achieved_duration" field.
func (pc *ProposalCreate) SetAchievedDuration(i int) *ProposalCreate {
	pc.mutation.SetAchievedDuration(i)
	return pc
}

// SetNillableAchievedDuration sets the "achieved_duration" field if the given value is not nil.
func (pc *ProposalCreate) SetNillableAchievedDuration(i *int) *ProposalCreate {
	if i != nil {
		pc.SetAchievedDuration(*i)
	}
	return pc
}

// SetStatus sets the "status" field.
func (pc *ProposalCreate) SetStatus(pr proposal.Status) *ProposalCreate {
	pc.mutation.SetStatus(pr)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *ProposalCreate) SetNillableStatus(pr *proposal.Status) *ProposalCreate {
	if pr != nil {
		pc.SetStatus(*pr)
	}
	return pc
}

// SetScheduledFor sets the "scheduled_for" field.
func (pc *ProposalCreate) SetScheduledFor(t time.Time) *ProposalCreate {
	pc.mutation.SetScheduledFor(t)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProposalCreate) SetCreatedAt(t time.Time) *ProposalCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProposalCreate) SetNillableCreatedAt(t *time.Time) *ProposalCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProposalCreate) SetUpdatedAt(t time.Time) *ProposalCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProposalCreate) SetNillableUpdatedAt(t *time.Time) *ProposalCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (pc *ProposalCreate) SetTaskID(id int) *ProposalCreate {
	pc.mutation.SetTaskID(id)
	return pc
}

// SetNillableTaskID sets the "task" edge to the Task entity by ID if the given value is not nil.
func (pc *ProposalCreate) SetNillableTaskID(id *int) *ProposalCreate {
	if id != nil {
		pc = pc.SetTaskID(*id)
	}
	return pc
}

// SetTask sets the "task" edge to the Task entity.
func (pc *ProposalCreate) SetTask(t *Task) *ProposalCreate {
	return pc.SetTaskID(t.ID)
}

// SetSkillID sets the "skill" edge to the Skill entity by ID.
func (pc *ProposalCreate) SetSkillID(id int) *ProposalCreate {
	pc.mutation.SetSkillID(id)
	return pc
}

// SetNillableSkillID sets the "skill" edge to the Skill entity by ID if the given value is not nil.
func (pc *ProposalCreate) SetNillableSkillID(id *int) *ProposalCreate {
	if id != nil {
		pc = pc.SetSkillID(*id)
	}
	return pc
}

// SetSkill sets the "skill" edge to the Skill entity.
func (pc *ProposalCreate) SetSkill(s *Skill) *ProposalCreate {
	return pc.SetSkillID(s.ID)
}

// Mutation returns the ProposalMutation object of the builder.
func (pc *ProposalCreate) Mutation() *ProposalMutation {
	return pc.mutation
}

// Save creates the Proposal in the database.
func (pc *ProposalCreate) Save(ctx context.Context) (*Proposal, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProposalCreate) SaveX(ctx context.Context) *Proposal {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProposalCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProposalCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProposalCreate) defaults() {
	if _, ok := pc.mutation.Status(); !ok {
		v := proposal.DefaultStatus
		pc.mutation.SetStatus(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := proposal.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := proposal.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProposalCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Proposal.name"`)}
	}
	if _, ok := pc.mutation.AllocatedDuration(); !ok {
		return &ValidationError{Name: "allocated_duration", err: errors.New(`ent: missing required field "Proposal.allocated_duration"`)}
	}
	if v, ok := pc.mutation.AllocatedDuration(); ok {
		if err := proposal.AllocatedDurationValidator(v); err != nil {
			return &ValidationError{Name: "allocated_duration", err: fmt.Errorf(`ent: validator failed for field "Proposal.allocated_duration": %w`, err)}
		}
	}
	if v, ok := pc.mutation.AchievedDuration(); ok {
		if err := proposal.AchievedDurationValidator(v); err != nil {
			return &ValidationError{Name: "achieved_duration", err: fmt.Errorf(`ent: validator failed for field "Proposal.achieved_duration": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Proposal.status"`)}
	}
	if v, ok := pc.mutation.Status(); ok {
		if err := proposal.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Proposal.status": %w`, err)}
		}
	}
	if _, ok := pc.mutation.ScheduledFor(); !ok {
		return &ValidationError{Name: "scheduled_for", err: errors.New(`ent: missing required field "Proposal.scheduled_for"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Proposal.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Proposal.updated_at"`)}
	}
	return nil
}

func (pc *ProposalCreate) sqlSave(ctx context.Context) (*Proposal, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProposalCreate) createSpec() (*Proposal, *sqlgraph.CreateSpec) {
	var (
		_node = &Proposal{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(proposal.Table, sqlgraph.NewFieldSpec(proposal.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(proposal.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.AllocatedDuration(); ok {
		_spec.SetField(proposal.FieldAllocatedDuration, field.TypeInt, value)
		_node.AllocatedDuration = value
	}
	if value, ok := pc.mutation.AchievedDuration(); ok {
		_spec.SetField(proposal.FieldAchievedDuration, field.TypeInt, value)
		_node.AchievedDuration = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(proposal.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := pc.mutation.ScheduledFor(); ok {
		_spec.SetField(proposal.FieldScheduledFor, field.TypeTime, value)
		_node.ScheduledFor = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(proposal.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(proposal.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   proposal.TaskTable,
			Columns: []string{proposal.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.task_proposals = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SkillIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   proposal.SkillTable,
			Columns: []string{proposal.SkillColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.skill_proposals = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProposalCreateBulk is the builder for creating many Proposal entities in bulk.
type ProposalCreateBulk struct {
	config
	builders []*ProposalCreate
}

// Save creates the Proposal entities in the database.
func (pcb *ProposalCreateBulk) Save(ctx context.Context) ([]*Proposal, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Proposal, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProposalMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProposalCreateBulk) SaveX(ctx context.Context) []*Proposal {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProposalCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProposalCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
