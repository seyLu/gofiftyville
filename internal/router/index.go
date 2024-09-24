package router

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func SetupRouter() *gin.Engine {
	app := gin.New()

	r := app.Group("/api/v1")
	r.Use(cors.New(cors.Options{
		AllowedOrigins: strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
		AllowedMethods: []string{"GET", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Content-Type", "Authorization"},
		ExposedHeaders: []string{"Content-Length"},
	}))
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
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", os.Getenv("ALLOW_ORIGINS"))
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Status(http.StatusOK)
	})
	addRoutes(r)

	return app
}

func addRoutes(s *gin.RouterGroup) {
	CrimeSceneReportsRouter(s)
	InterviewsRouter(s)
	BakerySecurityLogsRouter(s)
	AtmTransactionsRouter(s)
	People(s)
	PhoneCallsRouter(s)
	FlightsRouter(s)
	AirportsRouter(s)
	FinalAnswerRouter(s)
}
