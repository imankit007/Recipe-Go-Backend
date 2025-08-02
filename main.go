package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"sample.com/crud-api/config"
	"sample.com/crud-api/initializers"
	"sample.com/crud-api/router"
)

func init() {
	initializers.LoadEnvVariables()
	config.InitLog()
}

func main() {
	fmt.Print("Running....")

	init := config.Init()
	app := router.Init(init)

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Ankit",
		})
	})

	app.Run("0.0.0.0:8080")
}
