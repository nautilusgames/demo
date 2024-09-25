package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Player struct {
	ent.Schema
}

func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),

		field.String("tenant_id"),
		field.String("username").Unique(),
		field.String("hashed_password").NotEmpty(),

		field.String("currency"),

		field.String("display_name"),

		field.Time("created_at").Default(time.Now),
	}
}
