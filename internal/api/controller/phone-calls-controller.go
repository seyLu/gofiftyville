package controller

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/model"
)

type PhoneCall struct {
	Caller   string `json:"caller"`
	Receiver string `json:"receiver"`
	Date     string `json:"date"`
	Duration int    `json:"duration"`
}

func GetPhoneCalls(c *gin.Context) {
	request := c.Request.URL.Query()

	f := model.PhoneCallsFilter{
		Year:               -1,
		Month:              -1,
		Day:                -1,
		DurationInequality: strings.TrimSpace(request.Get("duration-inequality")),
		Duration:           -1,
		Callers:            nil,
	}

	callDate := strings.TrimSpace(request.Get("date"))
	if callDate != "" {
		parsedCallDate, err := ParseDate(callDate)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		f.Year, f.Month, f.Day = parsedCallDate.Year, parsedCallDate.Month, parsedCallDate.Day
	}

	duration := strings.TrimSpace(request.Get("duration"))
	if duration != "" {
		d, err := strconv.Atoi(duration)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		f.Duration = d
	}

	callers := request["caller"]
	for i, caller := range callers {
		callers[i] = strings.TrimSpace(caller)
	}
	if len(callers) != 0 {
		validIneq := []string{">", "<"}
		if len(f.DurationInequality) == 1 && slices.Contains(validIneq, f.DurationInequality) {
			f.Callers = callers
		} else {
			errMsg := fmt.Sprintf("Invalid Duration Inequality %s", f.DurationInequality)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return
		}
	}

	phoneCalls, err := model.PhoneCalls(f)
	if err != nil {
		errMsg := fmt.Sprintf("Error getting PhoneCalls (callDate %s, durationInequality %s, duration %s, callers %s): %v", callDate, f.DurationInequality, duration, callers, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	var calls []PhoneCall
	for _, call := range phoneCalls {
		dateFormatted := fmt.Sprintf("%s %d, %d", time.Month(call.Month).String(), call.Day, call.Year)

		calls = append(calls, PhoneCall{
			Caller:   call.Caller,
			Receiver: call.Receiver,
			Date:     dateFormatted,
			Duration: call.Duration,
		})
	}

	c.JSON(http.StatusOK, calls)
}
