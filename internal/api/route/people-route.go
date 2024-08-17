package route

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/api/controller"
)

func People(s *gin.RouterGroup) {
	router := s.Group("/people")
	{
		router.GET("", controller.GetPeople)
	}
}
