package res

import (
	"errors"
	"net/http"

	"go-bankmate/model/app_error"
)

const (
	DefaultSuccessCode    = "00"
	DefaultSuccessStatus  = "Success"
	DefaultSuccessMessage = "Success"

	DefaultErrorCode    = "X00"
	DefaultErrorStatus  = "Failed"
	DefaultErrorMessage = "something went wrong"
)

type AppHttpResponse interface {
	Send()
}

type ApiResponse struct {
	Code    string `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewSuccessMessage(httpCode int, code string, msg string, data interface{}) (httpStatusCode int, apiResponse ApiResponse) {
	if httpCode == 0 {
		httpStatusCode = http.StatusOK
	} else {
		httpStatusCode = httpCode
	}
	if code == "" {
		code = DefaultSuccessCode
	}
	if msg == "" {
		msg = DefaultSuccessMessage
	}
	apiResponse = ApiResponse{
		code, DefaultSuccessStatus, msg, data,
	}
	return
}

func NewFailedMessage(httpCode int, code string, err error) (httpStatusCode int, apiResponse ApiResponse) {
	if httpCode == 0 {
		httpStatusCode = http.StatusOK
	} else {
		httpStatusCode = httpCode
	}
	if code == "" {
		code = DefaultSuccessCode
	}
	var userError *app_error.AppError
	if errors.As(err, &userError) {
		apiResponse = ApiResponse{
			code, DefaultErrorStatus, userError.ErrorMessage, nil,
		}
	} else {
		apiResponse = ApiResponse{
			code, DefaultErrorStatus, DefaultErrorMessage, nil,
		}
	}
	return
}
