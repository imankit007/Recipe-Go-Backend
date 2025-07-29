package models

type User struct {
	Base
	FirstName  string
	MiddleName string
	LastName   string
	Email      string
}
