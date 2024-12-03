package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func connectToDatabase() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/toko_buku")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	fmt.Println("Connected to database.")
}
