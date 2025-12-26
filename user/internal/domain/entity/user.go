package entity

import (
	"clirzy/user/internal/domain/valueobject"
	"errors"
)

type User struct {
	id           valueobject.UserID
	email        valueobject.Email
	passwordHash valueobject.PasswordHash
	status       valueobject.Status
}

func NewUser(
	id valueobject.UserID,
	email valueobject.Email,
	passwordHash valueobject.PasswordHash,
) *User {
	return &User{
		id:           id,
		email:        email,
		passwordHash: passwordHash,
		status:       valueobject.StatusActive,
	}
}

func (u *User) ID() valueobject.UserID {
	return u.id
}

func (u *User) Email() valueobject.Email {
	return u.email
}

func (u *User) PasswordHash() valueobject.PasswordHash {
	return u.passwordHash
}

func (u *User) Status() valueobject.Status {
	return u.status
}

func (u *User) Block() error {
	if u.status == valueobject.StatusDeleted {
		return errors.New("cannot block deleted user")
	}
	u.status = valueobject.StatusBlocked
	return nil
}

func (u *User) Activate() {
	u.status = valueobject.StatusActive
}
