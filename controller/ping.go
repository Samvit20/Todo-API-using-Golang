package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Samvit20/Todo-API-using-Golang/views"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method ==  http.MethodGet {
			data := views.Response {
				Code: http.StatusOK,
				Body: "ping",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}