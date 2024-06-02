package utils

import e "github.com/pjebs/jsonerror"

// ErrorType is an enum type for various error categories
type ErrorType int

const (
	// Define the error types
	ErrInvalidInput ErrorType = iota
	ErrNotFound
	ErrInternal
	ErrUnauthorized
	ErrForbidden
	ErrAuthNotSetup
	ErrUsersSomethingWrong
	ErrAuthNotSupplied
)

var ApiErrors = map[ErrorType]map[int]e.JE{
	ErrAuthNotSupplied: {401: e.New(int(ErrAuthNotSupplied), "Auth has not been supplied", "")},
	ErrInvalidInput:    {403: e.New(int(ErrInvalidInput), "Invalid input provided", "")},
	ErrNotFound:        {403: e.New(int(ErrNotFound), "Resource not found", "")},
	ErrInternal:        {500: e.New(int(ErrInternal), "Internal server error", "")},
	ErrUnauthorized:    {403: e.New(int(ErrUnauthorized), "Unauthorized access", "")},
	ErrForbidden:       {403: e.New(int(ErrForbidden), "Forbidden access", "")},
	ErrAuthNotSetup:    {403: e.New(int(ErrAuthNotSetup), "Auth has not been setup yet", "")},
	// Users.go
	ErrUsersSomethingWrong: {500: e.New(int(ErrUsersSomethingWrong), "Something went wrong with logging in user", "")},
}
