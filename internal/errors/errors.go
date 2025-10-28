package errors

import "fmt"

type ErrorType string

const (
	ErrNotFound     ErrorType = "NOT_FOUND"
	ErrInvalidInput ErrorType = "INVALID_INPUT"
	ErrDatabase     ErrorType = "DATABASE_ERROR"
	ErrInternal     ErrorType = "INTERNAL_ERROR"
)

type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Type:    ErrNotFound,
		Message: message,
	}
}

func NewInvalidInputError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrInvalidInput,
		Message: message,
		Err:     err,
	}
}

func NewDatabaseError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrDatabase,
		Message: message,
		Err:     err,
	}
}

func NewInternalError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrInternal,
		Message: message,
		Err:     err,
	}
}
