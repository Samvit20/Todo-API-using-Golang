package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Samvit20/Todo-API-using-Golang/controller"
	"github.com/Samvit20/Todo-API-using-Golang/model"
	_ "github.com/go-sql-driver/mysql" //mysql driver
)

func main() {
	mux := controller.Register()	
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}