package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

type Airport struct {
	ID           int
	Abbreviation string
	FullNname    string
	City         string
}

type AtmTransaction struct {
	ID              int
	AccountNumber   int
	Year            int
	Month           int
	Day             int
	AtmLocation     string
	TransactionType string
	Amount          int
}

type BakerySecurityLog struct {
	ID           int
	Year         int
	Month        int
	Day          int
	Hour         int
	Minute       int
	Activity     string
	LicensePlate string
}

type BankAccount struct {
	AccountNumber int
	PersonId      int
	CreationYear  int
}

type CrimeSceneReport struct {
	ID          int
	Year        int
	Month       int
	Day         int
	Street      string
	Description string
}

type Flight struct {
	ID                   int
	OriginAirportId      int
	DestinationAirportId int
	Year                 int
	Month                int
	Day                  int
	Hour                 int
	Minute               int
}

type Interview struct {
	ID         int
	Name       string
	Year       int
	Month      int
	Day        int
	Transcript string
}

type Passenger struct {
	FlightId       int
	PassportNumber int
	Seat           string
}

type Person struct {
	ID             int
	Name           string
	PhoneNumber    string
	PassportNumber int
	LicensePlate   string
}

type PhoneCall struct {
	ID       int
	Caller   string
	Receiver string
	Year     int
	Month    int
	Day      int
	Duration int
}

var db *sql.DB

func CrimeSceneReports(year int, month int, day int, street string) ([]CrimeSceneReport, error) {
	var reports []CrimeSceneReport

	rows, err := db.Query(`
		SELECT *
		FROM crime_scene_reports
		WHERE year=? AND month=? AND day=? AND street=?
	`, year, month, day, street)
	if err != nil {
		return nil, fmt.Errorf("CrimeSceneReports [year %q month %q day %q street %q]: %w", year, month, day, street, err)
	}
	defer rows.Close()

	for rows.Next() {
		var report CrimeSceneReport
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
	db, err = sql.Open("sqlite3", "fiftyville.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	reports, err := CrimeSceneReports(2021, 1, 1, "Chamberlin Street")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Reports: %v\n", reports)

	router := gin.Default()

	router.Run("localhost:8080")
}
