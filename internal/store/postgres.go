package store

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func InitPostgresDB() error {
	host := os.Getenv("DB_HOST")
	port := 5432
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected to the database!")

	return nil
}
