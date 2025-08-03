package repositories

import (
	"gorm.io/gorm"
	"sample.com/crud-api/data/models"
)

type UserRepository interface {
	GetUserByID(ID string) (models.User, error)

	GetUserByEmail(email string) (models.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u *UserRepositoryImpl) GetUserByID(ID string) (models.User, error) {
	var user models.User

	res := u.db.First(&user, ID)

	return user, res.Error
}

func (u *UserRepositoryImpl) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	res := u.db.Where("email = ?", email).First(&user)

	return user, res.Error
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&models.User{})
	return &UserRepositoryImpl{
		db: db,
	}
}
