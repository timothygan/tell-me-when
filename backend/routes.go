package main

import (
	"net/http"
)

func setupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Query endpoints (all return 501 Not Implemented for now)
	mux.HandleFunc("/api/query", http.NotImplemented)
	mux.HandleFunc("/api/query/", http.NotImplemented) // for /api/query/{id}/...
	mux.HandleFunc("/api/queries", http.NotImplemented)

	// User endpoint
	mux.HandleFunc("/api/register", http.NotImplemented)

	return mux
}
