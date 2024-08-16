package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/model"
)

type Interview struct {
	Name          string `json:"name"`
	DateFormatted string `json:"dateFormatted"`
	Transcript    string `json:"transcript"`
}

func GetInterviews(c *gin.Context) {
	request := c.Request.URL.Query()

	year, month, day := -1, -1, -1
	parsedInterviewDate := ""

	interviewDate := strings.TrimSpace(request.Get("date"))
	if interviewDate != "" {
		layout := "January 2, 2006"
		parsedInterviewDate, err := time.Parse(layout, interviewDate)
		if err != nil {
			errMsg := fmt.Sprintf("Error parsing date %s : %v", parsedInterviewDate, err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return
		}

		year, month, day = parsedInterviewDate.Year(), int(parsedInterviewDate.Month()), parsedInterviewDate.Day()
	}

	interviewsArr, err := model.Interviews(year, month, day)
	if err != nil {
		errMsg := fmt.Sprintf("Error getting interviews date %s : %v", parsedInterviewDate, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	var interviews []Interview
	for _, interview := range interviewsArr {
		var i Interview
		i.Name = interview.Name
		i.DateFormatted = fmt.Sprintf("%s %d, %d", time.Month(interview.Month).String(), interview.Day, interview.Year)
		i.Transcript = interview.Transcript
		interviews = append(interviews, i)
	}

	c.JSON(http.StatusOK, interviews)
}
