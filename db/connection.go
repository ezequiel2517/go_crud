package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	serverSQL := os.Getenv("SERVIDOR_SQL")
	dbName := os.Getenv("DB")
	userDB := os.Getenv("USER_DB")
	passwordDB := os.Getenv("PASSWORD_DB")
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", userDB, passwordDB, serverSQL, dbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
