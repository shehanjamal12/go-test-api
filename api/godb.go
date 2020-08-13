package main

import (
	"fmt"
	"log"
	"strconv"
)

type dbinter interface {
	//getitem()
	additem(items Items)
	//deleteitem(w http.ResponseWriter, r *http.Request)
}
type dbstruc struct {
	dbint dbinter
}

func newdbstrc(dbint dbinter) dbstruc {
	che := dbstruc{
		dbint,
	}
	return che
}

func getitem() {
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
	itemList = newItemList
}
func (h strucdb) additem(item Items) {
	_, err := conn.Exec("INSERT INTO items (itemName,itemQuantity) VALUES ($1,$2)", item.ItemName, strconv.Itoa(item.ItemQuantity))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Successfully added")
}
func deleteitem(name string) {
	_, err := conn.Exec("DELETE FROM items WHERE itemName=$1", name)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Successfully Deleted!")
}
