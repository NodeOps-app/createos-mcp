package codemode

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Mode string

const (
	ModeSearch  Mode = "search"
	ModeExecute Mode = "execute"
)

type AuthCtx struct {
	APIKey string `json:"apiKey,omitempty"`
	Bearer string `json:"bearer,omitempty"`
}

type RunRequest struct {
	Mode    Mode     `json:"mode"`
	Code    string   `json:"code"`
	AuthCtx *AuthCtx `json:"authCtx,omitempty"`
}

type RunResult struct {
	Status    string          `json:"status"`
	Result    json.RawMessage `json:"result,omitempty"`
	Logs      json.RawMessage `json:"logs,omitempty"`
	JobID     string          `json:"jobId,omitempty"`
	Error     string          `json:"error,omitempty"`
	ErrorKind string          `json:"errorKind,omitempty"`
	Stack     string          `json:"stack,omitempty"`
}

type Client struct {
	baseURL string
	http    *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		http:    &http.Client{Timeout: 95 * time.Second},
	}
}

func (c *Client) Run(ctx context.Context, req RunRequest) (*RunResult, error) {
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/run", bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("workerd: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusServiceUnavailable {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("workerd unavailable: %s", string(body))
	}
	var out RunResult
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}
	if resp.StatusCode >= 500 && out.Status == "" {
		return nil, errors.New("workerd error: " + resp.Status)
	}
	return &out, nil
}

type HealthStats struct {
	JobsRunning  int `json:"jobsRunning"`
	JobStoreSize int `json:"jobStoreSize"`
}

func (c *Client) HealthStats(ctx context.Context) (*HealthStats, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/health", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("health %d", resp.StatusCode)
	}
	var out HealthStats
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

type PollRequest struct {
	JobID string `json:"jobId"`
}

func (c *Client) Poll(ctx context.Context, jobID string) (*RunResult, error) {
	buf, err := json.Marshal(PollRequest{JobID: jobID})
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/poll", bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("workerd: %w", err)
	}
	defer resp.Body.Close()
	var out RunResult
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) Health(ctx context.Context) error {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/health", nil)
	if err != nil {
		return err
	}
	resp, err := c.http.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("health %d: %s", resp.StatusCode, string(body))
	}
	return nil
}
