// notification-service/main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Notification struct {
	ID      string
	Message string
}

var notifications = make(map[string]Notification)
var notificationIDCounter int

func main() {
	r := gin.Default()

	r.POST("/send", SendNotification)

	r.Run(":8084")
}

func SendNotification(c *gin.Context) {
	var message struct {
		UserID string `json:"userID"`
		Text   string `json:"text"`
	}
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	notificationID := fmt.Sprintf("notification%d", notificationIDCounter)
	newNotification := Notification{
		ID:      notificationID,
		Message: message.Text,
	}
	notifications[notificationID] = newNotification
	notificationIDCounter++

	c.JSON(http.StatusOK, gin.H{"message": "Notification sent successfully", "notification": newNotification})
}
