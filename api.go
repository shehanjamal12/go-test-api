package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Items struct {
	ItemName     string `json:"itemName"`
	ItemQuantity int    `json:"itemQuantity"`
}

var itemList []Items = []Items{}

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/addItem", addItem).Methods("POST")
	route.HandleFunc("/viewAllItem", viewAllItem).Methods("GET")
	route.HandleFunc("/viewItem/{id}", viewItem).Methods("GET")
	route.HandleFunc("/delItem/{id}", delItem).Methods("DELETE")
	route.HandleFunc("/delItembyName/{name}", delItembyName).Methods("DELETE")
	http.ListenAndServe(":5000", route)

}
func delItembyName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var name = mux.Vars(r)["name"]

	var i int = 0
	var length = len(itemList)
	for _, item := range itemList {
		if item.ItemName == name {
			itemList = append(itemList[:i], itemList[i+1:]...)
			break
		}
		if i == length {
			w.Write([]byte("Error: Could Not Find Item"))
		}
		i++
	}
	json.NewEncoder(w).Encode(itemList)
}
func delItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var idparam = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparam)

	if err != nil {
		w.Write([]byte("Error: Id Could not be converted"))
	}
	if id > len(itemList) {
		w.Write([]byte("Error: Id Could Not Be found"))
	}
	itemList = append(itemList[:id], itemList[id+1:]...)
	json.NewEncoder(w).Encode(itemList)
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
	var idparam = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparam)

	if err != nil {
		w.Write([]byte("Error: Id Could not be converted"))
	}
	if id > len(itemList) {
		w.Write([]byte("Error: Id Could Not Be found"))
	}

	specficItem := itemList[id]
	json.NewEncoder(w).Encode(specficItem)

}

func viewAllItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(itemList)

}
