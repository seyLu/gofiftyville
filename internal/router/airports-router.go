package router

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/controller"
)

func AirportsRouter(s *gin.RouterGroup) {
	r := s.Group("/airports")
	{
		r.GET("", controller.GetAirports)
	}
}
