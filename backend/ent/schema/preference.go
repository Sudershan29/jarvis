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
		field.Bool("free_weekends").
            Default(false),
		field.Int("weekly_frequency").
            Positive().
			Optional(),
	}
}

// Edges of the Preference.
func (Preference) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("user", User.Type).
            Ref("preference").
            Unique(),
    }
}
