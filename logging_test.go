package main

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestLogRequests(t *testing.T) {
	tests := []struct {
		name        string
		verb        string
		path        string
		remoteAddr  string
		expectedLog logMessage
	}{
		{
			name:       "GET /hello",
			verb:       "GET",
			path:       "/hello",
			remoteAddr: "127.0.0.1:1234",
			expectedLog: logMessage{
				Level:   slog.LevelInfo,
				Message: "Request received",
				Values: []any{
					slog.String("method", "GET"),
					slog.String("path", "/hello"),
					slog.String("remote", "127.0.0.1:1234"),
				},
			},
		},
		{
			name:       "POST /foo",
			verb:       "POST",
			path:       "/foo",
			remoteAddr: "5.5.5.5",
			expectedLog: logMessage{
				Level:   slog.LevelInfo,
				Message: "Request received",
				Values: []any{
					slog.String("method", "POST"),
					slog.String("path", "/foo"),
					slog.String("remote", "5.5.5.5"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextCalled := false
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { nextCalled = true })
			logger := testLogger{}

			handler := logRequests(&logger, next)

			req, err := http.NewRequest(tt.verb, tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}
			req.RemoteAddr = tt.remoteAddr
			w := httptest.NewRecorder()

			handler.ServeHTTP(w, req)

			if !nextCalled {
				t.Error("expected next handler to be called")
			}

			if len(logger.logs) != 1 {
				t.Fatalf("expected exactly 1 log message, got %d: %+v", len(logger.logs), logger.logs)
			}

			actualLog := logger.logs[0]

			if !reflect.DeepEqual(actualLog, tt.expectedLog) {
				t.Errorf("expected message to be %+v, got %+v", tt.expectedLog, actualLog)
			}
		})
	}
}
