package route

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/api/controller"
)

func PhoneCallsRotue(s *gin.RouterGroup) {
	router := s.Group("/phone-calls")
	{
		router.GET("", controller.GetPhoneCalls)
	}
}
