package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:1212@localhost:5432/mar_portal?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	return db, err
}
