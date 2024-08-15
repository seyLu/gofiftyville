package route

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/api/controller"
)

func CrimeSceneReportsRoutes(s *gin.RouterGroup) {
	router := s.Group("/crime-scene-reports")
	{
		router.GET("", controller.GetCrimeSceneReports)
	}
}
