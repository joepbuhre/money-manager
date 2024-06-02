package utils

import (
	"net/http"

	e "github.com/pjebs/jsonerror"
)

// ErrorType is an enum type for various error categories
type ErrorType string

const (
	// Define the error types
	ErrInvalidInput        ErrorType = "INVALID_INPUT"
	ErrNotFound            ErrorType = "NOT_FOUND"
	ErrInternal            ErrorType = "INTERNAL"
	ErrUnauthorized        ErrorType = "UNAUTHORIZED"
	ErrForbidden           ErrorType = "FORBIDDEN"
	ErrAuthNotSetup        ErrorType = "AUTH_NOT_SETUP"
	ErrUsersSomethingWrong ErrorType = "USERS_SOMETHING_WRONG"
	ErrAuthNotSupplied     ErrorType = "AUTH_NOT_SUPPLIED"
)

func createError(statusCode int, errorCode ErrorType, message string) map[int]e.JE {
	var mp = make(map[int]e.JE)
	mp[statusCode] = e.New(statusCode, string(errorCode), message)
	return mp
}

var ApiErrors = map[ErrorType]map[int]e.JE{
	ErrAuthNotSupplied: createError(401, ErrAuthNotSupplied, "Auth has not been supplied"),
	ErrInvalidInput:    createError(403, ErrInvalidInput, "Invalid input provided"),
	ErrNotFound:        createError(403, ErrNotFound, "Resource not found"),
	ErrInternal:        createError(500, ErrInternal, "Internal server error"),
	ErrUnauthorized:    createError(403, ErrUnauthorized, "Unauthorized access"),
	ErrForbidden:       createError(403, ErrForbidden, "Forbidden access"),
	ErrAuthNotSetup:    createError(403, ErrAuthNotSetup, "Auth has not been setup yet"),
	// Users.go
	ErrUsersSomethingWrong: createError(http.StatusBadRequest, ErrUsersSomethingWrong, "Something went wrong when logging in user"),
}
