package builder

import (
	"clirzy/pkg/grpcserver"
	grpccontroller "clirzy/user/internal/infrastructure/adapter/controller/grpc"
	"clirzy/user/internal/infrastructure/config"
	pb "clirzy/user/proto"

	"google.golang.org/grpc"
)

func BuildGRPC(cfg *config.Config, uc *UseCases) (*grpcserver.Server, error) {
	srv, err := grpcserver.NewServer(":"+cfg.GrpcServer.Port, func(g *grpc.Server) {
		pb.RegisterUserServiceServer(
			g,
			grpccontroller.NewUserServiceServer(uc.RegisterUser),
		)
	})
	if err != nil {
		return nil, err
	}

	return srv, nil
}
