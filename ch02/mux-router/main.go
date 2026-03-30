// Hands-on: Build a RESTful API with Gorilla Mux
// Simple RESTful API using the Gorilla Mux package in Go.
// Path-Based vs Query-Based Matching
// Endpoints: one for path-based matching and another for query-based matching.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Register a new route with a matcher for the URL path and a handler function.
	//Path-based matching with variables and regular expressions
	r.HandleFunc("/path-based/articles/{category}/{id:[0-9]+}", ArticleHandler)

	// Query-based matching with variables and regular expressions
	r.HandleFunc("/query-based/articles", QueryHandler)

	// Reverse mapping: Generate a URL from the route name and variables
	r.HandleFunc("/reverse-mapping/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("articleRoute")
	// Generate a URL for the "articleRoute" with the specified variables
	url, err := r.Get("articleRoute").URL("category", "books", "id", "123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(url.Path) // Output: /reverse-mapping/articles/books/123

	// PathPrefix and StripPrefix
	// Use PathPrefix to match all routes that start with /static/ and serve files from the ./static directory
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/tmp/static"))))

	// strict slash: Redirects requests with a trailing slash to the same path without the slash, and vice versa.
	// For example, if you have a route defined as /articles and a request comes in for /articles/, it will redirect to /articles.
	r.StrictSlash(true)
	r.Path("/articles/").Handler(http.HandlerFunc(ArticleHandler))

	// Match encoded paths: By default, Gorilla Mux does not match encoded paths. You can enable this feature by setting the UseEncodedPath option to true.
	r.UseEncodedPath()
	r.NewRoute().Path("category/id")

	// Bind to a port and pass our router into our http.Server
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Starting server on :8000")
	log.Fatal(srv.ListenAndServe())
}

// ArticleHandler is the handler function for the /path-based/articles/{category}/{id} route
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	// Extract variables from the URL path using mux.Vars
	vars := mux.Vars(r)
	// Write the response with the extracted variables
	w.WriteHeader(http.StatusOK)
	// Use fmt.Fprintf to write the response with the category and id variables
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

// QueryHandler is the handler function for the /query-based/articles route with query parameters
func QueryHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters from the URL using r.URL.Query()
	queryParams := r.URL.Query()
	// Write the response with the extracted query parameters
	w.WriteHeader(http.StatusOK)
	// Use fmt.Fprintf to write the response with the query parameters
	fmt.Fprintf(w, "Got parameter id: %s!\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category: %s!", queryParams["category"][0])
}
