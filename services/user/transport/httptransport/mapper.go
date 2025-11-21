package httptransport

import "clirzy/services/user/internal/domain"

func CreateRequestToDomain(req CreateUserRequest) domain.CreateUserInput {
	return domain.CreateUserInput{
		Username: req.Username,
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

func DomainToUserResponse(domain *domain.User) UserResponse {
	return UserResponse{
		Id:       domain.ID,
		Username: domain.Username,
	}
}

func DomainToCreateUsersBatchResponse(domainUsers []*domain.User) CreateUsersBatchResponse {
	users := make([]UserResponse, len(domainUsers))
	for i, u := range domainUsers {
		users[i] = DomainToUserResponse(u)
	}

	return CreateUsersBatchResponse{Users: users}
}
