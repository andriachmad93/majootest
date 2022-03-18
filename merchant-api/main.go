package main

import (
	"merchant-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", controllers.Login)
	r.GET("/transactions", controllers.Auth, controllers.GetTransaction)
	r.GET("/profile", controllers.Auth, controllers.Profile)
	r.GET("/report", controllers.Auth, controllers.GetReport)
	r.Run(":8080")
}
