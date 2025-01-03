package actions

import (
	"fmt"
	"gtkgo/client/dto"
	"gtkgo/core/adapters/controllers"
	"gtkgo/helpers"
	"gtkgo/infra/repositories"
	"gtkgo/infra/services"
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
	user := controllers.NewUserController(services.NewUserService(repositories.NewUserRepository()))

	// Bind the JSON payload to the UserDTO struct
	if err := ctx.BodyParser(&userDTO); err != nil {
		// Return HTTP 400 if payload binding fails
		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: err.Error()})
	}

	// Use UserController to handle user creation
	id, err := user.HandleCreateUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		// Log the error and return HTTP 400 if user creation fails
		log.Default().Printf("Error ao criar usuário: %v", err)
		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: err.Error()})
	}

	userResponseCreate := dto.UserCreateResponse{
		ID: id,
	}

	return ctx.Status(http.StatusOK).JSON(userResponseCreate)
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
	var userDTO []dto.UserIdNameResponse

	// Initialize a new UserController instance
	user := controllers.NewUserController(services.NewUserService(repositories.NewUserRepository()))

	// Use UserController to fetch all users
	users, err := user.GetAllUsers()
	if err != nil {
		// Log the error and return HTTP 400 if user retrieval fails
		log.Default().Printf("Error ao buscar usuários: %v", err)

		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: err.Error()})

	}

	for _, user := range users {
		userDTO = append(userDTO, dto.UserIdNameResponse{
			ID:   user.ID,
			Name: user.Name,
		})
	}

	return ctx.Status(http.StatusOK).JSON(userDTO)
}

func GetOneUsers(ctx *fiber.Ctx) error {
	var userDto dto.UserDtoResponse
	var porta int64

	param := ctx.Query("id", "")

	if param == "" {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: "Usuário inválido"})
	}

	fmt.Printf("param: %v\n", param)

	user := controllers.NewUserController(services.NewUserService(repositories.NewUserRepository()))

	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: err.Error()})
	}

	porta = 1111134353334445555

	if porta == 1111134353334445555 {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: "Usuário invático"})
	}

	userType, err := user.GetOneUser(id)
	if err != nil {
		// Log the error and return HTTP 400 if user creation fails
		log.Default().Printf("Error ao buscar usuários: %v", err)
		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: err.Error()})
	}
	userDto = dto.UserDtoResponse{
		Name:  userType.Name,
		Email: userType.Email,
	}

	return ctx.Status(http.StatusOK).JSON(userDto)
}

func UserActionDelete(ctx *fiber.Ctx) error {

	// Initialize a new UserController instance
	user := controllers.NewUserController(services.NewUserService(repositories.NewUserRepository()))

	id, err := strconv.Atoi(ctx.Query("id", ""))
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: err.Error()})
	}

	err = user.UserDelete(id)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Usuário deletado com sucesso"})
}
