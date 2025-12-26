package domain

import "errors"

var (
	ErrInvalidUserID = errors.New("invalid user id")
	ErrNoUserFound   = errors.New("user not found")
)
