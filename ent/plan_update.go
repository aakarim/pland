// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/aakarim/pland/ent/arbitrarysection"
	"github.com/aakarim/pland/ent/day"
	"github.com/aakarim/pland/ent/header"
	"github.com/aakarim/pland/ent/plan"
	"github.com/aakarim/pland/ent/predicate"
	"github.com/aakarim/pland/ent/user"
)

// PlanUpdate is the builder for updating Plan entities.
type PlanUpdate struct {
	config
	hooks    []Hook
	mutation *PlanMutation
}

// Where appends a list predicates to the PlanUpdate builder.
func (pu *PlanUpdate) Where(ps ...predicate.Plan) *PlanUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetDigest sets the "digest" field.
func (pu *PlanUpdate) SetDigest(s string) *PlanUpdate {
	pu.mutation.SetDigest(s)
	return pu
}

// SetTxt sets the "txt" field.
func (pu *PlanUpdate) SetTxt(s string) *PlanUpdate {
	pu.mutation.SetTxt(s)
	return pu
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (pu *PlanUpdate) SetAuthorID(id int) *PlanUpdate {
	pu.mutation.SetAuthorID(id)
	return pu
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (pu *PlanUpdate) SetNillableAuthorID(id *int) *PlanUpdate {
	if id != nil {
		pu = pu.SetAuthorID(*id)
	}
	return pu
}

// SetAuthor sets the "author" edge to the User entity.
func (pu *PlanUpdate) SetAuthor(u *User) *PlanUpdate {
	return pu.SetAuthorID(u.ID)
}

// AddDayIDs adds the "days" edge to the Day entity by IDs.
func (pu *PlanUpdate) AddDayIDs(ids ...int) *PlanUpdate {
	pu.mutation.AddDayIDs(ids...)
	return pu
}

// AddDays adds the "days" edges to the Day entity.
func (pu *PlanUpdate) AddDays(d ...*Day) *PlanUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pu.AddDayIDs(ids...)
}

// AddArbitrarySectionIDs adds the "arbitrarySections" edge to the ArbitrarySection entity by IDs.
func (pu *PlanUpdate) AddArbitrarySectionIDs(ids ...int) *PlanUpdate {
	pu.mutation.AddArbitrarySectionIDs(ids...)
	return pu
}

// AddArbitrarySections adds the "arbitrarySections" edges to the ArbitrarySection entity.
func (pu *PlanUpdate) AddArbitrarySections(a ...*ArbitrarySection) *PlanUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.AddArbitrarySectionIDs(ids...)
}

// SetHeaderID sets the "header" edge to the Header entity by ID.
func (pu *PlanUpdate) SetHeaderID(id int) *PlanUpdate {
	pu.mutation.SetHeaderID(id)
	return pu
}

// SetNillableHeaderID sets the "header" edge to the Header entity by ID if the given value is not nil.
func (pu *PlanUpdate) SetNillableHeaderID(id *int) *PlanUpdate {
	if id != nil {
		pu = pu.SetHeaderID(*id)
	}
	return pu
}

// SetHeader sets the "header" edge to the Header entity.
func (pu *PlanUpdate) SetHeader(h *Header) *PlanUpdate {
	return pu.SetHeaderID(h.ID)
}

// SetPrevID sets the "prev" edge to the Plan entity by ID.
func (pu *PlanUpdate) SetPrevID(id int) *PlanUpdate {
	pu.mutation.SetPrevID(id)
	return pu
}

// SetNillablePrevID sets the "prev" edge to the Plan entity by ID if the given value is not nil.
func (pu *PlanUpdate) SetNillablePrevID(id *int) *PlanUpdate {
	if id != nil {
		pu = pu.SetPrevID(*id)
	}
	return pu
}

// SetPrev sets the "prev" edge to the Plan entity.
func (pu *PlanUpdate) SetPrev(p *Plan) *PlanUpdate {
	return pu.SetPrevID(p.ID)
}

// SetNextID sets the "next" edge to the Plan entity by ID.
func (pu *PlanUpdate) SetNextID(id int) *PlanUpdate {
	pu.mutation.SetNextID(id)
	return pu
}

// SetNillableNextID sets the "next" edge to the Plan entity by ID if the given value is not nil.
func (pu *PlanUpdate) SetNillableNextID(id *int) *PlanUpdate {
	if id != nil {
		pu = pu.SetNextID(*id)
	}
	return pu
}

