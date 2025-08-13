package models

type Recipe struct {
	Base
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	PrepTimeMin  int           `json:"prep_time"`
	CookTimeMin  int           `json:"cook_time"`
	Servings     int           `json:"servings"`
	ImageURL     string        `json:"image_url"`
	Ingredients  []Ingredient  `gorm:"many2many:recipe_ingredients;foreignKey:ID" json:"ingredients"`
	Categories   []Category    `gorm:"many2many:recipe_categories;foreignKey:ID" json:"categories"`
	Instructions []Instruction `gorm:"foreignKey:RecipeID" json:"instrcutions"`
}
