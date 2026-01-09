package builder

import (
	pkgapp "clirzy/pkg/app"
	"clirzy/pkg/clock"
	"clirzy/pkg/db"
	"clirzy/user/internal/infrastructure/adapter/gorm/repo"
	"clirzy/user/internal/infrastructure/adapter/security"
	"clirzy/user/internal/infrastructure/config"
)

func BuildApp(cfg *config.Config) (*pkgapp.Application, error) {
	application := pkgapp.New()

	conn, err := db.ConnectGorm(cfg.Database.DSN)
	if err != nil {
		return nil, err
	}
	application.AddCloser(func() error {
		return db.CloseGorm(conn)
	})

	uc := BuildUseCases(
		repo.NewUsersGormRepo(conn),
		security.NewBcryptPasswordHasher(cfg.Security.HashCost),
		clock.NewSystemClock(),
	)

	grpcSrv, err := BuildGRPC(cfg, uc)
	if err != nil {
		return nil, err
	}

	application.Add(pkgapp.NewGRPCComponent(grpcSrv))

	return application, nil
}
