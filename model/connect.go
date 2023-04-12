package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "samvit:samvit20@tcp(localhost:3306)/todo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database!")
	con = db
	return db
}
