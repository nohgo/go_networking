package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Pool *sql.DB

func CreatePool() {
	tempDB, err := sql.Open("postgres", "connString")
	if err != nil {
		log.Fatal(err)
	}

	Pool = tempDB
}

func ClosePool() {
	err := Pool.Close()
	if err != nil {
		log.Fatal(err)
	}
}
