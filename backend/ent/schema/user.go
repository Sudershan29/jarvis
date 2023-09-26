package schema

import (
	"time"
    "entgo.io/ent"
	"github.com/google/uuid"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/index"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
        field.String("name"),
        field.String("email_address").
            Unique(),
        field.String("password").
            Sensitive(),
		field.Time("created_at").
            Default(time.Now),
		field.UUID("uuid", uuid.UUID{}).
            Default(uuid.New),
        field.Bool("premium").
            Default(false),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("skills", Skill.Type),
        edge.To("categories", Category.Type),
        edge.To("preference", Preference.Type).
            Unique(),
    }
}

func (User) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("uuid"),
    }
}