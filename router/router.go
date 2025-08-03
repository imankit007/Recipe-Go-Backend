package router

import (
	"github.com/gin-gonic/gin"
	"sample.com/crud-api/config"
	docs "sample.com/crud-api/docs"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())


	//@BasePath /api/v1
	api := router.Group("/api/v1")
	{	
		docs.SwaggerInfo.BasePath = "/api/v1"
		authRouter := api.Group("/")

		authRouter.POST("/login", )


		user := api.Group("/user")
		user.GET("", func(c *gin.Context) {
			c.Writer.WriteString("Users API")
		})
		recipe := api.Group("/recipe")

		recipe.GET("", init.RecipeCtrl.GetAllRecipe)

	}

	return router

}
