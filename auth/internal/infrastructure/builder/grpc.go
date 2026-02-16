package builder

import (
	grpccontroller "clirzy/auth/internal/infrastructure/adapter/controller/grpc"
	"clirzy/auth/internal/infrastructure/config"
	pb "clirzy/auth/proto"
	"clirzy/pkg/grpcserver"

	"google.golang.org/grpc"
)

func BuildGRPC(cfg *config.Config, uc *UseCases) (*grpcserver.Server, error) {
	srv, err := grpcserver.NewServer(":"+cfg.GrpcServer.Port, func(g *grpc.Server) {
		pb.RegisterAuthServiceServer(
			g,
			grpccontroller.NewAuthServiceServer(uc.ValidateAccessToken),
		)
	})
	if err != nil {
		return nil, err
	}

	return srv, nil
}
