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

	AuthController controller.AuthController
	authService service.AuthService
	userRepo repositories.UserRepository

}

func NewInitialization(recipeRepo repositories.RecipeRepository,
	recipeSvc service.RecipeService,
	recipeCtrl controller.RecipeController,
	userRepo repositories.UserRepository,
	authSvc service.AuthService,
	authCtrl controller.AuthController,
	) *Initialization {
	return &Initialization{
		recipeRepo: recipeRepo,
		recipeSvc:  recipeSvc,
		RecipeCtrl: recipeCtrl,
		userRepo: userRepo,
		authService: authSvc,
		AuthController: authCtrl,
	}
}
