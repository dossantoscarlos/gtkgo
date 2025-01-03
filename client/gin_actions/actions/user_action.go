package actions

import (
	"gtkgo/client/dto"
	"gtkgo/core/adapters/controllers"
	"gtkgo/infra/repositories"
	"gtkgo/infra/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserActionCreate handles the HTTP POST request to create a new user.
// It expects a JSON payload with user details (name, email, password) in the request body.
// If the payload is valid, it will invoke the UserController to create a user and return
// a success message with the created user details. If there's an error during user creation,
// it returns an HTTP 400 status with the error message.
func UserActionCreate(ctx *gin.Context) {
	var userDTO dto.UserDTO

	// Initialize a new UserController instance
	user := controllers.NewUserController(services.NewUserService(repositories.NewUserRepository()))

	// Bind the JSON payload to the UserDTO struct
	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		// Return HTTP 400 if payload binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use UserController to handle user creation
	userType, err := user.HandleCreateUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		// Log the error and return HTTP 400 if user creation fails
		log.Default().Printf("Error ao criar usuário: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponseCreate := dto.UserCreateResponse{
		ID: userType,
	}

	// Return HTTP 200 with success message and created user details
	ctx.JSON(http.StatusOK, gin.H{"user": userResponseCreate})
}

// UserActionGetAll handles the HTTP GET request to retrieve all users.
// It invokes the UserController to fetch all users and returns an HTTP 200 status
// with a JSON payload containing the list of users. If there's an error during
// user retrieval, it logs the error and returns an HTTP 400 status with the error
// message.
//
// @Summary Get all users
// @Description Retrieve the list of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /users [get]
func UserActionGetAll(ctx *gin.Context) {
	var userResponse []dto.UserDtoResponse

	// Initialize a new UserController instance
	user := controllers.NewUserController(services.NewUserService(repositories.NewUserRepository()))

	// Use UserController to fetch all users
	users, err := user.GetAllUsers()
	if err != nil {
		// Log the error and return HTTP 400 if user retrieval fails
		log.Default().Printf("Error ao buscar usuários: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, user := range users {
		userResponse = append(userResponse, dto.UserDtoResponse{
			Name:  user.Name,
			Email: user.Email,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"users": userResponse})
}

func GetOneUsers(ctx *gin.Context) {

	var userResponse dto.UserDtoResponse

	param := ctx.Query("id")

	if param == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Usuário inválido"})
		return
	}

	controllers := controllers.NewUserController(services.NewUserService(repositories.NewUserRepository()))

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Default().Printf("Error ao buscar usuários: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Usuário inválido"})
		return
	}

	userType, err := controllers.GetOneUser(id)
	if err != nil {
		// Log the error and return HTTP 400 if user creation fails
		log.Default().Printf("Error ao buscar usuários: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userResponse = dto.UserDtoResponse{
		Name:  userType.Name,
		Email: userType.Email,
	}

	ctx.JSON(http.StatusOK, gin.H{"user": userResponse})
}
