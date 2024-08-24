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
	Name       string `json:"name"`
	Date       string `json:"date"`
	Transcript string `json:"transcript"`
}

func GetInterviews(c *gin.Context) {
	request := c.Request.URL.Query()

	f := model.InterviewsFilter{
		Year:  -1,
		Month: -1,
		Day:   -1,
	}

	interviewDate := strings.TrimSpace(request.Get("date"))
	if interviewDate != "" {
		parsedInterviewDate, err := ParseDate(interviewDate)
		if err != nil {
			errMsg := fmt.Sprintf("(1) controller.GetInterviews date %s : %v", interviewDate, err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": errMsg})
			return
		}

		f.Year, f.Month, f.Day = parsedInterviewDate.Year, parsedInterviewDate.Month, parsedInterviewDate.Day
	}

	interviewsArr, err := model.Interviews(f)
	if err != nil {
		errMsg := fmt.Sprintf("(2) controller.GetInterviews date %s : %v", interviewDate, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	var interviews []Interview
	for _, interview := range interviewsArr {
		date := fmt.Sprintf("%s %d, %d", time.Month(interview.Month).String(), interview.Day, interview.Year)

		interviews = append(interviews, Interview{
			Name:       interview.Name,
			Date:       date,
			Transcript: interview.Transcript,
		})
	}

	c.JSON(http.StatusOK, interviews)
}
