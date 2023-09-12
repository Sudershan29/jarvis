package schema

import (
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
		field.Bool("name"),
	}
}

// Edges of the Skill.
func (Skill) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("categories", Category.Type).
            Ref("skills"),
		edge.To("userskills", UserSkill.Type),
    }
}
