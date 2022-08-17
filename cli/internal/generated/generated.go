// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package generated

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// CreatePlanCreatePlan includes the requested fields of the GraphQL type Plan.
type CreatePlanCreatePlan struct {
	Id          string `json:"id"`
	Txt         string `json:"txt"`
	Digest      string `json:"digest"`
	HasConflict bool   `json:"hasConflict"`
}

// GetId returns CreatePlanCreatePlan.Id, and is useful for accessing the field via an interface.
func (v *CreatePlanCreatePlan) GetId() string { return v.Id }

// GetTxt returns CreatePlanCreatePlan.Txt, and is useful for accessing the field via an interface.
func (v *CreatePlanCreatePlan) GetTxt() string { return v.Txt }

// GetDigest returns CreatePlanCreatePlan.Digest, and is useful for accessing the field via an interface.
func (v *CreatePlanCreatePlan) GetDigest() string { return v.Digest }

// GetHasConflict returns CreatePlanCreatePlan.HasConflict, and is useful for accessing the field via an interface.
func (v *CreatePlanCreatePlan) GetHasConflict() bool { return v.HasConflict }

// CreatePlanResponse is returned by CreatePlan on success.
type CreatePlanResponse struct {
	CreatePlan CreatePlanCreatePlan `json:"createPlan"`
}

// GetCreatePlan returns CreatePlanResponse.CreatePlan, and is useful for accessing the field via an interface.
func (v *CreatePlanResponse) GetCreatePlan() CreatePlanCreatePlan { return v.CreatePlan }

// GetLatestPlanMeUser includes the requested fields of the GraphQL type User.
type GetLatestPlanMeUser struct {
	// The most recent plan for this user
	Plan GetLatestPlanMeUserPlan `json:"plan"`
}

// GetPlan returns GetLatestPlanMeUser.Plan, and is useful for accessing the field via an interface.
func (v *GetLatestPlanMeUser) GetPlan() GetLatestPlanMeUserPlan { return v.Plan }

// GetLatestPlanMeUserPlan includes the requested fields of the GraphQL type Plan.
type GetLatestPlanMeUserPlan struct {
	Id     string `json:"id"`
	Txt    string `json:"txt"`
	Digest string `json:"digest"`
}

// GetId returns GetLatestPlanMeUserPlan.Id, and is useful for accessing the field via an interface.
func (v *GetLatestPlanMeUserPlan) GetId() string { return v.Id }

// GetTxt returns GetLatestPlanMeUserPlan.Txt, and is useful for accessing the field via an interface.
func (v *GetLatestPlanMeUserPlan) GetTxt() string { return v.Txt }

// GetDigest returns GetLatestPlanMeUserPlan.Digest, and is useful for accessing the field via an interface.
func (v *GetLatestPlanMeUserPlan) GetDigest() string { return v.Digest }

// GetLatestPlanResponse is returned by GetLatestPlan on success.
type GetLatestPlanResponse struct {
	Me GetLatestPlanMeUser `json:"me"`
}

// GetMe returns GetLatestPlanResponse.Me, and is useful for accessing the field via an interface.
func (v *GetLatestPlanResponse) GetMe() GetLatestPlanMeUser { return v.Me }

// __CreatePlanInput is used internally by genqlient
type __CreatePlanInput struct {
	Txt  string    `json:"txt"`
	Date time.Time `json:"date"`
}

// GetTxt returns __CreatePlanInput.Txt, and is useful for accessing the field via an interface.
func (v *__CreatePlanInput) GetTxt() string { return v.Txt }

// GetDate returns __CreatePlanInput.Date, and is useful for accessing the field via an interface.
func (v *__CreatePlanInput) GetDate() time.Time { return v.Date }

func CreatePlan(
	ctx context.Context,
	client graphql.Client,
	txt string,
	date time.Time,
) (*CreatePlanResponse, error) {
	req := &graphql.Request{
		OpName: "CreatePlan",
		Query: `
mutation CreatePlan ($txt: String!, $date: Time!) {
	createPlan(input: {txt:$txt,date:$date}) {
		id
		txt
		digest
		hasConflict
	}
}
`,
		Variables: &__CreatePlanInput{
			Txt:  txt,
			Date: date,
		},
	}
	var err error

	var data CreatePlanResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func GetLatestPlan(
	ctx context.Context,
	client graphql.Client,
) (*GetLatestPlanResponse, error) {
	req := &graphql.Request{
		OpName: "GetLatestPlan",
		Query: `
query GetLatestPlan {
	me {
		plan {
			id
			txt
			digest
		}
	}
}
`,
	}
	var err error

	var data GetLatestPlanResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
