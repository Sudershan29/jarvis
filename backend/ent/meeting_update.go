// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/meeting"
	"backend/ent/predicate"
	"backend/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MeetingUpdate is the builder for updating Meeting entities.
type MeetingUpdate struct {
	config
	hooks    []Hook
	mutation *MeetingMutation
}

// Where appends a list predicates to the MeetingUpdate builder.
func (mu *MeetingUpdate) Where(ps ...predicate.Meeting) *MeetingUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetName sets the "name" field.
func (mu *MeetingUpdate) SetName(s string) *MeetingUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetDescription sets the "description" field.
func (mu *MeetingUpdate) SetDescription(s string) *MeetingUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mu *MeetingUpdate) SetNillableDescription(s *string) *MeetingUpdate {
	if s != nil {
		mu.SetDescription(*s)
	}
	return mu
}

// ClearDescription clears the value of the "description" field.
func (mu *MeetingUpdate) ClearDescription() *MeetingUpdate {
	mu.mutation.ClearDescription()
	return mu
}

// SetWhere sets the "where" field.
func (mu *MeetingUpdate) SetWhere(s string) *MeetingUpdate {
	mu.mutation.SetWhere(s)
	return mu
}

// SetNillableWhere sets the "where" field if the given value is not nil.
func (mu *MeetingUpdate) SetNillableWhere(s *string) *MeetingUpdate {
	if s != nil {
		mu.SetWhere(*s)
	}
	return mu
}

// ClearWhere clears the value of the "where" field.
func (mu *MeetingUpdate) ClearWhere() *MeetingUpdate {
	mu.mutation.ClearWhere()
	return mu
}

// SetWhom sets the "whom" field.
func (mu *MeetingUpdate) SetWhom(s string) *MeetingUpdate {
	mu.mutation.SetWhom(s)
	return mu
}

// SetNillableWhom sets the "whom" field if the given value is not nil.
func (mu *MeetingUpdate) SetNillableWhom(s *string) *MeetingUpdate {
	if s != nil {
		mu.SetWhom(*s)
	}
	return mu
}

// ClearWhom clears the value of the "whom" field.
func (mu *MeetingUpdate) ClearWhom() *MeetingUpdate {
	mu.mutation.ClearWhom()
	return mu
}

// SetDuration sets the "duration" field.
func (mu *MeetingUpdate) SetDuration(i int) *MeetingUpdate {
	mu.mutation.ResetDuration()
	mu.mutation.SetDuration(i)
	return mu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (mu *MeetingUpdate) SetNillableDuration(i *int) *MeetingUpdate {
	if i != nil {
		mu.SetDuration(*i)
	}
	return mu
}

// AddDuration adds i to the "duration" field.
func (mu *MeetingUpdate) AddDuration(i int) *MeetingUpdate {
	mu.mutation.AddDuration(i)
	return mu
}

// SetCreatedAt sets the "created_at" field.
func (mu *MeetingUpdate) SetCreatedAt(t time.Time) *MeetingUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mu *MeetingUpdate) SetNillableCreatedAt(t *time.Time) *MeetingUpdate {
	if t != nil {
		mu.SetCreatedAt(*t)
	}
	return mu
}

// SetWhen sets the "when" field.
func (mu *MeetingUpdate) SetWhen(t time.Time) *MeetingUpdate {
	mu.mutation.SetWhen(t)
	return mu
}

// SetNillableWhen sets the "when" field if the given value is not nil.
func (mu *MeetingUpdate) SetNillableWhen(t *time.Time) *MeetingUpdate {
	if t != nil {
		mu.SetWhen(*t)
	}
	return mu
}

