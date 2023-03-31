package repository

import (
	"database/sql"
	"errors"
	"fmt"

	pb "github.com/daniilmikhaylov2005/school_project/api"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

var NoUsersError = errors.New("no such user")

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

func (r *Repository) GetUserByLogin(login string) (*pb.GetUserResponse, error) {
	var user pb.GetUserResponse
	query := fmt.Sprintf("SELECT id, firstname, secondname, login, email, password FROM %s WHERE login=$1", usersTable)
	err := r.db.QueryRow(query, login).Scan(&user.Id, &user.FirstName, &user.SecondName, &user.Login, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NoUsersError
		}
		return nil, fmt.Errorf("failed to find user in db: %v", err)
	}
	return &user, nil
}
