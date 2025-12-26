package usecase

import (
	"clirzy/user/internal/domain"
	"clirzy/user/internal/domain/entity"
	"clirzy/user/internal/domain/port"
	"clirzy/user/internal/domain/valueobject"
	"context"
)

type RegisterUserCommand struct {
	Email    string
	Password string
}

type RegisterUserUseCase struct {
	users     port.UsersRepo
	passwords port.PasswordHasher
	clock     port.Clock
}

func NewRegisterUserUseCase(users port.UsersRepo, passwords port.PasswordHasher, clock port.Clock) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		users:     users,
		passwords: passwords,
		clock:     clock,
	}
}

func (uc *RegisterUserUseCase) Execute(
	ctx context.Context,
	cmd RegisterUserCommand,
) (*entity.User, error) {
	if uc.users.ExistsByEmail(ctx, cmd.Email) {
		return nil, domain.ErrEmailAlreadyRegistered
	}

	hash, err := uc.passwords.Hash(cmd.Password)
	if err != nil {
		return nil, err
	}

	email, err := valueobject.NewEmail(cmd.Email)
	if err != nil {
		return nil, domain.ErrInvalidEmail
	}

	passwordHash, err := valueobject.NewPasswordHash(hash)
	if err != nil {
		return nil, domain.ErrInvalidPassword
	}

	u := entity.NewUser(valueobject.NewUserID(), email, passwordHash)

	if err := uc.users.Save(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}
