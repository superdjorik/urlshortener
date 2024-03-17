package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/superdjorik/urlshortener/internal/app/randomizer"
	"github.com/superdjorik/urlshortener/internal/app/storage"
	"io"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Uncorrect method, only POST", http.StatusBadRequest)
		return
	}

	responseData, err := io.ReadAll(r.Body)
	if err != nil || string(responseData) == "" {
		http.Error(w, "Invalid POST", http.StatusBadRequest)
		return
	}

	incomeUrl := string(responseData)
	shortUrl := randomizer.Randomaizer(8)
	storage.UrlList[shortUrl] = incomeUrl
	w.WriteHeader(http.StatusCreated)
	_, errWrite := w.Write([]byte("http://" + r.Host + "/" + shortUrl))
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
