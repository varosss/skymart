package httptransport

type ErrorAuthResponse struct {
	Error string `json:"error"`
}

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegisterResponse struct {
	Id       uint
	Username string
}

type RefreshRequest struct {
	RefreshToken string
}
