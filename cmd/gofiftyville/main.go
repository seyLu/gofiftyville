package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/seyLu/gofiftyville/internal/model"
	"github.com/seyLu/gofiftyville/internal/store"
)

func CrimeSceneReports(year int, month int, day int, street string) ([]model.CrimeSceneReport, error) {
	var reports []model.CrimeSceneReport

	rows, err := store.db.Query(`
		SELECT *
		FROM crime_scene_reports
		WHERE year=? AND month=? AND day=? AND street=?
	`, year, month, day, street)
	if err != nil {
		return nil, fmt.Errorf("CrimeSceneReports [year %q month %q day %q street %q]: %w", year, month, day, street, err)
	}
	defer rows.Close()

	for rows.Next() {
		var report model.CrimeSceneReport
		if err := rows.Scan(&report.ID, &report.Year, &report.Month, &report.Day, &report.Street, &report.Description); err != nil {
			return nil, fmt.Errorf("CrimeSceneReports [year %q month %q day %q street %q]: %w", year, month, day, street, err)
		}
		reports = append(reports, report)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("CrimeSceneReports [year %q month %q day %q street %q]: %w", year, month, day, street, err)
	}

	return reports, nil
}

func main() {
	var err error
	db, err = store.initDatabase("fiftyville.db")
	if err != nil {
		log.Fatal(err)
	}

	reports, err := CrimeSceneReports(2021, 1, 1, "Chamberlin Street")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Reports: %v\n", reports)

	router := gin.Default()

	router.Run("localhost:8080")
}
