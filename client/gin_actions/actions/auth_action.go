package actions

import (
	"gtkgo/client/dto"
	"gtkgo/core/adapters/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserActionCreate handles the HTTP POST request to create a new user.
// It expects a JSON payload with user details (name, email, password) in the request body.
// If the payload is valid, it will invoke the UserController to create a user and return
// a success message with the created user details. If there's an error during user creation,
// it returns an HTTP 400 status with the error message.
func AuthenticationAction(ctx *gin.Context) {
	var authRequestDTO dto.AuthRequestDTO

	// Initialize a new UserController instance
	auth := controllers.NewAuthController()

	// Bind the JSON payload to the UserDTO struct
	if err := ctx.ShouldBindJSON(&authRequestDTO); err != nil {
		// Return HTTP 400 if payload binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use UserController to handle user creation
	authType, err := auth.Authenticate(authRequestDTO.Email, authRequestDTO.Password)
	if err != nil {
		// Log the error and return HTTP 400 if user creation fails
		log.Default().Printf("Error ao autenticar: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponserAuth := dto.AuthDTO{
		Username: authType.Username,
		Email:    authType.Email,
	}

	// Return HTTP 200 with success message and created user details
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenAuthentication(userResponserAuth),
	})
}

func tokenAuthentication(user dto.AuthDTO) dto.AuthDTO {
	return user
}
