package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "item"
)

func main() {
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

	fmt.Println("Successfully connected!")

	rows, err := db.Query("Select * from items")

	if err != nil {
		fmt.Print(err)
	}
	for rows.Next() {
		//assign values to variables
		var itemName string
		var itemQuantity string
		err := rows.Scan(&itemName, &itemQuantity)
		if err != nil {
			fmt.Print(err)
		}
		//print results to console
		fmt.Printf("%s %s\n", itemName, itemQuantity)
	}
}
