package controllers

import (
	"gtkgo/core/adapters/dto"

	"gtkgo/infra/services"
)

type AuthController struct {
	auth_service *services.AuthService
}

func NewAuthController() *AuthController {
	auth_service := services.NewAuthService()
	return &AuthController{auth_service: auth_service}
}

func (auth *AuthController) Authenticate(name string, password string) (*dto.AuthDTO, error) {

	user, err := auth.auth_service.AuthenticateService(name, password)
	if err != nil {
		return nil, err
	}

	response := dto.AuthDTO{
		Email:    user.Email,
		Username: user.Username,
	}

	return &response, nil
}
