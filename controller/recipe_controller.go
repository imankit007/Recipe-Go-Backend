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


// Recipe Controller godoc
//@BasePath /api/v1/recipe
// Recipe godoc
// @Summary Get All Recipes
// @ID get-all-recipes
// @Schemes
// @Description Get all Recipes
// @Tags Recipe
// @Accept json
// @Produce json
// @Success 200 {array} models.Recipe
// @Router /recipe [get]
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
