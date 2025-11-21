package db

import "clirzy/services/user/internal/domain"

func CreateUserInputToDB(input domain.CreateUserInput) User {
	return User{
		Username:     input.Username,
		PasswordHash: input.Password,
	}
}

func DBToDomain(db User) *domain.User {
	return &domain.User{
		ID:           db.ID,
		Username:     db.Username,
		PasswordHash: db.PasswordHash,
	}
}
