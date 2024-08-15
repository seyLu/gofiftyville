package controller

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/model"
)

type Report struct {
	DateFormatted string `json:"dateFormatted"`
	Street        string `json:"street"`
	Description   string `json:"description"`
}

func GetCrimeSceneReports(c *gin.Context) {
	request := c.Request.URL.Query()

	reportDate := strings.TrimSpace(request.Get("report-date"))
	reportDateLayout := "January 2, 2006"
	parsedReportDate, err := time.Parse(reportDateLayout, reportDate)
	if err != nil {
		fmt.Printf("Error parsing date %s : %v", reportDate, err)
	}

	street := strings.TrimSpace(request.Get("street"))

	crimeSceneReports, err := model.CrimeSceneReports(parsedReportDate.Year(), int(parsedReportDate.Month()), parsedReportDate.Day(), street)
	if err != nil {
		fmt.Printf("Error getting CrimeSceneReports (parsedReportDate %s, street %s): %v", parsedReportDate, street, err)
	}

	var reports []Report
	for _, report := range crimeSceneReports {
		var r Report
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
