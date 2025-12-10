package mcputils

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

var (
	// BaseURL is the base URL for the backend API
	BaseURL = getEnvOrDefault("API_BASE_URL", "https://autogen-v2-api.nodeops.network")
	// AuthToken is the authentication token for API requests
	AuthToken = os.Getenv("AUTH_TOKEN")
)

// getEnvOrDefault returns the environment variable value or a default
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Client returns a configured Resty client with base URL and auth token
func Client() *resty.Client {
	client := resty.New().
		SetBaseURL(BaseURL).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	return client
}

// Get makes a GET request
func Get(path string, queryParams map[string]string, apiKey string) (*resty.Response, error) {
	req := Client().R().SetHeader("X-Api-Key", apiKey)

	for key, value := range queryParams {
		req.SetQueryParam(key, value)
	}

	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("GET request failed: %w", err)
	}

	if resp.IsError() {
		return resp, fmt.Errorf("API error (status %d): %s", resp.StatusCode(), string(resp.Body()))
	}

	return resp, nil
}

// Post makes a POST request
func Post(path string, body interface{}, apiKey string) (*resty.Response, error) {
	resp, err := Client().R().
		SetBody(body).
		SetHeader("X-Api-Key", apiKey).
		Post(path)
	if err != nil {
		return nil, fmt.Errorf("POST request failed: %w", err)
	}

	if resp.IsError() {
		return resp, fmt.Errorf("API error (status %d): %s", resp.StatusCode(), string(resp.Body()))
	}

	return resp, nil
}

// Put makes a PUT request
func Put(path string, body interface{}, apiKey string) (*resty.Response, error) {
	resp, err := Client().R().
		SetHeader("X-Api-Key", apiKey).
		SetBody(body).
		Put(path)
	if err != nil {
		return nil, fmt.Errorf("PUT request failed: %w", err)
	}

	if resp.IsError() {
		return resp, fmt.Errorf("API error (status %d): %s", resp.StatusCode(), string(resp.Body()))
	}

	return resp, nil
}

// Patch makes a PATCH request
func Patch(path string, body interface{}, apiKey string) (*resty.Response, error) {
	resp, err := Client().R().
		SetBody(body).
		SetHeader("X-Api-Key", apiKey).
		Patch(path)
	if err != nil {
		return nil, fmt.Errorf("PATCH request failed: %w", err)
	}

	if resp.IsError() {
		return resp, fmt.Errorf("API error (status %d): %s", resp.StatusCode(), string(resp.Body()))
	}

	return resp, nil
}

// Delete makes a DELETE request
func Delete(path string, apiKey string) (*resty.Response, error) {
	resp, err := Client().R().
		SetHeader("X-Api-Key", apiKey).
		Delete(path)
	if err != nil {
		return nil, fmt.Errorf("DELETE request failed: %w", err)
	}

	if resp.IsError() {
		return resp, fmt.Errorf("API error (status %d): %s", resp.StatusCode(), string(resp.Body()))
	}

	return resp, nil
}
