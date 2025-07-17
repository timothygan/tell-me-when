package main

import (
	"net/http"
)

func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not Implemented"))
}

func setupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Query endpoints (all return 501 Not Implemented for now)
	mux.HandleFunc("/api/query", notImplementedHandler)
	mux.HandleFunc("/api/query/", notImplementedHandler) // for /api/query/{id}/...
	mux.HandleFunc("/api/queries", notImplementedHandler)

	// User endpoint
	mux.HandleFunc("/api/register", notImplementedHandler)

	return mux
}
