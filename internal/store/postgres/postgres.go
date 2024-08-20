package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	dbname   = "fiftyville"
	user     = "root"
	password = "root"
)

var DB *sql.DB

func InitDB() error {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		return err
	}
}
