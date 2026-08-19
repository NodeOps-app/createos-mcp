package mcputils

import (
	"fmt"

	"github.com/NodeOps-app/createos-mcp/config"
	"github.com/go-resty/resty/v2"
)

// BaseURL is the base URL for the backend API

// Client returns a configured Resty client with base URL and auth token
func Client() *resty.Client {
	client := resty.New().
		SetBaseURL(config.Cfg.APIBaseUrl).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	return client
}

func SandboxClient() (*resty.Client, error) {
	if config.Cfg.SandboxAPIBaseUrl == "" {
		return nil, fmt.Errorf("sandbox_api_base_url is required for sandbox tools")
	}

	client := resty.New().
		SetBaseURL(config.Cfg.SandboxAPIBaseUrl).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	return client, nil
}

func setAuth(req *resty.Request, authMethod string, authValue string) error {
	switch authMethod {
	case "api-key":
		req.SetHeader("X-Api-Key", authValue)
	case "bearer-token":
		req.SetHeader("X-Access-Token", authValue)
	default:
		return fmt.Errorf("unsupported auth method: %s", authMethod)
	}
	return nil
}

// Get makes a GET request with authentication
func Get(path string, queryParams map[string]string, authMethod string, authValue string) (*resty.Response, error) {
	req := Client().R()

	// Set the appropriate header based on auth method
	switch authMethod {
	case "api-key":
		req.SetHeader("X-Api-Key", authValue)
	case "bearer-token":
		req.SetHeader("X-Access-Token", authValue)
	default:
		return nil, fmt.Errorf("unsupported auth method: %s", authMethod)
	}

	for key, value := range queryParams {
		req.SetQueryParam(key, value)
	}

	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("GET request failed: %w", err)
	}

	if resp.IsError() {
		return resp, fmt.Errorf("API error (status %d)", resp.StatusCode())
	}

	return resp, nil
}

// Post makes a POST request with authentication
func Post(path string, body interface{}, authMethod string, authValue string) (*resty.Response, error) {
	req := Client().R().SetBody(body)

	// Set the appropriate header based on auth method
	switch authMethod {
	case "api-key":
		req.SetHeader("X-Api-Key", authValue)
	case "bearer-token":
		req.SetHeader("X-Access-Token", authValue)
	default:
		return nil, fmt.Errorf("unsupported auth method: %s", authMethod)
	}

	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("POST request failed: %w", err)
	}

	if resp.IsError() {
		return resp, fmt.Errorf("API error (status %d)", resp.StatusCode())
	}

	return resp, nil
}

// Put makes a PUT request with authentication
func Put(path string, body interface{}, authMethod string, authValue string) (*resty.Response, error) {
	req := Client().R().SetBody(body)

	// Set the appropriate header based on auth method
	switch authMethod {
	case "api-key":
		req.SetHeader("X-Api-Key", authValue)
	case "bearer-token":
		req.SetHeader("X-Access-Token", authValue)
	default:
		return nil, fmt.Errorf("unsupported auth method: %s", authMethod)
	}

	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("PUT request failed: %w", err)
	}

	if resp.IsError() {
		return resp, fmt.Errorf("API error (status %d)", resp.StatusCode())
	}

	return resp, nil
}

// Patch makes a PATCH request with authentication
func Patch(path string, body interface{}, authMethod string, authValue string) (*resty.Response, error) {
	req := Client().R().SetBody(body)

	// Set the appropriate header based on auth method
	switch authMethod {
	case "api-key":
		req.SetHeader("X-Api-Key", authValue)
	case "bearer-token":
		req.SetHeader("X-Access-Token", authValue)
	default:
		return nil, fmt.Errorf("unsupported auth method: %s", authMethod)
	}

	resp, err := req.Patch(path)
	if err != nil {
		return nil, fmt.Errorf("PATCH request failed: %w", err)
	}

	if resp.IsError() {
		return resp, fmt.Errorf("API error (status %d)", resp.StatusCode())
	}

	return resp, nil
}

