package requestid

import (
	"context"

	"github.com/google/uuid"
)

const Header = "X-Request-ID"

type contextKey struct{}

// New generates a server-owned correlation ID for each request.
func New() string { return uuid.NewString() }

func WithContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, contextKey{}, id)
}
func FromContext(ctx context.Context) string { id, _ := ctx.Value(contextKey{}).(string); return id }
