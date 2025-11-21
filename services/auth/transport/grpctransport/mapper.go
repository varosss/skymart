package grpctransport

import (
	"clirzy/services/auth/internal/domain"
	"clirzy/services/auth/proto"
)

func LoginRequestToDomain(req *proto.LoginRequest) domain.AuthInput {
	return domain.AuthInput{
		Username: req.Username,
		Password: req.Password,
	}
}

func domainToLoginResponse(domain domain.LoginOutput) *proto.LoginResponse {
	return &proto.LoginResponse{
		AccessToken:  domain.AccessToken,
		RefreshToken: domain.RefreshToken,
	}
}
