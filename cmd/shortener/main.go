package main

import (
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var urlList = make(map[string]string)
var randomChars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var randRange *rand.Rand

// Инициализация рандомайзера
func init() {
	source := rand.NewSource(time.Now().UnixNano())
	randRange = rand.New(source)
}

// Генератор случайной строки
func randomaizer(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = randomChars[randRange.Intn(len(randomChars))]
	}
	return string(b)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	d := strings.TrimPrefix(r.URL.Path, "/")

	if d == "" {
		if r.Method != http.MethodPost {
			http.Error(w, "Uncorrect method, only POST", http.StatusBadRequest)
			return
		}

		responseData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid POST", http.StatusBadRequest)
			return
		}

		incomeUrl := string(responseData)
		shortUrl := randomaizer(8)
		urlList[shortUrl] = incomeUrl
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("http://" + r.Host + "/" + shortUrl))
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Uncorrect method, only GET", http.StatusBadRequest)
		return
	}

	if full, ok := urlList[d]; ok {
		w.Header().Add("Location", full)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	http.Error(w, "Not found!", http.StatusBadRequest)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", mux)
}
