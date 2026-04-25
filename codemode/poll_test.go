package codemode

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Poll_Done(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/poll" {
			t.Fatalf("unexpected path %s", r.URL.Path)
		}
		var body PollRequest
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body.JobID != "j1" {
			t.Fatalf("got jobId %q", body.JobID)
		}
		_, _ = w.Write([]byte(`{"status":"done","result":"x"}`))
	}))
	defer srv.Close()
	c := NewClient(srv.URL)
	got, err := c.Poll(context.Background(), "j1")
	if err != nil {
		t.Fatal(err)
	}
	if got.Status != "done" {
		t.Fatalf("status=%q", got.Status)
	}
}

func TestClient_Poll_JobMissing(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"status":"error","errorKind":"jobMissing","error":"unknown jobId"}`))
	}))
	defer srv.Close()
	c := NewClient(srv.URL)
	got, err := c.Poll(context.Background(), "j1")
	if err != nil {
		t.Fatal(err)
	}
	if got.ErrorKind != "jobMissing" {
		t.Fatalf("errorKind=%q", got.ErrorKind)
	}
}
