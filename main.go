package main

import (
	"net/http"

	"github.com/Samvit20/Todo-API-using-Golang/controller"
	"github.com/Samvit20/Todo-API-using-Golang/model"
	_ "github.com/go-sql-driver/mysql" //mysql driver
)

func main() {
	mux := controller.Register()	
	db := model.Connect()
	defer db.Close()
	http.ListenAndServe("localhost:3000", mux)
}