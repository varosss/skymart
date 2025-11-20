package grpctransport

import (
	"clirzy/auth/internal/service"
	"clirzy/auth/proto"
	"context"
)

type GRPCHandler struct {
	proto.UnimplementedAuthServiceServer

	service *service.AuthService
}

func NewGRPCHandler(service *service.AuthService) *GRPCHandler {
	return &GRPCHandler{service: service}
}

func (s *GRPCHandler) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	accessToken, err := s.service.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &proto.LoginResponse{AccessToken: accessToken, RefreshToken: "", User: &proto.User{}}, nil
}
