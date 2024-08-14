package api

import (
	"github.com/gin-gonic/gin"
)

func StartServer(domain string, port string) {
	host := domain + port

	router := gin.Default()

	router.Run(host)
}
