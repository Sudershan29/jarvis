package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserSkill holds the schema definition for the UserSkill entity.
type UserSkill struct {
	ent.Schema
}

// Fields of the UserSkill.
func (UserSkill) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("level").
			Values("beginner", "intermediate", "advanced"),
		field.Int("progress").
			Default(0),
		field.Int("duration").	// Duration in days
			Default(0),
	}
}

/* 
	Edges of the UserSkill.

	NOTE: Foreign key were not required to be defined as fields, as they get auto generated when Edges are defined

*/
func (UserSkill) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("user", User.Type).
			Ref("skills").
            Unique(),
        edge.From("skill", Skill.Type).
			Ref("userskills").
			Unique(),
    }
}
