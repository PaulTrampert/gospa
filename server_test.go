package main

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		verb           string
		expectedStatus int
		expectedBody   string
	}{
		{
			verb:           "GET",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello World\n",
		},
		{
			verb:           "POST",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Method Not Allowed\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.verb, func(t *testing.T) {
			req, err := http.NewRequest(tt.verb, "/hello", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			newServer(slog.New(slog.DiscardHandler)).ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("Expected body %q, got %q", tt.expectedBody, rr.Body.String())
			}
		})
	}
}
