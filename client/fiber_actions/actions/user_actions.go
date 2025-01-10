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

// @Summary		Show an account
// @Description	get string by ID
// @ID				get-string-by-int
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Account ID"
// @Success		200	{object}	model.Account
// @Failure		400	{object}	http.Response
// @Failure		404	{object}	http.Response
// @Router			/accounts/{id} [get]
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

// @Summary		Get all users
// @Description	Retrieve the list of all users
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]interface{}
// @Failure		400	{object}	map[string]interface{}
// @Router			/users [get]
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
