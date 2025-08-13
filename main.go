package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	var count int16
	init := config.Init()
	app := router.Init(init)

	// url := ginSwagger.URL("http://localhost:9000/swagger/doc.json")

	app.GET("/ping", func(ctx *gin.Context) {
		count++
		ctx.JSON(200, gin.H{
			"message": count,
		})
	})


	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	err := app.Run("0.0.0.0:8080")

	if err != nil{
		log.Fatal(err)
	}

}