// SetNext sets the "next" edge to the Plan entity.
func (pu *PlanUpdate) SetNext(p *Plan) *PlanUpdate {
	return pu.SetNextID(p.ID)
}

// Mutation returns the PlanMutation object of the builder.
func (pu *PlanUpdate) Mutation() *PlanMutation {
	return pu.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (pu *PlanUpdate) ClearAuthor() *PlanUpdate {
	pu.mutation.ClearAuthor()
	return pu
}

// ClearDays clears all "days" edges to the Day entity.
func (pu *PlanUpdate) ClearDays() *PlanUpdate {
	pu.mutation.ClearDays()
	return pu
}

// RemoveDayIDs removes the "days" edge to Day entities by IDs.
func (pu *PlanUpdate) RemoveDayIDs(ids ...int) *PlanUpdate {
	pu.mutation.RemoveDayIDs(ids...)
	return pu
}

// RemoveDays removes "days" edges to Day entities.
func (pu *PlanUpdate) RemoveDays(d ...*Day) *PlanUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pu.RemoveDayIDs(ids...)
}

// ClearArbitrarySections clears all "arbitrarySections" edges to the ArbitrarySection entity.
func (pu *PlanUpdate) ClearArbitrarySections() *PlanUpdate {
	pu.mutation.ClearArbitrarySections()
	return pu
}

// RemoveArbitrarySectionIDs removes the "arbitrarySections" edge to ArbitrarySection entities by IDs.
func (pu *PlanUpdate) RemoveArbitrarySectionIDs(ids ...int) *PlanUpdate {
	pu.mutation.RemoveArbitrarySectionIDs(ids...)
	return pu
}

// RemoveArbitrarySections removes "arbitrarySections" edges to ArbitrarySection entities.
func (pu *PlanUpdate) RemoveArbitrarySections(a ...*ArbitrarySection) *PlanUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.RemoveArbitrarySectionIDs(ids...)
}

// ClearHeader clears the "header" edge to the Header entity.
func (pu *PlanUpdate) ClearHeader() *PlanUpdate {
	pu.mutation.ClearHeader()
	return pu
}

// ClearPrev clears the "prev" edge to the Plan entity.
func (pu *PlanUpdate) ClearPrev() *PlanUpdate {
	pu.mutation.ClearPrev()
	return pu
}

