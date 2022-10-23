package errors

import (
	"fmt"
	"net/http"
)

type ApiError interface {
	Message() string
	Code() string
	Status() uint
	Error() string
	NewApiError(message string, error string, status uint) ApiError
}

type ApiErr struct {
	ErrorMessage string `json:"message"`
	ErrorCode    string `json:"error"`
	ErrorStatus  uint   `json:"status"`
}

const INTERNAL_SERVER_ERROR = "Internal server error"

func (err ApiErr) Code() string {
	return err.ErrorCode
}

func (err ApiErr) Error() string {
	return fmt.Sprintf("Message: %s; Error Code: %s; Status: %d", err.ErrorMessage, err.ErrorCode, err.ErrorStatus)
}

func (err ApiErr) Status() uint {
	return err.ErrorStatus
}

func (err ApiErr) Message() string {
	return err.Message()
}

func (err ApiErr) NewApiError(message string, error string, status uint) ApiError {
	return ApiErr{message, error, status}
}

func NewNotFoundApiError(message string) ApiError {
	return ApiErr{message, "not_found", http.StatusNotFound}
}

func NewBadRequestApiError(message string) ApiError {
	return ApiErr{message, "bad_request", http.StatusBadRequest}
}

func NewInternalServerError() ApiError {
	return ApiErr{INTERNAL_SERVER_ERROR, "internal_error", http.StatusInternalServerError}
}
