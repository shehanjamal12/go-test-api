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
	fmt.Println("Succesfully Connected")
	return db
}

var conn *sql.DB = dbconnect()

func main() {
	//var apites test
	var check strucdb
	//var datbase dbinter
	route := mux.NewRouter()
	//api := newstructapi(apites)
	//databas := newdbstrc(datbase)

	// routes used in postman tomake request or get responses
	route.HandleFunc("/addItem", check.addItem).Methods("POST")
	//wait group used to run api on 2 ports

	http.ListenAndServe(":5000", route)

}
