// internal/infrastructure/adapter/gorm/mapper/user.go
package mapper

import (
	"clirzy/user/internal/domain/entity"
	"clirzy/user/internal/domain/valueobject"
	"clirzy/user/internal/infrastructure/adapter/gorm/model"
)

func ToDomain(m *model.User) (*entity.User, error) {
	id, err := valueobject.ParseUserID(m.ID)
	if err != nil {
		return nil, err
	}

	email, err := valueobject.NewEmail(m.Email)
	if err != nil {
		return nil, err
	}

	hash, err := valueobject.NewPasswordHash(m.PasswordHash)
	if err != nil {
		return nil, err
	}

	u := entity.NewUser(id, email, hash)

	switch m.Status {
	case string(valueobject.StatusActive):
		u.Activate()
	case string(valueobject.StatusBlocked):
		_ = u.Block()
	}

	return u, nil
}

func ToModel(u *entity.User) *model.User {
	return &model.User{
		ID:           u.ID().String(),
		Email:        u.Email().String(),
		PasswordHash: u.PasswordHash().String(),
		Status:       u.Status().String(),
	}
}
