package model

import (
	"fmt"
	"strings"

	"github.com/seyLu/gofiftyville/internal/store"
)

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

type BakerySecurityLogsFilter struct {
	Year     int
	Month    int
	Day      int
	Hour     int
	Minute   int
	Hour2    int
	Minute2  int
	Activity string
}

func BakerySecurityLogs(f BakerySecurityLogsFilter) ([]BakerySecurityLog, error) {
	var filters []string
	query := `
		SELECT
			id, year, month, day, hour, minute, activity, license_plate
		FROM bakery_security_logs
	`
	args := []any{}

	if f.Year != -1 && f.Month != -1 && f.Day != -1 {
		filters = append(filters, "year=? AND month=? AND day=?")
		args = append(args, f.Year, f.Month, f.Day)
	}
	if f.Hour != -1 && f.Minute != -1 && f.Hour2 != -1 && f.Minute2 != -1 {
		filters = append(filters, `
			(hour>? OR (hour=? AND minute>=?))
  			AND (hour<? OR (hour=? AND minute<=?))
		`)
		args = append(args,
			f.Hour, f.Hour, f.Minute,
			f.Hour2, f.Hour2, f.Minute2,
		)
	}
	if f.Activity != "" {
		filters = append(filters, "LOWER(activity)=?")
		args = append(args, strings.ToLower(f.Activity))
	}

	query = QueryWithFilters(query, filters)

	rows, err := store.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("BakerySecurityLogs [year %q month %q day %q hour %q minute %q hour2 %q minute2 %q activity %q]: %w", f.Year, f.Month, f.Day, f.Hour, f.Minute, f.Hour2, f.Minute2, f.Activity, err)
	}
	defer rows.Close()

	var securityLogs []BakerySecurityLog
	for rows.Next() {
		var securityLog BakerySecurityLog
		if err := rows.Scan(&securityLog.ID, &securityLog.Year, &securityLog.Month, &securityLog.Day, &securityLog.Hour, &securityLog.Minute, &securityLog.Activity, &securityLog.LicensePlate); err != nil {
			return nil, fmt.Errorf("BakerySecurityLogs [year %q month %q day %q hour %q minute %q hour2 %q minute2 %q activity %q]: %w", f.Year, f.Month, f.Day, f.Hour, f.Minute, f.Hour2, f.Minute2, f.Activity, err)
		}

		securityLogs = append(securityLogs, securityLog)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("BakerySecurityLogs [year %q month %q day %q hour %q minute %q hour2 %q minute2 %q activity %q]: %w", f.Year, f.Month, f.Day, f.Hour, f.Minute, f.Hour2, f.Minute2, f.Activity, err)
	}

	return securityLogs, nil
}
