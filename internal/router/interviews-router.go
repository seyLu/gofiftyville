package router

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/controller"
)

func InterviewsRouter(s *gin.RouterGroup) {
	r := s.Group("/interviews")
	{
		r.GET("", controller.GetInterviews)
	}
}
