// analytics-service/main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Analytics struct {
	ID      string
	Message string
}

var analyticsData = make(map[string]Analytics)
var analyticsIDCounter int

func main() {
	r := gin.Default()

	r.GET("/analytics", GetAnalyticsData)

	r.Run(":8085")
}

func GetAnalyticsData(c *gin.Context) {
	analyticsID := fmt.Sprintf("analytics%d", analyticsIDCounter)
	newAnalytics := Analytics{
		ID:      analyticsID,
		Message: "Sample analytics data",
	}
	analyticsData[analyticsID] = newAnalytics
	analyticsIDCounter++

	c.JSON(http.StatusOK, newAnalytics)
}
