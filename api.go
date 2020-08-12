package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//Items struct
type Items struct {
	ItemName     string `json:"itemName"` //json part so that in postman it comes as simple i not capital
	ItemQuantity int    `json:"itemQuantity"`
}

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
var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var itemList []Items = []Items{} //slice used to hold all the items which will be added dleted or updated

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello"))

}

func viewbyName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var name = mux.Vars(r)["name"]
	var specific Items
	var i int = 0
	var length = len(itemList)
	for _, item := range itemList {
		if item.ItemName == name {
			specific = itemList[i]
			break
		}
		i++
		if i > length {
			w.WriteHeader(404)
			w.Write([]byte("Error: Could Not Find Item"))
		}
	}
	json.NewEncoder(w).Encode(specific)
}
func delItembyName(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") //makes sure all content in the reponse is of json type
	var name = mux.Vars(r)["name"]

	var i int = 0
	var length = len(itemList) //get length of list
	for _, item := range itemList {
		if item.ItemName == name {
			itemList = append(itemList[:i], itemList[i+1:]...)
			break
		}
		i++
		if i > length {
			w.WriteHeader(404)
			w.Write([]byte("Error: Could Not Find Item"))
		}
	}
	//response of the fucntion
	json.NewEncoder(w).Encode(itemList)
	deleteitem(name)
}
func delItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var idparam = mux.Vars(r)["id"]  //get the id from postman request url
	id, err := strconv.Atoi(idparam) //vonversion from string to integer
	//validation
	if err != nil {
		w.Write([]byte("Error: Id Could not be converted"))
	}
	if id > len(itemList) {
		w.Write([]byte("Error: Id Could Not Be found"))
	}
	//deleting and shifting the slice so that it will be ordered again correctly
	iname := itemList[id].ItemName
	itemList = append(itemList[:id], itemList[id+1:]...)
	json.NewEncoder(w).Encode(itemList)
	deleteitem(iname)
}
func addItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	var item Items
	//used to get the request body and sets it to the variable
	json.NewDecoder(r.Body).Decode(&item) //pointer is used here so that it will diretcly address the variable
	itemList = append(itemList, item)

	json.NewEncoder(w).Encode(itemList)
	additem(item)
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
	w.WriteHeader(200)
	w.Write([]byte("items"))
	getitem()
	json.NewEncoder(w).Encode(itemList)

}
func updateItembyName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var name = mux.Vars(r)["name"]
	var updateitem Items
	var i int = 0
	var length = len(itemList)
	for _, item := range itemList {
		if item.ItemName == name {
			json.NewDecoder(r.Body).Decode(&updateitem)
			itemList[i] = updateitem
			break
		}
		i++
		if i > length {
			w.Write([]byte("Error: Could Not Find Item"))
		}
	}

	json.NewEncoder(w).Encode(itemList)
}

func updateItembyID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var idparam = mux.Vars(r)["id"]
	var updateitem Items
	id, err := strconv.Atoi(idparam)

	if err != nil {
		w.Write([]byte("Error: Id Could not be converted"))
	}
	if id > len(itemList) {
		w.Write([]byte("Error: Id Could Not Be found"))
	}
	json.NewDecoder(r.Body).Decode(&updateitem)
	itemList[id] = updateitem
	json.NewEncoder(w).Encode(itemList)
}
