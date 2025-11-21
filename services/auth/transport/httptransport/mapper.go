package httptransport

import "clirzy/services/auth/internal/domain"

func AuthRequestToDomain(req *AuthRequest) domain.AuthInput {
	return domain.AuthInput{
		Username: req.Username,
		Password: req.Password,
	}
}

func domainToLoginResponse(domain *domain.LoginOutput) LoginResponse {
	return LoginResponse{
		AccessToken:  domain.AccessToken,
		RefreshToken: domain.RefreshToken,
	}
}

func domainToRegisterResponse(domain *domain.User) RegisterResponse {
	return RegisterResponse{
		Id:       domain.Id,
		Username: domain.Username,
	}
}
