package main

import (
	"io/fs"
	"log/slog"
	"net/http"
)

func newServer(logger *slog.Logger, wwwroot fs.FS) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", newSPAHandler(wwwroot))

	server := logRequests(logger, mux)

	return server
}
