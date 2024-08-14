package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getCrimeSceneReports(c *gin.Context) {
	fmt.Println(c.Request.Body)
}
