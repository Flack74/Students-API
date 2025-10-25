package storage

import "github.com/Flack74/Students-API/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	DeleteStudentById(intId int64) error
	UpdateStudentById(intId int64, name string, email string, age int) error
}
