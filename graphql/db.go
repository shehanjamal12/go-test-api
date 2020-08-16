package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	_"github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "item"
)

func getitem() []Items {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("Select * from items")
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
func additem(Name string, Quantity int) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Added!")

	log.Fatal(db.Exec("INSERT INTO items (itemName,itemQuantity) VALUES ($1,$2)", Name, strconv.Itoa(Quantity)))

}
func deleteitem(name string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Fatal(db.Prepare("DELETE FROM items WHERE itemName='$name'"))

	fmt.Println("Successfully Deleted!")
}
