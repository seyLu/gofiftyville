package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func SetupDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	switch os.Getenv("DEV") {
	case "local":
		err = InitSqlite3DB()
	case "docker":
		err = InitPostgresDB()
	default:
		err = InitSqlite3DB()
	}
	if err != nil {
		log.Fatal(err)
	}
}

func SetupTestDB() {
	err := InitSqlite3DB()
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

func FindRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); !os.IsNotExist(err) {
			return wd, nil
		}

		parent := filepath.Dir(wd)
		if parent == wd {
			return "", fmt.Errorf("could not find project root")
		}
		wd = parent
	}
}
