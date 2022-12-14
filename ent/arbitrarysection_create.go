// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/aakarim/pland/ent/arbitrarysection"
	"github.com/aakarim/pland/ent/plan"
)

// ArbitrarySectionCreate is the builder for creating a ArbitrarySection entity.
type ArbitrarySectionCreate struct {
	config
	mutation *ArbitrarySectionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (asc *ArbitrarySectionCreate) SetCreatedAt(t time.Time) *ArbitrarySectionCreate {
	asc.mutation.SetCreatedAt(t)
	return asc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (asc *ArbitrarySectionCreate) SetNillableCreatedAt(t *time.Time) *ArbitrarySectionCreate {
	if t != nil {
		asc.SetCreatedAt(*t)
	}
	return asc
}

// SetToken sets the "token" field.
func (asc *ArbitrarySectionCreate) SetToken(s string) *ArbitrarySectionCreate {
	asc.mutation.SetToken(s)
	return asc
}

// SetTxt sets the "txt" field.
func (asc *ArbitrarySectionCreate) SetTxt(s string) *ArbitrarySectionCreate {
	asc.mutation.SetTxt(s)
	return asc
}

// AddPlanIDs adds the "plan" edge to the Plan entity by IDs.
func (asc *ArbitrarySectionCreate) AddPlanIDs(ids ...int) *ArbitrarySectionCreate {
	asc.mutation.AddPlanIDs(ids...)
	return asc
}

// AddPlan adds the "plan" edges to the Plan entity.
func (asc *ArbitrarySectionCreate) AddPlan(p ...*Plan) *ArbitrarySectionCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return asc.AddPlanIDs(ids...)
}

// Mutation returns the ArbitrarySectionMutation object of the builder.
func (asc *ArbitrarySectionCreate) Mutation() *ArbitrarySectionMutation {
	return asc.mutation
}

// Save creates the ArbitrarySection in the database.
func (asc *ArbitrarySectionCreate) Save(ctx context.Context) (*ArbitrarySection, error) {
	var (
		err  error
		node *ArbitrarySection
	)
	asc.defaults()
	if len(asc.hooks) == 0 {
		if err = asc.check(); err != nil {
			return nil, err
		}
		node, err = asc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArbitrarySectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = asc.check(); err != nil {
				return nil, err
			}
			asc.mutation = mutation
			if node, err = asc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(asc.hooks) - 1; i >= 0; i-- {
			if asc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, asc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (asc *ArbitrarySectionCreate) SaveX(ctx context.Context) *ArbitrarySection {
	v, err := asc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (asc *ArbitrarySectionCreate) Exec(ctx context.Context) error {
	_, err := asc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asc *ArbitrarySectionCreate) ExecX(ctx context.Context) {
	if err := asc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asc *ArbitrarySectionCreate) defaults() {
	if _, ok := asc.mutation.CreatedAt(); !ok {
		v := arbitrarysection.DefaultCreatedAt()
		asc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asc *ArbitrarySectionCreate) check() error {
	if _, ok := asc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ArbitrarySection.created_at"`)}
	}
	if _, ok := asc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "ArbitrarySection.token"`)}
	}
	if _, ok := asc.mutation.Txt(); !ok {
		return &ValidationError{Name: "txt", err: errors.New(`ent: missing required field "ArbitrarySection.txt"`)}
	}
	return nil
}

func (asc *ArbitrarySectionCreate) sqlSave(ctx context.Context) (*ArbitrarySection, error) {
	_node, _spec := asc.createSpec()
	if err := sqlgraph.CreateNode(ctx, asc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (asc *ArbitrarySectionCreate) createSpec() (*ArbitrarySection, *sqlgraph.CreateSpec) {
	var (
		_node = &ArbitrarySection{config: asc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: arbitrarysection.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: arbitrarysection.FieldID,
			},
		}
	)
	if value, ok := asc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: arbitrarysection.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := asc.mutation.Token(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: arbitrarysection.FieldToken,
		})
		_node.Token = value
	}
	if value, ok := asc.mutation.Txt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: arbitrarysection.FieldTxt,
		})
		_node.Txt = value
	}
	if nodes := asc.mutation.PlanIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ArbitrarySectionCreateBulk is the builder for creating many ArbitrarySection entities in bulk.
type ArbitrarySectionCreateBulk struct {
	config
	builders []*ArbitrarySectionCreate
}

// Save creates the ArbitrarySection entities in the database.
func (ascb *ArbitrarySectionCreateBulk) Save(ctx context.Context) ([]*ArbitrarySection, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ascb.builders))
	nodes := make([]*ArbitrarySection, len(ascb.builders))
	mutators := make([]Mutator, len(ascb.builders))
	for i := range ascb.builders {
		func(i int, root context.Context) {
			builder := ascb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArbitrarySectionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ascb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ascb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ascb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ascb *ArbitrarySectionCreateBulk) SaveX(ctx context.Context) []*ArbitrarySection {
	v, err := ascb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ascb *ArbitrarySectionCreateBulk) Exec(ctx context.Context) error {
	_, err := ascb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ascb *ArbitrarySectionCreateBulk) ExecX(ctx context.Context) {
	if err := ascb.Exec(ctx); err != nil {
		panic(err)
	}
}