// Delete makes a DELETE request with authentication
func Delete(path string, authMethod string, authValue string) (*resty.Response, error) {
	req := Client().R()

	// Set the appropriate header based on auth method
	switch authMethod {
	case "api-key":
		req.SetHeader("X-Api-Key", authValue)
	case "bearer-token":
		req.SetHeader("X-Access-Token", authValue)
	default:
		return nil, fmt.Errorf("unsupported auth method: %s", authMethod)
	}

	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("DELETE request failed: %w", err)
	}

	if resp.IsError() {
		return resp, fmt.Errorf("API error (status %d)", resp.StatusCode())
	}

	return resp, nil
}

// DeleteWithBody makes a DELETE request with a body and authentication
func DeleteWithBody(path string, body interface{}, authMethod string, authValue string) (*resty.Response, error) {
	req := Client().R().SetBody(body)

	// Set the appropriate header based on auth method
	switch authMethod {
	case "api-key":
		req.SetHeader("X-Api-Key", authValue)
	case "bearer-token":
		req.SetHeader("X-Access-Token", authValue)
	default:
		return nil, fmt.Errorf("unsupported auth method: %s", authMethod)
	}

	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("DELETE request failed: %w", err)
	}

	if resp.IsError() {
		return resp, fmt.Errorf("API error (status %d)", resp.StatusCode())
	}

	return resp, nil
}

// SandboxPost makes a POST request to the sandbox API with authentication.
func SandboxPost(path string, body interface{}, authMethod string, authValue string) (*resty.Response, error) {
	client, err := SandboxClient()
	if err != nil {
		return nil, err
	}
	req := client.R().SetBody(body)
	if err := setAuth(req, authMethod, authValue); err != nil {
		return nil, err
	}

	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("sandbox POST request failed: %w", err)
	}
	if resp.IsError() {
		return resp, fmt.Errorf("sandbox API error (status %d)", resp.StatusCode())
	}

	return resp, nil
}

// SandboxGet makes a GET request to the sandbox API with authentication.
func SandboxGet(path string, queryParams map[string]string, authMethod string, authValue string) (*resty.Response, error) {
	client, err := SandboxClient()
	if err != nil {
		return nil, err
	}
	req := client.R()
	if err := setAuth(req, authMethod, authValue); err != nil {
		return nil, err
	}
	for key, value := range queryParams {
		req.SetQueryParam(key, value)
	}

	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("sandbox GET request failed: %w", err)
	}
	if resp.IsError() {
		return resp, fmt.Errorf("sandbox API error (status %d)", resp.StatusCode())
	}

	return resp, nil
}

// SandboxPatch makes a PATCH request to the sandbox API with authentication.
func SandboxPatch(path string, body interface{}, authMethod string, authValue string) (*resty.Response, error) {
	client, err := SandboxClient()
	if err != nil {
		return nil, err
	}
	req := client.R().SetBody(body)
	if err := setAuth(req, authMethod, authValue); err != nil {
		return nil, err
	}

	resp, err := req.Patch(path)
	if err != nil {
		return nil, fmt.Errorf("sandbox PATCH request failed: %w", err)
	}
	if resp.IsError() {
		return resp, fmt.Errorf("sandbox API error (status %d)", resp.StatusCode())
	}

	return resp, nil
}

// SandboxDelete makes a DELETE request to the sandbox API with authentication.
func SandboxDelete(path string, queryParams map[string]string, authMethod string, authValue string) (*resty.Response, error) {
	client, err := SandboxClient()
	if err != nil {
		return nil, err
	}
	req := client.R()
	if err := setAuth(req, authMethod, authValue); err != nil {
		return nil, err
	}
	for key, value := range queryParams {
		req.SetQueryParam(key, value)
	}

	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("sandbox DELETE request failed: %w", err)
	}
	if resp.IsError() {
		return resp, fmt.Errorf("sandbox API error (status %d)", resp.StatusCode())
	}

	return resp, nil
}
