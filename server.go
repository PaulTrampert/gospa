package main

import (
	"io/fs"
	"log/slog"
	"net/http"

	"gospa/internal/spa"
)

func newServer(logger *slog.Logger, wwwroot fs.FS) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", spa.NewHandler(wwwroot))

	server := logRequests(logger, mux)

	return server
}
