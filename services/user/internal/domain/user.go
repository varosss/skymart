package domain

import "time"

type CreateUserInput struct {
	Email    string
	Password string
}

type User struct {
	ID        uint
	Email     string
	CreatedAt time.Time
}
