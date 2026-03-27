// Hands-On Restful Web Services with Go
// Simple Static File Server
package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Create a new router
	router := httprouter.New()

	// Mapping to method is possible with Http Router
	// Serve files from the "/api" directory, endpoint is "/api/*filepath"
	router.ServeFiles("/api/*filepath", http.Dir("api"))

	// Serve files from the root directory, endpoint is "/*filepath"
	router.ServeFiles("//*filepath", http.Dir("/"))

	// Enable CORS for all routes (optional)
	//handler := cors.Default().Handler(router)

	// Start the server
	port := ":8080"
	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))

	// If you want to use CORS, replace the above line with the following:
	//log.Fatal(http.ListenAndServe(port, handler))
}
