package app

import (
	"clirzy/services/auth/internal/db"
	"clirzy/services/auth/internal/service"
	"clirzy/services/auth/proto"
	"clirzy/services/auth/transport/grpctransport"
	"clirzy/services/auth/transport/httptransport"

	"clirzy/pkg/bootstrap"
	userclient "clirzy/pkg/client/user"
	"clirzy/pkg/config"
	"clirzy/pkg/consts"
	pkgdb "clirzy/pkg/db"
	"clirzy/pkg/server"
	"clirzy/pkg/utils"

	"google.golang.org/grpc"
)

type App struct {
	bootstrap *bootstrap.Server
}

func New(cfg *config.Config) (*App, error) {
	conn, err := pkgdb.Connect(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	userClient, err := userclient.New("clirzy-user:50051")
	if err != nil {
		return nil, err
	}

	authService := service.NewAuthService(
		db.NewTokensRepo(conn),
		utils.NewJWTManager(cfg.JWTSecret, consts.DEFAULT_ACCESS_TOKEN_LIFETIME),
		userClient,
	)

	grpcHandler := grpctransport.NewGRPCHandler(authService)
	httpHandler := httptransport.NewHTTPHandler(authService)

	grpcSrv := server.NewGRPCServer(":50051", func(s *grpc.Server) {
		proto.RegisterAuthServiceServer(s, grpcHandler)
	})

	httpSrv := server.NewHTTPServer(":80")
	httpSrv.AddRoute("POST", "/login", httpHandler.Login)
	httpSrv.AddRoute("POST", "/register", httpHandler.Register)
	httpSrv.AddRoute("POST", "/refresh", httpHandler.Refresh)

	bootsrapServerInst := bootstrap.NewServer(
		bootstrap.WithHTTP(httpSrv),
		bootstrap.WithGRPC(grpcSrv),
	)

	return &App{
		bootstrap: bootsrapServerInst,
	}, nil
}

func (a *App) Run() {
	a.bootstrap.Run()
}
