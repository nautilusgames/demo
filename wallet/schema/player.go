package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Player struct {
	ent.Schema
}

func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.String("currency"),
		field.Int64("balance").Default(0).Min(0),
	}
}

func (Player) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}
