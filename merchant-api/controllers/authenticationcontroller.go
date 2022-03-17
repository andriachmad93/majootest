package controllers

import (
	"fmt"
	"net/http"
	"time"

	"merchant-api/db"
	"merchant-api/models"
	"merchant-api/utils"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
)

func Auth(c *gin.Context) {
	tokenstring := c.Request.Header.Get("Authorization")

	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) ([]byte, error) {
		if jwt.GetSigningMethod("HS256") != t.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token == nil && err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		})

		return
	}
}

func Login(c *gin.Context) {
	var user models.User

	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "cannot bind struct",
		})
		c.Abort()
	}

	// get user data from email
	userdata, _ := models.FindByEmail(db.Connect(), user.Username)

	if user.Username != userdata.Username {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "wrong username",
		})

		return
	}

	if utils.GetMD5Hash(user.Password) != userdata.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "wrong password",
		})

		return
	}

	sign := jwt.New(jwt.GetSigningMethod("HS256"))

	sign.Claims["userid"] = userdata.ID
	sign.Claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	token, err := sign.SignedString([]byte("secret"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Profile(c *gin.Context) {
	var profile []models.User
	var merchant []models.Merchant
	//var data map[string]interface{}

	claim, flag := utils.ExtractClaim(c.Request.Header.Get("Authorization"))

	if !flag {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "No Claim",
		})

		return
	}

	resultclaim := claim.(map[string]interface{})

	db.Connect().Where("id = ?", resultclaim["userid"]).Find(&profile)

	if len(profile) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": "No Data Found",
		})

		return
	}

	db.Connect().Where("user_id = ?", resultclaim["userid"]).Find(&merchant)

	//data["profile"] = profile
	//data["merchant"] = merchant

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Get Data Profile",
		"data":    merchant,
	})
}
