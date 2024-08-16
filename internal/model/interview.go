package model

import (
	"fmt"

	"github.com/seyLu/gofiftyville/internal/store"
)

type Interview struct {
	ID         int
	Name       string
	Year       int
	Month      int
	Day        int
	Transcript string
}

func Interviews(year int, month int, day int) ([]Interview, error) {
	var filters []string
	query := `
		SELECT
			id, name, year, month, day, transcript
		FROM interviews
	`
	args := []any{}

	if year != -1 && month != -1 && day != -1 {
		filters = append(filters, "year=? AND month=? AND day=?")
		args = append(args, year, month, day)
	}

	query = QueryWithFilters(query, filters)

	rows, err := store.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("Interviews [%q year, %q month, %q day] : %w", year, month, day, err)
	}
	defer rows.Close()

	var interviews []Interview
	for rows.Next() {
		var interview Interview
		if err := rows.Scan(&interview.ID, &interview.Name, &interview.Year, &interview.Month, &interview.Day, &interview.Transcript); err != nil {
			return nil, fmt.Errorf("Interviews [%q year, %q month, %q day] : %w", year, month, day, err)
		}
		interviews = append(interviews, interview)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Interviews [%q year, %q month, %q day] : %w", year, month, day, err)
	}

	return interviews, nil
}
