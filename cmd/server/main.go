package main

import (
	"log"

	"github.com/seyLu/gofiftyville/internal/api"
	"github.com/seyLu/gofiftyville/internal/store"
)

func main() {
	err := store.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	api.StartServer()
}
