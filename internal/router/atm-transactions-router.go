package router

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/controller"
)

func AtmTransactionsRouter(s *gin.RouterGroup) {
	r := s.Group("/atm-transactions")
	{
		r.GET("", controller.GetAtmTransactions)
	}
}
