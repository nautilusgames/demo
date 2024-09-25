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

// SetGameSessionID sets the "game_session_id" field.
func (su *SessionUpdate) SetGameSessionID(s string) *SessionUpdate {
	su.mutation.SetGameSessionID(s)
	return su
}

// SetNillableGameSessionID sets the "game_session_id" field if the given value is not nil.
func (su *SessionUpdate) SetNillableGameSessionID(s *string) *SessionUpdate {
	if s != nil {
		su.SetGameSessionID(*s)
	}
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
	if value, ok := su.mutation.GameID(); ok {
		_spec.SetField(session.FieldGameID, field.TypeString, value)
	}
	if value, ok := su.mutation.GameSessionID(); ok {
		_spec.SetField(session.FieldGameSessionID, field.TypeString, value)
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

// SetGameSessionID sets the "game_session_id" field.
func (suo *SessionUpdateOne) SetGameSessionID(s string) *SessionUpdateOne {
	suo.mutation.SetGameSessionID(s)
	return suo
}

// SetNillableGameSessionID sets the "game_session_id" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableGameSessionID(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetGameSessionID(*s)
	}
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
	if value, ok := suo.mutation.GameID(); ok {
		_spec.SetField(session.FieldGameID, field.TypeString, value)
	}
	if value, ok := suo.mutation.GameSessionID(); ok {
		_spec.SetField(session.FieldGameSessionID, field.TypeString, value)
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
