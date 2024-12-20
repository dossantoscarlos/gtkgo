package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		return nil, fmt.Errorf("Error ao conectar ao banco de dados: %v", err)
	}

	err = createTable(db)
	if err != nil {
		return nil, fmt.Errorf("Error ao criar tabela: %v", err)
	}

	return db, nil
}

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("erro ao criar tabela: %v", err)
	}

	return nil
}

func CloseDB(db *sql.DB) {
	db.Close()
}
