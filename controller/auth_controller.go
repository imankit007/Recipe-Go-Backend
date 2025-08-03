package controller

import (
	"github.com/gin-gonic/gin"
	"sample.com/crud-api/service"
)

type AuthController interface {
	Login(c *gin.Context)
}

type AuthControllerImpl struct {
	authSvc service.AuthService
}

func (a *AuthControllerImpl) Login(c *gin.Context) {
	panic("TODO: Implement")
}

func AuthControllerInit(authSvc service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		authSvc: authSvc,
	}
}
