package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassw := os.Getenv("DB_PASSW")
	connection := fmt.Sprintf("user=%s dbname=loja password=%s host=localhost sslmode=disable", dbUser, dbPassw)

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}
