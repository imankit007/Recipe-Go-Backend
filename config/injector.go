//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"sample.com/crud-api/controller"
	"sample.com/crud-api/data/repositories"
	"sample.com/crud-api/initializers"
	"sample.com/crud-api/service"
)

var db = wire.NewSet(initializers.ConenctToDb)

var recipeServiceSet = wire.NewSet(service.RecipeServiceInit,
	wire.Bind(new(service.RecipeService), new(*service.RecipeServiceImpl)),
)

var recipeRepoSet = wire.NewSet(repositories.RecipeRepositoryInit,
	wire.Bind(new(repositories.RecipeRepository), new(*repositories.RecipeRepositoryImpl)),
)

var recipeCtrlSet = wire.NewSet(controller.RecipeControllerInit,
	wire.Bind(new(controller.RecipeController), new(*controller.RecipeControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, recipeRepoSet, recipeServiceSet, recipeCtrlSet)
	return nil
}
