package app

import (
	"clirzy/services/auth/internal/service"
	"clirzy/services/auth/proto"
	"clirzy/services/auth/transport/grpctransport"
	"clirzy/services/auth/transport/httptransport"

	"clirzy/pkg/bootstrap"
	"clirzy/pkg/server"

	"google.golang.org/grpc"
)

type App struct {
	bootstrap *bootstrap.Server
}

func New(authService *service.AuthService) *App {
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
