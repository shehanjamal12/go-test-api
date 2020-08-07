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
  dbname   = "test"
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

	rows, err := db.Query("Select id, name from person")
	
	if err != nil {
		fmt.Print(err)
	}
	for rows.Next() {
		//assign values to variables
		var fname string
		var lname string
		err := rows.Scan(&fname, &lname)
		if err != nil {
			fmt.Print(err)
		}
		//print results to console
		fmt.Printf("%s %s\n", fname, lname)
	}
}
