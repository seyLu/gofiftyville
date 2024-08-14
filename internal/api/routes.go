package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetCrimeSceneReports(c *gin.Context) {
	fmt.Println(c.Request.Body)
}
