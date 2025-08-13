package converters

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"sample.com/crud-api/data/models"
	"sample.com/crud-api/domain/dto"
	"sample.com/crud-api/utils/logger"
)

func RecipeDaoToModel(recipe dto.Recipe) *models.Recipe {

	var ingredients []models.Ingredient
	var categories []models.Category

	for _, ingredientId := range recipe.Ingredients {

		id, err := uuid.Parse(ingredientId)

		if err != nil {
			logger.Log.WithFields(log.Fields{
				"id": ingredientId,
			}).Error("Error parsing ID : ")
		}

		ingredients = append(ingredients, models.Ingredient{
			Base: models.Base{
				ID: id,
			},
		})
	}

	for _, categoryId := range recipe.Categories {

		id, err := uuid.Parse(categoryId)

		if err != nil {
			logger.Log.WithFields(log.Fields{
				"id": id,
			}).Error("Error parsing ID : ")
		}

		categories = append(categories, models.Category{
			Base: models.Base{
				ID: id,
			},
		})
	}

	var instructions []models.Instruction

	for step, instruction := range recipe.Instructions {

		instructions = append(instructions, models.Instruction{
			StepNumber: uint(step),
			StepText:   instruction.StepText,
		})
	}

	return &models.Recipe{
		Title:        recipe.Title,
		Description:  recipe.Description,
		PrepTimeMin:  recipe.PrepTimeMin,
		CookTimeMin:  recipe.CookTimeMin,
		Servings:     recipe.Servings,
		ImageURL:     recipe.ImageURL,
		Ingredients:  ingredients,
		Instructions: instructions,
		Categories:   categories,
	}

}
