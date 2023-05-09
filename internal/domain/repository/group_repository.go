package repository

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
)

type GroupRepository interface {
	AddGroup(ctx context.Context, group entities.Group) error
	GetAllGroup(ctx context.Context, name string) (*[]entities.Group, error)
	GetGroupById(ctx context.Context, groupId string) (*entities.Group, error)
	EditGroup(ctx context.Context, groupId string, group entities.Group) error
	DeleteGroup(ctx context.Context, groupId string) error
}
