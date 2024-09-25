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
	"github.com/nautilusgames/demo/auth/internal/ent/player"
	"github.com/nautilusgames/demo/auth/internal/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypePlayer = "Player"
)

// PlayerMutation represents an operation that mutates the Player nodes in the graph.
type PlayerMutation struct {
	config
	op              Op
	typ             string
	id              *int64
	tenant_id       *string
	username        *string
	hashed_password *string
	currency        *string
	display_name    *string
	created_at      *time.Time
	clearedFields   map[string]struct{}
	done            bool
	oldValue        func(context.Context) (*Player, error)
	predicates      []predicate.Player
}

var _ ent.Mutation = (*PlayerMutation)(nil)

// playerOption allows management of the mutation configuration using functional options.
type playerOption func(*PlayerMutation)

// newPlayerMutation creates new mutation for the Player entity.
func newPlayerMutation(c config, op Op, opts ...playerOption) *PlayerMutation {
	m := &PlayerMutation{
		config:        c,
		op:            op,
		typ:           TypePlayer,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPlayerID sets the ID field of the mutation.
func withPlayerID(id int64) playerOption {
	return func(m *PlayerMutation) {
		var (
			err   error
			once  sync.Once
			value *Player
		)
		m.oldValue = func(ctx context.Context) (*Player, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Player.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPlayer sets the old Player of the mutation.
func withPlayer(node *Player) playerOption {
	return func(m *PlayerMutation) {
		m.oldValue = func(context.Context) (*Player, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PlayerMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PlayerMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Player entities.
func (m *PlayerMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PlayerMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PlayerMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Player.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTenantID sets the "tenant_id" field.
func (m *PlayerMutation) SetTenantID(s string) {
	m.tenant_id = &s
}

// TenantID returns the value of the "tenant_id" field in the mutation.
func (m *PlayerMutation) TenantID() (r string, exists bool) {
	v := m.tenant_id
	if v == nil {
		return
	}
	return *v, true
}

// OldTenantID returns the old "tenant_id" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldTenantID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTenantID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTenantID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTenantID: %w", err)
	}
	return oldValue.TenantID, nil
}

// ResetTenantID resets all changes to the "tenant_id" field.
func (m *PlayerMutation) ResetTenantID() {
	m.tenant_id = nil
}

// SetUsername sets the "username" field.
func (m *PlayerMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *PlayerMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *PlayerMutation) ResetUsername() {
	m.username = nil
}

// SetHashedPassword sets the "hashed_password" field.
func (m *PlayerMutation) SetHashedPassword(s string) {
	m.hashed_password = &s
}

// HashedPassword returns the value of the "hashed_password" field in the mutation.
func (m *PlayerMutation) HashedPassword() (r string, exists bool) {
	v := m.hashed_password
	if v == nil {
		return
	}
	return *v, true
}

// OldHashedPassword returns the old "hashed_password" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldHashedPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldHashedPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldHashedPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldHashedPassword: %w", err)
	}
	return oldValue.HashedPassword, nil
}

// ResetHashedPassword resets all changes to the "hashed_password" field.
func (m *PlayerMutation) ResetHashedPassword() {
	m.hashed_password = nil
}

// SetCurrency sets the "currency" field.
func (m *PlayerMutation) SetCurrency(s string) {
	m.currency = &s
}

// Currency returns the value of the "currency" field in the mutation.
func (m *PlayerMutation) Currency() (r string, exists bool) {
	v := m.currency
	if v == nil {
		return
	}
	return *v, true
}

// OldCurrency returns the old "currency" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldCurrency(ctx context.Context) (v string, err error) {
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
func (m *PlayerMutation) ResetCurrency() {
	m.currency = nil
}

// SetDisplayName sets the "display_name" field.
func (m *PlayerMutation) SetDisplayName(s string) {
	m.display_name = &s
}

// DisplayName returns the value of the "display_name" field in the mutation.
func (m *PlayerMutation) DisplayName() (r string, exists bool) {
	v := m.display_name
	if v == nil {
		return
	}
	return *v, true
}

// OldDisplayName returns the old "display_name" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldDisplayName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDisplayName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDisplayName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDisplayName: %w", err)
	}
	return oldValue.DisplayName, nil
}

// ResetDisplayName resets all changes to the "display_name" field.
func (m *PlayerMutation) ResetDisplayName() {
	m.display_name = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *PlayerMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *PlayerMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
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
func (m *PlayerMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the PlayerMutation builder.
func (m *PlayerMutation) Where(ps ...predicate.Player) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the PlayerMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *PlayerMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Player, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *PlayerMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *PlayerMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Player).
func (m *PlayerMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PlayerMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.tenant_id != nil {
		fields = append(fields, player.FieldTenantID)
	}
	if m.username != nil {
		fields = append(fields, player.FieldUsername)
	}
	if m.hashed_password != nil {
		fields = append(fields, player.FieldHashedPassword)
	}
	if m.currency != nil {
		fields = append(fields, player.FieldCurrency)
	}
	if m.display_name != nil {
		fields = append(fields, player.FieldDisplayName)
	}
	if m.created_at != nil {
		fields = append(fields, player.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PlayerMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case player.FieldTenantID:
		return m.TenantID()
	case player.FieldUsername:
		return m.Username()
	case player.FieldHashedPassword:
		return m.HashedPassword()
	case player.FieldCurrency:
		return m.Currency()
	case player.FieldDisplayName:
		return m.DisplayName()
	case player.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PlayerMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case player.FieldTenantID:
		return m.OldTenantID(ctx)
	case player.FieldUsername:
		return m.OldUsername(ctx)
	case player.FieldHashedPassword:
		return m.OldHashedPassword(ctx)
	case player.FieldCurrency:
		return m.OldCurrency(ctx)
	case player.FieldDisplayName:
		return m.OldDisplayName(ctx)
	case player.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Player field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PlayerMutation) SetField(name string, value ent.Value) error {
	switch name {
	case player.FieldTenantID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTenantID(v)
		return nil
	case player.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case player.FieldHashedPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHashedPassword(v)
		return nil
	case player.FieldCurrency:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCurrency(v)
		return nil
	case player.FieldDisplayName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDisplayName(v)
		return nil
	case player.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Player field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PlayerMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PlayerMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PlayerMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Player numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PlayerMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PlayerMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PlayerMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Player nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PlayerMutation) ResetField(name string) error {
	switch name {
	case player.FieldTenantID:
		m.ResetTenantID()
		return nil
	case player.FieldUsername:
		m.ResetUsername()
		return nil
	case player.FieldHashedPassword:
		m.ResetHashedPassword()
		return nil
	case player.FieldCurrency:
		m.ResetCurrency()
		return nil
	case player.FieldDisplayName:
		m.ResetDisplayName()
		return nil
	case player.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Player field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PlayerMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PlayerMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PlayerMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PlayerMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PlayerMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PlayerMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PlayerMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Player unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PlayerMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Player edge %s", name)
}
