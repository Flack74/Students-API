package sqlite

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/Flack74/Students-API/internal/config"
	appErrors "github.com/Flack74/Students-API/internal/errors"
	"github.com/Flack74/Students-API/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func (s *Sqlite) Close() error {
	return s.Db.Close()
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
		id   INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		age  INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, appErrors.NewDatabaseError("failed to prepare insert statement", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, appErrors.NewDatabaseError("failed to insert student", err)
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, appErrors.NewDatabaseError("failed to get last insert id", err)
	}

	return lastId, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, appErrors.NewDatabaseError("failed to prepare select statement", err)
	}
	defer stmt.Close()

	var student types.Student
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, appErrors.NewNotFoundError(fmt.Sprintf("student with id %d not found", id))
		}
		return types.Student{}, appErrors.NewDatabaseError("failed to query student", err)
	}

	return student, nil
}

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM students")
	if err != nil {
		return nil, appErrors.NewDatabaseError("failed to prepare select statement", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, appErrors.NewDatabaseError("failed to query students", err)
	}
	defer rows.Close()

	var students []types.Student
	for rows.Next() {
		var student types.Student
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, appErrors.NewDatabaseError("failed to scan student row", err)
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, appErrors.NewDatabaseError("error iterating rows", err)
	}

	return students, nil
}

func (s *Sqlite) DeleteStudentById(id int64) error {
	stmt, err := s.Db.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		return appErrors.NewDatabaseError("failed to prepare delete statement", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return appErrors.NewDatabaseError("failed to delete student", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return appErrors.NewDatabaseError("failed to get affected rows", err)
	}

	if affected == 0 {
		return appErrors.NewNotFoundError(fmt.Sprintf("student with id %d not found", id))
	}

	return nil
}

func (s *Sqlite) UpdateStudentById(id int64, name string, email string, age int) error {
	stmt, err := s.Db.Prepare("UPDATE students SET name = ?, email = ?, age = ? WHERE id = ?")
	if err != nil {
		return appErrors.NewDatabaseError("failed to prepare update statement", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, email, age, id)
	if err != nil {
		return appErrors.NewDatabaseError("failed to update student", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return appErrors.NewDatabaseError("failed to get affected rows", err)
	}

	if rowsAffected == 0 {
		return appErrors.NewNotFoundError(fmt.Sprintf("student with id %d not found", id))
	}

	slog.Info("student updated", "id", id, "rows_affected", rowsAffected)
	return nil
}

// func (s *Sqlite) PartialUpdateStudentById(id int64) error {

// }
