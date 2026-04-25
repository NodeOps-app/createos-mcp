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
