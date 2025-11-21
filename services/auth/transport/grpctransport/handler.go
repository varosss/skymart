package grpctransport

import (
	"context"

	"clirzy/services/auth/internal/service"
	"clirzy/services/auth/proto"
)

type GRPCHandler struct {
	proto.UnimplementedAuthServiceServer

	service *service.AuthService
}

func NewGRPCHandler(service *service.AuthService) *GRPCHandler {
	return &GRPCHandler{service: service}
}

func (h *GRPCHandler) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	output, err := h.service.Login(ctx, LoginRequestToDomain(req))
	if err != nil {
		return nil, err
	}

	return domainToLoginResponse(*output), nil
}
