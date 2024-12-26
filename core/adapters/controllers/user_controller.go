package controllers

import (
	"gtkgo/core/domain/entities"
)

type IUserService interface {
	GetAllUsersService() ([]entities.User, error)
	GetOneUserService(id int) (entities.User, error)
	CreateUserService(name string, email string, password string) (id int, err error)
	UpdateUserService(id string, user entities.User) (entities.User, error)
	DeleteUserService(id string) error
}

type UserController struct {
	user_service IUserService
}

func NewUserController(userService IUserService) *UserController {
	return &UserController{user_service: userService}
}

func (uc *UserController) HandleCreateUser(name string, email string, pass string) (int, error) {
	// Chama o use case para criar o usuário
	id, err := uc.user_service.CreateUserService(name, email, pass)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetAllUsers retrieves all users from the user service.
// It returns a slice of User entities and an error, if any occurs.
func (uc *UserController) GetAllUsers() ([]entities.User, error) {

	user, err := uc.user_service.GetAllUsersService()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserController) GetOneUser(id int) (entities.User, error) {
	user, err := uc.user_service.GetOneUserService(id)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
