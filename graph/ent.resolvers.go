package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/aakarim/pland/ent"
	"github.com/aakarim/pland/graph/generated"
)

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.Client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) CharmID(ctx context.Context, obj *ent.User) (string, error) {
	return obj.CharmID.String(), nil
}

func (r *userWhereInputResolver) CharmID(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) CharmIDNeq(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) CharmIDIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) CharmIDNotIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) CharmIDGt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) CharmIDGte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) CharmIDLt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *userWhereInputResolver) CharmIDLte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// UserWhereInput returns generated.UserWhereInputResolver implementation.
func (r *Resolver) UserWhereInput() generated.UserWhereInputResolver {
	return &userWhereInputResolver{r}
}

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userWhereInputResolver struct{ *Resolver }
