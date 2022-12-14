// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (as *ArbitrarySection) Plan(ctx context.Context) ([]*Plan, error) {
	result, err := as.Edges.PlanOrErr()
	if IsNotLoaded(err) {
		result, err = as.QueryPlan().All(ctx)
	}
	return result, err
}

func (d *Day) Plan(ctx context.Context) ([]*Plan, error) {
	result, err := d.Edges.PlanOrErr()
	if IsNotLoaded(err) {
		result, err = d.QueryPlan().All(ctx)
	}
	return result, err
}

func (h *Header) Plan(ctx context.Context) (*Plan, error) {
	result, err := h.Edges.PlanOrErr()
	if IsNotLoaded(err) {
		result, err = h.QueryPlan().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pl *Plan) Author(ctx context.Context) (*User, error) {
	result, err := pl.Edges.AuthorOrErr()
	if IsNotLoaded(err) {
		result, err = pl.QueryAuthor().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pl *Plan) Days(ctx context.Context) ([]*Day, error) {
	result, err := pl.Edges.DaysOrErr()
	if IsNotLoaded(err) {
		result, err = pl.QueryDays().All(ctx)
	}
	return result, err
}

func (pl *Plan) ArbitrarySections(ctx context.Context) ([]*ArbitrarySection, error) {
	result, err := pl.Edges.ArbitrarySectionsOrErr()
	if IsNotLoaded(err) {
		result, err = pl.QueryArbitrarySections().All(ctx)
	}
	return result, err
}

func (pl *Plan) Header(ctx context.Context) (*Header, error) {
	result, err := pl.Edges.HeaderOrErr()
	if IsNotLoaded(err) {
		result, err = pl.QueryHeader().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pl *Plan) Prev(ctx context.Context) (*Plan, error) {
	result, err := pl.Edges.PrevOrErr()
	if IsNotLoaded(err) {
		result, err = pl.QueryPrev().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pl *Plan) Next(ctx context.Context) (*Plan, error) {
	result, err := pl.Edges.NextOrErr()
	if IsNotLoaded(err) {
		result, err = pl.QueryNext().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Plans(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *PlanOrder, where *PlanWhereInput,
) (*PlanConnection, error) {
	opts := []PlanPaginateOption{
		WithPlanOrder(orderBy),
		WithPlanFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[0][alias]
	if nodes, err := u.NamedPlans(alias); err == nil || hasTotalCount {
		pager, err := newPlanPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &PlanConnection{Edges: []*PlanEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryPlans().Paginate(ctx, after, first, before, last, opts...)
}
