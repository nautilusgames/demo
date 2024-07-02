// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/nautilusgames/demo/auth/internal/ent/player"
	"github.com/nautilusgames/demo/auth/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	playerFields := schema.Player{}.Fields()
	_ = playerFields
	// playerDescHashedPassword is the schema descriptor for hashed_password field.
	playerDescHashedPassword := playerFields[2].Descriptor()
	// player.HashedPasswordValidator is a validator for the "hashed_password" field. It is called by the builders before save.
	player.HashedPasswordValidator = playerDescHashedPassword.Validators[0].(func(string) error)
	// playerDescCreatedAt is the schema descriptor for created_at field.
	playerDescCreatedAt := playerFields[4].Descriptor()
	// player.DefaultCreatedAt holds the default value on creation for the created_at field.
	player.DefaultCreatedAt = playerDescCreatedAt.Default.(func() time.Time)
}
