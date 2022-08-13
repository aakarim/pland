package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aakarim/pland/ent"
	"github.com/aakarim/pland/ent/plan"
	"github.com/aakarim/pland/ent/user"
	"github.com/aakarim/pland/graph/generated"
	"github.com/aakarim/pland/graph/model"
	"github.com/aakarim/pland/ssh-backend/pkg/auth"
	xxhash "github.com/cespare/xxhash/v2"
)

func (r *mutationResolver) CreatePlan(ctx context.Context, input model.CreatePlanCLIInput) (*ent.Plan, error) {
	user := auth.GetUserFromContext(ctx)
	if user == nil {
		return nil, ErrAccessDenied
	}
	// convert timestamp to date
	y, m, d := input.Date.Date()
	setDate, err := time.Parse("2006-1-2", fmt.Sprintf("%d-%d-%d", y, m, d))
	if err != nil {
		return nil, fmt.Errorf("time.Parse(): %w", err)
	}
	digest := fmt.Sprintf("%d", xxhash.Sum64String(input.Txt))
	existingPlan, err := r.Client.Plan.Query().Where(
		plan.DateEQ(setDate),
		plan.DigestEQ(digest),
	).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("error retrieving existingPlan: %w", err)
	}
	if existingPlan != nil {
		log.Println("not saving existing plan", existingPlan.ID)
		return existingPlan, nil
	}
	ts := input.Timestamp
	if ts == nil {
		now := time.Now()
		ts = &now
	}
	log.Printf("storing plan date %d-%d-%d for user %v", input.Date.Year(), input.Date.Month(), input.Date.Day(), user.ID)
	return r.Client.Plan.Create().
		SetAuthor(user).
		SetDate(setDate).
		SetDigest(digest).
		SetTimestamp(*ts).
		SetCreatedAt(time.Now()).
		SetTxt(input.Txt).Save(ctx)
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
	plans, err := obj.QueryPlans().Order(ent.Desc(plan.FieldDate), ent.Desc(plan.FieldTimestamp)).All(ctx)
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
