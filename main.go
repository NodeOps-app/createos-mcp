package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Command-line flags
	transport := flag.String("transport", "stdio", "Transport mode: stdio or http")
	port := flag.String("port", "3000", "Port for HTTP transport (default: 3000)")
	flag.Parse()

	mcpServer := NewMCPServer()

	switch *transport {
	case "stdio":
		if err := server.ServeStdio(mcpServer); err != nil {
			log.Fatal(err)
		}
	case "http":
		httpServer := server.NewStreamableHTTPServer(mcpServer)

		var mcpHandler http.Handler = httpServer

		mcpHandler = corsMiddleware(mcpHandler)

		mux := http.NewServeMux()
		mux.Handle("/mcp", mcpHandler)

		addr := fmt.Sprintf(":%s", *port)
		fmt.Printf("Starting MCP server on http://0.0.0.0%s/mcp\n", addr)
		if err := http.ListenAndServe(addr, mux); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Invalid transport mode: %s. Use 'stdio' or 'http'", *transport)
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
