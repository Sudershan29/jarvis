package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		field.Int("duration").
			Default(0),
		field.Int("duration_achieved").
			Default(0),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Skill.
func (Skill) Edges() []ent.Edge {
	return []ent.Edge{
		// M2M
		edge.From("categories", Category.Type).
			Ref("skills"),
		// O2M
		edge.From("user", User.Type).
			Ref("skills").
			Unique(),

		// M2M
		edge.From("time_preferences", TimePreference.Type).
			Ref("skills"),

		// O2M
		edge.To("proposals", Proposal.Type),
	}
}
