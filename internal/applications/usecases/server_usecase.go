package usecases

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/services"
)

type ServerUseCase struct {
	serverService services.ServerService
}

func NewServerUseCase(su services.ServerService) ServerUseCase {
	return ServerUseCase{serverService: su}
}

func (u *ServerUseCase) AddServer(ctx context.Context, server entities.Server) (*entities.Server, error) {
	return u.serverService.CreateServer(ctx, server)
}

func (u *ServerUseCase) GetAllServerByGroup(ctx context.Context, groupId string) (*[]entities.Server, error) {
	return u.serverService.GetAllServer(ctx, groupId)
}

func (u *ServerUseCase) GetServerById(ctx context.Context, serverId string) (*entities.Server, error) {
	return u.serverService.GetServerById(ctx, serverId)
}

func (u *ServerUseCase) EditServer(ctx context.Context, serverId string, server entities.Server) error {
	return u.serverService.EditServer(ctx, serverId, server)
}

func (u *ServerUseCase) DeleteServer(ctx context.Context, serverId string) error {
	return u.serverService.DeleteServer(ctx, serverId)
}
