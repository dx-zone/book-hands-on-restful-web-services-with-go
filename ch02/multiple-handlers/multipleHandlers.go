// Hands-On Restful Web Services with Go: Building Scalable Web Services and RESTful APIs with Go
// Adding multiple handlers to the server with http.NewServeMux() and http.HandleFunc() to create endpoints for random float and random integer generation. The server listens on port 8000 and logs the generated values along with the client's remote address.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	newMux := http.NewServeMux()

	// Add a handler for random floats, endpoint: /randomFloat
	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		randomFloat := rand.Float64()
		fmt.Fprintln(w, randomFloat)
		log.Printf("Random float %f generated and sent to client: %s\n", randomFloat, r.RemoteAddr)
	})

	// Add another handler for random integers, endpoint: /randomInt
	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		randomInt := rand.Intn(100)
		fmt.Fprintln(w, randomInt)
		log.Printf("Random integer %d generated and sent to client: %s\n", randomInt, r.RemoteAddr)
	})

	// Start the server
	port := ":8000"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, newMux))
}
