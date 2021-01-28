package database

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func init() {
	dsn := os.Getenv("DSN")
	engine := os.Getenv("ENGINE")

	db, _ = sql.Open(engine, dsn)
}

//Connection returns a pool connection
func Connection() *sql.DB {
	return db
}

//Close close the connection with the database
func Close() error {
	return db.Close()
}
