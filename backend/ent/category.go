// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/category"
	"backend/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Category is the model entity for the Category schema.
type Category struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CategoryQuery when eager-loading is set.
	Edges           CategoryEdges `json:"edges"`
	user_categories *int
	selectValues    sql.SelectValues
}

// CategoryEdges holds the relations/edges for other nodes in the graph.
type CategoryEdges struct {
	// Skills holds the value of the skills edge.
	Skills []*Skill `json:"skills,omitempty"`
	// Tasks holds the value of the tasks edge.
	Tasks []*Task `json:"tasks,omitempty"`
	// Goals holds the value of the goals edge.
	Goals []*Goal `json:"goals,omitempty"`
	// Hobbies holds the value of the hobbies edge.
	Hobbies []*Hobby `json:"hobbies,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// SkillsOrErr returns the Skills value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) SkillsOrErr() ([]*Skill, error) {
	if e.loadedTypes[0] {
		return e.Skills, nil
	}
	return nil, &NotLoadedError{edge: "skills"}
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[1] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// GoalsOrErr returns the Goals value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) GoalsOrErr() ([]*Goal, error) {
	if e.loadedTypes[2] {
		return e.Goals, nil
	}
	return nil, &NotLoadedError{edge: "goals"}
}

// HobbiesOrErr returns the Hobbies value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) HobbiesOrErr() ([]*Hobby, error) {
	if e.loadedTypes[3] {
		return e.Hobbies, nil
	}
	return nil, &NotLoadedError{edge: "hobbies"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CategoryEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[4] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Category) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case category.FieldID:
			values[i] = new(sql.NullInt64)
		case category.FieldName:
			values[i] = new(sql.NullString)
		case category.ForeignKeys[0]: // user_categories
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Category fields.
func (c *Category) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case category.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case category.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case category.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_categories", value)
			} else if value.Valid {
				c.user_categories = new(int)
				*c.user_categories = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Category.
// This includes values selected through modifiers, order, etc.
func (c *Category) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QuerySkills queries the "skills" edge of the Category entity.
func (c *Category) QuerySkills() *SkillQuery {
	return NewCategoryClient(c.config).QuerySkills(c)
}

// QueryTasks queries the "tasks" edge of the Category entity.
func (c *Category) QueryTasks() *TaskQuery {
	return NewCategoryClient(c.config).QueryTasks(c)
}

// QueryGoals queries the "goals" edge of the Category entity.
func (c *Category) QueryGoals() *GoalQuery {
	return NewCategoryClient(c.config).QueryGoals(c)
}

// QueryHobbies queries the "hobbies" edge of the Category entity.
func (c *Category) QueryHobbies() *HobbyQuery {
	return NewCategoryClient(c.config).QueryHobbies(c)
}

// QueryUser queries the "user" edge of the Category entity.
func (c *Category) QueryUser() *UserQuery {
	return NewCategoryClient(c.config).QueryUser(c)
}

// Update returns a builder for updating this Category.
// Note that you need to call Category.Unwrap() before calling this method if this Category
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Category) Update() *CategoryUpdateOne {
	return NewCategoryClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Category entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Category) Unwrap() *Category {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Category is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Category) String() string {
	var builder strings.Builder
	builder.WriteString("Category(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Categories is a parsable slice of Category.
type Categories []*Category
