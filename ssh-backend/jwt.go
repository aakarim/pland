package main

import (
	"context"
	"crypto"
	"net/http"

	"github.com/aakarim/pland/ent"
	"github.com/aakarim/pland/ent/user"
	"github.com/aakarim/pland/ssh-backend/pkg/auth"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/google/uuid"
)

// JWTMiddleware creates a new middleware function that will validate JWT
// tokens based on the supplied public key.
func JWTMiddleware(pk crypto.PublicKey, iss string, aud []string) (func(http.Handler) http.Handler, error) {
	jm, err := jwtMiddlewareImpl(pk, iss, aud)
	if err != nil {
		return nil, err
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jm(next).ServeHTTP(w, r)
		})
	}, nil
}

func AuthMiddleware(client *ent.Client) (func(http.Handler) http.Handler, error) {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
			if ok {
				uuidClaim, err := uuid.Parse(claims.RegisteredClaims.Subject)
				if err != nil {
					w.WriteHeader(401)
					w.Write([]byte(err.Error()))
				}
				user, err := client.User.Query().Where(user.CharmID(uuidClaim)).First(r.Context())
				if err != nil {
					w.WriteHeader(401)
					w.Write([]byte(err.Error()))
				}
				r = r.Clone(context.WithValue(r.Context(), auth.UserContextKey{}, user))

			}
			next.ServeHTTP(w, r)
		})
	}, nil
}

func jwtMiddlewareImpl(pk crypto.PublicKey, iss string, aud []string) (func(http.Handler) http.Handler, error) {
	kf := func(ctx context.Context) (interface{}, error) {
		return pk, nil
	}
	v, err := validator.New(
		kf,
		validator.EdDSA,
		iss,
		aud,
	)
	if err != nil {
		return nil, err
	}
	mw := jwtmiddleware.New(v.ValidateToken,
		jwtmiddleware.WithCredentialsOptional(true))
	return mw.CheckJWT, nil
}
