package services

import (
	"gtkgo/core/domain/entities"
	"gtkgo/helpers"
	"gtkgo/infra/interfaces"
	"log"
)

type UserService struct {
	userRepo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (us *UserService) CreateUserService(name string, email string, password string) (int, error) {

	hashPassword, err := helpers.HashPassword(password)
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}

	user := entities.User{
		ID:       0,
		Username: name,
		Email:    email,
		Password: hashPassword,
	}

	id, err := us.userRepo.CreateUser(user)
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}

	return id, nil
}

func (us *UserService) GetAllUsersService() ([]entities.User, error) {
	return us.userRepo.GetAllUsers()
}

func (us *UserService) GetOneUserService(id int) (entities.User, error) {
	user, err := us.userRepo.GetUserById(id)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (us *UserService) UpdateUserService(id string, user entities.User) (entities.User, error) {
	return us.userRepo.UpdateUser(id, user)
}

func (us *UserService) DeleteUserService(id string) error {
	return us.userRepo.DeleteUser(id)
}
