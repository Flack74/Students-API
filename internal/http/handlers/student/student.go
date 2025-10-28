package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	appErrors "github.com/Flack74/Students-API/internal/errors"
	"github.com/Flack74/Students-API/internal/storage"
	"github.com/Flack74/Students-API/internal/types"
	"github.com/Flack74/Students-API/internal/utils/response"
	"github.com/Flack74/Students-API/internal/utils/sanitize"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.HandleError(w, appErrors.NewInvalidInputError("request body is empty", err))
			return
		}
		if err != nil {
			response.HandleError(w, appErrors.NewInvalidInputError("invalid JSON format", err))
			return
		}

		// request validation
		if err := validate.Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors) // Type casting
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		// sanatize user input before saving
		student.Name, student.Email = sanitize.SanitizeJsonItems(student.Name, student.Email)

		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)
		if err != nil {
			response.HandleError(w, err)
			return
		}

		slog.Info("student created successfully", slog.Int64("id", lastId))
		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		slog.Info("getting student", slog.String("id", id))

		sanId, err := sanitize.SanitizeAndParseInt(id)
		if err != nil {
			response.HandleError(w, appErrors.NewInvalidInputError("invalid student id", err))
			return
		}

		student, err := storage.GetStudentById(sanId)
		if err != nil {
			response.HandleError(w, err)
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("getting all students")

		students, err := storage.GetStudents()
		if err != nil {
			response.HandleError(w, err)
			return
		}

		response.WriteJson(w, http.StatusOK, students)
	}
}

func DeleteById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("deleting student", slog.String("id", id))

		sanId, err := sanitize.SanitizeAndParseInt(id)
		if err != nil {
			response.HandleError(w, appErrors.NewInvalidInputError("invalid student id", err))
			return
		}

		err = storage.DeleteStudentById(sanId)
		if err != nil {
			response.HandleError(w, err)
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]string{
			"message": "student deleted successfully",
		})
	}
}

func UpdateById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("updating student", slog.String("id", id))

		sanId, err := sanitize.SanitizeAndParseInt(id)
		if err != nil {
			response.HandleError(w, appErrors.NewInvalidInputError("invalid student id", err))
			return
		}

		var student types.Student
		err = json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.HandleError(w, appErrors.NewInvalidInputError("request body is empty", err))
			return
		}
		if err != nil {
			response.HandleError(w, appErrors.NewInvalidInputError("invalid JSON format", err))
			return
		}

		// request validation
		if err := validate.Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors) // Type casting
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		// sanatize user input before saving
		student.Name, student.Email = sanitize.SanitizeJsonItems(student.Name, student.Email)

		err = storage.UpdateStudentById(sanId, student.Name, student.Email, student.Age)
		if err != nil {
			response.HandleError(w, err)
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]string{"message": "student updated successfully"})
	}
}
