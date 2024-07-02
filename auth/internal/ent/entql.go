// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/nautilusgames/demo/auth/internal/ent/player"
	"github.com/nautilusgames/demo/auth/internal/ent/sample"

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
			Table:   player.Table,
			Columns: player.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: player.FieldID,
			},
		},
		Type: "Player",
		Fields: map[string]*sqlgraph.FieldSpec{
			player.FieldUsername:       {Type: field.TypeString, Column: player.FieldUsername},
			player.FieldHashedPassword: {Type: field.TypeString, Column: player.FieldHashedPassword},
			player.FieldName:           {Type: field.TypeString, Column: player.FieldName},
			player.FieldCreatedAt:      {Type: field.TypeTime, Column: player.FieldCreatedAt},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   sample.Table,
			Columns: sample.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sample.FieldID,
			},
		},
		Type: "Sample",
		Fields: map[string]*sqlgraph.FieldSpec{
			sample.FieldCreatedAt: {Type: field.TypeTime, Column: sample.FieldCreatedAt},
			sample.FieldUpdatedAt: {Type: field.TypeTime, Column: sample.FieldUpdatedAt},
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
func (pq *PlayerQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PlayerQuery builder.
func (pq *PlayerQuery) Filter() *PlayerFilter {
	return &PlayerFilter{config: pq.config, predicateAdder: pq}
}

// addPredicate implements the predicateAdder interface.
func (m *PlayerMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PlayerMutation builder.
func (m *PlayerMutation) Filter() *PlayerFilter {
	return &PlayerFilter{config: m.config, predicateAdder: m}
}

// PlayerFilter provides a generic filtering capability at runtime for PlayerQuery.
type PlayerFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PlayerFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int64 predicate on the id field.
func (f *PlayerFilter) WhereID(p entql.Int64P) {
	f.Where(p.Field(player.FieldID))
}

// WhereUsername applies the entql string predicate on the username field.
func (f *PlayerFilter) WhereUsername(p entql.StringP) {
	f.Where(p.Field(player.FieldUsername))
}

// WhereHashedPassword applies the entql string predicate on the hashed_password field.
func (f *PlayerFilter) WhereHashedPassword(p entql.StringP) {
	f.Where(p.Field(player.FieldHashedPassword))
}

// WhereName applies the entql string predicate on the name field.
func (f *PlayerFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(player.FieldName))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *PlayerFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(player.FieldCreatedAt))
}

// addPredicate implements the predicateAdder interface.
func (sq *SampleQuery) addPredicate(pred func(s *sql.Selector)) {
	sq.predicates = append(sq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the SampleQuery builder.
func (sq *SampleQuery) Filter() *SampleFilter {
	return &SampleFilter{config: sq.config, predicateAdder: sq}
}

// addPredicate implements the predicateAdder interface.
func (m *SampleMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the SampleMutation builder.
func (m *SampleMutation) Filter() *SampleFilter {
	return &SampleFilter{config: m.config, predicateAdder: m}
}

// SampleFilter provides a generic filtering capability at runtime for SampleQuery.
type SampleFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *SampleFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int64 predicate on the id field.
func (f *SampleFilter) WhereID(p entql.Int64P) {
	f.Where(p.Field(sample.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *SampleFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(sample.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *SampleFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(sample.FieldUpdatedAt))
}
