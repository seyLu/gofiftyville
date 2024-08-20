package sqlite3

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "../../internal/store/sqlite3/fiftyville.db")
	if err != nil {
		return err
	}
	return nil
}
