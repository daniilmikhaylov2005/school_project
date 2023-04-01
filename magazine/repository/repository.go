package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"time"

	pb "github.com/daniilmikhaylov2005/school_project/api"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/jmoiron/sqlx"
)

const (
	magazineTable = "magazine"
	kidTable      = "kid"
	usersTable    = "users"
	gradesTable   = "grades"
)

var ErrNotFound = errors.New("not found")

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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
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

func (r *Repository) GetClassGrades(magazine_code int64) (*pb.GetClassGradesResponse, error) {
	var class pb.GetClassGradesResponse
	class.MagazineCode = magazine_code
	query := fmt.Sprintf("SELECT id, fullname, age, graduate FROM %s WHERE magazine_code=$1", kidTable)
	rows, err := r.db.Query(query, magazine_code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to select ids from kid table: %v", err)
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

		// grades of kid
		grades := make([]*pb.Grade, 0)
		query2 := fmt.Sprintf("SELECT date, subject, grade FROM %s WHERE kid_id=$1", gradesTable)
		rows2, err := r.db.Query(query2, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrNotFound
			}
			return nil, fmt.Errorf("failed to select ids from grade table: %v", err)
		}

		for rows2.Next() {
			var (
				date    time.Time
				subject string
				grade   int64
			)
			if err := rows2.Scan(&date, &subject, &grade); err != nil {
				return nil, fmt.Errorf("failed to select grade from row: %v", err)
			}
			grades = append(grades, &pb.Grade{
				Date:    timestamppb.New(date),
				Subject: subject,
				Grade:   grade,
			})
		}

		var kid_grades pb.KidGrades
		kid_grades.Kid = &pb.Kid{
			Id:       int64(id),
			Fullname: fullname,
			Age:      int64(age),
			Graduate: int64(graduate),
		}

		kid_grades.Grades = grades
		class.ChildrenGrades = append(class.ChildrenGrades, &kid_grades)
	}
	return &class, nil
}
