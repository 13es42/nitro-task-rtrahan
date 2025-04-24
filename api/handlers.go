package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	// "strings"
	util "nitro-task-rtrahan/helpers"
)

// Create a singleton cache instance
var factorialCache = util.NewFactorialCache()

// HelloHandler handles GET /hello
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{"message": "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ReverseRequest is the expected JSON for the reverse endpoint
type ReverseRequest struct {
	Text string `json:"text"`
}

// ReverseHandler handles POST /reverse
func ReverseHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ReverseRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Reverse the string
	runes := []rune(req.Text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversed := string(runes)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"reversed": reversed})
}

// FactorialHandler handles GET /factorial/{n}
func FactorialHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the parameter using regexp
	matches := factorialRegex.FindStringSubmatch(r.URL.Path)
	log.Println("MATCHES", matches)
	if len(matches) < 2 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid URL format"})
		return
	}

	nStr := matches[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Input must be a valid integer"})
		return
	}

	if n < 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Input must be a non-negative integer"})
		return
	}

	// Task 1 implementation
	// result := factorial(n)

	// Task 3 implementation
	result := factorialCache.CalculateFactorial(n)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"factorial": result.String()})
}
