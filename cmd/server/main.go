package main

import (
	"log"

	"github.com/seyLu/gofiftyville/internal/api"
	"github.com/seyLu/gofiftyville/internal/store/postgres"
)

func main() {
	err := postgres.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	api.StartServer()
}
