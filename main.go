package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/NodeOps-app/createos-mcp/codemode"
	"github.com/NodeOps-app/createos-mcp/config"
	"github.com/NodeOps-app/createos-mcp/pkg/oauth"
	"github.com/mark3labs/mcp-go/server"
	"golang.org/x/time/rate"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	if configPath == nil {
		log.Fatal("config path is required")
	}

	if err := config.LoadConfig(*configPath); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	mcpServer := NewMCPServer()

	// Code Mode: workerd sidecar health probe.
	wd := os.Getenv("WORKERD_URL")
	if wd == "" {
		wd = "http://127.0.0.1:8787"
	}
	cmMonitor := codemode.NewHealthMonitor(codemode.NewClient(wd))
	cmMonitor.Start(context.Background(), 5*time.Second)

	switch config.Cfg.Transport {
	case "stdio":
		if err := server.ServeStdio(mcpServer); err != nil {
			log.Fatal(err)
		}
	case "http":
		httpServer := server.NewStreamableHTTPServer(mcpServer)

		var mcpHandler http.Handler = httpServer

		registerStore := newRateLimiterStore()

		mux := http.NewServeMux()
		mux.Handle("/.well-known/oauth-authorization-server", corsMiddleware(oauthAuthorizationServerHandler()))
		mux.Handle("/.well-known/oauth-protected-resource", corsMiddleware(prmMetadataHandler()))
		mux.Handle("/register", corsMiddleware(registerRateLimitMiddleware(registerStore, registerHandler(config.Cfg))))

		mcpHandler = corsMiddleware(mcpHandler)
		mcpHandler = authMiddleware(*config.Cfg, mcpHandler)
		mux.Handle("/mcp", mcpHandler)

		mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
			if cmMonitor.Ready() {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{"ok":true}`))
				return
			}
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte(`{"ok":false,"reason":"workerd sidecar not ready"}`))
		})

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

type ipLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type rateLimiterStore struct {
	mu       sync.Mutex
	limiters map[string]*ipLimiter
}

func newRateLimiterStore() *rateLimiterStore {
	store := &rateLimiterStore{
		limiters: make(map[string]*ipLimiter),
	}
	// Periodically clean up stale entries
	go func() {
		for range time.Tick(5 * time.Minute) {
			store.mu.Lock()
			for ip, l := range store.limiters {
				if time.Since(l.lastSeen) > 10*time.Minute {
					delete(store.limiters, ip)
				}
			}
			store.mu.Unlock()
		}
	}()
	return store
}

func (s *rateLimiterStore) get(ip string) *rate.Limiter {
	s.mu.Lock()
	defer s.mu.Unlock()
	l, ok := s.limiters[ip]
	if !ok {
		// 5 requests per minute, burst of 5
		l = &ipLimiter{limiter: rate.NewLimiter(rate.Every(time.Minute/5), 5)}
		s.limiters[ip] = l
	}
	l.lastSeen = time.Now()
	return l.limiter
}

func registerRateLimitMiddleware(store *rateLimiterStore, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}
		if !store.get(ip).Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
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
		log.Printf("RemoteAddr: %s", r.RemoteAddr)
		log.Printf("User-Agent: %s", r.UserAgent())
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
			w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Bearer realm="createos", resource_metadata="%s"`, prmURL))
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
			Resource:             config.Cfg.BaseURL,
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
	oauthClient := oauth.NewOAuthClient(cfg.APIBaseUrl, cfg.MCPServerToken)

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

		var registrationRequest ClientRegistrationRequest
		if err := json.Unmarshal(body, &registrationRequest); err != nil {
			http.Error(w, "Failed to unmarshal request body", http.StatusInternalServerError)
			return
		}
		log.Printf("Registration request received for client: %s", registrationRequest.ClientName)

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
		w.WriteHeader(http.StatusCreated)

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
