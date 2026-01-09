package builder

import (
	"clirzy/auth/internal/infrastructure/adapter/auth"
	"clirzy/auth/internal/infrastructure/adapter/gorm/repo"
	"clirzy/auth/internal/infrastructure/adapter/grpcclient"
	"clirzy/auth/internal/infrastructure/adapter/security"
	"clirzy/auth/internal/infrastructure/config"
	pkgapp "clirzy/pkg/app"
	"clirzy/pkg/clock"
	"clirzy/pkg/db"
	pkgsecurity "clirzy/pkg/security"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	publicKey, err := pkgsecurity.LoadPublicKey(cfg.Security.PublicKeyPath)
	if err != nil {
		return nil, err
	}
	privateKey, err := pkgsecurity.LoadPrivateKey(cfg.Security.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	userServiceConn, err := grpc.NewClient(
		cfg.UserService.GRPCAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	application.AddCloser(func() error {
		return userServiceConn.Close()
	})

	uc := BuildUseCases(
		grpcclient.NewUserServiceClient(userServiceConn),
		security.NewBcryptPasswordVerifier(),
		repo.NewRefreshTokensGormRepo(conn),
		auth.NewJWTSigner(
			privateKey,
			cfg.JWT.Issuer,
			cfg.JWT.AccessTTL,
			cfg.JWT.RefreshTTL,
		),
		auth.NewJWTVerifier(
			publicKey,
			cfg.JWT.Issuer,
		),
		clock.NewSystemClock(),
		cfg.JWT.RefreshTTL,
	)

	httpSrv := BuildHTTP(cfg, uc)
	grpcSrv, err := BuildGRPC(cfg, uc)
	if err != nil {
		return nil, err
	}

	application.Add(pkgapp.NewHTTPComponent(httpSrv))
	application.Add(pkgapp.NewGRPCComponent(grpcSrv))

	return application, nil
}
