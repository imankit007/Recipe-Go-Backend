package service

import (
	"github.com/gin-gonic/gin"
	"sample.com/crud-api/data/repositories"
)

type RecipeService interface {
	GetAllRecipe(c *gin.Context)
	AddRecipe(c *gin.Context)
	GetRecipeById(c *gin.Context)
	UpdateRecipe(c *gin.Context)
}

type RecipeServiceImpl struct {
	recipeRepo repositories.UserRepository
}

func (r RecipeServiceImpl) GetAllRecipe(c *gin.Context) {

}
