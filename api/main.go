package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	goapi"api/goapi"
)

func Dbconnect() *sql.DB {

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

var conn *sql.DB = Dbconnect()

func main() {
	var apites goapi.Test
	route := mux.NewRouter()
	api := goapi.Newstructapi(apites)
	//datbase := Newdbstrc(databasinter)

	// routes used in postman tomake request or get responses
	//route.HandleFunc("/", test12).Methods("GET")
	route.HandleFunc("/addItem", api.AddItem).Methods("POST")
	route.HandleFunc("/viewAllItem", api.ViewAllItem).Methods("GET")
	route.HandleFunc("/DelItem", api.DelItem).Methods("DELETE")

	//wait group used to run api on 2 ports

	http.ListenAndServe(":5000", route)

}
