package codemode

import (
	"context"
	"testing"
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
