package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// Skill holds the schema definition for the Skill entity.
type Skill struct {
	ent.Schema
}

// Fields of the Skill.
func (Skill) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("level"),
		field.Int("progress").
			Default(0),
		field.Int("duration").	// Duration in days
			Default(0),
		field.Time("created_at").
            Default(time.Now),
	}
}

// Edges of the Skill.
func (Skill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("categories", Category.Type).
			Ref("skills"),
        edge.From("user", User.Type).
			Ref("skills").
            Unique(),
    }
}
