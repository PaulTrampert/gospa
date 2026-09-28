package main

import (
	"fmt"
	"log"
	"net/http"
)

func newServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request for /hello from %s", r.RemoteAddr)
		fmt.Fprintln(w, "Hello World")
	})
	return mux
}
