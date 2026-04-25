package codemode

import (
	"context"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

// Auth flows:
//   - HTTP transport: mcp-go attaches request headers to mcp.CallToolRequest.
//     Use AuthFromRequest in the tool handler to extract them.
//   - stdio transport: middleware can stash headers on the context via
//     WithAuthHeaders so AuthFromContext finds them.
//   - AuthFromRequest is the source of truth and falls back to
//     AuthFromContext for stdio.
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
	out := authCtxFromHeaders(raw)
	if out.APIKey == "" && out.Bearer == "" {
		return nil
	}
	return out
}

// AuthFromRequest extracts caller auth headers from the MCP request, matching
// the native handlers/request.go GetAuthInfo behavior. Falls back to
// AuthFromContext for stdio transport.
func AuthFromRequest(ctx context.Context, req mcp.CallToolRequest) *AuthCtx {
	out := &AuthCtx{}
	if v := req.Header.Get("X-Api-Key"); v != "" {
		out.APIKey = v
	}
	if h := req.Header.Get("Authorization"); h != "" {
		parts := strings.SplitN(h, " ", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], "bearer") && parts[1] != "" {
			out.Bearer = parts[1]
		}
	}
	if out.APIKey != "" || out.Bearer != "" {
		return out
	}
	return AuthFromContext(ctx)
}

func authCtxFromHeaders(raw map[string]string) *AuthCtx {
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
	return out
}
