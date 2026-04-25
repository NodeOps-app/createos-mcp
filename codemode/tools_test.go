package codemode

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestHandler_Search_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"status":"done","result":["a","b"],"logs":[]}`))
	}))
	defer srv.Close()

	h := &Handler{Client: NewClient(srv.URL)}
	tool := NewSearchTool()
	if tool.Name != "search" {
		t.Fatalf("name = %q", tool.Name)
	}

	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]interface{}{"code": "async () => spec"}
	res, err := h.Search(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if res.IsError {
		t.Fatalf("unexpected error result: %#v", res.Content)
	}
	got, ok := res.Content[0].(mcp.TextContent)
	if !ok {
		t.Fatalf("not text content: %T", res.Content[0])
	}
	var parsed RunResult
	if err := json.Unmarshal([]byte(got.Text), &parsed); err != nil {
		t.Fatal(err)
	}
	if parsed.Status != "done" {
		t.Fatalf("status = %q", parsed.Status)
	}
}

func TestHandler_Execute_ForwardsAuth(t *testing.T) {
	var seen RunRequest
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewDecoder(r.Body).Decode(&seen)
		_, _ = w.Write([]byte(`{"status":"done","result":1}`))
	}))
	defer srv.Close()

	h := &Handler{Client: NewClient(srv.URL)}
	ctx := WithAuthHeaders(context.Background(), map[string]string{"X-Api-Key": "k"})
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]interface{}{"code": "async () => 1"}
	res, err := h.Execute(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	if res.IsError {
		t.Fatalf("unexpected error: %#v", res.Content)
	}
	if seen.Mode != ModeExecute {
		t.Fatalf("mode %q", seen.Mode)
	}
	if seen.AuthCtx == nil || seen.AuthCtx.APIKey != "k" {
		t.Fatalf("authCtx %+v", seen.AuthCtx)
	}
}

func TestHandler_Search_MissingCode(t *testing.T) {
	h := &Handler{Client: NewClient("http://invalid")}
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]interface{}{}
	res, err := h.Search(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if !res.IsError {
		t.Fatal("want error")
	}
	got := res.Content[0].(mcp.TextContent).Text
	if !strings.Contains(got, "code") {
		t.Fatalf("unexpected msg: %s", got)
	}
}
