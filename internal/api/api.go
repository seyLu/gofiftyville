package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer(domain string, port int) {
	router := gin.Default()

	router.GET("crime-scene-reports", getCrimeSceneReports)

	host := fmt.Sprintf("%v:%v", domain, port)
	router.Run(host)
}

func getCrimeSceneReports(c *gin.Context) {
	fmt.Println(c.Request.Body)
}
