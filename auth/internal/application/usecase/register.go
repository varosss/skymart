package usecase

import (
	aport "clirzy/auth/internal/application/port"
	"clirzy/auth/internal/domain/valueobject"
	"context"
)

type RegisterCommand struct {
	Email    string
	Password string
}

type RegisterResult struct {
	UserID valueobject.UserID
}

type RegisterUseCase struct {
	users aport.UserGateway
}

func NewRegisterUseCase(users aport.UserGateway) *RegisterUseCase {
	return &RegisterUseCase{
		users: users,
	}
}

func (uc *RegisterUseCase) Execute(
	ctx context.Context,
	cmd RegisterCommand,
) (*RegisterResult, error) {
	user, err := uc.users.RegisterUser(ctx, cmd.Email, cmd.Password)
	if err != nil {
		return nil, err
	}

	userID, err := valueobject.ParseUserID(user.ID)
	if err != nil {
		return nil, err
	}

	return &RegisterResult{
		UserID: userID,
	}, nil
}
