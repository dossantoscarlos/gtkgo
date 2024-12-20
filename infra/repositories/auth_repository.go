package repositories

import (
	"database/sql"
	"fmt"
	"gtkgo/core/domain/entities"
	"gtkgo/helpers"
	"gtkgo/infra/database"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository() *AuthRepository {
	db, err := database.InitDB()

	if err != nil {
		panic(err)
	}

	return &AuthRepository{db: db}
}

func (r *AuthRepository) Auth(email string, password string) (*entities.User, error) {
	var user entities.User

	query := "SELECT id, username, email, password FROM users WHERE email = ?"

	err := r.db.QueryRow(
		query,
		email).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password)

	if err != nil {
		return nil, err
	}

	fmt.Printf("campo: %v\n", user)

	fmt.Printf("campo: %v\n", user.Password)

	fmt.Printf("campo: %v\n", password)

	if !helpers.CheckPasswordHash([]byte(password), user.Password) {
		return nil, fmt.Errorf("Senha ou email incorretos")
	}

	return &user, nil
}
