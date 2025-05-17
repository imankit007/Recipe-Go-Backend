package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"sample.com/crud-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConenctToDb()
}

func main() {
	fmt.Print("Running....")

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
