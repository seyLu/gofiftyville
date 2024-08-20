package model

import (
	"fmt"
	"strings"

	"github.com/seyLu/gofiftyville/internal/store"
)

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

type PassengerFlight struct {
	PassportNumber     int
	Seat               string
	Year               int
	Month              int
	Day                int
	Hour               int
	Minute             int
	OriginAirport      string
	DestinationAirport string
}

type FlightsFilter struct {
	Year            int
	Month           int
	Day             int
	PassportNumbers []int
}

func Flights(f FlightsFilter) ([]PassengerFlight, error) {
	var filters []string
	query := `
		SELECT
			p.passport_number, p.seat,
			f.year, f.month, f.day, f.hour, f.minute,
			oa.full_name, da.full_name
		FROM flights AS f
		INNER JOIN passengers AS p
			ON p.flight_id=f.id
		INNER JOIN airports AS oa
			ON oa.id=f.origin_airport_id
		INNER JOIN airports AS da
			ON da.id=f.destination_airport_id
	`
	args := []any{}

	if f.Year != -1 && f.Month != -1 && f.Day != -1 {
		filters = append(filters, "year=? AND month=? AND day=?")
		args = append(args, f.Year, f.Month, f.Day)
	}
	if len(f.PassportNumbers) > 0 {
		placeholders := strings.Repeat("?, ", len(f.PassportNumbers))
		placeholders = strings.TrimSuffix(placeholders, ", ")
		filters = append(filters, fmt.Sprintf("passport_number IN (%v)", placeholders))
		for _, passportNumber := range f.PassportNumbers {
			args = append(args, passportNumber)
		}
	}

	query = QueryWithFilters(query, filters)

	rows, err := store.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("-> (1) model.Flights %+v: %w", f, err)
	}
	defer rows.Close()

	var flights []PassengerFlight
	for rows.Next() {
		var flight PassengerFlight
		if err := rows.Scan(
			&flight.PassportNumber, &flight.Seat,
			&flight.Year, &flight.Month, &flight.Day, &flight.Hour, &flight.Minute,
			&flight.OriginAirport, &flight.DestinationAirport,
		); err != nil {
			return nil, fmt.Errorf("-> (2) model.Flights %+v: %w", f, err)
		}

		flights = append(flights, flight)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("-> (3) model.Flights %+v: %w", f, err)
	}

	return flights, nil
}
