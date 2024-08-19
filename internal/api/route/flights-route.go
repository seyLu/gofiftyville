package route

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/api/controller"
)

func FlightsRoute(s *gin.RouterGroup) {
	router := s.Group("/flights")
	{
		router.GET("", controller.GetFlights)
	}
}
