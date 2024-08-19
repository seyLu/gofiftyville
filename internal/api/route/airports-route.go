package route

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/api/controller"
)

func AirportsRoute(s *gin.RouterGroup) {
	router := s.Group("/airports")
	{
		router.GET("", controller.GetAirports)
	}
}
