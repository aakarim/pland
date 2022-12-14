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
	"github.com/aakarim/pland/ent/plan"
	"github.com/aakarim/pland/ent/predicate"
)

// ArbitrarySectionUpdate is the builder for updating ArbitrarySection entities.
type ArbitrarySectionUpdate struct {
	config
	hooks    []Hook
	mutation *ArbitrarySectionMutation
}

// Where appends a list predicates to the ArbitrarySectionUpdate builder.
func (asu *ArbitrarySectionUpdate) Where(ps ...predicate.ArbitrarySection) *ArbitrarySectionUpdate {
	asu.mutation.Where(ps...)
	return asu
}

// SetToken sets the "token" field.
func (asu *ArbitrarySectionUpdate) SetToken(s string) *ArbitrarySectionUpdate {
	asu.mutation.SetToken(s)
	return asu
}

// SetTxt sets the "txt" field.
func (asu *ArbitrarySectionUpdate) SetTxt(s string) *ArbitrarySectionUpdate {
	asu.mutation.SetTxt(s)
	return asu
}

// AddPlanIDs adds the "plan" edge to the Plan entity by IDs.
func (asu *ArbitrarySectionUpdate) AddPlanIDs(ids ...int) *ArbitrarySectionUpdate {
	asu.mutation.AddPlanIDs(ids...)
	return asu
}

// AddPlan adds the "plan" edges to the Plan entity.
func (asu *ArbitrarySectionUpdate) AddPlan(p ...*Plan) *ArbitrarySectionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return asu.AddPlanIDs(ids...)
}

// Mutation returns the ArbitrarySectionMutation object of the builder.
func (asu *ArbitrarySectionUpdate) Mutation() *ArbitrarySectionMutation {
	return asu.mutation
}

// ClearPlan clears all "plan" edges to the Plan entity.
func (asu *ArbitrarySectionUpdate) ClearPlan() *ArbitrarySectionUpdate {
	asu.mutation.ClearPlan()
	return asu
}

// RemovePlanIDs removes the "plan" edge to Plan entities by IDs.
func (asu *ArbitrarySectionUpdate) RemovePlanIDs(ids ...int) *ArbitrarySectionUpdate {
	asu.mutation.RemovePlanIDs(ids...)
	return asu
}

// RemovePlan removes "plan" edges to Plan entities.
func (asu *ArbitrarySectionUpdate) RemovePlan(p ...*Plan) *ArbitrarySectionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return asu.RemovePlanIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (asu *ArbitrarySectionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(asu.hooks) == 0 {
		affected, err = asu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArbitrarySectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			asu.mutation = mutation
			affected, err = asu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(asu.hooks) - 1; i >= 0; i-- {
			if asu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, asu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (asu *ArbitrarySectionUpdate) SaveX(ctx context.Context) int {
	affected, err := asu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (asu *ArbitrarySectionUpdate) Exec(ctx context.Context) error {
	_, err := asu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asu *ArbitrarySectionUpdate) ExecX(ctx context.Context) {
	if err := asu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (asu *ArbitrarySectionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   arbitrarysection.Table,
			Columns: arbitrarysection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: arbitrarysection.FieldID,
			},
		},
	}
	if ps := asu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := asu.mutation.Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: arbitrarysection.FieldToken,
		})
	}
	if value, ok := asu.mutation.Txt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: arbitrarysection.FieldTxt,
		})
	}
	if asu.mutation.PlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   arbitrarysection.PlanTable,
			Columns: arbitrarysection.PlanPrimaryKey,
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
	if nodes := asu.mutation.RemovedPlanIDs(); len(nodes) > 0 && !asu.mutation.PlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   arbitrarysection.PlanTable,
			Columns: arbitrarysection.PlanPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asu.mutation.PlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   arbitrarysection.PlanTable,
			Columns: arbitrarysection.PlanPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, asu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{arbitrarysection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ArbitrarySectionUpdateOne is the builder for updating a single ArbitrarySection entity.
type ArbitrarySectionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArbitrarySectionMutation
}

