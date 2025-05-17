package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName  string
	MiddleName string
	LastName   string
}
