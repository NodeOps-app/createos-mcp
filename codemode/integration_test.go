//go:build integration

package codemode_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/NodeOps-app/createos-mcp/codemode"
)

func startBackend(t *testing.T) *httptest.Server {
	t.Helper()
	yamlBytes, err := os.ReadFile("testdata/openapi-mini.yaml")
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api-docs/openapi.yaml", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		_, _ = w.Write(yamlBytes)
	})
	mux.HandleFunc("/v1/foo", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Api-Key") != "test-key" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"missing api key"}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"items":[{"id":"f1"}]}`))
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}

func startWorkerd(t *testing.T, backendURL string) (string, func()) {
	t.Helper()
	wd, _ := filepath.Abs("workerd")
	// Build the bundled spec-loader so workerd can embed it.
	build := exec.Command("bun", "run", "build")
	build.Dir = wd
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("bun build failed: %v\n%s", err, out)
	}
	cmd := exec.Command("./node_modules/.bin/workerd", "serve", "--experimental", "config.capnp")
	cmd.Dir = wd
	cmd.Env = append(os.Environ(), "BACKEND_URL="+backendURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	cancel := func() {
		_ = cmd.Process.Kill()
		_, _ = cmd.Process.Wait()
	}
	t.Cleanup(cancel)
	deadline := time.Now().Add(15 * time.Second)
	c := codemode.NewClient("http://127.0.0.1:8787")
	for time.Now().Before(deadline) {
		if err := c.Health(context.Background()); err == nil {
			return "http://127.0.0.1:8787", cancel
		}
		time.Sleep(200 * time.Millisecond)
	}
	cancel()
	t.Fatal("workerd never became healthy")
	return "", cancel
}

func TestIntegration_Search(t *testing.T) {
	if os.Getenv("CODEMODE_INTEGRATION") == "" {
		t.Skip("set CODEMODE_INTEGRATION=1 to run")
	}
	backend := startBackend(t)
	url, _ := startWorkerd(t, backend.URL)
	c := codemode.NewClient(url)

	res, err := c.Run(context.Background(), codemode.RunRequest{
		Mode: codemode.ModeSearch,
		Code: "async () => Object.keys(spec.paths)",
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Status != "done" {
		t.Fatalf("status = %q error = %q", res.Status, res.Error)
	}
	var paths []string
	if err := json.Unmarshal(res.Result, &paths); err != nil {
		t.Fatal(err)
	}
	if len(paths) == 0 || !strings.HasPrefix(paths[0], "/") {
		t.Fatalf("unexpected paths %v", paths)
	}
}

func TestIntegration_Execute(t *testing.T) {
	if os.Getenv("CODEMODE_INTEGRATION") == "" {
		t.Skip("set CODEMODE_INTEGRATION=1 to run")
	}
	backend := startBackend(t)
	url, _ := startWorkerd(t, backend.URL)
	c := codemode.NewClient(url)

	res, err := c.Run(context.Background(), codemode.RunRequest{
		Mode:    codemode.ModeExecute,
		Code:    "async () => (await api.foo.listFoo()).data.items[0].id",
		AuthCtx: &codemode.AuthCtx{APIKey: "test-key"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Status != "done" {
		t.Fatalf("status=%q error=%q", res.Status, res.Error)
	}
	var id string
	if err := json.Unmarshal(res.Result, &id); err != nil {
		t.Fatal(err)
	}
	if id != "f1" {
		t.Fatalf("id=%q", id)
	}
}

func TestIntegration_Execute_AuthError(t *testing.T) {
	if os.Getenv("CODEMODE_INTEGRATION") == "" {
		t.Skip("set CODEMODE_INTEGRATION=1 to run")
	}
	backend := startBackend(t)
	url, _ := startWorkerd(t, backend.URL)
	c := codemode.NewClient(url)

	res, err := c.Run(context.Background(), codemode.RunRequest{
		Mode:    codemode.ModeExecute,
		Code:    "async () => { try { return (await api.foo.listFoo()).data; } catch (e) { return {kind: e.name, status: e.status}; } }",
		AuthCtx: &codemode.AuthCtx{APIKey: "wrong"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Status != "done" {
		t.Fatalf("status=%q error=%q", res.Status, res.Error)
	}
	var got struct {
		Kind   string `json:"kind"`
		Status int    `json:"status"`
	}
	if err := json.Unmarshal(res.Result, &got); err != nil {
		t.Fatal(err)
	}
	if got.Kind != "ApiError" || got.Status != 401 {
		t.Fatalf("unexpected %+v", got)
	}
}
