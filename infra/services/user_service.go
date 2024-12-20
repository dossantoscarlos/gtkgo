package services

import (
	"fmt"
	"gtkgo/core/domain/entities"
	"gtkgo/helpers"
	"gtkgo/infra/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService() *UserService {
	repository := repositories.NewUserRepository()
	return &UserService{userRepo: repository}
}

func (us *UserService) CreateUserService(name string, email string, password string) (*entities.User, error) {
	var err error

	hashPassword, err := helpers.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := entities.User{
		ID:       1,
		Username: name,
		Email:    email,
		Password: hashPassword,
	}

	fmt.Println(user)

	err = us.userRepo.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserService) GetAllUsersService() ([]entities.User, error) {
	return us.userRepo.GetAllUsers()
}
