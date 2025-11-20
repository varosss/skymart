package grpctransport

import (
	"clirzy/user/internal/domain"
	"clirzy/user/proto"
)

func CreateUserRequestToDomain(req *proto.CreateUsersRequest) []domain.CreateUserInput {
	out := make([]domain.CreateUserInput, len(req.Users))
	for i, u := range req.Users {
		out[i] = domain.CreateUserInput{Email: u.Email, Password: u.Password}
	}
	return out
}

func DomainToCreateUsersResponse(domainUsers []*domain.User) *proto.CreateUsersResponse {
	grpcUsers := make([]*proto.User, len(domainUsers))
	for i, domain := range domainUsers {
		grpcUsers[i] = &proto.User{
			Id:    int64(domain.ID),
			Email: domain.Email,
		}
	}

	return &proto.CreateUsersResponse{Users: grpcUsers}
}

func DomainToProtoUser(domain *domain.User) *proto.User {
	return &proto.User{Id: int64(domain.ID), Email: domain.Email}
}
