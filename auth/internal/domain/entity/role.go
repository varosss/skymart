package entity

import "clirzy/auth/internal/domain/valueobject"

type Role struct {
	id   valueobject.RoleID
	code valueobject.RoleCode
}

func NewRole(
	code valueobject.RoleCode,
) Role {
	return Role{
		id:   valueobject.NewRoleID(),
		code: code,
	}
}

func (r Role) Code() valueobject.RoleCode {
	return r.code
}
