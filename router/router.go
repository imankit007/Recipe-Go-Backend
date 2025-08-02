package router

import (
	"github.com/gin-gonic/gin"
	"sample.com/crud-api/config"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	{
		user := api.Group("/user")
		user.GET("", func(c *gin.Context) {
			c.Writer.WriteString("Users API")
		})

		recipe := api.Group("/recipe")

		recipe.GET("", init.RecipeCtrl.GetAllRecipe)

	}

	return router

}
