// transaction-service/main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var transactions = make(map[string]Transaction)
var transactionIDCounter int

type Transaction struct {
	ID          string
	Timestamp   string
	Amount      float64
	VehicleInfo float64
	SpotInfo    float64
}

func main() {
	r := gin.Default()

	r.POST("/pay", MakePayment)
	r.GET("/transactions", GetTransactions)

	r.Run(":8083")
}

func MakePayment(c *gin.Context) {
	var payment struct {
		UserID   string  `json:"userID"`
		Amount   float64 `json:"amount"`
		SpotInfo float64 `json:"spotInfo"`
	}
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionID := fmt.Sprintf("transaction%d", transactionIDCounter)
	newTransaction := Transaction{
		ID:          transactionID,
		Timestamp:   "2023-09-01 12:00:00",
		Amount:      payment.Amount,
		VehicleInfo: payment.SpotInfo,
		SpotInfo:    payment.SpotInfo,
	}
	transactions[transactionID] = newTransaction
	transactionIDCounter++

	c.JSON(http.StatusOK, gin.H{"message": "Payment successful", "transaction": newTransaction})
}

func GetTransactions(c *gin.Context) {
	transactionList := make([]Transaction, 0)
	for _, transaction := range transactions {
		transactionList = append(transactionList, transaction)
	}

	c.JSON(http.StatusOK, transactionList)
}
