package main

import "log/slog"

// Test logger. Records log messages as received.
type testLogger struct {
	logs []logMessage
}

// Adds an INFO level message to the logger.
func (l *testLogger) Info(message string, args ...any) {
	l.logs = append(l.logs, logMessage{
		Level:   slog.LevelInfo,
		Message: message,
		Values:  args,
	})
}

// Log message struct used to test the arguments passed to a log statement.
type logMessage struct {
	Level   slog.Level
	Message string
	Values  []any
}
