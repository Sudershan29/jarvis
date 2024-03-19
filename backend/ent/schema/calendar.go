package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Calendar holds the schema definition for the Calendar entity.
type Calendar struct {
	ent.Schema
}

// Fields of the Calendar.
func (Calendar) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("type"),
		field.String("token").
			Optional().
			Sensitive(),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Calendar.
func (Calendar) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M
		edge.From("user", User.Type).
			Ref("calendars").
			Unique(),
	}
}
