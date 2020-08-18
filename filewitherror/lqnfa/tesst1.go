package main

import (
	"encoding/json"
	"net/http"
	"sync"

	_ "github.com/lib/pq"
)

//Items struct
type Items struct {
	ItemName     string `json:"itemName"` //json part so that in postman it comes as simple i not capital
	ItemQuantity int    `json:"itemQuantity"`
}
type Test interface {
	//delItem(w http.ResponseWriter, r *http.Request)
	addItem(w http.ResponseWriter, r *http.Request)
	//viewAllItem(w http.ResponseWriter, r *http.Request)
}
type Structapi struct {
	tapi test
}
type Strucdb struct {
	repo dbinter
}

func newstructapi(tapi test) structapi {
	return structapi{
		tapi,
	}

}

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var itemList []Items = []Items{} //slice used to hold all the items which will be added dleted or updated

func(t* structapi) addItem(w http.ResponseWriter, r *http.Request) {
	var t1 dbinter
	t2 := newdbstrc(t1)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	var item Items
	//used to get the request body and sets it to the variable
	json.NewDecoder(r.Body).Decode(&item) //pointer is used here so that it will diretcly address the variable
	t2.add(item)
	itemList = append(itemList, item)
	json.NewEncoder(w).Encode(itemList)

}
