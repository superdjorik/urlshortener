package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/superdjorik/urlshortener/internal/app/config"
	"github.com/superdjorik/urlshortener/internal/app/handlers"
	"net/http"
)

func main() {
	appConfig := config.ParseFlags()
	println(appConfig)
	defaultRoute := "/"
	fmt.Println("Runs on: " + appConfig.Host)
	fmt.Println("Prefix: " + defaultRoute)

	r := chi.NewRouter()
	r.Route(defaultRoute, func(r chi.Router) {
		r.Post("/", handlers.MainHandler)
		r.Get("/{id}", handlers.ShortHandler)
	})
	http.ListenAndServe(appConfig.Host, r)
}
