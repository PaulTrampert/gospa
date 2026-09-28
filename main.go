package main

import (
	"log"
	"net/http"
)

func main() {
	mux := newServer()

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
