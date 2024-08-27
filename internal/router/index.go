package router

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	app := gin.New()

	r := app.Group("/api/v1")
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
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
	r.Use(gin.Recovery())
	addRoutes(r)

	return app
}

func addRoutes(s *gin.RouterGroup) {
	CrimeSceneReportsRouter(s)
	InterviewsRouter(s)
	BakerySecurityLogsRouter(s)
	AtmTransactionsRouter(s)
	People(s)
	PhoneCallsRotue(s)
	FlightsRouter(s)
	AirportsRouter(s)
	FinalAnswerRouter(s)
}
