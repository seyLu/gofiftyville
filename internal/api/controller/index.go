package controller

import (
	"fmt"
	"time"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

type Time struct {
	Hour   int
	Minute int
}

func ParseDate(dateFormatted string) (Date, error) {
	layout := "January 2, 2006"
	parsedDate, err := time.Parse(layout, dateFormatted)
	if err != nil {
		return Date{}, fmt.Errorf("Error parsing date %s : %w", dateFormatted, err)
	}

	return Date{
		Year:  parsedDate.Year(),
		Month: int(parsedDate.Month()),
		Day:   parsedDate.Day(),
	}, nil
}

func ParseTime(timeFormatted string) (Time, error) {
	layout := "03:04 PM"
	parsedTime, err := time.Parse(layout, timeFormatted)
	if err != nil {
		return Time{}, fmt.Errorf("Error parsing time %s : %w", timeFormatted, err)
	}

	return Time{
		Hour:   parsedTime.Hour(),
		Minute: parsedTime.Minute(),
	}, nil
}
