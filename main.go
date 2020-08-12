package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	
	route := mux.NewRouter()

	// routes used in postman tomake request or get responses
	route.HandleFunc("/", test).Methods("GET")
	route.HandleFunc("/addItem", addItem).Methods("POST")
	route.HandleFunc("/viewAllItem", viewAllItem).Methods("GET")
	route.HandleFunc("/viewItem/{id}", viewItem).Methods("GET")
	route.HandleFunc("/viewbyName/{name}", viewbyName).Methods("GET")
	route.HandleFunc("/delItem/{id}", delItem).Methods("DELETE")
	route.HandleFunc("/delItembyName/{name}", delItembyName).Methods("DELETE")
	route.HandleFunc("/updateItembyId/{id}", updateItembyID).Methods("PUT")
	route.HandleFunc("/updateItembyName/{name}", updateItembyName).Methods("PUT")
	//wait group used to run api on 2 ports

	http.ListenAndServe(":5400", route)

}
