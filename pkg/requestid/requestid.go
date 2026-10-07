package requestid

import (
	"context"

	"github.com/google/uuid"
)

const Header = "X-Request-ID"

type contextKey struct{}

// Resolve accepts bounded, log-safe caller IDs and replaces invalid or absent IDs.
func Resolve(value string) string {
	if len(value) == 0 || len(value) > 128 {
		return uuid.NewString()
	}
	for _, c := range value {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '-' || c == '_' || c == '.') {
			return uuid.NewString()
		}
	}
	return value
}
func WithContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, contextKey{}, id)
}
func FromContext(ctx context.Context) string { id, _ := ctx.Value(contextKey{}).(string); return id }
