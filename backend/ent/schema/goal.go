package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// Goal holds the schema definition for the Goal entity.
type Goal struct {
	ent.Schema
}

// Fields of the Goal.
func (Goal) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").
			Optional(),
		field.Time("created_at").
            Default(time.Now),
	}
}

// Edges of the Goal.
func (Goal) Edges() []ent.Edge {
	return []ent.Edge{
		// M2M
		edge.From("categories", Category.Type).
			Ref("goals"),
		// O2M
        edge.From("user", User.Type).
			Ref("goals").
            Unique(),
    }
}
