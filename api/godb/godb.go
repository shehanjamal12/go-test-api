package godb

import (
	"fmt"
	"log"
	"strconv"
	goapi"api/goapi"
	api"api"
)
var  dbstr goapi.Dbinter
type dbstruc struct {
	repo goapi.Dbinter
}
type Items struct{
	ItemName     string `json:"itemName"` //json part so that in postman it comes as simple i not capital
	ItemQuantity int    `json:"itemQuantity"`
}

func (h *dbstruc) Getitem() []Items {
	conn:=api.Dbconnect()
	fmt.Println("heyyyyyyyyyyyyy")
	rows, err := conn.Query("Select * from items")
	if err != nil {
		fmt.Print(err)
	}
	
	newItemList := []Items{}
	for rows.Next() {
		//assign values to variables
		var iName string
		var iQuantity string
		err := rows.Scan(&iName, &iQuantity)
		if err != nil {
			fmt.Print(err)
		}
		qua, err := strconv.Atoi(iQuantity)
		if err != nil {
			log.Println(err)
		}
		test := Items{
			ItemName:     iName,
			ItemQuantity: qua,
		}

		newItemList = append(newItemList, test)

	}
	return newItemList
}
func (h *dbstruc) Add(items Items) {

	//item variable comes through parameter as empty and gives a runtime error
	conn:=api.Dbconnect()
	_, err := conn.Exec("INSERT INTO items (itemName,itemQuantity) VALUES ($1,$2)", items.ItemName, strconv.Itoa(items.ItemQuantity))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Successfully added")
}

func (h *dbstruc) Deleteitem(name string) {
	conn:=api.Dbconnect()
	_, err := conn.Exec("DELETE FROM items WHERE itemName=$1", name)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Successfully Deleted!")
}
