// Code generated by ent, DO NOT EDIT.

package arbitrarysection

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/aakarim/pland/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// Txt applies equality check predicate on the "txt" field. It's identical to TxtEQ.
func Txt(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTxt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ArbitrarySection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ArbitrarySection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldToken), v))
	})
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.ArbitrarySection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldToken), v...))
	})
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.ArbitrarySection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldToken), v...))
	})
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldToken), v))
	})
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldToken), v))
	})
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldToken), v))
	})
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldToken), v))
	})
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldToken), v))
	})
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldToken), v))
	})
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldToken), v))
	})
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldToken), v))
	})
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldToken), v))
	})
}

// TxtEQ applies the EQ predicate on the "txt" field.
func TxtEQ(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTxt), v))
	})
}

// TxtNEQ applies the NEQ predicate on the "txt" field.
func TxtNEQ(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTxt), v))
	})
}

// TxtIn applies the In predicate on the "txt" field.
func TxtIn(vs ...string) predicate.ArbitrarySection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTxt), v...))
	})
}

// TxtNotIn applies the NotIn predicate on the "txt" field.
func TxtNotIn(vs ...string) predicate.ArbitrarySection {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTxt), v...))
	})
}

// TxtGT applies the GT predicate on the "txt" field.
func TxtGT(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTxt), v))
	})
}

// TxtGTE applies the GTE predicate on the "txt" field.
func TxtGTE(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTxt), v))
	})
}

// TxtLT applies the LT predicate on the "txt" field.
func TxtLT(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTxt), v))
	})
}

// TxtLTE applies the LTE predicate on the "txt" field.
func TxtLTE(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTxt), v))
	})
}

// TxtContains applies the Contains predicate on the "txt" field.
func TxtContains(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTxt), v))
	})
}

// TxtHasPrefix applies the HasPrefix predicate on the "txt" field.
func TxtHasPrefix(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTxt), v))
	})
}

// TxtHasSuffix applies the HasSuffix predicate on the "txt" field.
func TxtHasSuffix(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTxt), v))
	})
}

// TxtEqualFold applies the EqualFold predicate on the "txt" field.
func TxtEqualFold(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTxt), v))
	})
}

// TxtContainsFold applies the ContainsFold predicate on the "txt" field.
func TxtContainsFold(v string) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTxt), v))
	})
}

// HasPlan applies the HasEdge predicate on the "plan" edge.
func HasPlan() predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PlanTable, PlanPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlanWith applies the HasEdge predicate on the "plan" edge with a given conditions (other predicates).
func HasPlanWith(preds ...predicate.Plan) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PlanTable, PlanPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ArbitrarySection) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ArbitrarySection) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
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
func Not(p predicate.ArbitrarySection) predicate.ArbitrarySection {
	return predicate.ArbitrarySection(func(s *sql.Selector) {
		p(s.Not())
	})
}
