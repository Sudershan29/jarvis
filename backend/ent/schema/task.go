package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").
			Optional(),
		field.Int("duration").
			Default(0),
		field.Time("created_at").
            Default(time.Now),
		field.Time("deadline").
            Default(time.Now),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		// M2M
		edge.From("categories", Category.Type).
			Ref("tasks"),

		// O2M
        edge.From("user", User.Type).
			Ref("tasks").
            Unique(),
    }
}