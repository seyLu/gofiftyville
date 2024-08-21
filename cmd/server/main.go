package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/seyLu/gofiftyville/internal/api"
	"github.com/seyLu/gofiftyville/internal/store"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	switch os.Getenv("DEV") {
	case "local":
		err = store.InitSqlite3DB()
	case "docker":
		err = store.InitPostgresDB()
	default:
		err = store.InitSqlite3DB()
	}
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := store.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	api.StartServer()
}
