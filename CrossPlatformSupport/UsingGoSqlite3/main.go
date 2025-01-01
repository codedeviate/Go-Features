package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS example (id INTEGER, name TEXT)")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Database table created successfully!")
}
