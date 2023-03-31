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
	usersTable    = "users"
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

func (r *Repository) GetClass(magazine_code int64) (*pb.GetClassResponse, error) {
	var teacher_firstname, teacher_secondname string
	var class pb.GetClassResponse
	query := fmt.Sprintf("SELECT firstname, secondname FROM %s WHERE login = (SELECT teacher_login FROM %s WHERE magazine_code=$1)", usersTable, magazineTable)
	err := r.db.QueryRow(query, magazine_code).Scan(&teacher_firstname, &teacher_secondname)
	if err != nil {
		return nil, fmt.Errorf("failed to select teacher login from db: %v", err)
	}
	class.TeacherFullname = fmt.Sprintf("%s %s", teacher_firstname, teacher_secondname)

	query2 := fmt.Sprintf("SELECT id, fullname, age, graduate FROM %s WHERE magazine_code=$1", kidTable)
	rows, err := r.db.Query(query2, magazine_code)
	if err != nil {
		return nil, fmt.Errorf("failed to select children from db: %v", err)
	}

	for rows.Next() {
		var (
			id       int
			fullname string
			age      int
			graduate int
		)

		if err := rows.Scan(&id, &fullname, &age, &graduate); err != nil {
			return nil, fmt.Errorf("failed to select kid from row: %v", err)
		}

		class.Children = append(class.Children, &pb.Kid{Fullname: fullname, Age: int64(age), Id: int64(id), Graduate: int64(graduate)})
	}

	return &class, nil
}
