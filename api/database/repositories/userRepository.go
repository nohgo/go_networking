package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/nohgo/go_networking/api/database"
	"log"
)

var pool = db.Pool

func getAll() {
}
