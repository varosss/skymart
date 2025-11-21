package grpctransport

import (
	"clirzy/services/user/internal/domain"
	"clirzy/services/user/proto"
)

func CreateUserRequestToDomain(req *proto.CreateUsersRequest) []domain.CreateUserInput {
	out := make([]domain.CreateUserInput, len(req.Users))
	for i, u := range req.Users {
		out[i] = domain.CreateUserInput{Username: u.Username, Password: u.Password}
	}
	return out
}

func DomainToCreateUsersResponse(domainUsers []*domain.User) *proto.CreateUsersResponse {
	grpcUsers := make([]*proto.UserResponse, len(domainUsers))
	for i, domain := range domainUsers {
		grpcUsers[i] = &proto.UserResponse{
			Id:       int64(domain.ID),
			Username: domain.Username,
		}
	}

	return &proto.CreateUsersResponse{Users: grpcUsers}
}

func DomainToProtoUserResponse(domain *domain.User) *proto.UserResponse {
	return &proto.UserResponse{Id: int64(domain.ID), Username: domain.Username}
}
