package grpctransport

import (
	"clirzy/user/internal/service"
	"clirzy/user/proto"
	"context"
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

func (h *GRPCHandler) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.User, error) {
	user, err := h.service.GetUser(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return DomainToProtoUser(user), nil
}
