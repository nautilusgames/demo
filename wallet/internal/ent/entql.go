// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/nautilusgames/demo/wallet/internal/ent/session"
	"github.com/nautilusgames/demo/wallet/internal/ent/wallet"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: session.FieldID,
			},
		},
		Type: "Session",
		Fields: map[string]*sqlgraph.FieldSpec{
			session.FieldCreatedAt:     {Type: field.TypeTime, Column: session.FieldCreatedAt},
			session.FieldUpdatedAt:     {Type: field.TypeTime, Column: session.FieldUpdatedAt},
			session.FieldGameID:        {Type: field.TypeString, Column: session.FieldGameID},
			session.FieldGameSessionID: {Type: field.TypeString, Column: session.FieldGameSessionID},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   wallet.Table,
			Columns: wallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: wallet.FieldID,
			},
		},
		Type: "Wallet",
		Fields: map[string]*sqlgraph.FieldSpec{
			wallet.FieldCreatedAt: {Type: field.TypeTime, Column: wallet.FieldCreatedAt},
			wallet.FieldUpdatedAt: {Type: field.TypeTime, Column: wallet.FieldUpdatedAt},
			wallet.FieldCurrency:  {Type: field.TypeString, Column: wallet.FieldCurrency},
			wallet.FieldBalance:   {Type: field.TypeInt64, Column: wallet.FieldBalance},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (sq *SessionQuery) addPredicate(pred func(s *sql.Selector)) {
	sq.predicates = append(sq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the SessionQuery builder.
func (sq *SessionQuery) Filter() *SessionFilter {
	return &SessionFilter{config: sq.config, predicateAdder: sq}
}

// addPredicate implements the predicateAdder interface.
func (m *SessionMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the SessionMutation builder.
func (m *SessionMutation) Filter() *SessionFilter {
	return &SessionFilter{config: m.config, predicateAdder: m}
}

// SessionFilter provides a generic filtering capability at runtime for SessionQuery.
type SessionFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *SessionFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int64 predicate on the id field.
func (f *SessionFilter) WhereID(p entql.Int64P) {
	f.Where(p.Field(session.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *SessionFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(session.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *SessionFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(session.FieldUpdatedAt))
}

// WhereGameID applies the entql string predicate on the game_id field.
func (f *SessionFilter) WhereGameID(p entql.StringP) {
	f.Where(p.Field(session.FieldGameID))
}

// WhereGameSessionID applies the entql string predicate on the game_session_id field.
func (f *SessionFilter) WhereGameSessionID(p entql.StringP) {
	f.Where(p.Field(session.FieldGameSessionID))
}

// addPredicate implements the predicateAdder interface.
func (wq *WalletQuery) addPredicate(pred func(s *sql.Selector)) {
	wq.predicates = append(wq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the WalletQuery builder.
func (wq *WalletQuery) Filter() *WalletFilter {
	return &WalletFilter{config: wq.config, predicateAdder: wq}
}

// addPredicate implements the predicateAdder interface.
func (m *WalletMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the WalletMutation builder.
func (m *WalletMutation) Filter() *WalletFilter {
	return &WalletFilter{config: m.config, predicateAdder: m}
}

// WalletFilter provides a generic filtering capability at runtime for WalletQuery.
type WalletFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *WalletFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int64 predicate on the id field.
func (f *WalletFilter) WhereID(p entql.Int64P) {
	f.Where(p.Field(wallet.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *WalletFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(wallet.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *WalletFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(wallet.FieldUpdatedAt))
}

// WhereCurrency applies the entql string predicate on the currency field.
func (f *WalletFilter) WhereCurrency(p entql.StringP) {
	f.Where(p.Field(wallet.FieldCurrency))
}

// WhereBalance applies the entql int64 predicate on the balance field.
func (f *WalletFilter) WhereBalance(p entql.Int64P) {
	f.Where(p.Field(wallet.FieldBalance))
}
