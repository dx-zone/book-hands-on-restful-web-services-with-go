// Hands-On Restful Web Services with Go: Building Scalable Web Services and RESTful APIs with Go
// HealthCheck API returns date time to client
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/health", HealthCheck)
	port := 8080
	portStr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server. Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(portStr, nil))
}

// HealthCheck API returns date time to client
func HealthCheck(w http.ResponseWriter, req *http.Request) {
	// Allow any origin (common for development)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Handle preflight OPTIONS requests if Swagger sends them
	if req.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	currentTime := time.Now()
	io.WriteString(w, currentTime.String())
	log.Printf("Health check requested from %s at %s\n", req.RemoteAddr, currentTime.String())
}
