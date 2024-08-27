package router

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/controller"
)

func People(s *gin.RouterGroup) {
	r := s.Group("/people")
	{
		r.GET("", controller.GetPeople)
	}
}
