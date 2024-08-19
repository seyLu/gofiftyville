package model

import (
	"errors"
	"fmt"
	"strings"

	"github.com/seyLu/gofiftyville/internal/store"
)

type PhoneCall struct {
	ID       int
	Caller   string
	Receiver string
	Year     int
	Month    int
	Day      int
	Duration int
}

type PhoneCallsFilter struct {
	Year               int
	Month              int
	Day                int
	DurationInequality string
	Duration           int
	Callers            []string
}

func PhoneCalls(f PhoneCallsFilter) ([]PhoneCall, error) {
	var filters []string
	query := `
		SELECT
			id, caller, receiver, year, month, day, duration
		FROM phone_calls
	`
	args := []any{}

	if f.Year != -1 && f.Month != -1 && f.Day != -1 {
		filters = append(filters, "year=? AND month=? AND day=?")
		args = append(args, f.Year, f.Month, f.Day)
	}
	if f.Duration != -1 {
		var filter string
		switch f.DurationInequality {
		case ">":
			filter = "duration>?"
		case "<":
			filter = "duration<?"
		default:
			return nil, fmt.Errorf("PhoneCalls %+v: %w", f, errors.New("invalid 'duration inequality' filter"))
		}
		filters = append(filters, filter)
		args = append(args, f.Duration)
	}
	if len(f.Callers) > 0 {
		placeholders := strings.Repeat("?, ", len(f.Callers))
		placeholders = strings.TrimSuffix(placeholders, ", ")
		filters = append(filters, fmt.Sprintf("caller IN (%v)", placeholders))
		for _, caller := range f.Callers {
			args = append(args, caller)
		}
	}

	query = QueryWithFilters(query, filters)

	rows, err := store.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("PhoneCalls %+v: %w", f, err)
	}
	defer rows.Close()

	var calls []PhoneCall
	for rows.Next() {
		var call PhoneCall
		if err := rows.Scan(&call.ID, &call.Caller, &call.Receiver, &call.Year, &call.Month, &call.Day, &call.Duration); err != nil {
			return nil, fmt.Errorf("PhoneCalls %+v: %w", f, err)
		}
		calls = append(calls, call)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("PhoneCalls %+v: %w", f, err)
	}

	return calls, nil
}
