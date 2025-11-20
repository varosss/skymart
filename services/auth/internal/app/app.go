package app

import (
	"clirzy/auth/internal/service"
	"clirzy/auth/internal/transport/grpctransport"
	"clirzy/auth/internal/transport/httptransport"
	"clirzy/auth/proto"
	"clirzy/common/bootstrap"
	"clirzy/common/server"

	"google.golang.org/grpc"
)

type App struct {
	bootstrap *bootstrap.Server
}

func New() *App {
	authService := service.NewAuthService()

	grpcHandler := grpctransport.NewGRPCHandler(authService)
	httpHandler := httptransport.NewHTTPHandler(authService)

	grpcSrv := server.NewGRPCServer(":50051", func(s *grpc.Server) {
		proto.RegisterAuthServiceServer(s, grpcHandler)
	})

	httpSrv := server.NewHTTPServer(":8080")
	httpSrv.AddRoute("POST", "/login", httpHandler.Login)
	httpSrv.AddRoute("POST", "/register", httpHandler.Register)

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
