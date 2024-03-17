package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandler(t *testing.T) {
	type want struct {
		statusCode int
	}
	tests := []struct {
		name    string
		method  string
		request string
		body    string
		want    want
	}{
		{
			name:    "Post",
			request: "/",
			method:  http.MethodPost,
			body:    "https://praktikum.yandex.ru",
			want:    want{http.StatusCreated},
		},
		{
			name:    "GET to POST",
			request: "/",
			method:  http.MethodGet,
			want:    want{http.StatusBadRequest},
		},
		{
			name:    "GET to unknown id",
			request: "/X3",
			method:  http.MethodGet,
			want:    want{http.StatusBadRequest},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			body := strings.NewReader(tt.body)
			request := httptest.NewRequest(tt.method, tt.request, body)
			w := httptest.NewRecorder()
			MainHandler(w, request)
			res := w.Result()
			assert.Equal(t, res.StatusCode, tt.want.statusCode)
			defer res.Body.Close()
		})
	}
}
