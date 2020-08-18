package main

import (
	"fmt"
	"log"
	"strconv"
)

type dbinter interface {
	//getitem()
	add(items Items)
	//deleteitem(w http.ResponseWriter, r *http.Request)
}
type dbstruc struct {
	dbint dbinter
}

func newdbstrc(dbint dbinter) dbstruc {
	return dbstruc{
		dbint,
	}
	
}

func (h *dbstruc) add(items Items)  {
	
	//item variable comes through parameter as empty and gives a runtime error
	_, err := conn.Exec("INSERT INTO items (itemName,itemQuantity) VALUES ($1,$2)", items.ItemName, strconv.Itoa(items.ItemQuantity))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Successfully added")
}
