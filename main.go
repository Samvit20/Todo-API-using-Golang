package main

import (
	"encoding/json"
	"net/http"

	// "github.com/Samvit20/Todo-API-using-Golang/tree/master/structs"
)

func main() {
	mux := http.NewServeMux()	
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method ==  http.MethodGet {
			data := structs.Response {
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe("localhost:3000", mux)
}