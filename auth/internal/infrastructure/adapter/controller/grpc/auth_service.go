package grpc

import (
	"clirzy/auth/internal/application/usecase"
	"clirzy/auth/internal/domain"
	pb "clirzy/auth/proto"
	"context"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer

	validateAccessTokenUC *usecase.ValidateAccessTokenUseCase
}

func NewAuthServiceServer(
	validateAccessTokenUC *usecase.ValidateAccessTokenUseCase,
) *AuthServiceServer {
	return &AuthServiceServer{
		validateAccessTokenUC: validateAccessTokenUC,
	}
}

func (s *AuthServiceServer) ValidateAccessToken(ctx context.Context, req *pb.ValidateAccessTokenRequest) (*pb.ValidateAccessTokenResponse, error) {
	result, err := s.validateAccessTokenUC.Execute(ctx, usecase.ValidateAccessTokenCommand{
		AccessToken: req.AccessToken,
	})

	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	return &pb.ValidateAccessTokenResponse{UserId: result.UserID.String()}, nil
}
