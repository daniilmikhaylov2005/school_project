package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(firstname, secondname, login, email, password string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (firstname, secondname, login, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id", usersTable)
	err := r.db.QueryRow(query, firstname, secondname, login, email, password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create user in db: %v", err)
	}
	return id, nil
}
