// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreatePlanInput represents a mutation input for creating plans.
type CreatePlanInput struct {
	CreatedAt   time.Time
	HasConflict *bool
	Digest      string
	Txt         string
	AuthorID    *int
	PrevID      *int
	NextID      *int
}

// Mutate applies the CreatePlanInput on the PlanMutation builder.
func (i *CreatePlanInput) Mutate(m *PlanMutation) {
	m.SetCreatedAt(i.CreatedAt)
	if v := i.HasConflict; v != nil {
		m.SetHasConflict(*v)
	}
	m.SetDigest(i.Digest)
	m.SetTxt(i.Txt)
	if v := i.AuthorID; v != nil {
		m.SetAuthorID(*v)
	}
	if v := i.PrevID; v != nil {
		m.SetPrevID(*v)
	}
	if v := i.NextID; v != nil {
		m.SetNextID(*v)
	}
}

// SetInput applies the change-set in the CreatePlanInput on the PlanCreate builder.
func (c *PlanCreate) SetInput(i CreatePlanInput) *PlanCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdatePlanInput represents a mutation input for updating plans.
type UpdatePlanInput struct {
	Digest      *string
	Txt         *string
	ClearAuthor bool
	AuthorID    *int
	ClearPrev   bool
	PrevID      *int
	ClearNext   bool
	NextID      *int
}

// Mutate applies the UpdatePlanInput on the PlanMutation builder.
func (i *UpdatePlanInput) Mutate(m *PlanMutation) {
	if v := i.Digest; v != nil {
		m.SetDigest(*v)
	}
	if v := i.Txt; v != nil {
		m.SetTxt(*v)
	}
	if i.ClearAuthor {
		m.ClearAuthor()
	}
	if v := i.AuthorID; v != nil {
		m.SetAuthorID(*v)
	}
	if i.ClearPrev {
		m.ClearPrev()
	}
	if v := i.PrevID; v != nil {
		m.SetPrevID(*v)
	}
	if i.ClearNext {
		m.ClearNext()
	}
	if v := i.NextID; v != nil {
		m.SetNextID(*v)
	}
}

// SetInput applies the change-set in the UpdatePlanInput on the PlanUpdate builder.
func (c *PlanUpdate) SetInput(i UpdatePlanInput) *PlanUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdatePlanInput on the PlanUpdateOne builder.
func (c *PlanUpdateOne) SetInput(i UpdatePlanInput) *PlanUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
