// Code generated by ent, DO NOT EDIT.

package pubsubmessage

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/review-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// MessageID applies equality check predicate on the "message_id" field. It's identical to MessageIDEQ.
func MessageID(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessageID), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// RespToID applies equality check predicate on the "resp_to_id" field. It's identical to RespToIDEQ.
func RespToID(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRespToID), v))
	})
}

// UndoID applies equality check predicate on the "undo_id" field. It's identical to UndoIDEQ.
func UndoID(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUndoID), v))
	})
}

// Arguments applies equality check predicate on the "arguments" field. It's identical to ArgumentsEQ.
func Arguments(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArguments), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// MessageIDEQ applies the EQ predicate on the "message_id" field.
func MessageIDEQ(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessageID), v))
	})
}

// MessageIDNEQ applies the NEQ predicate on the "message_id" field.
func MessageIDNEQ(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessageID), v))
	})
}

// MessageIDIn applies the In predicate on the "message_id" field.
func MessageIDIn(vs ...string) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMessageID), v...))
	})
}

// MessageIDNotIn applies the NotIn predicate on the "message_id" field.
func MessageIDNotIn(vs ...string) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMessageID), v...))
	})
}

// MessageIDGT applies the GT predicate on the "message_id" field.
func MessageIDGT(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessageID), v))
	})
}

// MessageIDGTE applies the GTE predicate on the "message_id" field.
func MessageIDGTE(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessageID), v))
	})
}

// MessageIDLT applies the LT predicate on the "message_id" field.
func MessageIDLT(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessageID), v))
	})
}

// MessageIDLTE applies the LTE predicate on the "message_id" field.
func MessageIDLTE(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessageID), v))
	})
}

// MessageIDContains applies the Contains predicate on the "message_id" field.
func MessageIDContains(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessageID), v))
	})
}

// MessageIDHasPrefix applies the HasPrefix predicate on the "message_id" field.
func MessageIDHasPrefix(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessageID), v))
	})
}

// MessageIDHasSuffix applies the HasSuffix predicate on the "message_id" field.
func MessageIDHasSuffix(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessageID), v))
	})
}

// MessageIDIsNil applies the IsNil predicate on the "message_id" field.
func MessageIDIsNil() predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMessageID)))
	})
}

// MessageIDNotNil applies the NotNil predicate on the "message_id" field.
func MessageIDNotNil() predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMessageID)))
	})
}

// MessageIDEqualFold applies the EqualFold predicate on the "message_id" field.
func MessageIDEqualFold(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessageID), v))
	})
}

// MessageIDContainsFold applies the ContainsFold predicate on the "message_id" field.
func MessageIDContainsFold(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessageID), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldState), v))
	})
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldState), v))
	})
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldState), v))
	})
}

// StateIsNil applies the IsNil predicate on the "state" field.
func StateIsNil() predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldState)))
	})
}

// StateNotNil applies the NotNil predicate on the "state" field.
func StateNotNil() predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldState)))
	})
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldState), v))
	})
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldState), v))
	})
}

// RespToIDEQ applies the EQ predicate on the "resp_to_id" field.
func RespToIDEQ(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRespToID), v))
	})
}

// RespToIDNEQ applies the NEQ predicate on the "resp_to_id" field.
func RespToIDNEQ(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRespToID), v))
	})
}

// RespToIDIn applies the In predicate on the "resp_to_id" field.
func RespToIDIn(vs ...uuid.UUID) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRespToID), v...))
	})
}

// RespToIDNotIn applies the NotIn predicate on the "resp_to_id" field.
func RespToIDNotIn(vs ...uuid.UUID) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRespToID), v...))
	})
}

// RespToIDGT applies the GT predicate on the "resp_to_id" field.
func RespToIDGT(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRespToID), v))
	})
}

// RespToIDGTE applies the GTE predicate on the "resp_to_id" field.
func RespToIDGTE(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRespToID), v))
	})
}

// RespToIDLT applies the LT predicate on the "resp_to_id" field.
func RespToIDLT(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRespToID), v))
	})
}

// RespToIDLTE applies the LTE predicate on the "resp_to_id" field.
func RespToIDLTE(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRespToID), v))
	})
}

