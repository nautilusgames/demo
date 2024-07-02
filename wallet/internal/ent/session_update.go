// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nautilusgames/demo/wallet/internal/ent/predicate"
	"github.com/nautilusgames/demo/wallet/internal/ent/session"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks     []Hook
	mutation  *SessionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SessionUpdate) SetUpdatedAt(t time.Time) *SessionUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetPlayerID sets the "player_id" field.
func (su *SessionUpdate) SetPlayerID(s string) *SessionUpdate {
	su.mutation.SetPlayerID(s)
	return su
}

// SetNillablePlayerID sets the "player_id" field if the given value is not nil.
func (su *SessionUpdate) SetNillablePlayerID(s *string) *SessionUpdate {
	if s != nil {
		su.SetPlayerID(*s)
	}
	return su
}

// SetGameID sets the "game_id" field.
func (su *SessionUpdate) SetGameID(s string) *SessionUpdate {
	su.mutation.SetGameID(s)
	return su
}

// SetNillableGameID sets the "game_id" field if the given value is not nil.
func (su *SessionUpdate) SetNillableGameID(s *string) *SessionUpdate {
	if s != nil {
		su.SetGameID(*s)
	}
	return su
}

// SetSessionID sets the "session_id" field.
func (su *SessionUpdate) SetSessionID(i int64) *SessionUpdate {
	su.mutation.ResetSessionID()
	su.mutation.SetSessionID(i)
	return su
}

// SetNillableSessionID sets the "session_id" field if the given value is not nil.
func (su *SessionUpdate) SetNillableSessionID(i *int64) *SessionUpdate {
	if i != nil {
		su.SetSessionID(*i)
	}
	return su
}

// AddSessionID adds i to the "session_id" field.
func (su *SessionUpdate) AddSessionID(i int64) *SessionUpdate {
	su.mutation.AddSessionID(i)
	return su
}

// SetWalletType sets the "wallet_type" field.
func (su *SessionUpdate) SetWalletType(s string) *SessionUpdate {
	su.mutation.SetWalletType(s)
	return su
}

// SetNillableWalletType sets the "wallet_type" field if the given value is not nil.
func (su *SessionUpdate) SetNillableWalletType(s *string) *SessionUpdate {
	if s != nil {
		su.SetWalletType(*s)
	}
	return su
}

// ClearWalletType clears the value of the "wallet_type" field.
func (su *SessionUpdate) ClearWalletType() *SessionUpdate {
	su.mutation.ClearWalletType()
	return su
}

// SetBetAmount sets the "bet_amount" field.
func (su *SessionUpdate) SetBetAmount(i int64) *SessionUpdate {
	su.mutation.ResetBetAmount()
	su.mutation.SetBetAmount(i)
	return su
}

// SetNillableBetAmount sets the "bet_amount" field if the given value is not nil.
func (su *SessionUpdate) SetNillableBetAmount(i *int64) *SessionUpdate {
	if i != nil {
		su.SetBetAmount(*i)
	}
	return su
}

// AddBetAmount adds i to the "bet_amount" field.
func (su *SessionUpdate) AddBetAmount(i int64) *SessionUpdate {
	su.mutation.AddBetAmount(i)
	return su
}

// SetWinAmount sets the "win_amount" field.
func (su *SessionUpdate) SetWinAmount(i int64) *SessionUpdate {
	su.mutation.ResetWinAmount()
	su.mutation.SetWinAmount(i)
	return su
}

// SetNillableWinAmount sets the "win_amount" field if the given value is not nil.
func (su *SessionUpdate) SetNillableWinAmount(i *int64) *SessionUpdate {
	if i != nil {
		su.SetWinAmount(*i)
	}
	return su
}

// AddWinAmount adds i to the "win_amount" field.
func (su *SessionUpdate) AddWinAmount(i int64) *SessionUpdate {
	su.mutation.AddWinAmount(i)
	return su
}

// SetChange sets the "change" field.
func (su *SessionUpdate) SetChange(i int64) *SessionUpdate {
	su.mutation.ResetChange()
	su.mutation.SetChange(i)
	return su
}

// SetNillableChange sets the "change" field if the given value is not nil.
func (su *SessionUpdate) SetNillableChange(i *int64) *SessionUpdate {
	if i != nil {
		su.SetChange(*i)
	}
	return su
}

// AddChange adds i to the "change" field.
func (su *SessionUpdate) AddChange(i int64) *SessionUpdate {
	su.mutation.AddChange(i)
	return su
}

