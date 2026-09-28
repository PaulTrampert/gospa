package main

import (
	"log"
	"log/slog"
	"net/http"
)

func main() {
	server := newServer(slog.Default())

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}
