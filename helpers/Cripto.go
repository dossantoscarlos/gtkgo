package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Criptografa a senha usando bcrypt
func HashPassword(password string) (string, error) {
	// Gerar o hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Verifica se a senha fornecida corresponde ao hash armazenado
func CheckPasswordHash(password []byte, hashedPassword string) bool {
	// Comparar a senha com o hash armazenado
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())
		return false
	}
	return true // Senha correta
}