// SetNewBalance sets the "new_balance" field.
func (su *SessionUpdate) SetNewBalance(i int64) *SessionUpdate {
	su.mutation.ResetNewBalance()
	su.mutation.SetNewBalance(i)
	return su
}

// SetNillableNewBalance sets the "new_balance" field if the given value is not nil.
func (su *SessionUpdate) SetNillableNewBalance(i *int64) *SessionUpdate {
	if i != nil {
		su.SetNewBalance(*i)
	}
	return su
}

// AddNewBalance adds i to the "new_balance" field.
func (su *SessionUpdate) AddNewBalance(i int64) *SessionUpdate {
	su.mutation.AddNewBalance(i)
	return su
}

// ClearNewBalance clears the value of the "new_balance" field.
func (su *SessionUpdate) ClearNewBalance() *SessionUpdate {
	su.mutation.ClearNewBalance()
	return su
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SessionUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := session.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SessionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SessionUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt64))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(session.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.PlayerID(); ok {
		_spec.SetField(session.FieldPlayerID, field.TypeString, value)
	}
	if value, ok := su.mutation.GameID(); ok {
		_spec.SetField(session.FieldGameID, field.TypeString, value)
	}
	if value, ok := su.mutation.SessionID(); ok {
		_spec.SetField(session.FieldSessionID, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedSessionID(); ok {
		_spec.AddField(session.FieldSessionID, field.TypeInt64, value)
	}
	if value, ok := su.mutation.WalletType(); ok {
		_spec.SetField(session.FieldWalletType, field.TypeString, value)
	}
	if su.mutation.WalletTypeCleared() {
		_spec.ClearField(session.FieldWalletType, field.TypeString)
	}
	if value, ok := su.mutation.BetAmount(); ok {
		_spec.SetField(session.FieldBetAmount, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedBetAmount(); ok {
		_spec.AddField(session.FieldBetAmount, field.TypeInt64, value)
	}
	if value, ok := su.mutation.WinAmount(); ok {
		_spec.SetField(session.FieldWinAmount, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedWinAmount(); ok {
		_spec.AddField(session.FieldWinAmount, field.TypeInt64, value)
	}
	if value, ok := su.mutation.Change(); ok {
		_spec.SetField(session.FieldChange, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedChange(); ok {
		_spec.AddField(session.FieldChange, field.TypeInt64, value)
	}
	if value, ok := su.mutation.NewBalance(); ok {
		_spec.SetField(session.FieldNewBalance, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedNewBalance(); ok {
		_spec.AddField(session.FieldNewBalance, field.TypeInt64, value)
	}
	if su.mutation.NewBalanceCleared() {
		_spec.ClearField(session.FieldNewBalance, field.TypeInt64)
	}
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SessionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SessionUpdateOne) SetUpdatedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetPlayerID sets the "player_id" field.
func (suo *SessionUpdateOne) SetPlayerID(s string) *SessionUpdateOne {
	suo.mutation.SetPlayerID(s)
	return suo
}

// SetNillablePlayerID sets the "player_id" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillablePlayerID(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetPlayerID(*s)
	}
	return suo
}

// SetGameID sets the "game_id" field.
func (suo *SessionUpdateOne) SetGameID(s string) *SessionUpdateOne {
	suo.mutation.SetGameID(s)
	return suo
}

// SetNillableGameID sets the "game_id" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableGameID(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetGameID(*s)
	}
	return suo
}

// SetSessionID sets the "session_id" field.
func (suo *SessionUpdateOne) SetSessionID(i int64) *SessionUpdateOne {
	suo.mutation.ResetSessionID()
	suo.mutation.SetSessionID(i)
	return suo
}

// SetNillableSessionID sets the "session_id" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableSessionID(i *int64) *SessionUpdateOne {
	if i != nil {
		suo.SetSessionID(*i)
	}
	return suo
}

// AddSessionID adds i to the "session_id" field.
func (suo *SessionUpdateOne) AddSessionID(i int64) *SessionUpdateOne {
	suo.mutation.AddSessionID(i)
	return suo
}

// SetWalletType sets the "wallet_type" field.
func (suo *SessionUpdateOne) SetWalletType(s string) *SessionUpdateOne {
	suo.mutation.SetWalletType(s)
	return suo
}

// SetNillableWalletType sets the "wallet_type" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableWalletType(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetWalletType(*s)
	}
	return suo
}

// ClearWalletType clears the value of the "wallet_type" field.
func (suo *SessionUpdateOne) ClearWalletType() *SessionUpdateOne {
	suo.mutation.ClearWalletType()
	return suo
}

// SetBetAmount sets the "bet_amount" field.
func (suo *SessionUpdateOne) SetBetAmount(i int64) *SessionUpdateOne {
	suo.mutation.ResetBetAmount()
	suo.mutation.SetBetAmount(i)
	return suo
}

// SetNillableBetAmount sets the "bet_amount" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableBetAmount(i *int64) *SessionUpdateOne {
	if i != nil {
		suo.SetBetAmount(*i)
	}
	return suo
}

// AddBetAmount adds i to the "bet_amount" field.
func (suo *SessionUpdateOne) AddBetAmount(i int64) *SessionUpdateOne {
	suo.mutation.AddBetAmount(i)
	return suo
}

// SetWinAmount sets the "win_amount" field.
func (suo *SessionUpdateOne) SetWinAmount(i int64) *SessionUpdateOne {
	suo.mutation.ResetWinAmount()
	suo.mutation.SetWinAmount(i)
	return suo
}

// SetNillableWinAmount sets the "win_amount" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableWinAmount(i *int64) *SessionUpdateOne {
	if i != nil {
		suo.SetWinAmount(*i)
	}
	return suo
}

// AddWinAmount adds i to the "win_amount" field.
func (suo *SessionUpdateOne) AddWinAmount(i int64) *SessionUpdateOne {
	suo.mutation.AddWinAmount(i)
	return suo
}

// SetChange sets the "change" field.
func (suo *SessionUpdateOne) SetChange(i int64) *SessionUpdateOne {
	suo.mutation.ResetChange()
	suo.mutation.SetChange(i)
	return suo
}

// SetNillableChange sets the "change" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableChange(i *int64) *SessionUpdateOne {
	if i != nil {
		suo.SetChange(*i)
	}
	return suo
}

