package domain

type AuthInput struct {
	Username string
	Password string
}

type User struct {
	Id       uint
	Username string
}

type AuthPair struct {
	AccessToken  string
	RefreshToken string
}
