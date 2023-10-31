// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent/pubsubmessage"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent/review"

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
			Table:   pubsubmessage.Table,
			Columns: pubsubmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: pubsubmessage.FieldID,
			},
		},
		Type: "PubsubMessage",
		Fields: map[string]*sqlgraph.FieldSpec{
			pubsubmessage.FieldCreatedAt: {Type: field.TypeUint32, Column: pubsubmessage.FieldCreatedAt},
			pubsubmessage.FieldUpdatedAt: {Type: field.TypeUint32, Column: pubsubmessage.FieldUpdatedAt},
			pubsubmessage.FieldDeletedAt: {Type: field.TypeUint32, Column: pubsubmessage.FieldDeletedAt},
			pubsubmessage.FieldEntID:     {Type: field.TypeUUID, Column: pubsubmessage.FieldEntID},
			pubsubmessage.FieldMessageID: {Type: field.TypeString, Column: pubsubmessage.FieldMessageID},
			pubsubmessage.FieldState:     {Type: field.TypeString, Column: pubsubmessage.FieldState},
			pubsubmessage.FieldRespToID:  {Type: field.TypeUUID, Column: pubsubmessage.FieldRespToID},
			pubsubmessage.FieldUndoID:    {Type: field.TypeUUID, Column: pubsubmessage.FieldUndoID},
			pubsubmessage.FieldArguments: {Type: field.TypeString, Column: pubsubmessage.FieldArguments},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   review.Table,
			Columns: review.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: review.FieldID,
			},
		},
		Type: "Review",
		Fields: map[string]*sqlgraph.FieldSpec{
			review.FieldCreatedAt:  {Type: field.TypeUint32, Column: review.FieldCreatedAt},
			review.FieldUpdatedAt:  {Type: field.TypeUint32, Column: review.FieldUpdatedAt},
			review.FieldDeletedAt:  {Type: field.TypeUint32, Column: review.FieldDeletedAt},
			review.FieldEntID:      {Type: field.TypeUUID, Column: review.FieldEntID},
			review.FieldAppID:      {Type: field.TypeUUID, Column: review.FieldAppID},
			review.FieldReviewerID: {Type: field.TypeUUID, Column: review.FieldReviewerID},
			review.FieldDomain:     {Type: field.TypeString, Column: review.FieldDomain},
			review.FieldObjectID:   {Type: field.TypeUUID, Column: review.FieldObjectID},
			review.FieldTrigger:    {Type: field.TypeString, Column: review.FieldTrigger},
			review.FieldObjectType: {Type: field.TypeString, Column: review.FieldObjectType},
			review.FieldState:      {Type: field.TypeString, Column: review.FieldState},
			review.FieldMessage:    {Type: field.TypeString, Column: review.FieldMessage},
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
func (pmq *PubsubMessageQuery) addPredicate(pred func(s *sql.Selector)) {
	pmq.predicates = append(pmq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PubsubMessageQuery builder.
func (pmq *PubsubMessageQuery) Filter() *PubsubMessageFilter {
	return &PubsubMessageFilter{config: pmq.config, predicateAdder: pmq}
}

// addPredicate implements the predicateAdder interface.
func (m *PubsubMessageMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PubsubMessageMutation builder.
func (m *PubsubMessageMutation) Filter() *PubsubMessageFilter {
	return &PubsubMessageFilter{config: m.config, predicateAdder: m}
}

// PubsubMessageFilter provides a generic filtering capability at runtime for PubsubMessageQuery.
type PubsubMessageFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PubsubMessageFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *PubsubMessageFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(pubsubmessage.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *PubsubMessageFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(pubsubmessage.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *PubsubMessageFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(pubsubmessage.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *PubsubMessageFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(pubsubmessage.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *PubsubMessageFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(pubsubmessage.FieldEntID))
}

// WhereMessageID applies the entql string predicate on the message_id field.
func (f *PubsubMessageFilter) WhereMessageID(p entql.StringP) {
	f.Where(p.Field(pubsubmessage.FieldMessageID))
}

// WhereState applies the entql string predicate on the state field.
func (f *PubsubMessageFilter) WhereState(p entql.StringP) {
	f.Where(p.Field(pubsubmessage.FieldState))
}

// WhereRespToID applies the entql [16]byte predicate on the resp_to_id field.
func (f *PubsubMessageFilter) WhereRespToID(p entql.ValueP) {
	f.Where(p.Field(pubsubmessage.FieldRespToID))
}

// WhereUndoID applies the entql [16]byte predicate on the undo_id field.
func (f *PubsubMessageFilter) WhereUndoID(p entql.ValueP) {
	f.Where(p.Field(pubsubmessage.FieldUndoID))
}

// WhereArguments applies the entql string predicate on the arguments field.
func (f *PubsubMessageFilter) WhereArguments(p entql.StringP) {
	f.Where(p.Field(pubsubmessage.FieldArguments))
}

// addPredicate implements the predicateAdder interface.
func (rq *ReviewQuery) addPredicate(pred func(s *sql.Selector)) {
	rq.predicates = append(rq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ReviewQuery builder.
func (rq *ReviewQuery) Filter() *ReviewFilter {
	return &ReviewFilter{config: rq.config, predicateAdder: rq}
}

// addPredicate implements the predicateAdder interface.
func (m *ReviewMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ReviewMutation builder.
func (m *ReviewMutation) Filter() *ReviewFilter {
	return &ReviewFilter{config: m.config, predicateAdder: m}
}

// ReviewFilter provides a generic filtering capability at runtime for ReviewQuery.
type ReviewFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ReviewFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *ReviewFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(review.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *ReviewFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(review.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *ReviewFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(review.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *ReviewFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(review.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *ReviewFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(review.FieldEntID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *ReviewFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(review.FieldAppID))
}

// WhereReviewerID applies the entql [16]byte predicate on the reviewer_id field.
func (f *ReviewFilter) WhereReviewerID(p entql.ValueP) {
	f.Where(p.Field(review.FieldReviewerID))
}

// WhereDomain applies the entql string predicate on the domain field.
func (f *ReviewFilter) WhereDomain(p entql.StringP) {
	f.Where(p.Field(review.FieldDomain))
}

// WhereObjectID applies the entql [16]byte predicate on the object_id field.
func (f *ReviewFilter) WhereObjectID(p entql.ValueP) {
	f.Where(p.Field(review.FieldObjectID))
}

// WhereTrigger applies the entql string predicate on the trigger field.
func (f *ReviewFilter) WhereTrigger(p entql.StringP) {
	f.Where(p.Field(review.FieldTrigger))
}

// WhereObjectType applies the entql string predicate on the object_type field.
func (f *ReviewFilter) WhereObjectType(p entql.StringP) {
	f.Where(p.Field(review.FieldObjectType))
}

// WhereState applies the entql string predicate on the state field.
func (f *ReviewFilter) WhereState(p entql.StringP) {
	f.Where(p.Field(review.FieldState))
}

// WhereMessage applies the entql string predicate on the message field.
func (f *ReviewFilter) WhereMessage(p entql.StringP) {
	f.Where(p.Field(review.FieldMessage))
}
