package api

import (
	"log"
	"net/http"
	"regexp"
)

// Regular expression to extract the parameter from factorial path
// the regex was not capturing negative numbers
// var factorialRegex = regexp.MustCompile(`^/factorial/(\d+)$`)

// this regex captures negative numbers so i can demonstrate the error handling for negative numbers.
var factorialRegex = regexp.MustCompile(`^/factorial/(-?\d+)$`)

// SetupRoutes configures the API routes
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/hello", HelloHandler)
	mux.HandleFunc("/reverse", ReverseHandler)
	mux.HandleFunc("/factorial/", FactorialHandler) // Note the trailing slash

	return mux
}

// StartServer initializes and starts the HTTP server
func StartServer(port string) {
	router := SetupRoutes()
	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
