package route

import (
	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/api/controller"
)

func AtmTransactionsRoute(s *gin.RouterGroup) {
	router := s.Group("/atm-transactions")
	{
		router.GET("", controller.GetAtmTransactions)
	}
}
