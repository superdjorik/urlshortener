package main

import (
	"fmt"
	"github.com/superdjorik/urlshortener/internal/app/config"
	"github.com/superdjorik/urlshortener/internal/app/handlers"
	"net/http"
)

func main() {
	config.ParseFlags()
	fmt.Println("Server runs on: " + config.Host())
	fmt.Println("Location: " + config.Location())

	mux := http.NewServeMux()
	mux.HandleFunc("POST /", handlers.StoreUrl)
	mux.HandleFunc("GET /{id}", handlers.GetUrl)
	http.ListenAndServe(config.Host(), mux)
}
