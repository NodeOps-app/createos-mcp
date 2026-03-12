package oauth

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type OAuthClient struct {
	Client *resty.Client
}

func NewOAuthClient(baseURL string, token string) OAuthClient {
	client := resty.New().SetBaseURL(baseURL).SetAuthToken(token)
	return OAuthClient{
		Client: client,
	}
}

type DCRClientRegistrationRequest struct {
	ClientName    string   `json:"client_name,omitempty"`
	RedirectURIs  []string `json:"redirect_uris,omitempty"`
	GrantTypes    []string `json:"grant_types,omitempty"`
	ResponseTypes []string `json:"response_types,omitempty"`
	Scope         string   `json:"scope,omitempty"`
}

type DCRClientRegistrationResponse struct {
	ClientID              string   `json:"client_id"`
	ClientIDIssuedAt      int64    `json:"client_id_issued_at,omitempty"`
	ClientSecretExpiresAt int64    `json:"client_secret_expires_at,omitempty"`
	RedirectURIs          []string `json:"redirect_uris,omitempty"`
	GrantTypes            []string `json:"grant_types,omitempty"`
	ResponseTypes         []string `json:"response_types,omitempty"`
	ClientName            string   `json:"client_name,omitempty"`
}

type ResponseDataWrapper[T any] struct {
	Data T `json:"data"`
}

func (o OAuthClient) CreateDCRClientRegistration(request DCRClientRegistrationRequest) (DCRClientRegistrationResponse, error) {
	var response ResponseDataWrapper[DCRClientRegistrationResponse]

	resp, err := o.Client.R().
		SetBody(request).
		SetResult(&response).
		Post("/v1/-/oauth2/dcr")
	if err != nil {
		return response.Data, err
	}

	if resp.StatusCode() != http.StatusCreated {
		return response.Data, fmt.Errorf("failed to create DCR client registration (status %d)", resp.StatusCode())
	}

	return response.Data, nil
}
