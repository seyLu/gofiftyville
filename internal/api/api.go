package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/api"
)

func StartServer(domain string, port int) {
	router := gin.Default()

	router.GET("crime-scene-reports", api.GetCrimeSceneReports())

	host := fmt.Sprintf("%v:%v", domain, port)
	router.Run(host)
}
