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

func (r RecipeControllerImpl) 	AddRecipe(c *gin.Context) {
	r.svc.GetAllRecipe(c)
}

func (r RecipeControllerImpl) GetRecipeById(c *gin.Context) {
	r.svc.GetAllRecipe(c)
}

func (r RecipeControllerImpl) UpdateRecipe(c *gin.Context) {
	r.svc.GetAllRecipe(c)
}




func RecipeControllerInit(recipeSvc service.RecipeService) *RecipeControllerImpl {
	return &RecipeControllerImpl{
		svc: recipeSvc,
	}
}
