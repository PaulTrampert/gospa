// Package spa provides a handler for serving a SPA application.
package spa

import (
	"errors"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

// NewHandler returns a handler that serves static files according to SPA rules:
//   - Any request that matches a file in wwwroot serves that file
//   - Any request that matches a directory containing an index.html serves that index.html
//   - Any request that requests a file that doesn't exist will return 404
//   - Any other request will serve the wwwroot/index.html
//
// A "file" request is a path that has a file extension (e.g. '.js', '.html', etc.)
func NewHandler(wwwroot fs.FS) http.Handler {
	fileHandler := http.FileServerFS(wwwroot)

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		reqPath := strings.Trim(path.Clean(req.URL.Path), "/")
		if reqPath == "" {
			reqPath = "."
		}
		if shouldFallback(wwwroot, reqPath) {
			http.ServeFileFS(w, req, wwwroot, "index.html")
			return
		}

		fileHandler.ServeHTTP(w, req)
	})
}

// shouldFallback returns a boolean indicating whether the SPA fallback index.html should be served.
func shouldFallback(wwwroot fs.FS, cleanPath string) bool {
	hasExt := path.Ext(cleanPath) != ""
	fileInfo, err := fs.Stat(wwwroot, cleanPath)
	isDir := err == nil && fileInfo.IsDir()
	switch {
	case errors.Is(err, fs.ErrNotExist) && !hasExt:
		return true
	case isDir:
		_, dirIndexErr := fs.Stat(wwwroot, path.Join(cleanPath, "index.html"))
		return dirIndexErr != nil
	default:
		return false
	}
}
