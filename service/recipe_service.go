package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"sample.com/crud-api/data/repositories"
	"sample.com/crud-api/domain/dto"
)

type RecipeService interface {
	GetAllRecipe(c *gin.Context)
	AddRecipe(c *gin.Context)
	GetRecipeById(c *gin.Context)
	UpdateRecipe(c *gin.Context)
}

type RecipeServiceImpl struct {
	recipeRepo repositories.RecipeRepository
}

func (r RecipeServiceImpl) GetAllRecipe(c *gin.Context) {
	log.Info("Getting All Recipe")
	data, err := r.recipeRepo.FindAllRecipe()

	if err != nil {
		log.Error("Error fetching all recipes")

	}
	c.JSON(http.StatusOK, data)
}

func (r RecipeServiceImpl) AddRecipe(c *gin.Context) {

	var recipe *dto.Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	savedRecipe, err := r.recipeRepo.Save(recipe)
	if err != nil {
		log.WithFields(log.Fields{
			"recipe": recipe,
		}).Error("Error while saving recipe")
		return
	}

	c.JSON(http.StatusCreated, savedRecipe)
}

func (r RecipeServiceImpl) GetRecipeById(c *gin.Context) {

}

func (r RecipeServiceImpl) UpdateRecipe(c *gin.Context) {

}

func RecipeServiceInit(recipeRepo repositories.RecipeRepository) *RecipeServiceImpl {
	return &RecipeServiceImpl{
		recipeRepo: recipeRepo,
	}
}
