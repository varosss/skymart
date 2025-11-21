package domain

type AuthInput struct {
	Username string
	Password string
}

type User struct {
	Id       uint
	Username string
}

type LoginOutput struct {
	AccessToken  string
	RefreshToken string
	User         User
}
