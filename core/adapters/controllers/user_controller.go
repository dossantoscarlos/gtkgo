package controllers

import (
	"gtkgo/core/adapters/dto"
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

func (uc *UserController) HandleCreateUser(name string, email string, pass string) (userID *dto.UserCreateResponse, err error) {
	// Chama o use case para criar o usuário
	id, err := uc.user_service.CreateUserService(name, email, pass)
	if err != nil {
		return nil, err
	}
	userID = &dto.UserCreateResponse{ID: int(id)}

	return userID, nil
}

// GetAllUsers retrieves all users from the user service.
// It returns a slice of User entities and an error, if any occurs.
func (uc *UserController) GetAllUsers() ([]dto.UserDtoResponse, error) {

	var userResponse []dto.UserDtoResponse

	user, err := uc.user_service.GetAllUsersService()
	if err != nil {
		return nil, err
	}

	for _, u := range user {
		userResponse = append(userResponse, dto.UserDtoResponse{
			Name:  u.Username,
			Email: u.Email,
		})
	}

	return userResponse, nil
}

func (uc *UserController) GetOneUser(id int) (entities.User, error) {
	user, err := uc.user_service.GetOneUserService(id)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
