package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/api/route"
)

func StartServer(domain string, port int) {
	app := gin.New()

	router := app.Group("/api/v1")
	route.AddRoutes(router)

	host := fmt.Sprintf("%v:%v", domain, port)
	app.Run(host)
}
