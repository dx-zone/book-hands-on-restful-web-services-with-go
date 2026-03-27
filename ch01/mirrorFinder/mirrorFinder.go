// Hands-On Restful Web Services with Go: Building Scalable Web Services and RESTful APIs with Go
// mirrorFinder is a simple Go application that finds the fastest mirror from a list of URLs by measuring the latency of each mirror. It exposes an HTTP endpoint that returns the fastest mirror and its latency in JSON format.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mirrorfinder/mirrors"
	"net/http"
	"time"
)

type response struct {
	FastestURL string        `json:"fastest_url"`
	Latency    time.Duration `json:"latency"`
}

func main() {
	http.HandleFunc("/fastest-mirror", func(w http.ResponseWriter, r *http.Request) {
		// Allow Swagger UI to access this endpoint
		w.Header().Set("Access-Control-Allow-Origin", "*")

		response := findFastest(mirrors.MirrorList)
		respJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	})

	// Start the server
	port := 8081
	startServer(port)
}

func findFastest(urls []string) response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)

	for _, url := range urls {
		mirrorURL := url
		go func() {
			log.Println("Started probing: ", mirrorURL)
			start := time.Now()
			_, err := http.Get(mirrorURL + "/README")
			latency := time.Now().Sub(start) / time.Millisecond
			if err == nil {
				urlChan <- mirrorURL
				latencyChan <- latency
			}
			log.Printf("Got the best mirror: %s with latency: %s", mirrorURL, latency)
		}()
	}
	return response{<-urlChan, <-latencyChan}
}

func startServer(port int) {
	portStr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:           portStr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Starting server on port %d\n", port)
	log.Fatal(server.ListenAndServe())
}
