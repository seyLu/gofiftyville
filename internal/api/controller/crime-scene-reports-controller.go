package controller

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/model"
)

type CrimeSceneReport struct {
	DateFormatted string `json:"dateFormatted"`
	Street        string `json:"street"`
	Description   string `json:"description"`
}

func GetCrimeSceneReports(c *gin.Context) {
	request := c.Request.URL.Query()

	year, month, day := -1, -1, -1
	parsedReportDate := ""

	reportDate := strings.TrimSpace(request.Get("date"))
	if reportDate != "" {
		layout := "January 2, 2006"
		parsedReportDate, err := time.Parse(layout, reportDate)
		if err != nil {
			fmt.Printf("Error parsing date %s : %v", reportDate, err)
		}

		year, month, day = parsedReportDate.Year(), int(parsedReportDate.Month()), parsedReportDate.Day()
	}

	street := strings.TrimSpace(request.Get("street"))

	crimeSceneReports, err := model.CrimeSceneReports(year, month, day, street)
	if err != nil {
		fmt.Printf("Error getting CrimeSceneReports (parsedReportDate %s, street %s): %v", parsedReportDate, street, err)
	}

	var reports []CrimeSceneReport
	for _, report := range crimeSceneReports {
		var r CrimeSceneReport
		r.DateFormatted = fmt.Sprintf("%s %d, %d", time.Month(report.Month).String(), report.Day, report.Year)
		r.Street = report.Street
		r.Description = report.Description
		reports = append(reports, r)
	}

	reportsData, err := json.Marshal(reports)
	if err != nil {
		fmt.Printf("Could not marshal json: %s\n", err)
		return
	}

	fmt.Println(string(reportsData))
}
