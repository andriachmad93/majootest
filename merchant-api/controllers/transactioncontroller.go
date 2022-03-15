package controllers

import (
	"merchant-api/db"
	"merchant-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransaction(c *gin.Context) {
	var transactions []models.Transaction

	db.Connect().Find(&transactions)

	message := "Get Transaction"

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": message,
		"data":    transactions,
	})
}
