// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/category"
	"backend/ent/hobby"
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

// HobbyUpdate is the builder for updating Hobby entities.
type HobbyUpdate struct {
	config
	hooks    []Hook
	mutation *HobbyMutation
}

// Where appends a list predicates to the HobbyUpdate builder.
func (hu *HobbyUpdate) Where(ps ...predicate.Hobby) *HobbyUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetName sets the "name" field.
func (hu *HobbyUpdate) SetName(s string) *HobbyUpdate {
	hu.mutation.SetName(s)
	return hu
}

// SetDescription sets the "description" field.
func (hu *HobbyUpdate) SetDescription(s string) *HobbyUpdate {
	hu.mutation.SetDescription(s)
	return hu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (hu *HobbyUpdate) SetNillableDescription(s *string) *HobbyUpdate {
	if s != nil {
		hu.SetDescription(*s)
	}
	return hu
}

// ClearDescription clears the value of the "description" field.
func (hu *HobbyUpdate) ClearDescription() *HobbyUpdate {
	hu.mutation.ClearDescription()
	return hu
}

// SetCreatedAt sets the "created_at" field.
func (hu *HobbyUpdate) SetCreatedAt(t time.Time) *HobbyUpdate {
	hu.mutation.SetCreatedAt(t)
	return hu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hu *HobbyUpdate) SetNillableCreatedAt(t *time.Time) *HobbyUpdate {
	if t != nil {
		hu.SetCreatedAt(*t)
	}
	return hu
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (hu *HobbyUpdate) AddCategoryIDs(ids ...int) *HobbyUpdate {
	hu.mutation.AddCategoryIDs(ids...)
	return hu
}

// AddCategories adds the "categories" edges to the Category entity.
func (hu *HobbyUpdate) AddCategories(c ...*Category) *HobbyUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return hu.AddCategoryIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (hu *HobbyUpdate) SetUserID(id int) *HobbyUpdate {
	hu.mutation.SetUserID(id)
	return hu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (hu *HobbyUpdate) SetNillableUserID(id *int) *HobbyUpdate {
	if id != nil {
		hu = hu.SetUserID(*id)
	}
	return hu
}

// SetUser sets the "user" edge to the User entity.
func (hu *HobbyUpdate) SetUser(u *User) *HobbyUpdate {
	return hu.SetUserID(u.ID)
}

// Mutation returns the HobbyMutation object of the builder.
func (hu *HobbyUpdate) Mutation() *HobbyMutation {
	return hu.mutation
}

// ClearCategories clears all "categories" edges to the Category entity.
func (hu *HobbyUpdate) ClearCategories() *HobbyUpdate {
	hu.mutation.ClearCategories()
	return hu
}

// RemoveCategoryIDs removes the "categories" edge to Category entities by IDs.
func (hu *HobbyUpdate) RemoveCategoryIDs(ids ...int) *HobbyUpdate {
	hu.mutation.RemoveCategoryIDs(ids...)
	return hu
}

// RemoveCategories removes "categories" edges to Category entities.
func (hu *HobbyUpdate) RemoveCategories(c ...*Category) *HobbyUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return hu.RemoveCategoryIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (hu *HobbyUpdate) ClearUser() *HobbyUpdate {
	hu.mutation.ClearUser()
	return hu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HobbyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, hu.sqlSave, hu.mutation, hu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HobbyUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HobbyUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HobbyUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hu *HobbyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(hobby.Table, hobby.Columns, sqlgraph.NewFieldSpec(hobby.FieldID, field.TypeInt))
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.Name(); ok {
		_spec.SetField(hobby.FieldName, field.TypeString, value)
	}
	if value, ok := hu.mutation.Description(); ok {
		_spec.SetField(hobby.FieldDescription, field.TypeString, value)
	}
	if hu.mutation.DescriptionCleared() {
		_spec.ClearField(hobby.FieldDescription, field.TypeString)
	}
	if value, ok := hu.mutation.CreatedAt(); ok {
		_spec.SetField(hobby.FieldCreatedAt, field.TypeTime, value)
	}
	if hu.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hobby.CategoriesTable,
			Columns: hobby.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !hu.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hobby.CategoriesTable,
			Columns: hobby.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hobby.CategoriesTable,
			Columns: hobby.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hobby.UserTable,
			Columns: []string{hobby.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hobby.UserTable,
			Columns: []string{hobby.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hobby.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hu.mutation.done = true
	return n, nil
}

// HobbyUpdateOne is the builder for updating a single Hobby entity.
type HobbyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HobbyMutation
}

// SetName sets the "name" field.
func (huo *HobbyUpdateOne) SetName(s string) *HobbyUpdateOne {
	huo.mutation.SetName(s)
	return huo
}

// SetDescription sets the "description" field.
func (huo *HobbyUpdateOne) SetDescription(s string) *HobbyUpdateOne {
	huo.mutation.SetDescription(s)
	return huo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (huo *HobbyUpdateOne) SetNillableDescription(s *string) *HobbyUpdateOne {
	if s != nil {
		huo.SetDescription(*s)
	}
	return huo
}

// ClearDescription clears the value of the "description" field.
func (huo *HobbyUpdateOne) ClearDescription() *HobbyUpdateOne {
	huo.mutation.ClearDescription()
	return huo
}

// SetCreatedAt sets the "created_at" field.
func (huo *HobbyUpdateOne) SetCreatedAt(t time.Time) *HobbyUpdateOne {
	huo.mutation.SetCreatedAt(t)
	return huo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (huo *HobbyUpdateOne) SetNillableCreatedAt(t *time.Time) *HobbyUpdateOne {
	if t != nil {
		huo.SetCreatedAt(*t)
	}
	return huo
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (huo *HobbyUpdateOne) AddCategoryIDs(ids ...int) *HobbyUpdateOne {
	huo.mutation.AddCategoryIDs(ids...)
	return huo
}

// AddCategories adds the "categories" edges to the Category entity.
func (huo *HobbyUpdateOne) AddCategories(c ...*Category) *HobbyUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return huo.AddCategoryIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (huo *HobbyUpdateOne) SetUserID(id int) *HobbyUpdateOne {
	huo.mutation.SetUserID(id)
	return huo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (huo *HobbyUpdateOne) SetNillableUserID(id *int) *HobbyUpdateOne {
	if id != nil {
		huo = huo.SetUserID(*id)
	}
	return huo
}

// SetUser sets the "user" edge to the User entity.
func (huo *HobbyUpdateOne) SetUser(u *User) *HobbyUpdateOne {
	return huo.SetUserID(u.ID)
}

// Mutation returns the HobbyMutation object of the builder.
func (huo *HobbyUpdateOne) Mutation() *HobbyMutation {
	return huo.mutation
}

// ClearCategories clears all "categories" edges to the Category entity.
func (huo *HobbyUpdateOne) ClearCategories() *HobbyUpdateOne {
	huo.mutation.ClearCategories()
	return huo
}

// RemoveCategoryIDs removes the "categories" edge to Category entities by IDs.
func (huo *HobbyUpdateOne) RemoveCategoryIDs(ids ...int) *HobbyUpdateOne {
	huo.mutation.RemoveCategoryIDs(ids...)
	return huo
}

// RemoveCategories removes "categories" edges to Category entities.
func (huo *HobbyUpdateOne) RemoveCategories(c ...*Category) *HobbyUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return huo.RemoveCategoryIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (huo *HobbyUpdateOne) ClearUser() *HobbyUpdateOne {
	huo.mutation.ClearUser()
	return huo
}

// Where appends a list predicates to the HobbyUpdate builder.
func (huo *HobbyUpdateOne) Where(ps ...predicate.Hobby) *HobbyUpdateOne {
	huo.mutation.Where(ps...)
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HobbyUpdateOne) Select(field string, fields ...string) *HobbyUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated Hobby entity.
func (huo *HobbyUpdateOne) Save(ctx context.Context) (*Hobby, error) {
	return withHooks(ctx, huo.sqlSave, huo.mutation, huo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HobbyUpdateOne) SaveX(ctx context.Context) *Hobby {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HobbyUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HobbyUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (huo *HobbyUpdateOne) sqlSave(ctx context.Context) (_node *Hobby, err error) {
	_spec := sqlgraph.NewUpdateSpec(hobby.Table, hobby.Columns, sqlgraph.NewFieldSpec(hobby.FieldID, field.TypeInt))
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Hobby.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hobby.FieldID)
		for _, f := range fields {
			if !hobby.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != hobby.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.Name(); ok {
		_spec.SetField(hobby.FieldName, field.TypeString, value)
	}
	if value, ok := huo.mutation.Description(); ok {
		_spec.SetField(hobby.FieldDescription, field.TypeString, value)
	}
	if huo.mutation.DescriptionCleared() {
		_spec.ClearField(hobby.FieldDescription, field.TypeString)
	}
	if value, ok := huo.mutation.CreatedAt(); ok {
		_spec.SetField(hobby.FieldCreatedAt, field.TypeTime, value)
	}
	if huo.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hobby.CategoriesTable,
			Columns: hobby.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !huo.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hobby.CategoriesTable,
			Columns: hobby.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hobby.CategoriesTable,
			Columns: hobby.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hobby.UserTable,
			Columns: []string{hobby.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hobby.UserTable,
			Columns: []string{hobby.UserColumn},
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
	_node = &Hobby{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hobby.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	huo.mutation.done = true
	return _node, nil
}