// AddChange adds i to the "change" field.
func (suo *SessionUpdateOne) AddChange(i int64) *SessionUpdateOne {
	suo.mutation.AddChange(i)
	return suo
}

// SetNewBalance sets the "new_balance" field.
func (suo *SessionUpdateOne) SetNewBalance(i int64) *SessionUpdateOne {
	suo.mutation.ResetNewBalance()
	suo.mutation.SetNewBalance(i)
	return suo
}

// SetNillableNewBalance sets the "new_balance" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableNewBalance(i *int64) *SessionUpdateOne {
	if i != nil {
		suo.SetNewBalance(*i)
	}
	return suo
}

// AddNewBalance adds i to the "new_balance" field.
func (suo *SessionUpdateOne) AddNewBalance(i int64) *SessionUpdateOne {
	suo.mutation.AddNewBalance(i)
	return suo
}

// ClearNewBalance clears the value of the "new_balance" field.
func (suo *SessionUpdateOne) ClearNewBalance() *SessionUpdateOne {
	suo.mutation.ClearNewBalance()
	return suo
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (suo *SessionUpdateOne) Where(ps ...predicate.Session) *SessionUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SessionUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := session.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SessionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SessionUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt64))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(session.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.PlayerID(); ok {
		_spec.SetField(session.FieldPlayerID, field.TypeString, value)
	}
	if value, ok := suo.mutation.GameID(); ok {
		_spec.SetField(session.FieldGameID, field.TypeString, value)
	}
	if value, ok := suo.mutation.SessionID(); ok {
		_spec.SetField(session.FieldSessionID, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedSessionID(); ok {
		_spec.AddField(session.FieldSessionID, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.WalletType(); ok {
		_spec.SetField(session.FieldWalletType, field.TypeString, value)
	}
	if suo.mutation.WalletTypeCleared() {
		_spec.ClearField(session.FieldWalletType, field.TypeString)
	}
	if value, ok := suo.mutation.BetAmount(); ok {
		_spec.SetField(session.FieldBetAmount, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedBetAmount(); ok {
		_spec.AddField(session.FieldBetAmount, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.WinAmount(); ok {
		_spec.SetField(session.FieldWinAmount, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedWinAmount(); ok {
		_spec.AddField(session.FieldWinAmount, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.Change(); ok {
		_spec.SetField(session.FieldChange, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedChange(); ok {
		_spec.AddField(session.FieldChange, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.NewBalance(); ok {
		_spec.SetField(session.FieldNewBalance, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedNewBalance(); ok {
		_spec.AddField(session.FieldNewBalance, field.TypeInt64, value)
	}
	if suo.mutation.NewBalanceCleared() {
		_spec.ClearField(session.FieldNewBalance, field.TypeInt64)
	}
	_spec.AddModifiers(suo.modifiers...)
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
