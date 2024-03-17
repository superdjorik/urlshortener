package main

import (
	"github.com/superdjorik/urlshortener/internal/app/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.MainHandler)
	http.ListenAndServe(":8080", mux)
}
