package store

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var DB *sql.DB

func InitDatabase() error {
	var err error
	DB, err = sql.Open("sqlite3", "../../internal/store/fiftyville.db")
	if err != nil {
		return err
	}
	return nil
}
