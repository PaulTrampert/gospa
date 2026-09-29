package spa

import (
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"
)

func TestSpaHandler(t *testing.T) {
	rootIndexContent := "<html></html>"
	appJsContent := "console.log('hello world');"
	docsIndexContent := "<html><body>hello world!</body></html>"
	notFoundContent := "404 page not found\n"

	files := fstest.MapFS{
		"index.html":      {Data: []byte(rootIndexContent)},
		"assets/app.js":   {Data: []byte(appJsContent)},
		"docs/index.html": {Data: []byte(docsIndexContent)},
		"empty":           {Mode: fs.ModeDir},
	}
	tests := []struct {
		path             string
		expectedStatus   int
		expectedContent  string
		expectedLocation string
	}{
		{path: "/", expectedContent: rootIndexContent, expectedStatus: http.StatusOK, expectedLocation: ""},
		{path: "index.html", expectedContent: "", expectedStatus: http.StatusMovedPermanently, expectedLocation: "./"},
		{path: "assets/app.js", expectedContent: appJsContent, expectedStatus: http.StatusOK, expectedLocation: ""},
		{path: "assets/", expectedContent: rootIndexContent, expectedStatus: http.StatusOK, expectedLocation: ""},
		{path: "docs/", expectedContent: docsIndexContent, expectedStatus: http.StatusOK, expectedLocation: ""},
		{path: "docs/index.html", expectedContent: "", expectedStatus: http.StatusMovedPermanently, expectedLocation: "./"},
		{path: "empty", expectedContent: rootIndexContent, expectedStatus: http.StatusOK, expectedLocation: ""},
		{path: "/dashboard/settings", expectedContent: rootIndexContent, expectedStatus: http.StatusOK, expectedLocation: ""},
		{path: "/assets/missing.js", expectedContent: notFoundContent, expectedStatus: http.StatusNotFound, expectedLocation: ""},
	}

	for _, test := range tests {
		t.Run(test.path, func(t *testing.T) {
			req, _ := http.NewRequest("GET", test.path, nil)
			w := httptest.NewRecorder()

			handler := NewHandler(files)
			handler.ServeHTTP(w, req)

			if w.Code != test.expectedStatus {
				t.Errorf("Expected status %d, got %d", test.expectedStatus, w.Code)
			}

			if w.Body.String() != test.expectedContent {
				t.Errorf("Expected content %q, got %q", test.expectedContent, w.Body.String())
			}

			location := w.Header().Get("Location")
			if location != test.expectedLocation {
				t.Errorf("Expected location %q, got %q", test.expectedLocation, location)
			}
		})
	}
}
