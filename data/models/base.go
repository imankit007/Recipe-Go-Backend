package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID 		`gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time 		`json:"-"` 
	UpdatedAt time.Time 		`json:"-"`
	DeletedAt gorm.DeletedAt 	`gorm:"index;"`
}
