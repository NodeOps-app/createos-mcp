package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/NodeOps-app/autogen-backend-v2-mcp/config"
	"github.com/NodeOps-app/autogen-backend-v2-mcp/pkg/oauth"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to the config file")
	toolsetFlag := flag.String("toolset", "", "Toolset to expose: core or all")
	flag.Parse()

	if configPath == nil {
		log.Fatal("config path is required")
	}

	if err := config.LoadConfig(*configPath); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	toolset := "core"
	if config.Cfg.Toolset != "" {
		toolset = config.Cfg.Toolset
	}
	if envToolset := os.Getenv("MCP_TOOLSET"); envToolset != "" {
		toolset = envToolset
	}
	if toolsetFlag != nil && *toolsetFlag != "" {
		toolset = *toolsetFlag
	}

	mcpServer := NewMCPServer(toolset)

	switch config.Cfg.Transport {
	case "stdio":
		if err := server.ServeStdio(mcpServer); err != nil {
			log.Fatal(err)
		}
	case "http":
		httpServer := server.NewStreamableHTTPServer(mcpServer)

		var mcpHandler http.Handler = httpServer

		mux := http.NewServeMux()
		mux.Handle("/.well-known/oauth-authorization-server", corsMiddleware(oauthAuthorizationServerHandler()))
		mux.Handle("/.well-known/oauth-protected-resource", corsMiddleware(prmMetadataHandler()))
		mux.Handle("/register", corsMiddleware(registerHandler(config.Cfg)))

		mcpHandler = corsMiddleware(mcpHandler)
		mcpHandler = authMiddleware(*config.Cfg, mcpHandler)
		mux.Handle("/mcp", mcpHandler)

		// Wrap the entire mux with logging middleware to log ALL requests
		rootHandler := loggingMiddleware(mux)

		addr := fmt.Sprintf(":%d", config.Cfg.Port)
		fmt.Printf("Starting MCP server on http://0.0.0.0%s/mcp\n", addr)
		if err := http.ListenAndServe(addr, rootHandler); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Invalid transport mode: %s. Use 'stdio' or 'http'", config.Cfg.Transport)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Api-Key, mcp-session-id")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Continue with the next handler
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("=== Request ===")
		log.Printf("Method: %s", r.Method)
		log.Printf("Path: %s", r.URL.Path)
		log.Printf("Query: %s", r.URL.RawQuery)
		log.Printf("RemoteAddr: %s", r.RemoteAddr)
		log.Printf("User-Agent: %s", r.UserAgent())
		if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
			// Read body for logging (we'll need to restore it)
			body, err := io.ReadAll(r.Body)
			if err == nil {
				log.Printf("Body: %s", string(body))
				// Restore body for the next handler
				r.Body = io.NopCloser(strings.NewReader(string(body)))
			}
		}
		log.Printf("Headers: %v", r.Header)
		log.Printf("===============")
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(cfg config.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-Api-Key")
		bearerToken := r.Header.Get("Authorization")
		if bearerToken != "" {
			parts := strings.Split(bearerToken, " ")
			if len(parts) < 2 || parts[1] == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		// either api key or bearer token is required
		// If both are missing, send 401 with WWW-Authenticate header pointing to PRM endpoint
		if apiKey == "" && bearerToken == "" {
			prmURL := fmt.Sprintf("%s/.well-known/oauth-protected-resource", cfg.BaseURL)
			w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Bearer realm="mcp", resource_metadata="%s"`, prmURL))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// PRMMetadata represents the Protected Resource Metadata document (RFC 9728)
type PRMMetadata struct {
	Resource             string   `json:"resource"`
	AuthorizationServers []string `json:"authorization_servers"`
	ScopesSupported      []string `json:"scopes_supported"`
}

func prmMetadataHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("PRM endpoint accessed: %s %s", r.Method, r.URL.Path)

		if r.Method != http.MethodGet && r.Method != http.MethodOptions {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		prm := PRMMetadata{
			Resource:             fmt.Sprintf("%s/mcp", config.Cfg.BaseURL),
			AuthorizationServers: []string{config.Cfg.BaseURL},
			ScopesSupported:      config.Cfg.SupportedScopes,
		}

		// Set content type and return JSON
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(prm); err != nil {
			log.Printf("Error encoding PRM metadata: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
}

// OAuthAuthorizationServerMetadata represents OAuth 2.0 Authorization Server Metadata (RFC 8414)
type OAuthAuthorizationServerMetadata struct {
	Issuer                            string   `json:"issuer"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	RevocationEndpoint                string   `json:"revocation_endpoint"`
	SupportedScopes                   []string `json:"scopes_supported"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	GrantTypesSupported               []string `json:"grant_types_supported"`
	RegistrationEndpoint              string   `json:"registration_endpoint"`
	CodeChallengeMethodsSupported     []string `json:"code_challenge_methods_supported"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
}

// oauthAuthorizationServerHandler returns a handler for OAuth Authorization Server Metadata endpoint
func oauthAuthorizationServerHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("OAuth Authorization Server endpoint accessed: %s %s", r.Method, r.URL.Path)

		// Only allow GET requests
		if r.Method != http.MethodGet && r.Method != http.MethodOptions {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// Create OAuth Authorization Server Metadata
		metadata := OAuthAuthorizationServerMetadata{
			Issuer:                            config.Cfg.AuthorizationServerUrl,
			AuthorizationEndpoint:             config.Cfg.AuthorizationEndpoint,
			TokenEndpoint:                     config.Cfg.TokenEndpoint,
			ResponseTypesSupported:            config.Cfg.ResponseTypesSupported,
			GrantTypesSupported:               config.Cfg.GrantTypesSupported,
			CodeChallengeMethodsSupported:     config.Cfg.CodeChallengeMethodsSupported,
			RevocationEndpoint:                config.Cfg.RevokeEndpoint,
			SupportedScopes:                   config.Cfg.SupportedScopes,
			RegistrationEndpoint:              fmt.Sprintf("%s/register", config.Cfg.BaseURL),
			TokenEndpointAuthMethodsSupported: config.Cfg.TokenEndpointAuthMethodsSupported,
		}

		// Set content type and return JSON
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(metadata); err != nil {
			log.Printf("Error encoding OAuth Authorization Server metadata: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		log.Printf("OAuth Authorization Server metadata served successfully")
	})
}

// ClientRegistrationRequest represents Dynamic Client Registration request (RFC 7591)
type ClientRegistrationRequest struct {
	ClientName    string   `json:"client_name,omitempty"`
	RedirectURIs  []string `json:"redirect_uris,omitempty"`
	GrantTypes    []string `json:"grant_types,omitempty"`
	ResponseTypes []string `json:"response_types,omitempty"`
	Scope         string   `json:"scope,omitempty"`
}

// ClientRegistrationResponse represents Dynamic Client Registration response (RFC 7591)
type ClientRegistrationResponse struct {
	ClientID              string   `json:"client_id"`
	ClientSecret          string   `json:"client_secret,omitempty"`
	ClientIDIssuedAt      int64    `json:"client_id_issued_at,omitempty"`
	ClientSecretExpiresAt int64    `json:"client_secret_expires_at,omitempty"`
	RedirectURIs          []string `json:"redirect_uris,omitempty"`
	GrantTypes            []string `json:"grant_types,omitempty"`
	ResponseTypes         []string `json:"response_types,omitempty"`
	ClientName            string   `json:"client_name,omitempty"`
}

func registerHandler(cfg *config.Config) http.Handler {
	oauthClient := oauth.NewOAuthClient(cfg.APIBaseUrl, "")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("OAuth Authorization Server endpoint accessed: %s %s", r.Method, r.URL.Path)

		// Only allow GET requests
		if r.Method != http.MethodPost && r.Method != http.MethodOptions {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// print request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		fmt.Printf("Request body ---->>>>>>: %s", string(body))

		var registrationRequest ClientRegistrationRequest
		if err := json.Unmarshal(body, &registrationRequest); err != nil {
			http.Error(w, "Failed to unmarshal request body", http.StatusInternalServerError)
			return
		}
		log.Printf("Registration request: %+v", registrationRequest)

		if len(registrationRequest.Scope) == 0 {
			registrationRequest.Scope = strings.Join(cfg.SupportedScopes, " ")
		}

		registrationResponse, err := oauthClient.CreateDCRClientRegistration(oauth.DCRClientRegistrationRequest{
			ClientName:    registrationRequest.ClientName,
			RedirectURIs:  registrationRequest.RedirectURIs,
			GrantTypes:    registrationRequest.GrantTypes,
			ResponseTypes: registrationRequest.ResponseTypes,
			Scope:         registrationRequest.Scope,
		})
		if err != nil {
			log.Printf("Error creating DCR client registration: %v", err)
			http.Error(w, "Failed to create DCR client registration", http.StatusInternalServerError)
			return
		}

		// Set headers BEFORE writing status code
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Write status code
		w.WriteHeader(http.StatusCreated) // 201 Created for registration

		// Encode JSON response directly to response writer
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "") // No indentation to avoid extra formatting
		if err := encoder.Encode(registrationResponse); err != nil {
			log.Printf("Error encoding client registration response: %v", err)
			// Don't write error here as we already wrote status code
			return
		}
	})
}