// ClearWhen clears the value of the "when" field.
func (mu *MeetingUpdate) ClearWhen() *MeetingUpdate {
	mu.mutation.ClearWhen()
	return mu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (mu *MeetingUpdate) SetUserID(id int) *MeetingUpdate {
	mu.mutation.SetUserID(id)
	return mu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (mu *MeetingUpdate) SetNillableUserID(id *int) *MeetingUpdate {
	if id != nil {
		mu = mu.SetUserID(*id)
	}
	return mu
}

// SetUser sets the "user" edge to the User entity.
func (mu *MeetingUpdate) SetUser(u *User) *MeetingUpdate {
	return mu.SetUserID(u.ID)
}

// Mutation returns the MeetingMutation object of the builder.
func (mu *MeetingUpdate) Mutation() *MeetingMutation {
	return mu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (mu *MeetingUpdate) ClearUser() *MeetingUpdate {
	mu.mutation.ClearUser()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MeetingUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MeetingUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MeetingUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MeetingUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MeetingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(meeting.Table, meeting.Columns, sqlgraph.NewFieldSpec(meeting.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(meeting.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.SetField(meeting.FieldDescription, field.TypeString, value)
	}
	if mu.mutation.DescriptionCleared() {
		_spec.ClearField(meeting.FieldDescription, field.TypeString)
	}
	if value, ok := mu.mutation.GetWhere(); ok {
		_spec.SetField(meeting.FieldWhere, field.TypeString, value)
	}
	if mu.mutation.WhereCleared() {
		_spec.ClearField(meeting.FieldWhere, field.TypeString)
	}
	if value, ok := mu.mutation.Whom(); ok {
		_spec.SetField(meeting.FieldWhom, field.TypeString, value)
	}
	if mu.mutation.WhomCleared() {
		_spec.ClearField(meeting.FieldWhom, field.TypeString)
	}
	if value, ok := mu.mutation.Duration(); ok {
		_spec.SetField(meeting.FieldDuration, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedDuration(); ok {
		_spec.AddField(meeting.FieldDuration, field.TypeInt, value)
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.SetField(meeting.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.When(); ok {
		_spec.SetField(meeting.FieldWhen, field.TypeTime, value)
	}
	if mu.mutation.WhenCleared() {
		_spec.ClearField(meeting.FieldWhen, field.TypeTime)
	}
	if mu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meeting.UserTable,
			Columns: []string{meeting.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meeting.UserTable,
			Columns: []string{meeting.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{meeting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MeetingUpdateOne is the builder for updating a single Meeting entity.
type MeetingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MeetingMutation
}

// SetName sets the "name" field.
func (muo *MeetingUpdateOne) SetName(s string) *MeetingUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetDescription sets the "description" field.
func (muo *MeetingUpdateOne) SetDescription(s string) *MeetingUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (muo *MeetingUpdateOne) SetNillableDescription(s *string) *MeetingUpdateOne {
	if s != nil {
		muo.SetDescription(*s)
	}
	return muo
}

// ClearDescription clears the value of the "description" field.
func (muo *MeetingUpdateOne) ClearDescription() *MeetingUpdateOne {
	muo.mutation.ClearDescription()
	return muo
}

// SetWhere sets the "where" field.
func (muo *MeetingUpdateOne) SetWhere(s string) *MeetingUpdateOne {
	muo.mutation.SetWhere(s)
	return muo
}

// SetNillableWhere sets the "where" field if the given value is not nil.
func (muo *MeetingUpdateOne) SetNillableWhere(s *string) *MeetingUpdateOne {
	if s != nil {
		muo.SetWhere(*s)
	}
	return muo
}

// ClearWhere clears the value of the "where" field.
func (muo *MeetingUpdateOne) ClearWhere() *MeetingUpdateOne {
	muo.mutation.ClearWhere()
	return muo
}

// SetWhom sets the "whom" field.
func (muo *MeetingUpdateOne) SetWhom(s string) *MeetingUpdateOne {
	muo.mutation.SetWhom(s)
	return muo
}

// SetNillableWhom sets the "whom" field if the given value is not nil.
func (muo *MeetingUpdateOne) SetNillableWhom(s *string) *MeetingUpdateOne {
	if s != nil {
		muo.SetWhom(*s)
	}
	return muo
}

// ClearWhom clears the value of the "whom" field.
func (muo *MeetingUpdateOne) ClearWhom() *MeetingUpdateOne {
	muo.mutation.ClearWhom()
	return muo
}

// SetDuration sets the "duration" field.
func (muo *MeetingUpdateOne) SetDuration(i int) *MeetingUpdateOne {
	muo.mutation.ResetDuration()
	muo.mutation.SetDuration(i)
	return muo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (muo *MeetingUpdateOne) SetNillableDuration(i *int) *MeetingUpdateOne {
	if i != nil {
		muo.SetDuration(*i)
	}
	return muo
}

// AddDuration adds i to the "duration" field.
func (muo *MeetingUpdateOne) AddDuration(i int) *MeetingUpdateOne {
	muo.mutation.AddDuration(i)
	return muo
}

// SetCreatedAt sets the "created_at" field.
func (muo *MeetingUpdateOne) SetCreatedAt(t time.Time) *MeetingUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (muo *MeetingUpdateOne) SetNillableCreatedAt(t *time.Time) *MeetingUpdateOne {
	if t != nil {
		muo.SetCreatedAt(*t)
	}
	return muo
}

// SetWhen sets the "when" field.
func (muo *MeetingUpdateOne) SetWhen(t time.Time) *MeetingUpdateOne {
	muo.mutation.SetWhen(t)
	return muo
}

// SetNillableWhen sets the "when" field if the given value is not nil.
func (muo *MeetingUpdateOne) SetNillableWhen(t *time.Time) *MeetingUpdateOne {
	if t != nil {
		muo.SetWhen(*t)
	}
	return muo
}

// ClearWhen clears the value of the "when" field.
func (muo *MeetingUpdateOne) ClearWhen() *MeetingUpdateOne {
	muo.mutation.ClearWhen()
	return muo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (muo *MeetingUpdateOne) SetUserID(id int) *MeetingUpdateOne {
	muo.mutation.SetUserID(id)
	return muo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (muo *MeetingUpdateOne) SetNillableUserID(id *int) *MeetingUpdateOne {
	if id != nil {
		muo = muo.SetUserID(*id)
	}
	return muo
}

// SetUser sets the "user" edge to the User entity.
func (muo *MeetingUpdateOne) SetUser(u *User) *MeetingUpdateOne {
	return muo.SetUserID(u.ID)
}

// Mutation returns the MeetingMutation object of the builder.
func (muo *MeetingUpdateOne) Mutation() *MeetingMutation {
	return muo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (muo *MeetingUpdateOne) ClearUser() *MeetingUpdateOne {
	muo.mutation.ClearUser()
	return muo
}

// Where appends a list predicates to the MeetingUpdate builder.
func (muo *MeetingUpdateOne) Where(ps ...predicate.Meeting) *MeetingUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MeetingUpdateOne) Select(field string, fields ...string) *MeetingUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Meeting entity.
func (muo *MeetingUpdateOne) Save(ctx context.Context) (*Meeting, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MeetingUpdateOne) SaveX(ctx context.Context) *Meeting {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MeetingUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MeetingUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MeetingUpdateOne) sqlSave(ctx context.Context) (_node *Meeting, err error) {
	_spec := sqlgraph.NewUpdateSpec(meeting.Table, meeting.Columns, sqlgraph.NewFieldSpec(meeting.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Meeting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, meeting.FieldID)
		for _, f := range fields {
			if !meeting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != meeting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(meeting.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.SetField(meeting.FieldDescription, field.TypeString, value)
	}
	if muo.mutation.DescriptionCleared() {
		_spec.ClearField(meeting.FieldDescription, field.TypeString)
	}
	if value, ok := muo.mutation.GetWhere(); ok {
		_spec.SetField(meeting.FieldWhere, field.TypeString, value)
	}
	if muo.mutation.WhereCleared() {
		_spec.ClearField(meeting.FieldWhere, field.TypeString)
	}
	if value, ok := muo.mutation.Whom(); ok {
		_spec.SetField(meeting.FieldWhom, field.TypeString, value)
	}
	if muo.mutation.WhomCleared() {
		_spec.ClearField(meeting.FieldWhom, field.TypeString)
	}
	if value, ok := muo.mutation.Duration(); ok {
		_spec.SetField(meeting.FieldDuration, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedDuration(); ok {
		_spec.AddField(meeting.FieldDuration, field.TypeInt, value)
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.SetField(meeting.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.When(); ok {
		_spec.SetField(meeting.FieldWhen, field.TypeTime, value)
	}
	if muo.mutation.WhenCleared() {
		_spec.ClearField(meeting.FieldWhen, field.TypeTime)
	}
	if muo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meeting.UserTable,
			Columns: []string{meeting.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meeting.UserTable,
			Columns: []string{meeting.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Meeting{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{meeting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
