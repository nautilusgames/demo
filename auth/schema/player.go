package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Player struct {
	ent.Schema
}

func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),

		field.String("username").Unique(),
		field.String("hashed_password").NotEmpty(),

		field.String("name"),

		field.Time("created_at").Default(time.Now),
	}
}

func (Player) Edges() []ent.Edge {
	return nil
}

func (Player) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
	}
}
