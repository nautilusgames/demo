package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Session struct {
	ent.Schema
}

func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("game_id"),
		field.String("game_session_id"),
	}
}

func (Session) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}
