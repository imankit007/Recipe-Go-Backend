package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPass), err
}


func CheckPasswordHash(password string, hash string)  (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash) , []byte(password))
	return err == nil, err
}