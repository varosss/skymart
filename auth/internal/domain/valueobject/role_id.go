package valueobject

import "github.com/google/uuid"

type RoleID string

func NewRoleID() RoleID {
	return RoleID(uuid.NewString())
}

func (id RoleID) String() string {
	return string(id)
}
