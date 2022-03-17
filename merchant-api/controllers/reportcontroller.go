package controllers

import (
	"merchant-api/db"
	"merchant-api/models"
	"merchant-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReport(c *gin.Context) {
	claim, flag := utils.ExtractClaim(c.Request.Header.Get("Authorization"))

	if !flag {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "No Claim",
		})

		return
	}

	tokenresult := claim.(map[string]interface{})

	result, _ := models.GetReport(db.Connect(), tokenresult["userid"].(float64))

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": "No Data Found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Get Report",
		"data":    result,
	})
}
