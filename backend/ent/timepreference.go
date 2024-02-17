// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/timepreference"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// TimePreference is the model entity for the TimePreference schema.
type TimePreference struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Day holds the value of the "day" field.
	Day string `json:"day,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TimePreferenceQuery when eager-loading is set.
	Edges        TimePreferenceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TimePreferenceEdges holds the relations/edges for other nodes in the graph.
type TimePreferenceEdges struct {
	// Skills holds the value of the skills edge.
	Skills []*Skill `json:"skills,omitempty"`
	// Tasks holds the value of the tasks edge.
	Tasks []*Task `json:"tasks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SkillsOrErr returns the Skills value or an error if the edge
// was not loaded in eager-loading.
func (e TimePreferenceEdges) SkillsOrErr() ([]*Skill, error) {
	if e.loadedTypes[0] {
		return e.Skills, nil
	}
	return nil, &NotLoadedError{edge: "skills"}
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e TimePreferenceEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[1] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TimePreference) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case timepreference.FieldID:
			values[i] = new(sql.NullInt64)
		case timepreference.FieldDay:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TimePreference fields.
func (tp *TimePreference) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case timepreference.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tp.ID = int(value.Int64)
		case timepreference.FieldDay:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field day", values[i])
			} else if value.Valid {
				tp.Day = value.String
			}
		default:
			tp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TimePreference.
// This includes values selected through modifiers, order, etc.
func (tp *TimePreference) Value(name string) (ent.Value, error) {
	return tp.selectValues.Get(name)
}

// QuerySkills queries the "skills" edge of the TimePreference entity.
func (tp *TimePreference) QuerySkills() *SkillQuery {
	return NewTimePreferenceClient(tp.config).QuerySkills(tp)
}

// QueryTasks queries the "tasks" edge of the TimePreference entity.
func (tp *TimePreference) QueryTasks() *TaskQuery {
	return NewTimePreferenceClient(tp.config).QueryTasks(tp)
}

// Update returns a builder for updating this TimePreference.
// Note that you need to call TimePreference.Unwrap() before calling this method if this TimePreference
// was returned from a transaction, and the transaction was committed or rolled back.
func (tp *TimePreference) Update() *TimePreferenceUpdateOne {
	return NewTimePreferenceClient(tp.config).UpdateOne(tp)
}

// Unwrap unwraps the TimePreference entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tp *TimePreference) Unwrap() *TimePreference {
	_tx, ok := tp.config.driver.(*txDriver)
	if !ok {
		panic("ent: TimePreference is not a transactional entity")
	}
	tp.config.driver = _tx.drv
	return tp
}

// String implements the fmt.Stringer.
func (tp *TimePreference) String() string {
	var builder strings.Builder
	builder.WriteString("TimePreference(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tp.ID))
	builder.WriteString("day=")
	builder.WriteString(tp.Day)
	builder.WriteByte(')')
	return builder.String()
}

// TimePreferences is a parsable slice of TimePreference.
type TimePreferences []*TimePreference
