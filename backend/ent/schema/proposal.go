package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Proposal holds the schema definition for the Proposal entity.
type Proposal struct {
	ent.Schema
}

// Fields of the Proposal.
func (Proposal) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("allocated_duration").Positive(),
		field.Int("achieved_duration").Optional().Min(0),
		field.Enum("status").Values("pending", "done", "deleted").Default("pending"),
		field.Time("scheduled_for"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the Proposal.
func (Proposal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).
			Ref("proposals").
			Unique(),

		edge.From("skill", Skill.Type).
			Ref("proposals").
			Unique(),
	}
}
