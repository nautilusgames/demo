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
	"github.com/nautilusgames/demo/wallet/internal/ent/session"
)

// SessionCreate is the builder for creating a Session entity.
type SessionCreate struct {
	config
	mutation *SessionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (sc *SessionCreate) SetCreatedAt(t time.Time) *SessionCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SessionCreate) SetNillableCreatedAt(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SessionCreate) SetUpdatedAt(t time.Time) *SessionCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SessionCreate) SetNillableUpdatedAt(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetGameID sets the "game_id" field.
func (sc *SessionCreate) SetGameID(s string) *SessionCreate {
	sc.mutation.SetGameID(s)
	return sc
}

// SetGameSessionID sets the "game_session_id" field.
func (sc *SessionCreate) SetGameSessionID(s string) *SessionCreate {
	sc.mutation.SetGameSessionID(s)
	return sc
}

// SetID sets the "id" field.
func (sc *SessionCreate) SetID(i int64) *SessionCreate {
	sc.mutation.SetID(i)
	return sc
}

// Mutation returns the SessionMutation object of the builder.
func (sc *SessionCreate) Mutation() *SessionMutation {
	return sc.mutation
}

// Save creates the Session in the database.
func (sc *SessionCreate) Save(ctx context.Context) (*Session, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SessionCreate) SaveX(ctx context.Context) *Session {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SessionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SessionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SessionCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := session.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := session.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SessionCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Session.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Session.updated_at"`)}
	}
	if _, ok := sc.mutation.GameID(); !ok {
		return &ValidationError{Name: "game_id", err: errors.New(`ent: missing required field "Session.game_id"`)}
	}
	if _, ok := sc.mutation.GameSessionID(); !ok {
		return &ValidationError{Name: "game_session_id", err: errors.New(`ent: missing required field "Session.game_session_id"`)}
	}
	return nil
}

