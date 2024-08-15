package route

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(s *gin.RouterGroup) {
	CrimeSceneReportsRoutes(s)
}
