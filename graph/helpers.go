package graph

import (
	"context"
	"fmt"
	"time"

	"github.com/aakarim/pland/ent"
	planEntity "github.com/aakarim/pland/pkg/plan"
)

func (r *mutationResolver) savePlanSections(ctx context.Context, p *planEntity.PlanFile) ([]*ent.Day, []*ent.ArbitrarySection, *ent.Header, error) {
	pDays := make([]*ent.Day, 0, len(p.Days))
	pArbSections := make([]*ent.ArbitrarySection, 0, len(p.ArbitrarySections))
	pHeader, err := r.Client.Header.Create().SetTxt(p.Header.Contents).Save(ctx)
	if err != nil {
		return pDays, pArbSections, nil, fmt.Errorf("Header: %w", err)
	}
	for _, d := range p.Days {
		day, err := r.Client.Day.Create().
			SetDate(d.Date).
			SetTxt(d.Contents).
			Save(ctx)
		if err != nil {
			return pDays, pArbSections, nil, fmt.Errorf("Day: %w", err)
		}
		pDays = append(pDays, day)
	}
	for _, a := range p.ArbitrarySections {
		arb, err := r.Client.ArbitrarySection.Create().
			SetTxt(a.Contents).
			SetToken(a.Token).Save(ctx)
		if err != nil {
			return pDays, pArbSections, nil, fmt.Errorf("Arb: %w", err)
		}
		pArbSections = append(pArbSections, arb)
	}
	return pDays, pArbSections, pHeader, nil
}
func (r *mutationResolver) makePlan(ctx context.Context, user *ent.User, p *planEntity.PlanFile) (*ent.PlanCreate, error) {
	days, arbitrarySections, header, err := r.savePlanSections(ctx, p)
	if err != nil {
		return nil, fmt.Errorf("savePlanSections: %w", err)
	}
	return r.Client.Plan.Create().
		SetAuthor(user).
		SetDigest(p.Digest()).
		SetCreatedAt(time.Now()).
		SetHasConflict(false).
		SetTxt(p.String()).
		AddDays(days...).
		AddArbitrarySections(arbitrarySections...).
		SetHeader(header), nil
}
