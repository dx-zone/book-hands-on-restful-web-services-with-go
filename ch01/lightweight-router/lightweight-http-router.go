// Hands-On Restful Web Services with Go
// Lightweight Router Example with julienschmidt/httprouter
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Create a new router
	router := httprouter.New()

	// Register a handler for the GET method and the /api/v1/go-version path
	router.GET("/api/v1/go-version", goVersion)

	// Register a handler for the GET method and the /api/v1/show-file/:name path (:name is a path parameter)
	router.GET("/api/v1/show-file/:name", getFileContent)

	// Register a handler for the GET method and the /api/v1/list-dir path
	router.GET("/api/v1/list-dir", listDir)

	// Start the HTTP server on port 8000 and use the router to handle incoming requests
	port := ":8000"
	log.Printf("Starting server on %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// getCommandOutput executes a command with the given arguments and returns its output as a string
func getCommandOutput(command string, arguments ...string) string {
	out, _ := exec.Command(command, arguments...).Output()
	return string(out)
}

// goVersion is a handler function that executes the "go version" command and writes its output to the HTTP response
func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := getCommandOutput("/usr/local/go/bin/go", "version")
	log.Printf("Client requested Go version: %s", r.RemoteAddr)
	io.WriteString(w, response)
}

// getFileContent is a handler function that executes the "cat" command with the file name provided as a path parameter and writes its output to the HTTP response
func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Printf("Client %s requested file content for %s", r.RemoteAddr, params.ByName("name"))
	fmt.Fprintf(w, getCommandOutput("/bin/cat", params.ByName("name")))
}

// listDir is a handler function that executes the "ls -l" command and writes its output to the HTTP response
func listDir(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := getCommandOutput("/usr/bin/ls", "-l")
	log.Printf("Client requested directory listing: %s", r.RemoteAddr)
	io.WriteString(w, response)
}
