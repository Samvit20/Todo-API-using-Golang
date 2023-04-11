package main

import (
	"encoding/json"
	"net/http"

	"github.com/Samvit20/Todo-API-using-Golang/structs"
)

func main() {
	mux := http.NewServeMux()	
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method ==  http.MethodGet {
			data := views.Response {
				Code: http.StatusOK,
				Body: "ping",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe("localhost:3000", mux)
}