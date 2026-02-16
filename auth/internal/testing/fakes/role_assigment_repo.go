package fakes

import (
	"context"

	"clirzy/auth/internal/domain/valueobject"
)

type FakeRoleAssignmentRepo struct {
	Roles []valueobject.RoleCode
	Err   error
}

func (f *FakeRoleAssignmentRepo) GetRoles(ctx context.Context, userID valueobject.UserID) ([]valueobject.RoleCode, error) {
	return f.Roles, f.Err
}

func (f *FakeRoleAssignmentRepo) AssignRole(ctx context.Context, userID valueobject.UserID, roleCode valueobject.RoleCode) error {
	return f.Err
}

func (f *FakeRoleAssignmentRepo) RevokeRole(ctx context.Context, userID valueobject.UserID, roleCode valueobject.RoleCode) error {
	return f.Err
}
