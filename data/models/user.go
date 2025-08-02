package models

type User struct {
	Base
	FirstName  string
	MiddleName string
	LastName   string
	Username   string `gorm:"unique"`
	Email      string
	Password   string
}
