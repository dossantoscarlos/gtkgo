package actions

import (
	"fmt"
	"gtkgo/core/adapters/controllers"
	"gtkgo/core/adapters/dto"
	"gtkgo/helpers"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// UserActionCreate handles the HTTP POST request to create a new user.
// It expects a JSON payload with user details (name, email, password) in the request body.
// If the payload is valid, it will invoke the UserController to create a user and return
// a success message with the created user details. If there's an error during user creation,
// it returns an HTTP 400 status with the error message.
func UserActionCreate(ctx *fiber.Ctx) error {
	var userDTO dto.UserDTO

	// Initialize a new UserController instance
	user := controllers.NewUserController()

	// Bind the JSON payload to the UserDTO struct
	if err := ctx.BodyParser(&userDTO); err != nil {
		// Return HTTP 400 if payload binding fails
		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: err.Error()})
	}

	// Use UserController to handle user creation
	userType, err := user.HandleCreateUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		// Log the error and return HTTP 400 if user creation fails
		log.Default().Printf("Error ao criar usuário: %v", err)
		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: err.Error()})
	}

	// Return HTTP 200 with success message and created user details
	//ctx.JSON(http.StatusOK, gin.H{"message": "Usuário criado com sucesso", "user": userType})

	return ctx.Status(http.StatusOK).JSON(userType)
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
func UserActionGetAll(ctx *fiber.Ctx) error {
	// Initialize a new UserController instance
	user := controllers.NewUserController()

	// Use UserController to fetch all users
	users, err := user.GetAllUsers()
	if err != nil {
		// Log the error and return HTTP 400 if user retrieval fails
		log.Default().Printf("Error ao buscar usuários: %v", err)

		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: err.Error()})

	}

	return ctx.Status(http.StatusOK).JSON(users)
}

func GetOneUsers(ctx *fiber.Ctx) error {
	param := ctx.Query("id", "")

	if param == "" {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: "Usuário inválido"})
	}

	fmt.Printf("param: %v\n", param)

	user := controllers.NewUserController()

	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: err.Error()})
	}

	userType, err := user.GetOneUser(id)
	if err != nil {
		// Log the error and return HTTP 400 if user creation fails
		log.Default().Printf("Error ao buscar usuários: %v", err)
		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(userType)
}
