package context

import (
	"context"
	"net/http"

	"github.com/3WDeveloper-GM/billings/cmd/pkg/domain"
)

type ContextKey string

const userKey = ContextKey("user")

func NewContext() *ContextKey {
  newContext := userKey
  return &newContext
}

func (c *ContextKey) ContextSetUser(r *http.Request, user *domain.Users) *http.Request {
	ctx := context.WithValue(r.Context(), userKey, user)
	return r.WithContext(ctx)
}

func (c *ContextKey) ContextGetUser(r *http.Request) *domain.Users {
	user, ok := r.Context().Value(userKey).(*domain.Users)
	if !ok {
		panic("missing user value in request context")
	}

	return user
}
