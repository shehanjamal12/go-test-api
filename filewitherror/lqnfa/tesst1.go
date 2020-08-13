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
type test interface {
	//delItem(w http.ResponseWriter, r *http.Request)
	addItem(item Items)
	//viewAllItem(w http.ResponseWriter, r *http.Request)
}
type structapi struct {
	tapi test
}
type strucdb struct {
	repo dbinter
}

func newstructapi(tapi test) *structapi {
	che := &structapi{
		tapi,
	}
	return che
}

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var itemList []Items = []Items{} //slice used to hold all the items which will be added dleted or updated

func (h strucdb) addItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	var item Items
	//used to get the request body and sets it to the variable
	json.NewDecoder(r.Body).Decode(&item) //pointer is used here so that it will diretcly address the variable
	h.repo.additem(item)
	itemList = append(itemList, item)
	json.NewEncoder(w).Encode(itemList)

}
