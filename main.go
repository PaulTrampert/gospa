package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	server := newServer(slog.Default(), os.DirFS("./wwwroot"))

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}
