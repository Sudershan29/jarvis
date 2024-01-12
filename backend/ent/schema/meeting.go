package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// Meeting holds the schema definition for the Meeting entity.
type Meeting struct {
	ent.Schema
}

// Fields of the Meeting.
func (Meeting) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").
			Optional(),
		field.String("where").
			Optional(),
		field.String("whom").
			Optional(),
		field.Int("duration").
			Default(0),
		field.Time("created_at").
            Default(time.Now),
		field.Time("when").
			Optional(),
	}
}

// Edges of the Meeting.
func (Meeting) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M
        edge.From("user", User.Type).
			Ref("meetings").
            Unique(),
    }
}
