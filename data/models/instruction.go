package models

import "github.com/google/uuid"

type Instruction struct {
	Base
	RecipeID   uuid.UUID  `gorm:"type:uuid"`
	StepNumber uint
	StepText   string
}
