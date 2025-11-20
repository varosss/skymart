package httptransport

import "clirzy/user/internal/domain"

func CreateRequestToDomain(req CreateUserRequest) domain.CreateUserInput {
	return domain.CreateUserInput{
		Email:    req.Email,
		Password: req.Password,
	}
}

func CreateUsersBatchRequestToDomain(req CreateUsersBatchRequest) []domain.CreateUserInput {
	inputs := make([]domain.CreateUserInput, len(req.Users))
	for i, u := range req.Users {
		inputs[i] = CreateRequestToDomain(u)
	}

	return inputs
}

func DomainToCreateUserResponse(domain *domain.User) UserResponse {
	return UserResponse{
		Id:    domain.ID,
		Email: domain.Email,
	}
}

func DomainToCreateUsersBatchResponse(domainUsers []*domain.User) CreateUsersBatchResponse {
	users := make([]UserResponse, len(domainUsers))
	for i, u := range domainUsers {
		users[i] = DomainToCreateUserResponse(u)
	}

	return CreateUsersBatchResponse{Users: users}
}
