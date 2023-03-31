package repository

import (
	"fmt"
	"math/rand"

	pb "github.com/daniilmikhaylov2005/school_project/api"

	"github.com/jmoiron/sqlx"
)

const (
	magazineTable = "magazine"
	kidTable      = "kid"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateClass(children []*pb.Kid, teacher_login string, graduate int64) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to create transaction: %v", err)
	}

	magazineCode := rand.Intn(1000)
	query := fmt.Sprintf("INSERT INTO %s (magazine_code, teacher_login) VALUES ($1, $2)", magazineTable)
	_, err = tx.Exec(query, magazineCode, teacher_login)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to insert magazine: %v", err)
	}

	for _, v := range children {
		query := fmt.Sprintf("INSERT INTO %s (magazine_code, fullname, age, graduate) VALUES ($1, $2, $3, $4)", kidTable)
		_, err := tx.Exec(query, magazineCode, v.GetFullname(), v.GetAge(), graduate)
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("failed to insert kid: %v", err)
		}
	}

	tx.Commit()
	return magazineCode, nil
}
