package main

import (
	"log"

	"sample.com/crud-api/initializers"
	"sample.com/crud-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConenctToDb()

}

func main() {
	var err error
	log.Default().Println("Start Migrate Database...")
	log.Println("Migrating model User...")
	err = initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration Failed")
	}
	log.Println("Migration Complete!!!")
}
