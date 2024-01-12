package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Preference holds the schema definition for the Preference entity.
type Preference struct {
	ent.Schema
}

// Fields of the Preference.
func (Preference) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("free_weekends").	// Represents if users want a free weekend
            Default(false),
		field.Int("weekly_frequency").	// Represents how many hours per week they are commiting to work
            Positive().
			Optional(),
	}
}

// Edges of the Preference.
func (Preference) Edges() []ent.Edge {
	return []ent.Edge{
		// O2O
        edge.From("user", User.Type).
            Ref("preference").
            Unique(),
    }
}
