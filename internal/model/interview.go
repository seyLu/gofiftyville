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

type InterviewsFilter struct {
	Year  int
	Month int
	Day   int
}

func Interviews(f InterviewsFilter) ([]Interview, error) {
	var filters []string
	query := `
		SELECT
			id, name, year, month, day, transcript
		FROM interviews
	`
	args := []any{}

	if f.Year != -1 && f.Month != -1 && f.Day != -1 {
		filters = append(filters, "year=? AND month=? AND day=?")
		args = append(args, f.Year, f.Month, f.Day)
	}

	query = QueryWithFilters(query, filters)

	rows, err := store.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("Interviews %+v : %w", f, err)
	}
	defer rows.Close()

	var interviews []Interview
	for rows.Next() {
		var interview Interview
		if err := rows.Scan(&interview.ID, &interview.Name, &interview.Year, &interview.Month, &interview.Day, &interview.Transcript); err != nil {
			return nil, fmt.Errorf("Interviews %+v : %w", f, err)
		}
		interviews = append(interviews, interview)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Interviews %+v : %w", f, err)
	}

	return interviews, nil
}
