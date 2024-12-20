package services

import (
	"gtkgo/core/domain/entities"
	"gtkgo/infra/repositories"
)

type AuthService struct {
	authRepo *repositories.AuthRepository
}

func NewAuthService() *AuthService {
	repository := repositories.NewAuthRepository()
	return &AuthService{authRepo: repository}
}

func (a *AuthService) AuthenticateService(name string, password string) (*entities.User, error) {
	user, err := a.authRepo.Auth(name, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
