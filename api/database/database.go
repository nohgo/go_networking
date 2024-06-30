package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Pool *sql.DB

func CreatePool() {
	tempDB, err := sql.Open("postgres", os.Getenv("GO_CONN_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	Pool = tempDB

	err = Pool.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func ClosePool() {
	err := Pool.Close()
	if err != nil {
		log.Fatal(err)
	}
}
