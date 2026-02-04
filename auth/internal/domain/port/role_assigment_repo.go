package port

import (
	"context"

	"clirzy/auth/internal/domain/valueobject"
)

type RoleAssignmentRepo interface {
	GetRoles(ctx context.Context, userID valueobject.UserID) ([]valueobject.RoleCode, error)
	AssignRole(ctx context.Context, userID valueobject.UserID, roleCode valueobject.RoleCode) error
	RevokeRole(ctx context.Context, userID valueobject.UserID, roleCode valueobject.RoleCode) error
}
