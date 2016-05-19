package main

import (
	"fmt"
	"oracle"
)

func Query(db oracle.DB, queryString string) string {
	connectionID := db.Connect()
	defer db.Close(connectionID)
	return db.Exec(queryString, connectionID)
}

func main() {
	q, db := "SELECT * FROM happy_fun_times;", oracle.New()
	fmt.Println(Query(db, q))
}
