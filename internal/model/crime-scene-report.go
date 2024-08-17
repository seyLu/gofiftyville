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

type CrimeSceneReportsFilter struct {
	Year   int
	Month  int
	Day    int
	Street string
}

func CrimeSceneReports(f CrimeSceneReportsFilter) ([]CrimeSceneReport, error) {
	var filters []string
	query := `
		SELECT
			id, year, month, day, street, description
		FROM crime_scene_reports
	`
	args := []any{}

	if f.Year != -1 && f.Month != -1 && f.Day != -1 {
		filters = append(filters, "year=? AND month=? AND day=?")
		args = append(args, f.Year, f.Month, f.Day)
	}
	if f.Street != "" {
		filters = append(filters, "LOWER(street)=?")
		args = append(args, strings.ToLower(f.Street))
	}

	query = QueryWithFilters(query, filters)

	rows, err := store.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("CrimeSceneReports %+v: %w", f, err)
	}
	defer rows.Close()

	var reports []CrimeSceneReport
	for rows.Next() {
		var report CrimeSceneReport
		if err := rows.Scan(&report.ID, &report.Year, &report.Month, &report.Day, &report.Street, &report.Description); err != nil {
			return nil, fmt.Errorf("CrimeSceneReports %+v: %w", f, err)
		}
		reports = append(reports, report)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("CrimeSceneReports %+v: %w", f, err)
	}

	return reports, nil
}
