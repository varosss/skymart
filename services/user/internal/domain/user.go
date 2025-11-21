package domain

type CreateUserInput struct {
	Username string
	Password string
}

type User struct {
	ID           uint
	Username     string
	PasswordHash string
}
