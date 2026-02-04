package domain

import "errors"

var (
	ErrInvalidUserID          = errors.New("invalid user id")
	ErrInvalidEmail           = errors.New("invalid email")
	ErrEmptyEmail             = errors.New("email is empty")
	ErrInvalidPassword        = errors.New("password is invalid")
	ErrUserNotFound           = errors.New("user not found")
	ErrEmailAlreadyRegistered = errors.New("user email is already registered")
)
