package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Items struct {
	ItemName string 
	ItemQuantity int
}

var itemList []Items = []Items{}
func main() {
	route := mux.NewRouter()

	route.HandleFunc("/addItem", addItem).Methods("POST")
	route.HandleFunc("/viewItem", viewItem)

	http.ListenAndServe(":5000", route)

}

func addItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Items
	json.NewDecoder(r.Body).Decode(&item)	
	itemList = append(itemList, item)

	json.NewEncoder(w).Encode(itemList)

}
func viewItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(itemList)

}
