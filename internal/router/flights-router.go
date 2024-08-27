package router

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/controller"
)

func FlightsRouter(s *gin.RouterGroup) {
	r := s.Group("/flights")
	{
		r.GET("", controller.GetFlights)
	}
}
