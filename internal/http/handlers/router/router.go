package router

import (
	"net/http"

	"github.com/Flack74/Students-API/internal/http/handlers/student"
	"github.com/Flack74/Students-API/internal/storage"
)

func New(storage storage.Storage) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("POST /api/students", student.New(storage))
	r.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	r.HandleFunc("GET /api/students", student.GetList(storage))
	r.HandleFunc("PUT /api/students/{id}", student.UpdateById(storage))
	r.HandleFunc("DELETE /api/students/{id}", student.DeleteById(storage))

	return r
}
