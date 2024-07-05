// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/nautilusgames/demo/wallet/internal/ent/predicate"
	"github.com/nautilusgames/demo/wallet/internal/ent/session"
	"github.com/nautilusgames/demo/wallet/internal/ent/wallet"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeSession = "Session"
	TypeWallet  = "Wallet"
)

// SessionMutation represents an operation that mutates the Session nodes in the graph.
type SessionMutation struct {
	config
	op                 Op
	typ                string
	id                 *int64
	created_at         *time.Time
	updated_at         *time.Time
	game_id            *string
	game_session_id    *int64
	addgame_session_id *int64
	clearedFields      map[string]struct{}
	done               bool
	oldValue           func(context.Context) (*Session, error)
	predicates         []predicate.Session
}

var _ ent.Mutation = (*SessionMutation)(nil)

// sessionOption allows management of the mutation configuration using functional options.
type sessionOption func(*SessionMutation)

// newSessionMutation creates new mutation for the Session entity.
func newSessionMutation(c config, op Op, opts ...sessionOption) *SessionMutation {
	m := &SessionMutation{
		config:        c,
		op:            op,
		typ:           TypeSession,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSessionID sets the ID field of the mutation.
func withSessionID(id int64) sessionOption {
	return func(m *SessionMutation) {
		var (
			err   error
			once  sync.Once
			value *Session
		)
		m.oldValue = func(ctx context.Context) (*Session, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Session.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSession sets the old Session of the mutation.
func withSession(node *Session) sessionOption {
	return func(m *SessionMutation) {
		m.oldValue = func(context.Context) (*Session, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SessionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SessionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Session entities.
func (m *SessionMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SessionMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SessionMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Session.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *SessionMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *SessionMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *SessionMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *SessionMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *SessionMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *SessionMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetGameID sets the "game_id" field.
func (m *SessionMutation) SetGameID(s string) {
	m.game_id = &s
}

// GameID returns the value of the "game_id" field in the mutation.
func (m *SessionMutation) GameID() (r string, exists bool) {
	v := m.game_id
	if v == nil {
		return
	}
	return *v, true
}

// OldGameID returns the old "game_id" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldGameID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGameID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGameID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGameID: %w", err)
	}
	return oldValue.GameID, nil
}

// ResetGameID resets all changes to the "game_id" field.
func (m *SessionMutation) ResetGameID() {
	m.game_id = nil
}

// SetGameSessionID sets the "game_session_id" field.
func (m *SessionMutation) SetGameSessionID(i int64) {
	m.game_session_id = &i
	m.addgame_session_id = nil
}

// GameSessionID returns the value of the "game_session_id" field in the mutation.
func (m *SessionMutation) GameSessionID() (r int64, exists bool) {
	v := m.game_session_id
	if v == nil {
		return
	}
	return *v, true
}

// OldGameSessionID returns the old "game_session_id" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldGameSessionID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGameSessionID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGameSessionID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGameSessionID: %w", err)
	}
	return oldValue.GameSessionID, nil
}

// AddGameSessionID adds i to the "game_session_id" field.
func (m *SessionMutation) AddGameSessionID(i int64) {
	if m.addgame_session_id != nil {
		*m.addgame_session_id += i
	} else {
		m.addgame_session_id = &i
	}
}

// AddedGameSessionID returns the value that was added to the "game_session_id" field in this mutation.
func (m *SessionMutation) AddedGameSessionID() (r int64, exists bool) {
	v := m.addgame_session_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetGameSessionID resets all changes to the "game_session_id" field.
func (m *SessionMutation) ResetGameSessionID() {
	m.game_session_id = nil
	m.addgame_session_id = nil
}

// Where appends a list predicates to the SessionMutation builder.
func (m *SessionMutation) Where(ps ...predicate.Session) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the SessionMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *SessionMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Session, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *SessionMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *SessionMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Session).
func (m *SessionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SessionMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.created_at != nil {
		fields = append(fields, session.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, session.FieldUpdatedAt)
	}
	if m.game_id != nil {
		fields = append(fields, session.FieldGameID)
	}
	if m.game_session_id != nil {
		fields = append(fields, session.FieldGameSessionID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SessionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case session.FieldCreatedAt:
		return m.CreatedAt()
	case session.FieldUpdatedAt:
		return m.UpdatedAt()
	case session.FieldGameID:
		return m.GameID()
	case session.FieldGameSessionID:
		return m.GameSessionID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SessionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case session.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case session.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case session.FieldGameID:
		return m.OldGameID(ctx)
	case session.FieldGameSessionID:
		return m.OldGameSessionID(ctx)
	}
	return nil, fmt.Errorf("unknown Session field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SessionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case session.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case session.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case session.FieldGameID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGameID(v)
		return nil
	case session.FieldGameSessionID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGameSessionID(v)
		return nil
	}
	return fmt.Errorf("unknown Session field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SessionMutation) AddedFields() []string {
	var fields []string
	if m.addgame_session_id != nil {
		fields = append(fields, session.FieldGameSessionID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SessionMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case session.FieldGameSessionID:
		return m.AddedGameSessionID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SessionMutation) AddField(name string, value ent.Value) error {
	switch name {
	case session.FieldGameSessionID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddGameSessionID(v)
		return nil
	}
	return fmt.Errorf("unknown Session numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SessionMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SessionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SessionMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Session nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SessionMutation) ResetField(name string) error {
	switch name {
	case session.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case session.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case session.FieldGameID:
		m.ResetGameID()
		return nil
	case session.FieldGameSessionID:
		m.ResetGameSessionID()
		return nil
	}
	return fmt.Errorf("unknown Session field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SessionMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SessionMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SessionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SessionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SessionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SessionMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SessionMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Session unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SessionMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Session edge %s", name)
}

// WalletMutation represents an operation that mutates the Wallet nodes in the graph.
type WalletMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	created_at    *time.Time
	updated_at    *time.Time
	currency      *string
	balance       *int64
	addbalance    *int64
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Wallet, error)
	predicates    []predicate.Wallet
}

var _ ent.Mutation = (*WalletMutation)(nil)

// walletOption allows management of the mutation configuration using functional options.
type walletOption func(*WalletMutation)

// newWalletMutation creates new mutation for the Wallet entity.
func newWalletMutation(c config, op Op, opts ...walletOption) *WalletMutation {
	m := &WalletMutation{
		config:        c,
		op:            op,
		typ:           TypeWallet,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withWalletID sets the ID field of the mutation.
func withWalletID(id int64) walletOption {
	return func(m *WalletMutation) {
		var (
			err   error
			once  sync.Once
			value *Wallet
		)
		m.oldValue = func(ctx context.Context) (*Wallet, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Wallet.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withWallet sets the old Wallet of the mutation.
func withWallet(node *Wallet) walletOption {
	return func(m *WalletMutation) {
		m.oldValue = func(context.Context) (*Wallet, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m WalletMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m WalletMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Wallet entities.
func (m *WalletMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *WalletMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *WalletMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Wallet.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *WalletMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *WalletMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Wallet entity.
// If the Wallet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WalletMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *WalletMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *WalletMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *WalletMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Wallet entity.
// If the Wallet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WalletMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *WalletMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetCurrency sets the "currency" field.
func (m *WalletMutation) SetCurrency(s string) {
	m.currency = &s
}

// Currency returns the value of the "currency" field in the mutation.
func (m *WalletMutation) Currency() (r string, exists bool) {
	v := m.currency
	if v == nil {
		return
	}
	return *v, true
}

// OldCurrency returns the old "currency" field's value of the Wallet entity.
// If the Wallet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WalletMutation) OldCurrency(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCurrency is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCurrency requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCurrency: %w", err)
	}
	return oldValue.Currency, nil
}

// ResetCurrency resets all changes to the "currency" field.
func (m *WalletMutation) ResetCurrency() {
	m.currency = nil
}

// SetBalance sets the "balance" field.
func (m *WalletMutation) SetBalance(i int64) {
	m.balance = &i
	m.addbalance = nil
}

// Balance returns the value of the "balance" field in the mutation.
func (m *WalletMutation) Balance() (r int64, exists bool) {
	v := m.balance
	if v == nil {
		return
	}
	return *v, true
}

// OldBalance returns the old "balance" field's value of the Wallet entity.
// If the Wallet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WalletMutation) OldBalance(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBalance is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBalance requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBalance: %w", err)
	}
	return oldValue.Balance, nil
}

// AddBalance adds i to the "balance" field.
func (m *WalletMutation) AddBalance(i int64) {
	if m.addbalance != nil {
		*m.addbalance += i
	} else {
		m.addbalance = &i
	}
}

// AddedBalance returns the value that was added to the "balance" field in this mutation.
func (m *WalletMutation) AddedBalance() (r int64, exists bool) {
	v := m.addbalance
	if v == nil {
		return
	}
	return *v, true
}

// ResetBalance resets all changes to the "balance" field.
func (m *WalletMutation) ResetBalance() {
	m.balance = nil
	m.addbalance = nil
}

// Where appends a list predicates to the WalletMutation builder.
func (m *WalletMutation) Where(ps ...predicate.Wallet) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the WalletMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *WalletMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Wallet, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *WalletMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *WalletMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Wallet).
func (m *WalletMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *WalletMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.created_at != nil {
		fields = append(fields, wallet.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, wallet.FieldUpdatedAt)
	}
	if m.currency != nil {
		fields = append(fields, wallet.FieldCurrency)
	}
	if m.balance != nil {
		fields = append(fields, wallet.FieldBalance)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *WalletMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case wallet.FieldCreatedAt:
		return m.CreatedAt()
	case wallet.FieldUpdatedAt:
		return m.UpdatedAt()
	case wallet.FieldCurrency:
		return m.Currency()
	case wallet.FieldBalance:
		return m.Balance()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *WalletMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case wallet.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case wallet.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case wallet.FieldCurrency:
		return m.OldCurrency(ctx)
	case wallet.FieldBalance:
		return m.OldBalance(ctx)
	}
	return nil, fmt.Errorf("unknown Wallet field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WalletMutation) SetField(name string, value ent.Value) error {
	switch name {
	case wallet.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case wallet.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case wallet.FieldCurrency:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCurrency(v)
		return nil
	case wallet.FieldBalance:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBalance(v)
		return nil
	}
	return fmt.Errorf("unknown Wallet field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *WalletMutation) AddedFields() []string {
	var fields []string
	if m.addbalance != nil {
		fields = append(fields, wallet.FieldBalance)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *WalletMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case wallet.FieldBalance:
		return m.AddedBalance()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WalletMutation) AddField(name string, value ent.Value) error {
	switch name {
	case wallet.FieldBalance:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddBalance(v)
		return nil
	}
	return fmt.Errorf("unknown Wallet numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *WalletMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *WalletMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *WalletMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Wallet nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *WalletMutation) ResetField(name string) error {
	switch name {
	case wallet.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case wallet.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case wallet.FieldCurrency:
		m.ResetCurrency()
		return nil
	case wallet.FieldBalance:
		m.ResetBalance()
		return nil
	}
	return fmt.Errorf("unknown Wallet field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *WalletMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *WalletMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *WalletMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *WalletMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *WalletMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *WalletMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *WalletMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Wallet unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *WalletMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Wallet edge %s", name)
}
