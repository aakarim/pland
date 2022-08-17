package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aakarim/pland/ent"
	"github.com/aakarim/pland/ent/plan"
	"github.com/aakarim/pland/ent/user"
	"github.com/aakarim/pland/graph/generated"
	"github.com/aakarim/pland/graph/model"
	planEntity "github.com/aakarim/pland/pkg/plan"
	"github.com/aakarim/pland/ssh-backend/pkg/auth"
)

func (r *mutationResolver) CreatePlan(ctx context.Context, input model.CreatePlanCLIInput) (*ent.Plan, error) {
	user := auth.GetUserFromContext(ctx)
	if user == nil {
		return nil, ErrAccessDenied
	}

	p, err := planEntity.Parse(ctx, input.Txt)
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}
	// get the tip
	tail, err := r.Client.Plan.Query().
		Where(plan.Not(plan.HasNext())).
		WithPrev().
		Only(ctx)
	noTail := ent.IsNotFound(err)
	if err != nil && !noTail {
		return nil, fmt.Errorf("could not get tail: %w", err)
	}
	// if no tail then this is the first one
	if noTail {
		tail, err := r.Client.Plan.Create().
			SetAuthor(user).
			SetDigest(p.Digest()).
			SetCreatedAt(time.Now()).
			SetHasConflict(false).
			SetTxt(p.String()).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not save plan: %w", err)
		}
		return tail, nil
	}
	// if they don't match then we need to do a merge
	tailPrev, err := tail.Prev(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("tail.Prev: %w", err)
	}
	// this is second one, no conflict
	if tailPrev == nil {
		tail, err := r.Client.Plan.Create().
			SetAuthor(user).
			SetDigest(p.Digest()).
			SetCreatedAt(time.Now()).
			SetPrev(tail).
			SetHasConflict(false).
			SetTxt(p.String()).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not save plan: %w", err)
		}
		return tail, nil
	}

	// tailPrev being nil means that the tail is the first entry
	var conflict bool
	if tailPrev.ID != p.ParentVersion && tail.Digest != p.Digest() {
		pTail, err := planEntity.Parse(ctx, tail.Txt)
		if err != nil {
			return nil, fmt.Errorf("could not parse tail: %w", err)
		}
		p, err = planEntity.Diff(p, pTail)
		if err != nil && !errors.Is(err, planEntity.ErrConflict) {
			return nil, fmt.Errorf("could not diff: %w", err)
		}
		conflict = errors.Is(err, planEntity.ErrConflict)

	}
	// save plan file
	p.ParentVersion = tail.ID
	tail, err = r.Client.Plan.Create().
		SetAuthor(user).
		SetDigest(p.Digest()).
		SetCreatedAt(time.Now()).
		SetHasConflict(conflict).
		SetPrev(tail).
		SetTxt(p.String()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not save plan: %w", err)
	}
	// TODO: build new derived data from plan/ fire event
	return tail, nil
}

func (r *queryResolver) User(ctx context.Context, name string) (*ent.User, error) {
	return r.Client.User.Query().Where(user.Name(name)).First(ctx)
}

func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	user := auth.GetUserFromContext(ctx)
	if user == nil {
		return nil, ErrAccessDenied
	}
	return user, nil
}

func (r *queryResolver) Fyp(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor) (*model.Fyp, error) {
	// return r.Client.Plan.Query().Order(ent.Desc(plan.FieldDate)).All()
	return &model.Fyp{}, nil
}

func (r *userResolver) Plan(ctx context.Context, obj *ent.User) (*ent.Plan, error) {
	plans, err := obj.QueryPlans().Order(ent.Desc(plan.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(plans) == 0 {
		return nil, nil
	}
	return plans[0], nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
