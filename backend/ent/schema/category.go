package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		// M2M
        edge.To("skills", Skill.Type),
        edge.To("tasks", Task.Type),
        edge.To("goals", Goal.Type),
        edge.To("hobbies", Hobby.Type),

		// O2M
		edge.From("user", User.Type).
			Ref("categories").
            Unique(),
    }
}
