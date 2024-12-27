package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handle the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Go API!")
	})

	// Handle the /api/hello path
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go API!")
	})

	// Start the server on port 8086
	fmt.Println("Starting server on :8086...")
	if err := http.ListenAndServe(":8086", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
