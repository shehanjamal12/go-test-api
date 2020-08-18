package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func dbconnect() *sql.DB {

	host := "localhost"
	port := 5432
	user := "postgres"
	password := "123"
	dbname := "item"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Succesfully Connected on port 5432")
	return db
}

var conn *sql.DB = dbconnect()

func main() {
	var apites test
	route := mux.NewRouter()
	api := newstructapi(apites)

	// routes used in postman tomake request or get responses
	route.HandleFunc("/addItem", api.addItem).Methods("POST")
	//wait group used to run api on 2 ports

	http.ListenAndServe(":5000", route)
	fmt.Println("Listening on port 5000")

}
