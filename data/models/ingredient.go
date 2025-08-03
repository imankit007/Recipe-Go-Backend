package models

type Ingredient struct {
	Base
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}
