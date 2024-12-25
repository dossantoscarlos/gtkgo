package controllers

import (
	"gtkgo/core/adapters/dto"
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
