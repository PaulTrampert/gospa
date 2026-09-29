package main

import (
	"errors"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

func newSPAHandler(wwwroot fs.FS) http.Handler {
	fileHandler := http.FileServerFS(wwwroot)
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		reqPath := strings.Trim(path.Clean(req.URL.Path), "/")
		isFile := len(path.Ext(reqPath)) > 0
		fileInfo, err := fs.Stat(wwwroot, reqPath)
		if errors.Is(err, fs.ErrNotExist) && !isFile {
			http.ServeFileFS(w, req, wwwroot, "index.html")
			return
		}

		if fileInfo.IsDir() {
			_, err := fs.Stat(wwwroot, path.Join(reqPath, "index.html"))
			if errors.Is(err, fs.ErrNotExist) {
				http.ServeFileFS(w, req, wwwroot, "index.html")
				return
			}
		}

		fileHandler.ServeHTTP(w, req)
	})
}
