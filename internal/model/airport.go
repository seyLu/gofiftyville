package model

import (
	"fmt"

	"github.com/seyLu/gofiftyville/internal/store"
)

type Airport struct {
	ID           int
	Abbreviation string
	FullName     string
	City         string
}

type AirportsFilter struct {
	FullName string
	Hour     int
	Minute   int
}

func Airports(f AirportsFilter) ([]Airport, error) {
	var filters []string
	query := `
		SELECT DISTINCT
			a.id, a.abbreviation, a.full_name, a.city
		FROM airports AS a
		INNER JOIN flights AS f
			ON f.destination_airport_id=a.id
	`
	args := []any{}

	if f.FullName != "" {
		filters = append(filters, "full_name=?")
		args = append(args, f.FullName)
	}

	if f.Hour != -1 && f.Minute != -1 {
		filters = append(filters, "hour=? AND minute=?")
		args = append(args, f.Hour, f.Minute)
	}

	query = QueryWithFilters(query, filters)

	rows, err := store.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("Airports %+v: %w", f, err)
	}
	defer rows.Close()

	var airports []Airport
	for rows.Next() {
		var airport Airport
		if err := rows.Scan(&airport.ID, &airport.Abbreviation, &airport.FullName, &airport.City); err != nil {
			return nil, fmt.Errorf("Airports %+v: %w", f, err)
		}

		airports = append(airports, airport)
	}

	return airports, nil
}
