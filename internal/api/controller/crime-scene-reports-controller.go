package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/model"
)

type CrimeSceneReport struct {
	Date        string `json:"date"`
	Street      string `json:"street"`
	Description string `json:"description"`
}

func GetCrimeSceneReports(c *gin.Context) {
	request := c.Request.URL.Query()

	f := model.CrimeSceneReportsFilter{
		Year:   -1,
		Month:  -1,
		Day:    -1,
		Street: strings.TrimSpace(request.Get("street")),
	}

	reportDate := strings.TrimSpace(request.Get("date"))
	if reportDate != "" {
		parsedReportDate, err := ParseDate(reportDate)
		if err != nil {
			errMsg := fmt.Sprintf("(1) controller.GetCrimeSceneReports (reportDate %s, street %s): %v", reportDate, f.Street, err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": errMsg})
			return
		}

		f.Year, f.Month, f.Day = parsedReportDate.Year, parsedReportDate.Month, parsedReportDate.Day
	}

	crimeSceneReports, err := model.CrimeSceneReports(f)
	if err != nil {
		errMsg := fmt.Sprintf("(2) controller.GetCrimeSceneReports (reportDate %s, street %s): %v", reportDate, f.Street, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	var reports []CrimeSceneReport
	for _, report := range crimeSceneReports {
		date := fmt.Sprintf("%s %d, %d", time.Month(report.Month).String(), report.Day, report.Year)

		reports = append(reports, CrimeSceneReport{
			Date:        date,
			Street:      report.Street,
			Description: report.Description,
		})
	}

	c.JSON(http.StatusOK, reports)
}
