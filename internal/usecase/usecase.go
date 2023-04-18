package usecase

import (
	"errors"
)

var (
	ErrUsecaseInternal       = errors.New("internal server error")
	ErrUsecaseNoData         = errors.New("no data")
	ErrUsecaseEmptyUsername  = errors.New("`username` cannot be empty")
	ErrUsecaseEmptyPassword  = errors.New("`password` cannot be empty")
	ErrUsecaseEmptyFname     = errors.New("`first_name` cannot be empty")
	ErrUsecaseExistsUsername = errors.New("`username` already exists")
	ErrUsecaseInvalidAuth    = errors.New("`username` or `password` wrong")
	ErrInvalidFileData       = errors.New("invalid file data")
	ErrInvalidUserID         = errors.New("invalid user id")
	ErrInvalidQuery          = errors.New("invalid query")
	ErrWalletAlreadyExists   = errors.New("wallet already exists")
	ErrInvalidEmail          = errors.New("invalid email")
	ErrUserNotFound          = errors.New("User not found")
	ErrInvalidAuth           = errors.New("invalid User Credential")
	ErrNotFound              = errors.New("wallet not found")
	ErrInsufficientBalance   = errors.New("insufficient balance")
)
