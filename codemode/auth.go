package codemode

import (
	"context"
	"strings"
)

// AuthFromContext returns the caller's auth context to forward to workerd.
//
// Source of truth:
//   - HTTP transport: middleware stores headers via WithAuthHeaders.
//   - stdio transport: env CREATEOS_API_KEY / CREATEOS_BEARER fall back into
//     the same context key at startup.
//
// Returns nil if no auth present.
type ctxKey int

const authHeadersKey ctxKey = 1

func WithAuthHeaders(ctx context.Context, headers map[string]string) context.Context {
	return context.WithValue(ctx, authHeadersKey, headers)
}

func AuthFromContext(ctx context.Context) *AuthCtx {
	raw, ok := ctx.Value(authHeadersKey).(map[string]string)
	if !ok {
		return nil
	}
	out := &AuthCtx{}
	for k, v := range raw {
		switch strings.ToLower(k) {
		case "x-api-key":
			out.APIKey = v
		case "authorization":
			if strings.HasPrefix(strings.ToLower(v), "bearer ") {
				out.Bearer = v[len("Bearer "):]
			}
		}
	}
	if out.APIKey == "" && out.Bearer == "" {
		return nil
	}
	return out
}
