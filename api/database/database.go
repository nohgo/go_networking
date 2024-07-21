// The database package contains the connection pool for the postgres server
// Call [db.CreatePool] and defer [db.ClosePool] in your main function then access the global variable [db.Pool]
package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// The connection pool for the postgres server
var Pool *sql.DB

// Initialzes the pool with enviroment variable "GO_CONN_STRING"
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

// Closes the open pool
func ClosePool() {
	err := Pool.Close()
	if err != nil {
		log.Fatal(err)
	}
}
