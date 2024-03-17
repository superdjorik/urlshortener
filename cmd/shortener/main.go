package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/superdjorik/urlshortener/internal/app/handlers"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Post("/", handlers.MainHandler)
		r.Get("/{id}", handlers.ShortHandler)
	})
	http.ListenAndServe(":8080", r)
}
