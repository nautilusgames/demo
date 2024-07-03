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
		field.String("player_id"),
		field.String("game_id"),
		field.Int64("session_id"),

		// name empty means cash
		field.String("wallet_type").Optional(),

		field.Int64("bet_amount"),
		field.Int64("win_amount"),
		field.Int64("change"),
		field.Int64("new_balance").Optional(),
	}
}

func (Session) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Base{},
	}
}
