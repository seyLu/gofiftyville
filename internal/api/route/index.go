package route

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(s *gin.RouterGroup) {
	CrimeSceneReportsRoute(s)
	InterviewsRoute(s)
	BakerySecurityLogsRoute(s)
	AtmTransactionsRoute(s)
	People(s)
	PhoneCallsRotue(s)
	FlightsRoute(s)
	AirportsRoute(s)
}