func (sc *SessionCreate) sqlSave(ctx context.Context) (*Session, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SessionCreate) createSpec() (*Session, *sqlgraph.CreateSpec) {
	var (
		_node = &Session{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(session.Table, sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(session.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(session.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.GameID(); ok {
		_spec.SetField(session.FieldGameID, field.TypeString, value)
		_node.GameID = value
	}
	if value, ok := sc.mutation.GameSessionID(); ok {
		_spec.SetField(session.FieldGameSessionID, field.TypeString, value)
		_node.GameSessionID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Session.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SessionUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (sc *SessionCreate) OnConflict(opts ...sql.ConflictOption) *SessionUpsertOne {
	sc.conflict = opts
	return &SessionUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Session.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SessionCreate) OnConflictColumns(columns ...string) *SessionUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SessionUpsertOne{
		create: sc,
	}
}

type (
	// SessionUpsertOne is the builder for "upsert"-ing
	//  one Session node.
	SessionUpsertOne struct {
		create *SessionCreate
	}

	// SessionUpsert is the "OnConflict" setter.
	SessionUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *SessionUpsert) SetUpdatedAt(v time.Time) *SessionUpsert {
	u.Set(session.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SessionUpsert) UpdateUpdatedAt() *SessionUpsert {
	u.SetExcluded(session.FieldUpdatedAt)
	return u
}

// SetGameID sets the "game_id" field.
func (u *SessionUpsert) SetGameID(v string) *SessionUpsert {
	u.Set(session.FieldGameID, v)
	return u
}

// UpdateGameID sets the "game_id" field to the value that was provided on create.
func (u *SessionUpsert) UpdateGameID() *SessionUpsert {
	u.SetExcluded(session.FieldGameID)
	return u
}

// SetGameSessionID sets the "game_session_id" field.
func (u *SessionUpsert) SetGameSessionID(v string) *SessionUpsert {
	u.Set(session.FieldGameSessionID, v)
	return u
}

// UpdateGameSessionID sets the "game_session_id" field to the value that was provided on create.
func (u *SessionUpsert) UpdateGameSessionID() *SessionUpsert {
	u.SetExcluded(session.FieldGameSessionID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Session.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(session.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SessionUpsertOne) UpdateNewValues() *SessionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(session.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(session.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Session.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SessionUpsertOne) Ignore() *SessionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SessionUpsertOne) DoNothing() *SessionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SessionCreate.OnConflict
// documentation for more info.
func (u *SessionUpsertOne) Update(set func(*SessionUpsert)) *SessionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SessionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SessionUpsertOne) SetUpdatedAt(v time.Time) *SessionUpsertOne {
	return u.Update(func(s *SessionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SessionUpsertOne) UpdateUpdatedAt() *SessionUpsertOne {
	return u.Update(func(s *SessionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetGameID sets the "game_id" field.
func (u *SessionUpsertOne) SetGameID(v string) *SessionUpsertOne {
	return u.Update(func(s *SessionUpsert) {
		s.SetGameID(v)
	})
}

// UpdateGameID sets the "game_id" field to the value that was provided on create.
func (u *SessionUpsertOne) UpdateGameID() *SessionUpsertOne {
	return u.Update(func(s *SessionUpsert) {
		s.UpdateGameID()
	})
}

// SetGameSessionID sets the "game_session_id" field.
func (u *SessionUpsertOne) SetGameSessionID(v string) *SessionUpsertOne {
	return u.Update(func(s *SessionUpsert) {
		s.SetGameSessionID(v)
	})
}

// UpdateGameSessionID sets the "game_session_id" field to the value that was provided on create.
func (u *SessionUpsertOne) UpdateGameSessionID() *SessionUpsertOne {
	return u.Update(func(s *SessionUpsert) {
		s.UpdateGameSessionID()
	})
}

// Exec executes the query.
func (u *SessionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SessionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SessionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SessionUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SessionUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SessionCreateBulk is the builder for creating many Session entities in bulk.
type SessionCreateBulk struct {
	config
	err      error
	builders []*SessionCreate
	conflict []sql.ConflictOption
}

// Save creates the Session entities in the database.
func (scb *SessionCreateBulk) Save(ctx context.Context) ([]*Session, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Session, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SessionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SessionCreateBulk) SaveX(ctx context.Context) []*Session {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SessionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SessionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Session.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SessionUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (scb *SessionCreateBulk) OnConflict(opts ...sql.ConflictOption) *SessionUpsertBulk {
	scb.conflict = opts
	return &SessionUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Session.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SessionCreateBulk) OnConflictColumns(columns ...string) *SessionUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SessionUpsertBulk{
		create: scb,
	}
}

// SessionUpsertBulk is the builder for "upsert"-ing
// a bulk of Session nodes.
type SessionUpsertBulk struct {
	create *SessionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Session.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(session.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SessionUpsertBulk) UpdateNewValues() *SessionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(session.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(session.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Session.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SessionUpsertBulk) Ignore() *SessionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SessionUpsertBulk) DoNothing() *SessionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SessionCreateBulk.OnConflict
// documentation for more info.
func (u *SessionUpsertBulk) Update(set func(*SessionUpsert)) *SessionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SessionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SessionUpsertBulk) SetUpdatedAt(v time.Time) *SessionUpsertBulk {
	return u.Update(func(s *SessionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SessionUpsertBulk) UpdateUpdatedAt() *SessionUpsertBulk {
	return u.Update(func(s *SessionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetGameID sets the "game_id" field.
func (u *SessionUpsertBulk) SetGameID(v string) *SessionUpsertBulk {
	return u.Update(func(s *SessionUpsert) {
		s.SetGameID(v)
	})
}

// UpdateGameID sets the "game_id" field to the value that was provided on create.
func (u *SessionUpsertBulk) UpdateGameID() *SessionUpsertBulk {
	return u.Update(func(s *SessionUpsert) {
		s.UpdateGameID()
	})
}

// SetGameSessionID sets the "game_session_id" field.
func (u *SessionUpsertBulk) SetGameSessionID(v string) *SessionUpsertBulk {
	return u.Update(func(s *SessionUpsert) {
		s.SetGameSessionID(v)
	})
}

// UpdateGameSessionID sets the "game_session_id" field to the value that was provided on create.
func (u *SessionUpsertBulk) UpdateGameSessionID() *SessionUpsertBulk {
	return u.Update(func(s *SessionUpsert) {
		s.UpdateGameSessionID()
	})
}

// Exec executes the query.
func (u *SessionUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SessionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SessionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SessionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
