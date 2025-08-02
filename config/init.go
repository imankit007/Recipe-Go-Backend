package config

import (
	"sample.com/crud-api/controller"
	"sample.com/crud-api/data/repositories"
	"sample.com/crud-api/service"
)

type Initialization struct {
	recipeRepo repositories.RecipeRepository
	recipeSvc  service.RecipeService
	RecipeCtrl controller.RecipeController
}

func NewInitialization(recipeRepo repositories.RecipeRepository,
	recipeSvc service.RecipeService,
	recipeCtrl controller.RecipeController) *Initialization {
	return &Initialization{
		recipeRepo: recipeRepo,
		recipeSvc:  recipeSvc,
		RecipeCtrl: recipeCtrl,
	}
}
