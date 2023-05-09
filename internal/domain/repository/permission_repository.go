package repository

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
)

type PermissionRepository interface {
	AddPermission(ctx context.Context, permission entities.Permission) error
	GetAllPermission(ctx context.Context, name string) (*[]entities.Permission, error)
	GetPermissionById(ctx context.Context, permissionId string) (*entities.Permission, error)
	EditPermission(ctx context.Context, permissionId string, permission entities.Permission) error
	DeletePermission(ctx context.Context, permissionId string) error
}