// RespToIDIsNil applies the IsNil predicate on the "resp_to_id" field.
func RespToIDIsNil() predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRespToID)))
	})
}

// RespToIDNotNil applies the NotNil predicate on the "resp_to_id" field.
func RespToIDNotNil() predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRespToID)))
	})
}

// UndoIDEQ applies the EQ predicate on the "undo_id" field.
func UndoIDEQ(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUndoID), v))
	})
}

// UndoIDNEQ applies the NEQ predicate on the "undo_id" field.
func UndoIDNEQ(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUndoID), v))
	})
}

// UndoIDIn applies the In predicate on the "undo_id" field.
func UndoIDIn(vs ...uuid.UUID) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUndoID), v...))
	})
}

// UndoIDNotIn applies the NotIn predicate on the "undo_id" field.
func UndoIDNotIn(vs ...uuid.UUID) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUndoID), v...))
	})
}

// UndoIDGT applies the GT predicate on the "undo_id" field.
func UndoIDGT(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUndoID), v))
	})
}

// UndoIDGTE applies the GTE predicate on the "undo_id" field.
func UndoIDGTE(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUndoID), v))
	})
}

// UndoIDLT applies the LT predicate on the "undo_id" field.
func UndoIDLT(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUndoID), v))
	})
}

// UndoIDLTE applies the LTE predicate on the "undo_id" field.
func UndoIDLTE(v uuid.UUID) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUndoID), v))
	})
}

// UndoIDIsNil applies the IsNil predicate on the "undo_id" field.
func UndoIDIsNil() predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUndoID)))
	})
}

// UndoIDNotNil applies the NotNil predicate on the "undo_id" field.
func UndoIDNotNil() predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUndoID)))
	})
}

// ArgumentsEQ applies the EQ predicate on the "arguments" field.
func ArgumentsEQ(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArguments), v))
	})
}

// ArgumentsNEQ applies the NEQ predicate on the "arguments" field.
func ArgumentsNEQ(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldArguments), v))
	})
}

// ArgumentsIn applies the In predicate on the "arguments" field.
func ArgumentsIn(vs ...string) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldArguments), v...))
	})
}

// ArgumentsNotIn applies the NotIn predicate on the "arguments" field.
func ArgumentsNotIn(vs ...string) predicate.PubsubMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldArguments), v...))
	})
}

// ArgumentsGT applies the GT predicate on the "arguments" field.
func ArgumentsGT(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldArguments), v))
	})
}

// ArgumentsGTE applies the GTE predicate on the "arguments" field.
func ArgumentsGTE(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldArguments), v))
	})
}

// ArgumentsLT applies the LT predicate on the "arguments" field.
func ArgumentsLT(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldArguments), v))
	})
}

// ArgumentsLTE applies the LTE predicate on the "arguments" field.
func ArgumentsLTE(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldArguments), v))
	})
}

// ArgumentsContains applies the Contains predicate on the "arguments" field.
func ArgumentsContains(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldArguments), v))
	})
}

// ArgumentsHasPrefix applies the HasPrefix predicate on the "arguments" field.
func ArgumentsHasPrefix(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldArguments), v))
	})
}

// ArgumentsHasSuffix applies the HasSuffix predicate on the "arguments" field.
func ArgumentsHasSuffix(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldArguments), v))
	})
}

// ArgumentsIsNil applies the IsNil predicate on the "arguments" field.
func ArgumentsIsNil() predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldArguments)))
	})
}

// ArgumentsNotNil applies the NotNil predicate on the "arguments" field.
func ArgumentsNotNil() predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldArguments)))
	})
}

// ArgumentsEqualFold applies the EqualFold predicate on the "arguments" field.
func ArgumentsEqualFold(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldArguments), v))
	})
}

// ArgumentsContainsFold applies the ContainsFold predicate on the "arguments" field.
func ArgumentsContainsFold(v string) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldArguments), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PubsubMessage) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PubsubMessage) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PubsubMessage) predicate.PubsubMessage {
	return predicate.PubsubMessage(func(s *sql.Selector) {
		p(s.Not())
	})
}
