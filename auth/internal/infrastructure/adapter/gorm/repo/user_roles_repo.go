package repo

import (
	"context"
	"errors"

	"clirzy/auth/internal/domain"
	"clirzy/auth/internal/domain/valueobject"
	"clirzy/auth/internal/infrastructure/adapter/gorm/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoleGormRepo struct {
	db *gorm.DB
}

func NewUserRoleGormRepo(db *gorm.DB) *UserRoleGormRepo {
	return &UserRoleGormRepo{db: db}
}

func (r *UserRoleGormRepo) GetRoles(ctx context.Context, userID valueobject.UserID) ([]valueobject.RoleCode, error) {
	var userRoles []model.UserRole
	if err := r.db.WithContext(ctx).Where("user_id = ?", string(userID)).Find(&userRoles).Error; err != nil {
		return nil, err
	}

	if len(userRoles) == 0 {
		return nil, nil
	}

	roleIDs := make([]string, 0, len(userRoles))
	for _, ur := range userRoles {
		roleIDs = append(roleIDs, ur.RoleID)
	}

	var roles []model.Role
	if err := r.db.WithContext(ctx).Where("id IN ?", roleIDs).Find(&roles).Error; err != nil {
		return nil, err
	}

	roleCodes := make([]valueobject.RoleCode, 0, len(roles))
	for _, role := range roles {
		roleCodes = append(roleCodes, valueobject.RoleCode(role.Code))
	}

	return roleCodes, nil
}

func (r *UserRoleGormRepo) AssignRole(ctx context.Context, userID valueobject.UserID, roleCode valueobject.RoleCode) error {
	var role model.Role
	if err := r.db.WithContext(ctx).Where("code = ?", string(roleCode)).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrRoleNotFound
		}
		return err
	}

	userRole := model.UserRole{
		ID:     uuid.New().String(),
		UserID: userID.String(),
		RoleID: role.ID,
	}

	return r.db.WithContext(ctx).Create(&userRole).Error
}

func (r *UserRoleGormRepo) RevokeRole(ctx context.Context, userID valueobject.UserID, roleCode valueobject.RoleCode) error {
	var role model.Role
	if err := r.db.WithContext(ctx).Where("code = ?", string(roleCode)).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrRoleNotFound
		}
		return err
	}

	return r.db.WithContext(ctx).
		Where("user_id = ? AND role_id = ?", string(userID), role.ID).
		Delete(&model.UserRole{}).Error
}
