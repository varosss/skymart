package domain

import "errors"

var (
	ErrInvalidUserID = errors.New("invalid user id")
	ErrUserNotFound  = errors.New("user not found")
)
