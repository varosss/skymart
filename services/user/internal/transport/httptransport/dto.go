package httptransport

type ErrorUsersResponse struct {
	Error string `json:"error"`
}

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Id    uint   `json:"id" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type CreateUsersBatchRequest struct {
	Users []CreateUserRequest `json:"users" binding:"required,min=1,dive"`
}

type CreateUsersBatchResponse struct {
	Users []UserResponse `json:"users" binding:"required,min=1,dive"`
}
