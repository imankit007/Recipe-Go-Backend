package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"sample.com/crud-api/data/models"
)

func init() {
	// initializers.LoadEnvVariables()
	// initializers.ConenctToDb()
}

func main() {
	fmt.Print("Running....")
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Ankit",
		})
	})

	r.POST("/recipe" , func(ctx *gin.Context) {
		var recipe models.Recipe ;

		if err := ctx.BindJSON(&recipe); err != nil{
			log.Panic("Parsing Failed for Request")
			return 
		}

		fmt.Print(recipe);


	 })

	r.Run("0.0.0.0:8080")
}
