package main

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"
)

func TestStaticAssets(t *testing.T) {
	files := fstest.MapFS{
		"index.html":    {Data: []byte("<html></html>")},
		"assets/app.js": {Data: []byte("console.log('hi');")},
	}

	tests := []struct {
		path                string
		expectedStatus      int
		expectedContentType string
		expectedBody        string
	}{
		{
			path:                "/index.html",
			expectedStatus:      http.StatusMovedPermanently,
			expectedContentType: "",
			expectedBody:        "",
		},
		{
			path:                "/",
			expectedStatus:      http.StatusOK,
			expectedContentType: "text/html; charset=utf-8",
			expectedBody:        "<html></html>",
		},
		{
			path:                "/assets/app.js",
			expectedStatus:      http.StatusOK,
			expectedContentType: "text/javascript; charset=utf-8",
			expectedBody:        "console.log('hi');",
		},
		{
			path:                "/not/found.html",
			expectedStatus:      http.StatusNotFound,
			expectedContentType: "text/plain; charset=utf-8",
			expectedBody:        "",
		},
	}

	for _, test := range tests {
		t.Run(test.path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, test.path, nil)
			w := httptest.NewRecorder()

			newServer(slog.New(slog.DiscardHandler), files).ServeHTTP(w, req)

			if w.Code != test.expectedStatus {
				t.Errorf("Expected status %d, got %d", test.expectedStatus, w.Code)
			}

			if w.Header().Get("Content-Type") != test.expectedContentType {
				t.Errorf("Expected content type %q, got %q", test.expectedContentType, w.Header().Get("Content-Type"))
			}

			if test.expectedBody != "" && w.Body.String() != test.expectedBody {
				t.Errorf("Expected body %q, got %q", test.expectedBody, w.Body.String())
			}
		})
	}
}
