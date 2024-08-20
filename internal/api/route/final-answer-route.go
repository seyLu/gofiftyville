package route

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/api/controller"
)

func FinalAnswerRoute(c *gin.RouterGroup) {
	router := c.Group("/final-answer")
	{
		router.GET("", controller.GetFinalAnswer)
	}
}
