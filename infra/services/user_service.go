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

func (us *UserService) CreateUserService(name string, email string, password string) (id int, err error) {

	hashPassword, err := helpers.HashPassword(password)
	if err != nil {
		return
	}

	user := entities.User{
		Username: name,
		Email:    email,
		Password: hashPassword,
	}

	fmt.Println(user)

	id, err = us.userRepo.CreateUser(user)
	if err != nil {
		return
	}

	return id, nil
}

func (us *UserService) GetAllUsersService() ([]entities.User, error) {
	return us.userRepo.GetAllUsers()
}

func (us *UserService) GetOneUserService(id int) (entities.User, error) {
	user, err := us.userRepo.GetOneUser(id)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
