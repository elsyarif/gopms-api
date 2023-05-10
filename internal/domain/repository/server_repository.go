package repository

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
)

type ServerRepository interface {
	AddServer(ctx context.Context, server entities.Server) error
	VerifyServerGroup(ctx context.Context, serverId, groupId string) error
	GetAllServerByGroup(ctx context.Context, groupId string) (*[]entities.Server, error)
	GetServerById(ctx context.Context, serverId string) (*entities.Server, error)
	EditServer(ctx context.Context, serverId string, server entities.Server) error
	DeleteServer(ctx context.Context, serverId string) error
}
