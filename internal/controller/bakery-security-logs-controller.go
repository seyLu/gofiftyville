package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/model"
)

type BakerySecurityLog struct {
	Date          string `json:"date"`
	TimeFormatted string `json:"time"`
	Activity      string `json:"activity"`
	LicensePlate  string `json:"licensePlate"`
}

func GetBakerySecurityLogs(c *gin.Context) {
	req := c.Request.URL.Query()

	f := model.BakerySecurityLogsFilter{
		Year:     -1,
		Month:    -1,
		Day:      -1,
		Hour:     -1,
		Minute:   -1,
		Hour2:    -1,
		Minute2:  -1,
		Activity: strings.TrimSpace(req.Get("activity")),
	}

	logDate := strings.TrimSpace(req.Get("date"))
	if logDate != "" {
		parsedLogDate, err := ParseDate(logDate)
		if err != nil {
			errMsg := fmt.Sprintf("(1) controller.GetBakerySecurityLogs (date %s, time %s, time-2 %s, activity %s): %v", logDate, "", "", f.Activity, err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": errMsg})
			return
		}

		f.Year, f.Month, f.Day = parsedLogDate.Year, parsedLogDate.Month, parsedLogDate.Day
	}

	logTime := strings.TrimSpace(req.Get("time"))
	if logTime != "" {
		parsedLogTime, err := ParseTime(logTime)
		if err != nil {
			errMsg := fmt.Sprintf("(2) controller.GetBakerySecurityLogs (date %s, time %s, time-2 %s, activity %s): %v", logDate, logTime, "", f.Activity, err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": errMsg})
			return
		}

		f.Hour, f.Minute = parsedLogTime.Hour, parsedLogTime.Minute
	}

	logTime2 := strings.TrimSpace(req.Get("time-2"))
	if logTime2 != "" {
		parsedLogTime2, err := ParseTime(logTime2)
		if err != nil {
			errMsg := fmt.Sprintf("(3) controller.GetBakerySecurityLogs (date %s, time %s, time-2 %s, activity %s): %v", logDate, logTime, logTime2, f.Activity, err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": errMsg})
			return
		}

		f.Hour2, f.Minute2 = parsedLogTime2.Hour, parsedLogTime2.Minute
	}

	bakerySecurityLogs, err := model.BakerySecurityLogs(f)
	if err != nil {
		errMsg := fmt.Sprintf("(4) controller.GetBakerySecurityLogs (date %s, time %s, time-2 %s, activity %s): %v", logDate, logTime, logTime2, f.Activity, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	var securityLogs []BakerySecurityLog
	for _, securityLog := range bakerySecurityLogs {
		date := fmt.Sprintf("%s %d, %d", time.Month(securityLog.Month).String(), securityLog.Day, securityLog.Year)
		timeSuffix := "AM"
		if securityLog.Hour >= 12 {
			timeSuffix = "PM"
		}
		hour := securityLog.Hour % 12
		if hour == 0 {
			hour = 12
		}
		time := fmt.Sprintf("%02d:%02d %s", hour, securityLog.Minute, timeSuffix)

		securityLogs = append(securityLogs, BakerySecurityLog{
			Date:          date,
			TimeFormatted: time,
			Activity:      securityLog.Activity,
			LicensePlate:  securityLog.LicensePlate,
		})
	}

	c.JSON(http.StatusOK, securityLogs)
}
