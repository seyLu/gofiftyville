package controller

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/model"
)

func GetCrimeSceneReports(c *gin.Context) {
	request := c.Request.URL.Query()

	reportDate := strings.TrimSpace(request.Get("report-date"))
	reportDateLayout := "January 2, 2006"
	parsedReportDate, err := time.Parse(reportDateLayout, reportDate)
	if err != nil {
		fmt.Printf("Error parsing date %s : %v", reportDate, err)
	}

	street := strings.TrimSpace(request.Get("street"))

	reports, err := model.CrimeSceneReports(parsedReportDate.Year(), int(parsedReportDate.Month()), parsedReportDate.Day(), street)
	if err != nil {
		fmt.Printf("Error getting CrimeSceneReports (parsedReportDate %s, street %s): %v", parsedReportDate, street, err)
	}
	fmt.Printf("%v\n", reports)
}
