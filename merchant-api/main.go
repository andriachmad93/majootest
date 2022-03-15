package main

import (
	"merchant-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/transactions", controllers.GetTransaction)
	r.Run()
}
