package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/model"
)

type Airport struct {
	Abbreviation string `json:"abbreviation"`
	FullName     string `json:"fullName"`
	City         string `json:"city"`
}

func GetAirports(c *gin.Context) {
	request := c.Request.URL.Query()

	f := model.AirportsFilter{
		FullName: strings.TrimSpace(request.Get("full-name")),
		Hour:     -1,
		Minute:   -1,
	}

	flightTime := strings.TrimSpace(request.Get("flight-time"))
	if flightTime != "" {
		parsedFlightTime, err := ParseTime(flightTime)
		if err != nil {
			errMsg := fmt.Sprintf("-> (1) controller.Airports (fullName %s, time %s): %v", f.FullName, flightTime, err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": errMsg})
			return
		}

		f.Hour, f.Minute = parsedFlightTime.Hour, parsedFlightTime.Minute
	}

	airportsArr, err := model.Airports(f)
	if err != nil {
		errMsg := fmt.Sprintf("-> (2) controller.Airports (fullName %s, time %s): %v", f.FullName, flightTime, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	var airports []Airport
	for _, airport := range airportsArr {
		airports = append(airports, Airport{
			Abbreviation: airport.Abbreviation,
			FullName:     airport.FullName,
			City:         airport.City,
		})
	}

	c.JSON(http.StatusOK, airports)
}
