package codemode

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestAuthFromContext_APIKey(t *testing.T) {
	ctx := WithAuthHeaders(context.Background(), map[string]string{"X-Api-Key": "k1"})
	got := AuthFromContext(ctx)
	if got == nil || got.APIKey != "k1" {
		t.Fatalf("got %+v", got)
	}
}

func TestAuthFromContext_Bearer(t *testing.T) {
	ctx := WithAuthHeaders(context.Background(), map[string]string{"Authorization": "Bearer t1"})
	got := AuthFromContext(ctx)
	if got == nil || got.Bearer != "t1" {
		t.Fatalf("got %+v", got)
	}
}

func TestAuthFromContext_Both(t *testing.T) {
	ctx := WithAuthHeaders(context.Background(), map[string]string{
		"X-Api-Key":     "k1",
		"Authorization": "Bearer t1",
	})
	got := AuthFromContext(ctx)
	if got == nil || got.APIKey != "k1" || got.Bearer != "t1" {
		t.Fatalf("got %+v", got)
	}
}

func TestAuthFromContext_Missing(t *testing.T) {
	if AuthFromContext(context.Background()) != nil {
		t.Fatal("want nil")
	}
}

func TestAuthFromRequest_APIKey(t *testing.T) {
	req := mcp.CallToolRequest{}
	req.Header = map[string][]string{"X-Api-Key": {"k1"}}
	got := AuthFromRequest(context.Background(), req)
	if got == nil || got.APIKey != "k1" {
		t.Fatalf("got %+v", got)
	}
}

func TestAuthFromRequest_Bearer(t *testing.T) {
	req := mcp.CallToolRequest{}
	req.Header = map[string][]string{"Authorization": {"Bearer tok123"}}
	got := AuthFromRequest(context.Background(), req)
	if got == nil || got.Bearer != "tok123" {
		t.Fatalf("got %+v", got)
	}
}

func TestAuthFromRequest_FallsBackToContext(t *testing.T) {
	req := mcp.CallToolRequest{}
	ctx := WithAuthHeaders(context.Background(), map[string]string{"X-Api-Key": "ctx-key"})
	got := AuthFromRequest(ctx, req)
	if got == nil || got.APIKey != "ctx-key" {
		t.Fatalf("got %+v", got)
	}
}
