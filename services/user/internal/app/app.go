package app

import (
	"clirzy/pkg/bootstrap"
	"clirzy/pkg/config"
	"clirzy/pkg/consts"
	pkgdb "clirzy/pkg/db"
	"clirzy/pkg/middleware"
	"clirzy/pkg/server"
	"clirzy/pkg/utils"
	"clirzy/services/user/internal/db"
	"clirzy/services/user/internal/service"
	"clirzy/services/user/proto"
	"clirzy/services/user/transport/grpctransport"
	"clirzy/services/user/transport/httptransport"

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

	userService := service.NewUserService(db.NewUsersRepo(conn))

	grpcHandler := grpctransport.NewGRPCHandler(userService)
	httpHandler := httptransport.NewHTTPHandler(userService)

	grpcSrv := server.NewGRPCServer(":50051", func(s *grpc.Server) {
		proto.RegisterUserServiceServer(s, grpcHandler)
	})

	authMiddleware := middleware.AuthMiddleware(utils.NewJWTManager("secret", consts.DEFAULT_ACCESS_TOKEN_LIFETIME))

	httpSrv := server.NewHTTPServer(":80")
	httpSrv.AddRoute("POST", "/users", httpHandler.CreateUsers)
	httpSrv.AddRoute("GET", "/users/:id", authMiddleware, httpHandler.GetUserById)

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
