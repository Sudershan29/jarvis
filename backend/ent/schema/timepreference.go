package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TimePreference holds the schema definition for the TimePreference entity.
type TimePreference struct {
	ent.Schema
}

/*

	Idea: If there is no Time Preference, then the task can be scheduled at any time, and its easier to just set Weekday and Weekend flags

*/
// Fields of the TimePreference.
func (TimePreference) Fields() []ent.Field {
	return []ent.Field{
		field.String("day"),
		// TODO: Not ask for more specific information about afternoon and everything
		// field.Enum("time_of_day").Values("morning", "afternoon", "evening", "night"),
	}
}

// Edges of the TimePreference.
func (TimePreference) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("skills", Skill.Type),
		edge.To("tasks", Task.Type),
	}
}
