package interfaces

import "gtkgo/core/domain/entities"

type UserRepository interface {
	GetAllUsers() ([]entities.User, error)
	GetUserById(id int) (entities.User, error)
	CreateUser(user entities.User) (int, error)
	UpdateUser(id string, user entities.User) (entities.User, error)
	DeleteUser(id string) error
}

type UserService interface {
	GetAllUsersService() ([]entities.User, error)
	GetOneUserService(id int) (entities.User, error)
	CreateUserService(name string, email string, password string) (id int, err error)
	UpdateUserService(id string, user entities.User) (entities.User, error)
	DeleteUserService(id string) error
}

type AuthService interface {
	AuthenticateService(name string, password string) (*entities.User, error)
}

type AuthRepository interface {
	Auth(email string, password string) (*entities.User, error)
}
