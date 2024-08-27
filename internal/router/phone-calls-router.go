package router

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/controller"
)

func PhoneCallsRotue(s *gin.RouterGroup) {
	r := s.Group("/phone-calls")
	{
		r.GET("", controller.GetPhoneCalls)
	}
}
