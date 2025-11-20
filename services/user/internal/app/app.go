package app

import (
	"clirzy/common/bootstrap"
	"clirzy/common/consts"
	"clirzy/common/middleware"
	"clirzy/common/server"
	"clirzy/common/utils"
	"clirzy/user/internal/service"
	"clirzy/user/internal/transport/grpctransport"
	"clirzy/user/internal/transport/httptransport"
	"clirzy/user/proto"

	"google.golang.org/grpc"
)

type App struct {
	bootstrap *bootstrap.Server
}

func New() *App {
	userService := service.NewUserService()

	grpcHandler := grpctransport.NewGRPCHandler(userService)
	httpHandler := httptransport.NewHTTPHandler(userService)

	grpcSrv := server.NewGRPCServer(":50051", func(s *grpc.Server) {
		proto.RegisterUserServiceServer(s, grpcHandler)
	})

	authMiddleware := middleware.AuthMiddleware(utils.NewJWTManager("secret", consts.DEFAULT_TOKEN_LIFETIME))

	httpSrv := server.NewHTTPServer(":8080")
	httpSrv.AddRoute("POST", "/users", httpHandler.CreateUsers)
	httpSrv.AddRoute("GET", "/users/:id", authMiddleware, httpHandler.GetUser)

	bootsrapServerInst := bootstrap.NewServer(
		bootstrap.WithHTTP(httpSrv),
		bootstrap.WithGRPC(grpcSrv),
	)

	return &App{
		bootstrap: bootsrapServerInst,
	}
}

func (a *App) Run() {
	a.bootstrap.Run()
}
