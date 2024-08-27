package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type FinalAnswerFilter struct {
	Thief      string
	City       string
	Accomplice string
}

var answerMap map[string]string = map[string]string{
	"thief":      "Bruce",
	"city":       "New York City",
	"accomplice": "Robin",
}

func GetFinalAnswer(c *gin.Context) {
	req := c.Request.URL.Query()

	f := FinalAnswerFilter{
		Thief:      strings.TrimSpace(req.Get("thief")),
		City:       strings.TrimSpace(req.Get("city")),
		Accomplice: strings.TrimSpace(req.Get("accomplice")),
	}

	isThief := strings.EqualFold(answerMap["thief"], f.Thief)
	isCity := strings.EqualFold(answerMap["city"], f.City)
	isAccomplice := strings.EqualFold(answerMap["accomplice"], f.Accomplice)

	if !isThief || !isAccomplice || !isCity {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong Answer. Please try again."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Congratulations! You found the culprit!"})
}