// SetToken sets the "token" field.
func (asuo *ArbitrarySectionUpdateOne) SetToken(s string) *ArbitrarySectionUpdateOne {
	asuo.mutation.SetToken(s)
	return asuo
}

// SetTxt sets the "txt" field.
func (asuo *ArbitrarySectionUpdateOne) SetTxt(s string) *ArbitrarySectionUpdateOne {
	asuo.mutation.SetTxt(s)
	return asuo
}

// AddPlanIDs adds the "plan" edge to the Plan entity by IDs.
func (asuo *ArbitrarySectionUpdateOne) AddPlanIDs(ids ...int) *ArbitrarySectionUpdateOne {
	asuo.mutation.AddPlanIDs(ids...)
	return asuo
}

// AddPlan adds the "plan" edges to the Plan entity.
func (asuo *ArbitrarySectionUpdateOne) AddPlan(p ...*Plan) *ArbitrarySectionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return asuo.AddPlanIDs(ids...)
}

// Mutation returns the ArbitrarySectionMutation object of the builder.
func (asuo *ArbitrarySectionUpdateOne) Mutation() *ArbitrarySectionMutation {
	return asuo.mutation
}

// ClearPlan clears all "plan" edges to the Plan entity.
func (asuo *ArbitrarySectionUpdateOne) ClearPlan() *ArbitrarySectionUpdateOne {
	asuo.mutation.ClearPlan()
	return asuo
}

// RemovePlanIDs removes the "plan" edge to Plan entities by IDs.
func (asuo *ArbitrarySectionUpdateOne) RemovePlanIDs(ids ...int) *ArbitrarySectionUpdateOne {
	asuo.mutation.RemovePlanIDs(ids...)
	return asuo
}

// RemovePlan removes "plan" edges to Plan entities.
func (asuo *ArbitrarySectionUpdateOne) RemovePlan(p ...*Plan) *ArbitrarySectionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return asuo.RemovePlanIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (asuo *ArbitrarySectionUpdateOne) Select(field string, fields ...string) *ArbitrarySectionUpdateOne {
	asuo.fields = append([]string{field}, fields...)
	return asuo
}

// Save executes the query and returns the updated ArbitrarySection entity.
func (asuo *ArbitrarySectionUpdateOne) Save(ctx context.Context) (*ArbitrarySection, error) {
	var (
		err  error
		node *ArbitrarySection
	)
	if len(asuo.hooks) == 0 {
		node, err = asuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArbitrarySectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			asuo.mutation = mutation
			node, err = asuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(asuo.hooks) - 1; i >= 0; i-- {
			if asuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, asuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ArbitrarySection)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ArbitrarySectionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (asuo *ArbitrarySectionUpdateOne) SaveX(ctx context.Context) *ArbitrarySection {
	node, err := asuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (asuo *ArbitrarySectionUpdateOne) Exec(ctx context.Context) error {
	_, err := asuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asuo *ArbitrarySectionUpdateOne) ExecX(ctx context.Context) {
	if err := asuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (asuo *ArbitrarySectionUpdateOne) sqlSave(ctx context.Context) (_node *ArbitrarySection, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   arbitrarysection.Table,
			Columns: arbitrarysection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: arbitrarysection.FieldID,
			},
		},
	}
	id, ok := asuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ArbitrarySection.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := asuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, arbitrarysection.FieldID)
		for _, f := range fields {
			if !arbitrarysection.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != arbitrarysection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := asuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := asuo.mutation.Token(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: arbitrarysection.FieldToken,
		})
	}
	if value, ok := asuo.mutation.Txt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: arbitrarysection.FieldTxt,
		})
	}
	if asuo.mutation.PlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   arbitrarysection.PlanTable,
			Columns: arbitrarysection.PlanPrimaryKey,
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
	if nodes := asuo.mutation.RemovedPlanIDs(); len(nodes) > 0 && !asuo.mutation.PlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   arbitrarysection.PlanTable,
			Columns: arbitrarysection.PlanPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asuo.mutation.PlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   arbitrarysection.PlanTable,
			Columns: arbitrarysection.PlanPrimaryKey,
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
	_node = &ArbitrarySection{config: asuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, asuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{arbitrarysection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
