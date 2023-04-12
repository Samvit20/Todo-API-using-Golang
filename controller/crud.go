package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Samvit20/Todo-API-using-Golang/model"
	"github.com/Samvit20/Todo-API-using-Golang/views"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			//take some data
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			//save it!
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			if name == "" {
				data, err := model.ReadAll()
				if err != nil {
					w.Write([]byte(err.Error()))
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(data)
			} else {
				data, err := model.ReadByName(name)
				if err != nil {
					w.Write([]byte(err.Error()))
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(data)
			}
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]
			if err := model.DeleteTodo(name); err != nil {
				w.Write([]byte("Some error occurred while deleting!"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:"status"`
			}{"Item deleted!"})
		}
	}
}
