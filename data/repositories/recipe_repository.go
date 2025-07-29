package repositories

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sample.com/crud-api/data/models"
)

type UserRepository interface {
	FindAllRecipe() ([]models.Recipe, error)
	Save(recipe *models.Recipe) (models.Recipe, error)
	FindById(id int) (models.Recipe, error)
	DeleteById(id int) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) FindAllRecipe() ([]models.Recipe, error) {
	var recipes []models.Recipe

	var err = u.db.Find(&recipes).Error
	if err != nil {
		log.Error("Error fetching Recipes", err)
		return nil, err
	}

	return recipes, nil
}