// ClearNext clears the "next" edge to the Plan entity.
func (pu *PlanUpdate) ClearNext() *PlanUpdate {
	pu.mutation.ClearNext()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlanUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlanUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlanUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlanUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PlanUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   plan.Table,
			Columns: plan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: plan.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Digest(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: plan.FieldDigest,
		})
	}
	if value, ok := pu.mutation.Txt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: plan.FieldTxt,
		})
	}
	if pu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   plan.AuthorTable,
			Columns: []string{plan.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   plan.AuthorTable,
			Columns: []string{plan.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.DaysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.DaysTable,
			Columns: plan.DaysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: day.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedDaysIDs(); len(nodes) > 0 && !pu.mutation.DaysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.DaysTable,
			Columns: plan.DaysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: day.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.DaysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.DaysTable,
			Columns: plan.DaysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: day.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ArbitrarySectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.ArbitrarySectionsTable,
			Columns: plan.ArbitrarySectionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: arbitrarysection.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedArbitrarySectionsIDs(); len(nodes) > 0 && !pu.mutation.ArbitrarySectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.ArbitrarySectionsTable,
			Columns: plan.ArbitrarySectionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: arbitrarysection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ArbitrarySectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.ArbitrarySectionsTable,
			Columns: plan.ArbitrarySectionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: arbitrarysection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.HeaderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   plan.HeaderTable,
			Columns: []string{plan.HeaderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: header.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.HeaderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   plan.HeaderTable,
			Columns: []string{plan.HeaderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: header.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.PrevCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   plan.PrevTable,
			Columns: []string{plan.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PrevIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   plan.PrevTable,
			Columns: []string{plan.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.NextCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   plan.NextTable,
			Columns: []string{plan.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.NextIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   plan.NextTable,
			Columns: []string{plan.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{plan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PlanUpdateOne is the builder for updating a single Plan entity.
type PlanUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlanMutation
}

// SetDigest sets the "digest" field.
func (puo *PlanUpdateOne) SetDigest(s string) *PlanUpdateOne {
	puo.mutation.SetDigest(s)
	return puo
}

// SetTxt sets the "txt" field.
func (puo *PlanUpdateOne) SetTxt(s string) *PlanUpdateOne {
	puo.mutation.SetTxt(s)
	return puo
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (puo *PlanUpdateOne) SetAuthorID(id int) *PlanUpdateOne {
	puo.mutation.SetAuthorID(id)
	return puo
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (puo *PlanUpdateOne) SetNillableAuthorID(id *int) *PlanUpdateOne {
	if id != nil {
		puo = puo.SetAuthorID(*id)
	}
	return puo
}

// SetAuthor sets the "author" edge to the User entity.
func (puo *PlanUpdateOne) SetAuthor(u *User) *PlanUpdateOne {
	return puo.SetAuthorID(u.ID)
}

// AddDayIDs adds the "days" edge to the Day entity by IDs.
func (puo *PlanUpdateOne) AddDayIDs(ids ...int) *PlanUpdateOne {
	puo.mutation.AddDayIDs(ids...)
	return puo
}

// AddDays adds the "days" edges to the Day entity.
func (puo *PlanUpdateOne) AddDays(d ...*Day) *PlanUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return puo.AddDayIDs(ids...)
}

// AddArbitrarySectionIDs adds the "arbitrarySections" edge to the ArbitrarySection entity by IDs.
func (puo *PlanUpdateOne) AddArbitrarySectionIDs(ids ...int) *PlanUpdateOne {
	puo.mutation.AddArbitrarySectionIDs(ids...)
	return puo
}

// AddArbitrarySections adds the "arbitrarySections" edges to the ArbitrarySection entity.
func (puo *PlanUpdateOne) AddArbitrarySections(a ...*ArbitrarySection) *PlanUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.AddArbitrarySectionIDs(ids...)
}

// SetHeaderID sets the "header" edge to the Header entity by ID.
func (puo *PlanUpdateOne) SetHeaderID(id int) *PlanUpdateOne {
	puo.mutation.SetHeaderID(id)
	return puo
}

// SetNillableHeaderID sets the "header" edge to the Header entity by ID if the given value is not nil.
func (puo *PlanUpdateOne) SetNillableHeaderID(id *int) *PlanUpdateOne {
	if id != nil {
		puo = puo.SetHeaderID(*id)
	}
	return puo
}

// SetHeader sets the "header" edge to the Header entity.
func (puo *PlanUpdateOne) SetHeader(h *Header) *PlanUpdateOne {
	return puo.SetHeaderID(h.ID)
}

// SetPrevID sets the "prev" edge to the Plan entity by ID.
func (puo *PlanUpdateOne) SetPrevID(id int) *PlanUpdateOne {
	puo.mutation.SetPrevID(id)
	return puo
}

// SetNillablePrevID sets the "prev" edge to the Plan entity by ID if the given value is not nil.
func (puo *PlanUpdateOne) SetNillablePrevID(id *int) *PlanUpdateOne {
	if id != nil {
		puo = puo.SetPrevID(*id)
	}
	return puo
}

// SetPrev sets the "prev" edge to the Plan entity.
func (puo *PlanUpdateOne) SetPrev(p *Plan) *PlanUpdateOne {
	return puo.SetPrevID(p.ID)
}

// SetNextID sets the "next" edge to the Plan entity by ID.
func (puo *PlanUpdateOne) SetNextID(id int) *PlanUpdateOne {
	puo.mutation.SetNextID(id)
	return puo
}

// SetNillableNextID sets the "next" edge to the Plan entity by ID if the given value is not nil.
func (puo *PlanUpdateOne) SetNillableNextID(id *int) *PlanUpdateOne {
	if id != nil {
		puo = puo.SetNextID(*id)
	}
	return puo
}

// SetNext sets the "next" edge to the Plan entity.
func (puo *PlanUpdateOne) SetNext(p *Plan) *PlanUpdateOne {
	return puo.SetNextID(p.ID)
}

// Mutation returns the PlanMutation object of the builder.
func (puo *PlanUpdateOne) Mutation() *PlanMutation {
	return puo.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (puo *PlanUpdateOne) ClearAuthor() *PlanUpdateOne {
	puo.mutation.ClearAuthor()
	return puo
}

// ClearDays clears all "days" edges to the Day entity.
func (puo *PlanUpdateOne) ClearDays() *PlanUpdateOne {
	puo.mutation.ClearDays()
	return puo
}

// RemoveDayIDs removes the "days" edge to Day entities by IDs.
func (puo *PlanUpdateOne) RemoveDayIDs(ids ...int) *PlanUpdateOne {
	puo.mutation.RemoveDayIDs(ids...)
	return puo
}

// RemoveDays removes "days" edges to Day entities.
func (puo *PlanUpdateOne) RemoveDays(d ...*Day) *PlanUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return puo.RemoveDayIDs(ids...)
}

// ClearArbitrarySections clears all "arbitrarySections" edges to the ArbitrarySection entity.
func (puo *PlanUpdateOne) ClearArbitrarySections() *PlanUpdateOne {
	puo.mutation.ClearArbitrarySections()
	return puo
}

// RemoveArbitrarySectionIDs removes the "arbitrarySections" edge to ArbitrarySection entities by IDs.
func (puo *PlanUpdateOne) RemoveArbitrarySectionIDs(ids ...int) *PlanUpdateOne {
	puo.mutation.RemoveArbitrarySectionIDs(ids...)
	return puo
}

// RemoveArbitrarySections removes "arbitrarySections" edges to ArbitrarySection entities.
func (puo *PlanUpdateOne) RemoveArbitrarySections(a ...*ArbitrarySection) *PlanUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.RemoveArbitrarySectionIDs(ids...)
}

// ClearHeader clears the "header" edge to the Header entity.
func (puo *PlanUpdateOne) ClearHeader() *PlanUpdateOne {
	puo.mutation.ClearHeader()
	return puo
}

// ClearPrev clears the "prev" edge to the Plan entity.
func (puo *PlanUpdateOne) ClearPrev() *PlanUpdateOne {
	puo.mutation.ClearPrev()
	return puo
}

// ClearNext clears the "next" edge to the Plan entity.
func (puo *PlanUpdateOne) ClearNext() *PlanUpdateOne {
	puo.mutation.ClearNext()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlanUpdateOne) Select(field string, fields ...string) *PlanUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Plan entity.
func (puo *PlanUpdateOne) Save(ctx context.Context) (*Plan, error) {
	var (
		err  error
		node *Plan
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Plan)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PlanMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlanUpdateOne) SaveX(ctx context.Context) *Plan {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlanUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlanUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PlanUpdateOne) sqlSave(ctx context.Context) (_node *Plan, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   plan.Table,
			Columns: plan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: plan.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Plan.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, plan.FieldID)
		for _, f := range fields {
			if !plan.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != plan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Digest(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: plan.FieldDigest,
		})
	}
	if value, ok := puo.mutation.Txt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: plan.FieldTxt,
		})
	}
	if puo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   plan.AuthorTable,
			Columns: []string{plan.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   plan.AuthorTable,
			Columns: []string{plan.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.DaysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.DaysTable,
			Columns: plan.DaysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: day.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedDaysIDs(); len(nodes) > 0 && !puo.mutation.DaysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.DaysTable,
			Columns: plan.DaysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: day.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.DaysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.DaysTable,
			Columns: plan.DaysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: day.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ArbitrarySectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.ArbitrarySectionsTable,
			Columns: plan.ArbitrarySectionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: arbitrarysection.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedArbitrarySectionsIDs(); len(nodes) > 0 && !puo.mutation.ArbitrarySectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.ArbitrarySectionsTable,
			Columns: plan.ArbitrarySectionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: arbitrarysection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ArbitrarySectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.ArbitrarySectionsTable,
			Columns: plan.ArbitrarySectionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: arbitrarysection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.HeaderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   plan.HeaderTable,
			Columns: []string{plan.HeaderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: header.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.HeaderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   plan.HeaderTable,
			Columns: []string{plan.HeaderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: header.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.PrevCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   plan.PrevTable,
			Columns: []string{plan.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PrevIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   plan.PrevTable,
			Columns: []string{plan.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.NextCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   plan.NextTable,
			Columns: []string{plan.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.NextIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   plan.NextTable,
			Columns: []string{plan.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: plan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Plan{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{plan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
