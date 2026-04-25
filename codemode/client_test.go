package codemode

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClient_Run_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/run" || r.Method != "POST" {
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
		var req RunRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatal(err)
		}
		if req.Mode != ModeSearch {
			t.Fatalf("want search, got %s", req.Mode)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"done","result":[1,2,3],"logs":[]}`))
	}))
	defer srv.Close()

	c := NewClient(srv.URL)
	got, err := c.Run(context.Background(), RunRequest{Mode: ModeSearch, Code: "async () => [1,2,3]"})
	if err != nil {
		t.Fatal(err)
	}
	if got.Status != "done" {
		t.Fatalf("want done, got %s", got.Status)
	}
	if !strings.Contains(string(got.Result), "1") {
		t.Fatalf("unexpected result %s", got.Result)
	}
}

func TestClient_Run_503(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte("spec not loaded"))
	}))
	defer srv.Close()

	c := NewClient(srv.URL)
	_, err := c.Run(context.Background(), RunRequest{Mode: ModeSearch, Code: "x"})
	if err == nil {
		t.Fatal("want error")
	}
	if !strings.Contains(err.Error(), "unavailable") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestClient_Health_OK(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	c := NewClient(srv.URL)
	if err := c.Health(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func TestClient_Health_503(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte("not ready"))
	}))
	defer srv.Close()

	c := NewClient(srv.URL)
	if err := c.Health(context.Background()); err == nil {
		t.Fatal("want error")
	}
}
