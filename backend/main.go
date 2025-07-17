package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := setupRoutes()
	port := 3001
	fmt.Printf("Go backend listening at http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
