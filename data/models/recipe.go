package models

type Recipe struct {
	Base
	Title        string `json:"title"`
	Description  string `json:"Description"`
	PrepTimeMin  int
	CookTimeMin  int
	Servings     int
	ImageURL     string
	Ingredients  []Ingredient  `gorm:"many2many:recipe_ingredients;foreignKey:ID"`
	Categories   []Category    `gorm:"many2many:recipe_categories;foreignKey:ID"`
	Instructions []Instruction `gorm:"foreignKey:RecipeID"`
}
