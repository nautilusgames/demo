package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Sample struct {
	ent.Schema
}

func (Sample) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
	}
}

func (Sample) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}
