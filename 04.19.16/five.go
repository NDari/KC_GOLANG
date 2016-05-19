package main

import (
	"fmt"
	"oracle"
)

type database interface {
	Connect() string
	Exec(string, string) string
	Close(string)
}

func Query(db database, queryString string) string { /* ... */ } // same as before

func main() {
	q := "SELECT * FROM happy_fun_times;"
	var db database   // notice this thing
	db = oracle.New() // implicit casting... same as db := database(oracle.New())
	fmt.Println(Query(db, q))
}
