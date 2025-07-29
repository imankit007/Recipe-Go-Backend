package models

import "github.com/google/uuid"

type Instruction struct {
	Base
	RecipeID   uuid.UUID
	StepNumber uint
	StepText   string
}
