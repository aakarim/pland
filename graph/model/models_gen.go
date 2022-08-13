// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/aakarim/pland/ent"
)

type CreatePlanCLIInput struct {
	Txt       string     `json:"txt"`
	Date      time.Time  `json:"date"`
	Timestamp *time.Time `json:"timestamp"`
}

type Fyp struct {
	Edges    []*ent.PlanEdge `json:"edges"`
	PageInfo *ent.PageInfo   `json:"pageInfo"`
}