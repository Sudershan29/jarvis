package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// Hobby holds the schema definition for the Hobby entity.
type Hobby struct {
	ent.Schema
}

// Fields of the Hobby.
func (Hobby) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").
			Optional(),
		field.Time("created_at").
            Default(time.Now),
	}
}

// Edges of the Hobby.
func (Hobby) Edges() []ent.Edge {
	return []ent.Edge{
		// M2M
		edge.From("categories", Category.Type).
			Ref("hobbies"),
		// O2M
        edge.From("user", User.Type).
			Ref("hobbies").
            Unique(),
    }
}
