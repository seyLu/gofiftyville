package store

import (
	"database/sql"
	"path"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func InitSqlite3DB() error {
	rootPath, err := FindRoot()
	if err != nil {
		return err
	}

	DB, err = sql.Open("sqlite3", path.Join(rootPath, "internal/store/fiftyville.db"))
	if err != nil {
		return err
	}
	return nil
}
