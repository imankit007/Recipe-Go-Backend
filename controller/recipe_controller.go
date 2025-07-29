package controller

import (
	"github.com/gin-gonic/gin"
	"sample.com/crud-api/service"
)

type RecipeController interface {
	GetAllRecipe(c *gin.Context)
	AddRecipe(c *gin.Context)
	GetRecipeById(c *gin.Context)
	UpdateRecipe(c *gin.Context)
}

type RecipeControllerImpl struct {
	svc service.RecipeService
}

func (r RecipeControllerImpl) GetAllRecipe(c *gin.Context) {
	r.svc.GetAllRecipe(c)
}
