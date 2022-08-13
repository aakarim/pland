package plans

import (
	"context"

	"github.com/cespare/xxhash/v2"
	"github.com/charmbracelet/charm/server"
)

type PlanStore struct {
	charmConfig *server.Config
}

type Plan struct {
	TS   int64         `json:"timestamp"`
	Hash xxhash.Digest `json:"hash"` // base64
	Txt  string        `json:"txt"`
}

func (p *PlanStore) GetPlansForUser(ctx context.Context, userID string) (plans []Plan, err error) {

	return plans, err
}
