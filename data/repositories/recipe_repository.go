package repositories

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sample.com/crud-api/data/models"
	"sample.com/crud-api/domain/dto"
	"sample.com/crud-api/utils/converters"
	"sample.com/crud-api/utils/logger"
)

type RecipeRepository interface {
	FindAllRecipe() ([]models.Recipe, error)
	Save(recipe *dto.Recipe) (models.Recipe, error)
	FindById(id int) (models.Recipe, error)
	DeleteById(id int) error
}

type RecipeRepositoryImpl struct {
	db *gorm.DB
}

func (r RecipeRepositoryImpl) FindAllRecipe() ([]models.Recipe, error) {

	var recipes []models.Recipe

	err := r.db.Preload("Instructions").Preload("Categories").Preload("Ingredients").Find(&recipes).Error
	if err != nil {
		log.Error("Error fetching Recipes", err)
		return nil, err
	}

	return recipes, nil
}

func (r RecipeRepositoryImpl) Save(recipe *dto.Recipe) (models.Recipe, error) {

	log.Info("Save Recipe : ")
	log.Print(recipe)

	var newRecipe *models.Recipe = converters.RecipeDaoToModel(*recipe)

	if err := r.db.Omit("Ingredients.*", "Categories.*").Create(&newRecipe).Error; err != nil {
		logger.Log.WithFields(log.Fields{
			"recipe": newRecipe,
		}).Error("Error saving recipe : ")
		return models.Recipe{}, err
	}

	return *newRecipe, nil
}

func (r RecipeRepositoryImpl) FindById(id int) (models.Recipe, error) {

	log.Info("Get Recipe : ", id)

	var recipe models.Recipe = models.Recipe{}
	return recipe, nil
}

func (r RecipeRepositoryImpl) DeleteById(id int) error {

	log.Info("Delete Recipe : ", id)

	return nil
}

func RecipeRepositoryInit(db *gorm.DB) *RecipeRepositoryImpl {
	db.AutoMigrate(&models.Recipe{})
	return &RecipeRepositoryImpl{
		db: db,
	}
}
