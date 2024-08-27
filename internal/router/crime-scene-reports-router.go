package router

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/controller"
)

func CrimeSceneReportsRouter(s *gin.RouterGroup) {
	r := s.Group("/crime-scene-reports")
	{
		r.GET("", controller.GetCrimeSceneReports)
	}
}
