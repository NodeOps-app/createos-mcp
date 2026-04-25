package codemode

import (
	"context"
	"os"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

// Auth flows:
//   - HTTP transport: mcp-go attaches request headers to mcp.CallToolRequest.
//     AuthFromRequest reads them directly. HTTP middleware also mirrors them
//     into ctx via WithAuthHeaders as a backup path.
//   - stdio transport: there is no per-request HTTP header. Auth comes from
//     environment variables: CREATEOS_API_KEY and/or CREATEOS_BEARER are
//     read at fall-through.
//   - AuthFromRequest is the source of truth and cascades: request header →
//     ctx headers → env.
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

// AuthFromEnv reads CREATEOS_API_KEY / CREATEOS_BEARER for stdio transport.
// Returns nil if neither is set.
func AuthFromEnv() *AuthCtx {
	out := &AuthCtx{
		APIKey: os.Getenv("CREATEOS_API_KEY"),
		Bearer: os.Getenv("CREATEOS_BEARER"),
	}
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
	if a := AuthFromContext(ctx); a != nil {
		return a
	}
	return AuthFromEnv()
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
