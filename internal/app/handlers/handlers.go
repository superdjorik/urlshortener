package handlers

import (
	"github.com/superdjorik/urlshortener/internal/app/randomizer"
	"github.com/superdjorik/urlshortener/internal/app/storage"
	"io"
	"net/http"
	"strings"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	d := strings.TrimPrefix(r.URL.Path, "/")

	if d == "" {
		if r.Method != http.MethodPost {
			http.Error(w, "Incorrect method, only POST", http.StatusBadRequest)
			return
		}

		responseData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid POST", http.StatusBadRequest)
			return
		}

		incomeUrl := string(responseData)
		shortUrl := randomizer.Randomaizer(8)
		storage.UrlList[shortUrl] = incomeUrl
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("http://" + r.Host + "/" + shortUrl))
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Incorrect method, only GET", http.StatusBadRequest)
		return
	}

	if full, ok := storage.UrlList[d]; ok {
		w.Header().Add("Location", full)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	http.Error(w, "Not found!", http.StatusBadRequest)
}
