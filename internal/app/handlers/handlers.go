package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/superdjorik/urlshortener/internal/app/config"
	"github.com/superdjorik/urlshortener/internal/app/randomizer"
	"github.com/superdjorik/urlshortener/internal/app/storage"
	"io"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	appConfig := config.ParseFlags()
	urlPrefix := appConfig.Location
	if len(urlPrefix) > 1 {
		lastChar := string(urlPrefix[len(urlPrefix)-1])
		if lastChar != "/" {
			urlPrefix += "/"
		}
	}
	if urlPrefix == "" {
		urlPrefix = "/"
	}
	switch r.Method {
	case http.MethodPost:

	}

	if r.Method != http.MethodPost {
		http.Error(w, "Uncorrect method, only POST", http.StatusBadRequest)
		return
	}

	responseData, err := io.ReadAll(r.Body)
	if err != nil || string(responseData) == "" {
		http.Error(w, "Invalid POST", http.StatusBadRequest)
		return
	}

	incomeURL := string(responseData)
	shortURL := randomizer.Randomaizer(8)
	storage.UrlList[shortURL] = incomeURL
	w.WriteHeader(http.StatusCreated)

	UrlPrefix := appConfig.Location
	if len(UrlPrefix) > 1 {
		lastChar := string(UrlPrefix[len(UrlPrefix)-1])
		if lastChar != "/" {
			UrlPrefix += "/"
		}
	}
	if UrlPrefix == "" {
		UrlPrefix = "/"
	}

	_, errWrite := w.Write([]byte("http://" + r.Host + UrlPrefix + shortURL))
	if errWrite != nil {
		panic(errWrite)
	}
}

func ShortHandler(w http.ResponseWriter, r *http.Request) {
	d := chi.URLParam(r, "id")

	if full, ok := storage.UrlList[d]; ok {
		w.Header().Add("Location", full)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	http.Error(w, "Not found!", http.StatusBadRequest)
}
