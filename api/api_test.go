package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloHandler)
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	if response["message"] != "Hello, World!" {
		t.Errorf("handler returned unexpected body: got %v want %v", response["message"], "Hello, World!")
	}
}

func TestReverseHandler(t *testing.T) {
	tests := []struct {
		name           string
		inputText      string
		expectedOutput string
		statusCode     int
	}{
		{"Normal string", "overwolf", "flowrevo", http.StatusOK},
		{"Empty string", "", "", http.StatusOK},
		{"Palindrome", "radar", "radar", http.StatusOK},
		{"Special characters", "a!b@c#", "#c@b!a", http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(map[string]string{"text": tt.inputText})
			req, err := http.NewRequest("POST", "/reverse", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(ReverseHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.statusCode)
			}

			var response map[string]string
			err = json.Unmarshal(rr.Body.Bytes(), &response)
			if err != nil {
				t.Fatalf("could not unmarshal response: %v", err)
			}

			if response["reversed"] != tt.expectedOutput {
				t.Errorf("handler returned unexpected body: got %v want %v", response["reversed"], tt.expectedOutput)
			}
		})
	}
}

func TestFactorialHandler(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		expectedCode int
		checkBody    func(*testing.T, []byte)
	}{
		{
			name:         "Valid input - 5",
			url:          "/factorial/5",
			expectedCode: http.StatusOK,
			checkBody: func(t *testing.T, body []byte) {
				var response map[string]string // Changed from int to string
				err := json.Unmarshal(body, &response)
				if err != nil {
					t.Fatalf("could not unmarshal response: %v", err)
				}
				if response["factorial"] != "120" { // Compare with string "120" instead of int 120
					t.Errorf("expected factorial 120, got %s", response["factorial"])
				}
			},
		},
		{
			name:         "Zero input",
			url:          "/factorial/0",
			expectedCode: http.StatusOK,
			checkBody: func(t *testing.T, body []byte) {
				var response map[string]string // Changed from int to string
				err := json.Unmarshal(body, &response)
				if err != nil {
					t.Fatalf("could not unmarshal response: %v", err)
				}
				if response["factorial"] != "1" { // Compare with string "1" instead of int 1
					t.Errorf("expected factorial 1, got %s", response["factorial"])
				}
			},
		},
		{
			name:         "Invalid path",
			url:          "/factorial/abc",
			expectedCode: http.StatusBadRequest,
			checkBody: func(t *testing.T, body []byte) {
				var response map[string]string
				err := json.Unmarshal(body, &response)
				if err != nil {
					t.Fatalf("could not unmarshal response: %v", err)
				}
				if response["error"] == "" {
					t.Errorf("expected error message, got empty string")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(FactorialHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedCode)
			}

			tt.checkBody(t, rr.Body.Bytes())
		})
	}
}
