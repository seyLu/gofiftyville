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
		return Time{}, fmt.Errorf("error parsing time %s : %w", timeFormatted, err)
	}

	return Time{
		Hour:   parsedTime.Hour(),
		Minute: parsedTime.Minute(),
	}, nil
}

func FormatDate(year int, month int, day int) string {
	return fmt.Sprintf("%s %d, %d", time.Month(month).String(), day, year)
}

func FormatTime(hour int, minute int) string {
	timeSuffix := "AM"
	if hour >= 12 {
		timeSuffix = "PM"
	}
	hour = hour % 12
	if hour == 0 {
		hour = 12
	}
	return fmt.Sprintf("%02d:%02d %s", hour, minute, timeSuffix)
}
