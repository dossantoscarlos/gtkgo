package controllers

import (
	"gtkgo/core/domain/entities"

	"gtkgo/infra/services"
)

type AuthController struct {
	auth_service *services.AuthService
}

func NewAuthController() *AuthController {
	auth_service := services.NewAuthService()
	return &AuthController{auth_service: auth_service}
}

func (auth *AuthController) Authenticate(name string, password string) (*entities.User, error) {

	user, err := auth.auth_service.AuthenticateService(name, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
