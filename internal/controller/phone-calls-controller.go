package controller

import (
	"errors"
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
	req := c.Request.URL.Query()

	f := model.PhoneCallsFilter{
		Year:               -1,
		Month:              -1,
		Day:                -1,
		DurationInequality: strings.TrimSpace(req.Get("duration-inequality")),
		Duration:           -1,
		Callers:            nil,
	}

	callDate := strings.TrimSpace(req.Get("date"))
	if callDate != "" {
		parsedCallDate, err := ParseDate(callDate)
		if err != nil {
			errMsg := fmt.Sprintf("(1) controller.GetPhoneCalls (callDate %s, durationInequality %s, duration %s, callers %s): %v", callDate, f.DurationInequality, "", "", err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": errMsg})
			return
		}

		f.Year, f.Month, f.Day = parsedCallDate.Year, parsedCallDate.Month, parsedCallDate.Day
	}

	duration := strings.TrimSpace(req.Get("duration"))
	if duration != "" {
		d, err := strconv.Atoi(duration)
		if err != nil {
			errMsg := fmt.Sprintf("(2) controller.GetPhoneCalls (callDate %s, durationInequality %s, duration %s, callers %s): %v", callDate, f.DurationInequality, duration, "", err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": errMsg})
			return
		}
		f.Duration = d
	}

	callers := req["caller"]
	for i, caller := range callers {
		callers[i] = strings.TrimSpace(caller)
	}
	if len(callers) != 0 {
		validIneq := []string{">", "<"}
		if len(f.DurationInequality) == 1 && slices.Contains(validIneq, f.DurationInequality) {
			f.Callers = callers
		} else {
			err := errors.New("invalid durationInequality. valid values are '<' or '>.")
			errMsg := fmt.Sprintf("(3) controller.GetPhoneCalls (callDate %s, durationInequality %s, duration %s, callers %s): %v", callDate, f.DurationInequality, duration, callers, err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return
		}
	}

	phoneCalls, err := model.PhoneCalls(f)
	if err != nil {
		errMsg := fmt.Sprintf("(4) controller.GetPhoneCalls (callDate %s, durationInequality %s, duration %v, callers %v): %v", callDate, f.DurationInequality, f.Duration, f.Callers, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	var calls []PhoneCall
	for _, call := range phoneCalls {
		date := fmt.Sprintf("%s %d, %d", time.Month(call.Month).String(), call.Day, call.Year)

		calls = append(calls, PhoneCall{
			Caller:   call.Caller,
			Receiver: call.Receiver,
			Date:     date,
			Duration: call.Duration,
		})
	}

	c.JSON(http.StatusOK, calls)
}
