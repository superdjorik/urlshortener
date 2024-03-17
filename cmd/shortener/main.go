package main

import (
	"fmt"
	"github.com/superdjorik/urlshortener/internal/app/config"
	"github.com/superdjorik/urlshortener/internal/app/handlers"
	"net/http"
)

func main() {
	appConfig := config.ParseFlags()
	fmt.Println("Server runs on: " + appConfig.Host)
	fmt.Println("Location: " + appConfig.Location)

	mux := http.NewServeMux()
	mux.HandleFunc(appConfig.Location, handlers.MainHandler)
	http.ListenAndServe(appConfig.Host, mux)
}
