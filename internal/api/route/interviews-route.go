package route

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/api/controller"
)

func InterviewsRoute(s *gin.RouterGroup) {
	router := s.Group("/interviews")
	{
		router.GET("", controller.GetInterviews)
	}
}
