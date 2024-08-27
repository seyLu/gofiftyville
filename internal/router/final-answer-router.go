package router

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/controller"
)

func FinalAnswerRouter(s *gin.RouterGroup) {
	r := s.Group("/final-answer")
	{
		r.GET("", controller.GetFinalAnswer)
	}
}
