package main

import (
	"log/slog"
	"net/http"
)

type infoLogger interface {
	Info(msg string, keysAndValues ...any)
}

// Logs incoming requests
func logRequests(logger infoLogger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Request received",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("remote", r.RemoteAddr),
		)
		next.ServeHTTP(w, r)
	})
}
