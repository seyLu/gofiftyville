package route

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/api/controller"
)

func BakerySecurityLogsRoute(s *gin.RouterGroup) {
	router := s.Group("/bakery-security-logs")
	{
		router.GET("", controller.GetBakerySecurityLogs)
	}
}
