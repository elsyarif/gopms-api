package repositories

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/repository"
	"github.com/jmoiron/sqlx"
)

type PermissionRepositoryPostgres struct {
	DB *sqlx.DB
}

func NewPermissionRepositoryPostgres(db *sqlx.DB) repository.PermissionRepository {
	return &PermissionRepositoryPostgres{
		DB: db,
	}
}

func (p *PermissionRepositoryPostgres) AddPermission(ctx context.Context, permission entities.Permission) error {
	//TODO implement me
	panic("implement me")
}

func (p *PermissionRepositoryPostgres) GetAllPermission(ctx context.Context, name string) (*[]entities.Permission, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PermissionRepositoryPostgres) GetPermissionById(ctx context.Context, permissionId string) (*entities.Permission, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PermissionRepositoryPostgres) EditPermission(ctx context.Context, permissionId string, permission entities.Permission) error {
	//TODO implement me
	panic("implement me")
}

func (p *PermissionRepositoryPostgres) DeletePermission(ctx context.Context, permissionId string) error {
	//TODO implement me
	panic("implement me")
}
