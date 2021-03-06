package controllers

import (
	"merchant-api/db"
	"merchant-api/models"
	"merchant-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransaction(c *gin.Context) {
	var transactions []models.Transaction
	claim, flag := utils.ExtractClaim(c.Request.Header.Get("Authorization"))

	if !flag {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "No Claim",
		})

		return
	}

	tokenresult := claim.(map[string]interface{})

	db.Connect().Where("merchant_id = ?", tokenresult["userid"]).Find(&transactions)

	message := "Get Transaction"

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": message,
		"data":    transactions,
	})
}
