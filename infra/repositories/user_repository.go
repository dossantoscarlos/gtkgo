package repositories

import (
	"database/sql"
	"fmt"
	"gtkgo/core/domain/entities"
	"gtkgo/infra/database"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	db, err := database.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]entities.User, error) {
	var users []entities.User
	rows, err := r.db.Query("SELECT id, username, email, password FROM users")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
			log.Fatal(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) CreateUser(user entities.User) (id int, err error) {
	log.Default().Printf("Inserindo usuário: %v\n", user)
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	result, err := r.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Usuário inserido com sucesso: %v\n", result)

	// Verifica se a execução afetou alguma linha (se o banco foi alterado)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return
	}

	id = int(lastInsertId)

	// Exibe o número de linhas afetadas (para depuração)
	fmt.Printf("Linhas afetadas: %d\n", rowsAffected)

	return id, nil
}

func (r *UserRepository) GetUserById(id int) (entities.User, error) {
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

	return user, nil
}

func (r *UserRepository) UpdateUser(id string, user entities.User) (entities.User, error) {
	return user, nil
}

func (r *UserRepository) DeleteUser(id int) error {

	query := "delete from users where id = ?"

	result, err := r.db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Verifica se a execução afetou alguma linha (se o banco foi alterado)
	if _, err := result.RowsAffected(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
