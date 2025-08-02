package repositories

import (
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sample.com/crud-api/data/models"
)

type RecipeRepository interface {
	FindAllRecipe() ([]models.Recipe, error)
	Save(recipe *models.Recipe) (models.Recipe, error)
	FindById(id int) (models.Recipe, error)
	DeleteById(id int) error
}

type RecipeRepositoryImpl struct {
	db *gorm.DB
}

func (r RecipeRepositoryImpl) FindAllRecipe() ([]models.Recipe, error) {
	var recipes []models.Recipe = []models.Recipe{
		{	
			Base: models.Base{
				ID: uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Title: "Recipe 1",	
			CookTimeMin: 10,
			Ingredients: []models.Ingredient{
				{
					Base: models.Base{
						ID: uuid.New(),
					},
					Name: "Carrot",
				},
			},
			Categories: []models.Category{
				{
				Base:models.Base{
				ID: uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "VEG",
				},
			},	
			
			Instructions: []models.Instruction{
				{
					Base: models.Base{
						ID: uuid.New(),
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),	
					},
					RecipeID: uuid.New(),
					StepText: "Take 500 ml of Water",
				},
			},
		},
		{
			Title: "Recipe 2",
		},
	}

	// var err = r.db.Find(&recipes).Error
	// if err != nil {
	// 	log.Error("Error fetching Recipes", err)
	// 	return nil, err
	// }

	return recipes, nil
}

func (r RecipeRepositoryImpl) Save(recipe *models.Recipe) (models.Recipe, error){

	log.Info("Save Recipe : ")

	log.Print(recipe)
	return *recipe, nil
} 


func (r RecipeRepositoryImpl) FindById(id int) (models.Recipe, error){

	log.Info("Get Recipe : " , id)

	var recipe models.Recipe = models.Recipe{
		
	}
	return recipe, nil
}

func (r RecipeRepositoryImpl) DeleteById(id int) (error){

	log.Info("Delete Recipe : " , id)

	return nil
}




func RecipeRepositoryInit(db *gorm.DB) *RecipeRepositoryImpl {
 db.AutoMigrate(&models.Recipe{})
 return &RecipeRepositoryImpl{
  db: db,
 }
}