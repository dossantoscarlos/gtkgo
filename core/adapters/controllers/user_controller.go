package controllers

import (
	"gtkgo/core/domain/entities"
	"gtkgo/infra/interfaces"
)

type UserController struct {
	user_service interfaces.IUserService
}

func NewUserController(userService interfaces.IUserService) *UserController {
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
func (uc *UserController) GetAllUsers() ([]entities.UserNotPassword, error) {
	var userType []entities.UserNotPassword

	user, err := uc.user_service.GetAllUsersService()
	if err != nil {
		return nil, err
	}

	for _, u := range user {
		userType = append(userType, entities.UserNotPassword{
			ID:    u.ID,
			Name:  u.Username,
			Email: u.Email,
		})
	}

	return userType, nil
}

func (uc *UserController) GetOneUser(id int) (entities.UserNotPassword, error) {
	var userType entities.UserNotPassword

	user, err := uc.user_service.GetOneUserService(id)

	if err != nil {
		return entities.UserNotPassword{}, err
	}

	userType = entities.UserNotPassword{
		ID:    user.ID,
		Name:  user.Username,
		Email: user.Email,
	}

	return userType, nil
}

func (uc *UserController) UserDelete(id int) error {
	err := uc.user_service.DeleteUserService(id)
	if err != nil {
		return err
	}

	return nil
}
