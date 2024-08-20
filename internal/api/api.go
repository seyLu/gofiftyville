package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seyLu/gofiftyville/internal/api/route"
)

func StartServer() {
	app := gin.New()

	router := app.Group("/api/v1")
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	route.AddRoutes(router)

	app.Run()
}
