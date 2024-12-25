package controllers_test

import (
	"database/sql"
	"fmt"
	"gtkgo/core/adapters/controllers"
	"gtkgo/core/domain/entities"
	"log"
	"os"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MyTestSuite struct {
	suite.Suite
	mock *MockUserService
	db   *sql.DB
	tx   *sql.Tx
}

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUserService(name string, email string, password string) (int, error) {
	args := m.Called(name, email, password)
	return args.Int(0), args.Error(1)
}

func (m *MockUserService) GetOneUserService(id int) (entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockUserService) GetAllUsersService() ([]entities.User, error) {
	args := m.Called()
	return args.Get(0).([]entities.User), args.Error(1)
}

func (m *MockUserService) UpdateUserService(id string, user entities.User) (entities.User, error) {
	args := m.Called(id, user)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockUserService) DeleteUserService(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (suite *MyTestSuite) SetupTest() {
	mockUserService := new(MockUserService)

	mockUserService.On("CreateUserService", "John Doe", "johndoe@example.com", "password123").Return(1, nil)

	mockUserService.On("GetOneUserService", 1).Return(entities.User{
		ID:       1,
		Username: "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
	}, nil)

	mockUserService.On("GetAllUsersService").Return([]entities.User{
		{ID: 1, Username: "John Doe", Email: "john.doe@example.com", Password: "password123"},
		{ID: 2, Username: "Jane Doe", Email: "jane.doe@example.com", Password: "password456"},
	}, nil)

	mockUserService.On("UpdateUserService", "1", entities.User{
		ID:       1,
		Username: "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
	}).Return(entities.User{
		ID:       1,
		Username: "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
	}, nil)

	mockUserService.On("DeleteUserService", "1").Return(nil)

	suite.mock = mockUserService
}

func (suite *MyTestSuite) TestCreateUser() {
	controller := controllers.NewUserController(suite.mock)

	data, err := controller.HandleCreateUser("John Doe", "johndoe@example.com", "password123")

	fmt.Println(data.ID)

	assert.NoError(suite.T(), err)
	assert.EqualValues(suite.T(), 1, data.ID)
}

func (suite *MyTestSuite) TestGetOneUser() {
	controller := controllers.NewUserController(suite.mock)

	user, err := controller.GetOneUser(1)

	fmt.Println(user)

	assert.NoError(suite.T(), err)
	assert.EqualValues(suite.T(), "John Doe", user.Username)
}

func (suite *MyTestSuite) TestGetAllUsers() {
	controller := controllers.NewUserController(suite.mock)

	users, err := controller.GetAllUsers()

	fmt.Println(users)

	assert.NoError(suite.T(), err)
	assert.EqualValues(suite.T(), 2, len(users))

	for _, user := range users {
		entity := reflect.ValueOf(user)

		password := entity.FieldByName("Password")
		if password.IsValid() {
			log.Fatalln("password is valid")
		}
	}

}

func TestMyTestSuite(t *testing.T) {
	suite.Run(t, new(MyTestSuite))
	arq := "./database.db"
	if _, err := os.Stat(arq); err == nil {
		os.Remove(arq)
	}
}
