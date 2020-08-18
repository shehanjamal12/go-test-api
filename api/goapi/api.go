package goapi

import (
	"encoding/json"
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
type Test interface {
	DelItem(w http.ResponseWriter, r *http.Request)
	AddItem(w http.ResponseWriter, r *http.Request)
	ViewAllItem(w http.ResponseWriter, r *http.Request)
}
type structapi struct {
	tapi Test
}
type Strucdb struct {
	repo Dbinter
}

func Newdbstrc(repo Dbinter) Strucdb {
	return Strucdb{
		repo,
	}

}
func Newstructapi(tapi Test) structapi {
	return structapi{
		tapi,
	}

}

var t1 Dbinter

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var itemList []Items = []Items{} //slice used to hold all the items which will be added dleted or updated

func test12(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello"))

}

func ViewbyName(w http.ResponseWriter, r *http.Request) {
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
func DelItembyName(w http.ResponseWriter, r *http.Request) {

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

}
func (t *structapi) DelItem(w http.ResponseWriter, r *http.Request) {
	t2 := Newdbstrc(t1)
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
	t2.repo.Deleteitem(iname)
	itemList = append(itemList[:id], itemList[id+1:]...)
	json.NewEncoder(w).Encode(itemList)

}
func (t *structapi) AddItem(w http.ResponseWriter, r *http.Request) {
	t2 := Newdbstrc(t1)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	var item Items
	//used to get the request body and sets it to the variable
	json.NewDecoder(r.Body).Decode(&item) //pointer is used here so that it will diretcly address the variable
	t2.repo.Add(item)
	itemList = append(itemList, item)
	json.NewEncoder(w).Encode(itemList)

}
func ViewItem(w http.ResponseWriter, r *http.Request) {
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

func (t *structapi) ViewAllItem(w http.ResponseWriter, r *http.Request) {
	t2 := Newdbstrc(t1)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("items"))
	itemList = t2.repo.Getitem()
	json.NewEncoder(w).Encode(itemList)

}
func UpdateItembyName(w http.ResponseWriter, r *http.Request) {
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

func UpdateItembyID(w http.ResponseWriter, r *http.Request) {
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
