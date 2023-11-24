package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	dsn := "postgres://postgres:miContrasenia@localhost:15432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
