package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Wallet struct {
	ent.Schema
}

func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.String("currency"),
		field.Int64("balance").Default(0).Min(0),
	}
}

func (Wallet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}
