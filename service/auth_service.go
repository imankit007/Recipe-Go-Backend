package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"sample.com/crud-api/data/repositories"
	"sample.com/crud-api/domain/dto"
)

type AuthService interface {
	Login(c *gin.Context)
}

type AuthServiceImpl struct {
	userRepo repositories.UserRepository
}

func (a *AuthServiceImpl) Login(c *gin.Context) {
	// TODO :

	var request dto.LoginRequest

	err := c.BindJSON(&request)
	if err != nil {
		log.Error("Error parsing Login Request Body")
	}
	user, err := a.userRepo.GetUserByEmail(request.Email)

	if err != nil {
		log.WithFields(log.Fields{
			"email": request.Email,
		}).Error("Error fetching user with email")

		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    "AUTH_FAILED",
			Message: "Invalid Email or Password",
		})
	}
	c.JSON(http.StatusOK, user)
}

func AuthServiceInit(userRepo repositories.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepo: userRepo,
	}
}
