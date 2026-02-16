package usecase

import (
	aport "clirzy/auth/internal/application/port"
	"clirzy/auth/internal/domain/port"
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
	roles port.RoleAssignmentRepo
}

func NewRegisterUseCase(users aport.UserGateway, roles port.RoleAssignmentRepo) *RegisterUseCase {
	return &RegisterUseCase{
		users: users,
		roles: roles,
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

	if err := uc.roles.AssignRole(ctx, user.ID, valueobject.UserRole); err != nil {
		return nil, err
	}

	return &RegisterResult{
		UserID: user.ID,
	}, nil
}
