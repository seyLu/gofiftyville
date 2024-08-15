package model

import (
	"fmt"
	"strings"

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

	var filters []string
	query := `
		SELECT
			id, year, month, day, street, description
		FROM crime_scene_reports
	`
	args := []any{}

	if year != -1 && month != -1 && day != -1 {
		filters = append(filters, "year=? AND month=? AND day=?")
		args = append(args, year, month, day)
	}
	if street != "" {
		filters = append(filters, "LOWER(street)=?")
		args = append(args, strings.ToLower(street))
	}

	for i, filter := range filters {
		switch i {
		case 0:
			query += fmt.Sprintf(" WHERE %s ", filter)
		default:
			query += fmt.Sprintf(" AND %s ", filter)
		}
	}

	rows, err := store.DB.Query(query, args...)
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
