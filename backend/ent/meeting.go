// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/meeting"
	"backend/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Meeting is the model entity for the Meeting schema.
type Meeting struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Where holds the value of the "where" field.
	Where string `json:"where,omitempty"`
	// Whom holds the value of the "whom" field.
	Whom string `json:"whom,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration int `json:"duration,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// When holds the value of the "when" field.
	When time.Time `json:"when,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MeetingQuery when eager-loading is set.
	Edges         MeetingEdges `json:"edges"`
	user_meetings *int
	selectValues  sql.SelectValues
}

// MeetingEdges holds the relations/edges for other nodes in the graph.
type MeetingEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MeetingEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Meeting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case meeting.FieldID, meeting.FieldDuration:
			values[i] = new(sql.NullInt64)
		case meeting.FieldName, meeting.FieldDescription, meeting.FieldWhere, meeting.FieldWhom:
			values[i] = new(sql.NullString)
		case meeting.FieldCreatedAt, meeting.FieldWhen:
			values[i] = new(sql.NullTime)
		case meeting.ForeignKeys[0]: // user_meetings
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Meeting fields.
func (m *Meeting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case meeting.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case meeting.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case meeting.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				m.Description = value.String
			}
		case meeting.FieldWhere:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field where", values[i])
			} else if value.Valid {
				m.Where = value.String
			}
		case meeting.FieldWhom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field whom", values[i])
			} else if value.Valid {
				m.Whom = value.String
			}
		case meeting.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				m.Duration = int(value.Int64)
			}
		case meeting.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case meeting.FieldWhen:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field when", values[i])
			} else if value.Valid {
				m.When = value.Time
			}
		case meeting.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_meetings", value)
			} else if value.Valid {
				m.user_meetings = new(int)
				*m.user_meetings = int(value.Int64)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Meeting.
// This includes values selected through modifiers, order, etc.
func (m *Meeting) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Meeting entity.
func (m *Meeting) QueryUser() *UserQuery {
	return NewMeetingClient(m.config).QueryUser(m)
}

// Update returns a builder for updating this Meeting.
// Note that you need to call Meeting.Unwrap() before calling this method if this Meeting
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Meeting) Update() *MeetingUpdateOne {
	return NewMeetingClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Meeting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Meeting) Unwrap() *Meeting {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Meeting is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Meeting) String() string {
	var builder strings.Builder
	builder.WriteString("Meeting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(m.Description)
	builder.WriteString(", ")
	builder.WriteString("where=")
	builder.WriteString(m.Where)
	builder.WriteString(", ")
	builder.WriteString("whom=")
	builder.WriteString(m.Whom)
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", m.Duration))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("when=")
	builder.WriteString(m.When.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Meetings is a parsable slice of Meeting.
type Meetings []*Meeting