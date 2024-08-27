package router

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/controller"
)

func BakerySecurityLogsRouter(s *gin.RouterGroup) {
	r := s.Group("/bakery-security-logs")
	{
		r.GET("", controller.GetBakerySecurityLogs)
	}
}
