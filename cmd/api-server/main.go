package main

import (
	"fmt"
	"log"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/seyLu/gofiftyville/internal/api"
	"github.com/seyLu/gofiftyville/internal/model"
	"github.com/seyLu/gofiftyville/internal/store"
)

func main() {
	var err error
	err = store.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	reports, err := model.CrimeSceneReports(2021, 1, 1, "Chamberlin Street")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Reports: %v\n", reports)

	api.StartServer("localhost", 3000)
}
