// user-service/main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = make(map[string]User)
var userIDCounter int

type User struct {
	ID       string
	Username string
	Password string
}

func main() {
	r := gin.Default()
	r.POST("/register", RegisterUser)
	r.POST("/login", AuthenticateUser)
	r.Run(":8081")
}

func RegisterUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := fmt.Sprintf("user%d", userIDCounter)
	newUser.ID = userID
	users[userID] = newUser
	userIDCounter++
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": newUser})
}

func AuthenticateUser(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, exists := users[credentials.Username]
	if !exists || user.Password != credentials.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Authentication successful", "user": user})
}
