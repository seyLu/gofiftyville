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

func GetFinalAnswer(c *gin.Context) {
	req := c.Request.URL.Query()

	f := FinalAnswerFilter{
		Thief:      strings.TrimSpace(req.Get("thief")),
		City:       strings.TrimSpace(req.Get("city")),
		Accomplice: strings.TrimSpace(req.Get("accomplice")),
	}

	isThief := strings.EqualFold(AnswerMap["thief"], f.Thief)
	isCity := strings.EqualFold(AnswerMap["city"], f.City)
	isAccomplice := strings.EqualFold(AnswerMap["accomplice"], f.Accomplice)

	if !isThief || !isAccomplice || !isCity {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong Answer. Please try again."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Congratulations! You found the culprit!"})
}
