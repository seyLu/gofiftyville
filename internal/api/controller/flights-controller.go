package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/model"
)

type PassengerFlight struct {
	PassportNumber     int    `json:"passportNumber"`
	Seat               string `json:"seat"`
	Date               string `json:"date"`
	TimeFormatted      string `json:"timeFormatted"`
	OriginAirport      string `json:"originAirport"`
	DestinationAirport string `json:"destinationAirport"`
}

func GetFlights(c *gin.Context) {
	request := c.Request.URL.Query()

	f := model.FlightsFilter{
		Year:            -1,
		Month:           -1,
		Day:             -1,
		PassportNumbers: nil,
	}

	flightDate := strings.TrimSpace(request.Get("date"))
	if flightDate != "" {
		parsedFlightDate, err := ParseDate(flightDate)
		if err != nil {
			errMsg := fmt.Sprintf("(1) controller.GetFlights (date %s, passportNumbers %v): %v", flightDate, nil, err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": errMsg})
			return
		}

		f.Year, f.Month, f.Day = parsedFlightDate.Year, parsedFlightDate.Month, parsedFlightDate.Day
	}

	passportNumbersReq := request["passport-number"]
	var passportNumbers []int
	for _, passportNumber := range passportNumbersReq {
		pN, err := strconv.Atoi(strings.TrimSpace(passportNumber))
		if err != nil {
			errMsg := fmt.Sprintf("(2) controller.GetFlights (date %s, passportNumbers %v): %v", flightDate, passportNumbersReq, err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": errMsg})
			return
		}
		passportNumbers = append(passportNumbers, pN)
	}
	if len(passportNumbers) != 0 {
		f.PassportNumbers = passportNumbers
	}

	flights, err := model.Flights(f)
	if err != nil {
		errMsg := fmt.Sprintf("(3) controller.Flights (date %s, passportNumbers %v): %v", flightDate, f.PassportNumbers, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	var passengerFlights []PassengerFlight
	for _, flight := range flights {
		date := fmt.Sprintf("%s %d, %d", time.Month(flight.Month).String(), flight.Day, flight.Year)
		timeSuffix := "AM"
		if flight.Hour >= 12 {
			timeSuffix = "PM"
		}
		hour := flight.Hour % 12
		if hour == 0 {
			hour = 12
		}
		timeFormatted := fmt.Sprintf("%02d:%02d %s", hour, flight.Minute, timeSuffix)

		passengerFlights = append(passengerFlights, PassengerFlight{
			PassportNumber:     flight.PassportNumber,
			Seat:               flight.Seat,
			Date:               date,
			TimeFormatted:      timeFormatted,
			OriginAirport:      flight.OriginAirport,
			DestinationAirport: flight.DestinationAirport,
		})
	}

	c.JSON(http.StatusOK, passengerFlights)
}
