package actions

import (
	"gtkgo/client/dto"
	"gtkgo/core/adapters/controllers"
	"gtkgo/helpers"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationAction(ctx *fiber.Ctx) error {
	var authRequestDTO dto.AuthRequestDTO

	// Initialize a new UserController instance
	auth := controllers.NewAuthController()

	// Bind the JSON payload to the UserDTO struct
	if err := ctx.BodyParser(&authRequestDTO); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(helpers.LogError{Error: err.Error()})
	}

	// Use UserController to handle user creation
	authType, err := auth.Authenticate(authRequestDTO.Email, authRequestDTO.Password)
	if err != nil {
		// Log the error and return HTTP 400 if user creation fails
		log.Default().Printf("Error ao autenticar: %v\n", err)
		return ctx.Status(http.StatusBadRequest).JSON(helpers.LogError{Error: err.Error()})
	}

	userResponserAuth := dto.AuthDTO{
		Username: authType.Username,
		Email:    authType.Email,
	}

	// Return HTTP 200 with success message and created user details
	return ctx.Status(http.StatusOK).JSON(map[string]interface{}{"token": tokenAuthentication(userResponserAuth)})
}

func tokenAuthentication(user dto.AuthDTO) dto.AuthDTO {
	return user
}
