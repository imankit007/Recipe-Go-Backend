package main

import (
	"log"

	"sample.com/crud-api/data/models"
	"sample.com/crud-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConenctToDb()

}

func main() {
	var err error
	log.Default().Println("Start Migrate Database...")
	err = initializers.DB.AutoMigrate(&models.User{}, &models.Category{},&models.Ingredient{},&models.Recipe{},&models.Instruction{})
	if err != nil {
		log.Fatal("Migration Failed")
	}
	log.Println("Migration Complete!!!")
}
