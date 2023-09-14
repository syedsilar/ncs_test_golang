// parking-service/main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var parkingSpots = make(map[int]ParkingSpot)
var spotIDCounter int

type ParkingSpot struct {
	ID           int
	Type         string
	Availability bool
	Floor        int
	Location     string
}

func main() {
	r := gin.Default()

	r.GET("/spots", GetParkingSpots)
	r.POST("/reserve/:spotID", ReserveParkingSpot)
	r.POST("/release/:spotID", ReleaseParkingSpot)

	r.Run(":8082")
}

func GetParkingSpots(c *gin.Context) {
	// Simulate retrieving a list of parking spots
	spots := make([]ParkingSpot, 0)
	for _, spot := range parkingSpots {
		spots = append(spots, spot)
	}

	c.JSON(http.StatusOK, spots)
}

func ReserveParkingSpot(c *gin.Context) {
	spotID := c.Param("spotID")
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Spot %s reserved successfully", spotID)})
}

func ReleaseParkingSpot(c *gin.Context) {
	spotID := c.Param("spotID")
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Spot %s released successfully", spotID)})
}
