// Code generated by ent, DO NOT EDIT.

package plan

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/aakarim/pland/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// HasConflict applies equality check predicate on the "has_conflict" field. It's identical to HasConflictEQ.
func HasConflict(v bool) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHasConflict), v))
	})
}

// Digest applies equality check predicate on the "digest" field. It's identical to DigestEQ.
func Digest(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDigest), v))
	})
}

// Txt applies equality check predicate on the "txt" field. It's identical to TxtEQ.
func Txt(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTxt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasConflictEQ applies the EQ predicate on the "has_conflict" field.
func HasConflictEQ(v bool) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHasConflict), v))
	})
}

// HasConflictNEQ applies the NEQ predicate on the "has_conflict" field.
func HasConflictNEQ(v bool) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHasConflict), v))
	})
}

// DigestEQ applies the EQ predicate on the "digest" field.
func DigestEQ(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDigest), v))
	})
}

// DigestNEQ applies the NEQ predicate on the "digest" field.
func DigestNEQ(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDigest), v))
	})
}

// DigestIn applies the In predicate on the "digest" field.
func DigestIn(vs ...string) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDigest), v...))
	})
}

// DigestNotIn applies the NotIn predicate on the "digest" field.
func DigestNotIn(vs ...string) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDigest), v...))
	})
}

// DigestGT applies the GT predicate on the "digest" field.
func DigestGT(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDigest), v))
	})
}

// DigestGTE applies the GTE predicate on the "digest" field.
func DigestGTE(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDigest), v))
	})
}

// DigestLT applies the LT predicate on the "digest" field.
func DigestLT(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDigest), v))
	})
}

// DigestLTE applies the LTE predicate on the "digest" field.
func DigestLTE(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDigest), v))
	})
}

// DigestContains applies the Contains predicate on the "digest" field.
func DigestContains(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDigest), v))
	})
}

// DigestHasPrefix applies the HasPrefix predicate on the "digest" field.
func DigestHasPrefix(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDigest), v))
	})
}

// DigestHasSuffix applies the HasSuffix predicate on the "digest" field.
func DigestHasSuffix(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDigest), v))
	})
}

// DigestEqualFold applies the EqualFold predicate on the "digest" field.
func DigestEqualFold(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDigest), v))
	})
}

// DigestContainsFold applies the ContainsFold predicate on the "digest" field.
func DigestContainsFold(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDigest), v))
	})
}

// TxtEQ applies the EQ predicate on the "txt" field.
func TxtEQ(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTxt), v))
	})
}

// TxtNEQ applies the NEQ predicate on the "txt" field.
func TxtNEQ(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTxt), v))
	})
}

// TxtIn applies the In predicate on the "txt" field.
func TxtIn(vs ...string) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTxt), v...))
	})
}

// TxtNotIn applies the NotIn predicate on the "txt" field.
func TxtNotIn(vs ...string) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTxt), v...))
	})
}

// TxtGT applies the GT predicate on the "txt" field.
func TxtGT(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTxt), v))
	})
}

// TxtGTE applies the GTE predicate on the "txt" field.
func TxtGTE(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTxt), v))
	})
}

// TxtLT applies the LT predicate on the "txt" field.
func TxtLT(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTxt), v))
	})
}

// TxtLTE applies the LTE predicate on the "txt" field.
func TxtLTE(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTxt), v))
	})
}

// TxtContains applies the Contains predicate on the "txt" field.
func TxtContains(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTxt), v))
	})
}

// TxtHasPrefix applies the HasPrefix predicate on the "txt" field.
func TxtHasPrefix(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTxt), v))
	})
}

// TxtHasSuffix applies the HasSuffix predicate on the "txt" field.
func TxtHasSuffix(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTxt), v))
	})
}

// TxtEqualFold applies the EqualFold predicate on the "txt" field.
func TxtEqualFold(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTxt), v))
	})
}

// TxtContainsFold applies the ContainsFold predicate on the "txt" field.
func TxtContainsFold(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTxt), v))
	})
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.User) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDays applies the HasEdge predicate on the "days" edge.
func HasDays() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DaysTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, DaysTable, DaysPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDaysWith applies the HasEdge predicate on the "days" edge with a given conditions (other predicates).
func HasDaysWith(preds ...predicate.Day) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DaysInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, DaysTable, DaysPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasArbitrarySections applies the HasEdge predicate on the "arbitrarySections" edge.
func HasArbitrarySections() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArbitrarySectionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ArbitrarySectionsTable, ArbitrarySectionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArbitrarySectionsWith applies the HasEdge predicate on the "arbitrarySections" edge with a given conditions (other predicates).
func HasArbitrarySectionsWith(preds ...predicate.ArbitrarySection) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArbitrarySectionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ArbitrarySectionsTable, ArbitrarySectionsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHeader applies the HasEdge predicate on the "header" edge.
func HasHeader() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HeaderTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, HeaderTable, HeaderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHeaderWith applies the HasEdge predicate on the "header" edge with a given conditions (other predicates).
func HasHeaderWith(preds ...predicate.Header) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HeaderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, HeaderTable, HeaderColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPrev applies the HasEdge predicate on the "prev" edge.
func HasPrev() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrevTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PrevTable, PrevColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrevWith applies the HasEdge predicate on the "prev" edge with a given conditions (other predicates).
func HasPrevWith(preds ...predicate.Plan) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PrevTable, PrevColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNext applies the HasEdge predicate on the "next" edge.
func HasNext() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NextTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, NextTable, NextColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNextWith applies the HasEdge predicate on the "next" edge with a given conditions (other predicates).
func HasNextWith(preds ...predicate.Plan) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, NextTable, NextColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Plan) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Plan) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
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
func Not(p predicate.Plan) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		p(s.Not())
	})
}
