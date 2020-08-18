package main

import (
	"fmt"
	"log"
	"strconv"
)

type Dbinter interface {
	//getitem()
	add(items Items)
	//deleteitem(w http.ResponseWriter, r *http.Request)
}
type Dbstruc struct {
	dbint Dbinter
}

func Newdbstrc(dbint Dbinter) Dbstruc {
	return Dbstruc{
		dbint,
	}
	
}

func (h *Dbstruc) add(items Items)  {
	
	//item variable comes through parameter as empty and gives a runtime error
	_, err := conn.Exec("INSERT INTO items (itemName,itemQuantity) VALUES ($1,$2)", items.ItemName, strconv.Itoa(items.ItemQuantity))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Successfully added")
}
