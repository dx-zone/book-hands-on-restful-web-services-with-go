// Hands-On Restful Web Services with Go: Building Scalable Web Services and RESTful APIs with Go
// UUID Generator: A simple HTTP server that generates random UUIDs on request
package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
)

// UUID is a custom multiplexer that implements the http.Handler interface
type UUID struct {
}

// main function sets up the HTTP server and listens for requests
func main() {
	mux := &UUID{}
	port := "8080"
	portStr := fmt.Sprintf(":%s", port)
	fmt.Println("Starting server on", portStr)
	err := http.ListenAndServe(portStr, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// ServeHTTP handles incoming HTTP requests and generates a random UUID
func (p *UUID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 1. Add CORS headers to every response
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// 2. Handle "Preflight" OPTIONS requests
	// Browsers send an OPTIONS request before the actual GET request to check permissions
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// 3. Handle GET requests to the root path (endpoint for generating UUIDs)
	if r.URL.Path == "/" {
		giveRandomUUID(w, r)
		return
	}
	http.NotFound(w, r)
}

// giveRandomUUID generates a random UUID and writes it to the response
func giveRandomUUID(w http.ResponseWriter, r *http.Request) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		log.Println("Error generating random bytes:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%x", b))
}
