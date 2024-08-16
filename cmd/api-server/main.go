package main

import (
	"log"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/seyLu/gofiftyville/internal/api"
	"github.com/seyLu/gofiftyville/internal/store"
)

func main() {
	err := store.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	api.StartServer("localhost", 3000)
}
