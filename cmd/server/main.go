package main

import (
	"log"

	"github.com/seyLu/gofiftyville/internal/api"
	"github.com/seyLu/gofiftyville/internal/store"
)

func main() {
	err := store.InitPostgresDB()
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
