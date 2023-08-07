package app_error

import (
	"fmt"
	"net/http"
)

type AppError struct {
	ErrorMessage string
	ErrorType    int
}

func (e *AppError) Error() string {
	return fmt.Sprintf("type: %d, err: %s", e.ErrorType, e.ErrorMessage)
}

func InvalidError(msg string) error {
	if msg == "" {
		return &AppError{
			ErrorMessage: "invalid input",
			ErrorType:    http.StatusBadRequest,
		}
	} else {
		return &AppError{
			ErrorMessage: msg,
			ErrorType:    http.StatusBadRequest,
		}
	}
}

func UnauthorizedError(msg string) error {
	if msg == "" {
		return &AppError{
			ErrorMessage: "unauthorized user",
			ErrorType:    http.StatusUnauthorized,
		}
	} else {
		return &AppError{
			ErrorMessage: msg,
			ErrorType:    http.StatusUnauthorized,
		}
	}
}

func DataNotFoundError(msg string) error {
	if msg == "" {
		return &AppError{
			ErrorMessage: "no data found",
		}
	} else {
		return &AppError{
			ErrorMessage: msg,
		}
	}
}

func UnknownError(msg string) error {
	if msg == "" {
		return &AppError{
			ErrorMessage: "something went wrong",
		}
	} else {
		return &AppError{
			ErrorMessage: msg,
		}
	}
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
