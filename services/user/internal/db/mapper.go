package db

import "clirzy/user/internal/domain"

func CreateUserInputToDB(input domain.CreateUserInput) User {
	return User{
		Email:        input.Email,
		PasswordHash: input.Password,
	}
}

func DBToDomain(db User) *domain.User {
	return &domain.User{
		ID:        db.ID,
		Email:     db.Email,
		CreatedAt: db.CreatedAt,
	}
}
