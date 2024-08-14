package model

import (
	"fmt"
	"github.com/seyLu/gofiftyville/internal/store"
)

type CrimeSceneReport struct {
	ID          int
	Year        int
	Month       int
	Day         int
	Street      string
	Description string
}

func CrimeSceneReports(year int, month int, day int, street string) ([]CrimeSceneReport, error) {
	var reports []CrimeSceneReport

	rows, err := store.DB.Query(`
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
