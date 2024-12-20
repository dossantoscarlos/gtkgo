package controllers

import (
	"gtkgo/core/domain/entities"
	"gtkgo/infra/services"
)

type UserController struct {
	user_service *services.UserService
}

func NewUserController() *UserController {
	user_service := services.NewUserService()
	return &UserController{user_service: user_service}
}

// HandleCreateUser handles the creation of a new user by first hashing the password
// and then calling the user service to create the user. It returns the created user
// entity and an error if any occurs during the process.
// Parameters:
//   - name: the name of the user
//   - email: the email of the user
//   - pass: the raw password of the user
//
// Returns:
//   - result: the created User entity
//   - err: an error if the password hashing or user creation fails
func (uc *UserController) HandleCreateUser(name string, email string, pass string) (result *entities.User, err error) {
	// Chama o use case para criar o usuário
	result, err = uc.user_service.CreateUserService(name, email, pass)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAllUsers retrieves all users from the user service.
// It returns a slice of User entities and an error, if any occurs.
func (uc *UserController) GetAllUsers() ([]entities.User, error) {
	return uc.user_service.GetAllUsersService()
}
