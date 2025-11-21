package grpctransport

import (
	"context"

	"clirzy/services/user/internal/service"
	"clirzy/services/user/proto"
)

type GRPCHandler struct {
	proto.UnimplementedUserServiceServer

	service *service.UserService
}

func NewGRPCHandler(service *service.UserService) *GRPCHandler {
	return &GRPCHandler{service: service}
}

func (h *GRPCHandler) CreateUsers(ctx context.Context, req *proto.CreateUsersRequest) (*proto.CreateUsersResponse, error) {
	domainUsers, err := h.service.CreateMany(ctx, CreateUserRequestToDomain(req))
	if err != nil {
		return nil, err
	}

	return DomainToCreateUsersResponse(domainUsers), nil
}

func (h *GRPCHandler) GetUserById(ctx context.Context, req *proto.GetUserByIdRequest) (*proto.UserResponse, error) {
	user, err := h.service.GetUserById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return DomainToProtoUserResponse(user), nil
}
