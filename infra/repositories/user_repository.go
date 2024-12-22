package repositories

import (
	"database/sql"
	"fmt"
	"gtkgo/core/domain/entities"
	"gtkgo/infra/database"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	db, err := database.InitDB()

	if err != nil {
		panic(err)
	}

	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]entities.User, error) {
	var users []entities.User
	rows, err := r.db.Query("SELECT id, username, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) CreateUser(user entities.User) (id int, err error) {

	fmt.Printf("Inserindo usuário: %v\n", user) // logando inseri usuario

	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	result, err := r.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return
	}

	fmt.Printf("Usuário inserido com sucesso: %v\n", result)

	// Verifica se a execução afetou alguma linha (se o banco foi alterado)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return
	}

	id = int(lastInsertId)

	defer r.db.Close()

	// Exibe o número de linhas afetadas (para depuração)
	fmt.Printf("Linhas afetadas: %d\n", rowsAffected)

	return id, nil
}

func (r *UserRepository) GetOneUser(id int) (entities.User, error) {
	var user entities.User
	var err error

	query := "SELECT id, username, email, password FROM users WHERE id = ?"
	err = r.db.QueryRow(
		query,
		id).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password)

	if err != nil {
		return entities.User{}, err
	}

	defer r.db.Close()

	return user, nil
}
