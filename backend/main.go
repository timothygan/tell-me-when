package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World from the Go backend!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := ":3001"
	fmt.Printf("Go backend listening at http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
