package main

import (
	"book-pai/database"
)

func main() {
	db, err := database.NewDatabasePostgres()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
