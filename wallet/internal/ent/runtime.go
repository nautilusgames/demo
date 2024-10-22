// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/nautilusgames/demo/wallet/internal/ent/session"
	"github.com/nautilusgames/demo/wallet/internal/ent/wallet"
	"github.com/nautilusgames/demo/wallet/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	sessionMixin := schema.Session{}.Mixin()
	sessionMixinFields0 := sessionMixin[0].Fields()
	_ = sessionMixinFields0
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescCreatedAt is the schema descriptor for created_at field.
	sessionDescCreatedAt := sessionMixinFields0[1].Descriptor()
	// session.DefaultCreatedAt holds the default value on creation for the created_at field.
	session.DefaultCreatedAt = sessionDescCreatedAt.Default.(func() time.Time)
	// sessionDescUpdatedAt is the schema descriptor for updated_at field.
	sessionDescUpdatedAt := sessionMixinFields0[2].Descriptor()
	// session.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	session.DefaultUpdatedAt = sessionDescUpdatedAt.Default.(func() time.Time)
	// session.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	session.UpdateDefaultUpdatedAt = sessionDescUpdatedAt.UpdateDefault.(func() time.Time)
	walletMixin := schema.Wallet{}.Mixin()
	walletMixinFields0 := walletMixin[0].Fields()
	_ = walletMixinFields0
	walletFields := schema.Wallet{}.Fields()
	_ = walletFields
	// walletDescCreatedAt is the schema descriptor for created_at field.
	walletDescCreatedAt := walletMixinFields0[1].Descriptor()
	// wallet.DefaultCreatedAt holds the default value on creation for the created_at field.
	wallet.DefaultCreatedAt = walletDescCreatedAt.Default.(func() time.Time)
	// walletDescUpdatedAt is the schema descriptor for updated_at field.
	walletDescUpdatedAt := walletMixinFields0[2].Descriptor()
	// wallet.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	wallet.DefaultUpdatedAt = walletDescUpdatedAt.Default.(func() time.Time)
	// wallet.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	wallet.UpdateDefaultUpdatedAt = walletDescUpdatedAt.UpdateDefault.(func() time.Time)
	// walletDescBalance is the schema descriptor for balance field.
	walletDescBalance := walletFields[1].Descriptor()
	// wallet.DefaultBalance holds the default value on creation for the balance field.
	wallet.DefaultBalance = walletDescBalance.Default.(int64)
	// wallet.BalanceValidator is a validator for the "balance" field. It is called by the builders before save.
	wallet.BalanceValidator = walletDescBalance.Validators[0].(func(int64) error)
}
