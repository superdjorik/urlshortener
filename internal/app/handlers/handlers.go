package handlers

import (
	"fmt"
	"github.com/superdjorik/urlshortener/internal/app/config"
	"github.com/superdjorik/urlshortener/internal/app/randomizer"
	"github.com/superdjorik/urlshortener/internal/app/storage"
	"io"
	"net/http"
	"strings"
)

func StoreUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		url, _ := io.ReadAll(r.Body)
		id := randomizer.Randomaizer(8)
		storage.UrlList[id] = string(url)
		resp := config.Location() + id
		fmt.Print(resp)
		//w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(resp))
		//} else if r.Method == http.MethodGet {
		//	GetUrl(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := strings.TrimPrefix(r.URL.Path, "/")
		url := storage.UrlList[id]
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusTemporaryRedirect)
		w.Write([]byte(""))
	} else {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
}
