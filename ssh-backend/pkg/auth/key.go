package auth

import (
	"context"

	"github.com/aakarim/pland/ent"
)

type UserContextKey struct{}

func GetUserFromContext(ctx context.Context) *ent.User {
	v, ok := ctx.Value(UserContextKey{}).(*ent.User)
	if ok {
		return v
	}
	return nil
}
