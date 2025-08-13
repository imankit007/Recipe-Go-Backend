package dto

type Recipe struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	PrepTimeMin  int    `json:"prep_time"`
	CookTimeMin  int    `json:"cook_time"`
	Servings     int    `json:"servings"`
	ImageURL     string `json:"image_url"`
	Ingredients  []string
	Categories   []string
	Instructions []Instruction
}
