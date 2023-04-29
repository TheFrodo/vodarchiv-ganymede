// Code generated by ent, DO NOT EDIT.

package channel

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldID, id))
}

// ExtID applies equality check predicate on the "ext_id" field. It's identical to ExtIDEQ.
func ExtID(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldExtID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldName, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldDisplayName, v))
}

// ImagePath applies equality check predicate on the "image_path" field. It's identical to ImagePathEQ.
func ImagePath(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldImagePath, v))
}

// Retention applies equality check predicate on the "retention" field. It's identical to RetentionEQ.
func Retention(v bool) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldRetention, v))
}

// RetentionDays applies equality check predicate on the "retention_days" field. It's identical to RetentionDaysEQ.
func RetentionDays(v int64) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldRetentionDays, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldCreatedAt, v))
}

// ExtIDEQ applies the EQ predicate on the "ext_id" field.
func ExtIDEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldExtID, v))
}

// ExtIDNEQ applies the NEQ predicate on the "ext_id" field.
func ExtIDNEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldExtID, v))
}

// ExtIDIn applies the In predicate on the "ext_id" field.
func ExtIDIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldExtID, vs...))
}

// ExtIDNotIn applies the NotIn predicate on the "ext_id" field.
func ExtIDNotIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldExtID, vs...))
}

// ExtIDGT applies the GT predicate on the "ext_id" field.
func ExtIDGT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldExtID, v))
}

// ExtIDGTE applies the GTE predicate on the "ext_id" field.
func ExtIDGTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldExtID, v))
}

// ExtIDLT applies the LT predicate on the "ext_id" field.
func ExtIDLT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldExtID, v))
}

// ExtIDLTE applies the LTE predicate on the "ext_id" field.
func ExtIDLTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldExtID, v))
}

// ExtIDContains applies the Contains predicate on the "ext_id" field.
func ExtIDContains(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContains(FieldExtID, v))
}

// ExtIDHasPrefix applies the HasPrefix predicate on the "ext_id" field.
func ExtIDHasPrefix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasPrefix(FieldExtID, v))
}

// ExtIDHasSuffix applies the HasSuffix predicate on the "ext_id" field.
func ExtIDHasSuffix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasSuffix(FieldExtID, v))
}

// ExtIDIsNil applies the IsNil predicate on the "ext_id" field.
func ExtIDIsNil() predicate.Channel {
	return predicate.Channel(sql.FieldIsNull(FieldExtID))
}

// ExtIDNotNil applies the NotNil predicate on the "ext_id" field.
func ExtIDNotNil() predicate.Channel {
	return predicate.Channel(sql.FieldNotNull(FieldExtID))
}

// ExtIDEqualFold applies the EqualFold predicate on the "ext_id" field.
func ExtIDEqualFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEqualFold(FieldExtID, v))
}

// ExtIDContainsFold applies the ContainsFold predicate on the "ext_id" field.
func ExtIDContainsFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContainsFold(FieldExtID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContainsFold(FieldName, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContainsFold(FieldDisplayName, v))
}

// ImagePathEQ applies the EQ predicate on the "image_path" field.
func ImagePathEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldImagePath, v))
}

// ImagePathNEQ applies the NEQ predicate on the "image_path" field.
func ImagePathNEQ(v string) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldImagePath, v))
}

// ImagePathIn applies the In predicate on the "image_path" field.
func ImagePathIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldImagePath, vs...))
}

// ImagePathNotIn applies the NotIn predicate on the "image_path" field.
func ImagePathNotIn(vs ...string) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldImagePath, vs...))
}

// ImagePathGT applies the GT predicate on the "image_path" field.
func ImagePathGT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldImagePath, v))
}

// ImagePathGTE applies the GTE predicate on the "image_path" field.
func ImagePathGTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldImagePath, v))
}

// ImagePathLT applies the LT predicate on the "image_path" field.
func ImagePathLT(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldImagePath, v))
}

// ImagePathLTE applies the LTE predicate on the "image_path" field.
func ImagePathLTE(v string) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldImagePath, v))
}

// ImagePathContains applies the Contains predicate on the "image_path" field.
func ImagePathContains(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContains(FieldImagePath, v))
}

// ImagePathHasPrefix applies the HasPrefix predicate on the "image_path" field.
func ImagePathHasPrefix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasPrefix(FieldImagePath, v))
}

// ImagePathHasSuffix applies the HasSuffix predicate on the "image_path" field.
func ImagePathHasSuffix(v string) predicate.Channel {
	return predicate.Channel(sql.FieldHasSuffix(FieldImagePath, v))
}

// ImagePathEqualFold applies the EqualFold predicate on the "image_path" field.
func ImagePathEqualFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldEqualFold(FieldImagePath, v))
}

// ImagePathContainsFold applies the ContainsFold predicate on the "image_path" field.
func ImagePathContainsFold(v string) predicate.Channel {
	return predicate.Channel(sql.FieldContainsFold(FieldImagePath, v))
}

// RetentionEQ applies the EQ predicate on the "retention" field.
func RetentionEQ(v bool) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldRetention, v))
}

// RetentionNEQ applies the NEQ predicate on the "retention" field.
func RetentionNEQ(v bool) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldRetention, v))
}

// RetentionDaysEQ applies the EQ predicate on the "retention_days" field.
func RetentionDaysEQ(v int64) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldRetentionDays, v))
}

// RetentionDaysNEQ applies the NEQ predicate on the "retention_days" field.
func RetentionDaysNEQ(v int64) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldRetentionDays, v))
}

// RetentionDaysIn applies the In predicate on the "retention_days" field.
func RetentionDaysIn(vs ...int64) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldRetentionDays, vs...))
}

// RetentionDaysNotIn applies the NotIn predicate on the "retention_days" field.
func RetentionDaysNotIn(vs ...int64) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldRetentionDays, vs...))
}

// RetentionDaysGT applies the GT predicate on the "retention_days" field.
func RetentionDaysGT(v int64) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldRetentionDays, v))
}

// RetentionDaysGTE applies the GTE predicate on the "retention_days" field.
func RetentionDaysGTE(v int64) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldRetentionDays, v))
}

// RetentionDaysLT applies the LT predicate on the "retention_days" field.
func RetentionDaysLT(v int64) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldRetentionDays, v))
}

// RetentionDaysLTE applies the LTE predicate on the "retention_days" field.
func RetentionDaysLTE(v int64) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldRetentionDays, v))
}

// RetentionDaysIsNil applies the IsNil predicate on the "retention_days" field.
func RetentionDaysIsNil() predicate.Channel {
	return predicate.Channel(sql.FieldIsNull(FieldRetentionDays))
}

// RetentionDaysNotNil applies the NotNil predicate on the "retention_days" field.
func RetentionDaysNotNil() predicate.Channel {
	return predicate.Channel(sql.FieldNotNull(FieldRetentionDays))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Channel {
	return predicate.Channel(sql.FieldLTE(FieldCreatedAt, v))
}

// HasVods applies the HasEdge predicate on the "vods" edge.
func HasVods() predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VodsTable, VodsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVodsWith applies the HasEdge predicate on the "vods" edge with a given conditions (other predicates).
func HasVodsWith(preds ...predicate.Vod) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := newVodsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLive applies the HasEdge predicate on the "live" edge.
func HasLive() predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LiveTable, LiveColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLiveWith applies the HasEdge predicate on the "live" edge with a given conditions (other predicates).
func HasLiveWith(preds ...predicate.Live) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		step := newLiveStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
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
func Not(p predicate.Channel) predicate.Channel {
	return predicate.Channel(func(s *sql.Selector) {
		p(s.Not())
	})
}
