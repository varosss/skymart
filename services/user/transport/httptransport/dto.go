package httptransport

type ErrorUsersResponse struct {
	Error string `json:"error"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Id       uint   `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type CreateUsersBatchRequest struct {
	Users []CreateUserRequest `json:"users" binding:"required,min=1,dive"`
}

type CreateUsersBatchResponse struct {
	Users []UserResponse `json:"users" binding:"required,min=1,dive"`
}
