package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func newServer(logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	server := logRequests(logger, mux)

	return server
}
