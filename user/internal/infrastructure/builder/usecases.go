package builder

import (
	"clirzy/user/internal/application/usecase"
	"clirzy/user/internal/domain/port"
)

type UseCases struct {
	RegisterUser *usecase.RegisterUserUseCase
}

func BuildUseCases(
	users port.UsersRepo,
	passwords port.PasswordHasher,
	clock port.Clock,
) *UseCases {
	return &UseCases{
		RegisterUser: usecase.NewRegisterUserUseCase(
			users,
			passwords,
			clock,
		),
	}
}
