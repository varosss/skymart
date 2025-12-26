package valueobject

import "errors"

type PasswordHash string

func NewPasswordHash(hash string) (PasswordHash, error) {
	if hash == "" {
		return "", errors.New("password hash is empty")
	}

	return PasswordHash(hash), nil
}
