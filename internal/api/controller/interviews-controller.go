package controller

import (
	"encoding/json"
	"fmt"
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
			fmt.Printf("Error parsing date %s : %v", parsedInterviewDate, err)
		}

		year, month, day = parsedInterviewDate.Year(), int(parsedInterviewDate.Month()), parsedInterviewDate.Day()
	}

	interviews, err := model.Interviews(year, month, day)
	if err != nil {
		fmt.Printf("Error getting interviews date %s : %v", parsedInterviewDate, err)
	}

	var interviewsArr []Interview
	for _, interview := range interviews {
		var i Interview
		i.Name = interview.Name
		i.DateFormatted = fmt.Sprintf("%s %d, %d", time.Month(interview.Month).String(), interview.Day, interview.Year)
		i.Transcript = interview.Transcript
		interviewsArr = append(interviewsArr, i)
	}

	interviewsData, err := json.Marshal(interviewsArr)
	if err != nil {
		fmt.Printf("Could not marshal json: %s\n", err)
		return
	}

	fmt.Println(string(interviewsData))
}
