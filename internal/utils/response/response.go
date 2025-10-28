package response

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	appErrors "github.com/Flack74/Students-API/internal/errors"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func HandleError(w http.ResponseWriter, err error) {
	var appErr *appErrors.AppError
	if e, ok := err.(*appErrors.AppError); ok {
		appErr = e
	} else {
		appErr = appErrors.NewInternalError("an unexpected error occurred", err)
	}

	slog.Error("request error", "type", appErr.Type, "message", appErr.Message, "error", appErr.Err)

	var status int
	var message string

	switch appErr.Type {
	case appErrors.ErrNotFound:
		status = http.StatusNotFound
		message = appErr.Message
	case appErrors.ErrInvalidInput:
		status = http.StatusBadRequest
		message = appErr.Message
	case appErrors.ErrDatabase:
		status = http.StatusInternalServerError
		message = "database operation failed"
	default:
		status = http.StatusInternalServerError
		message = "an unexpected error occurred"
	}

	WriteJson(w, status, Response{
		Status: StatusError,
		Error:  message,
	})
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required field", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
