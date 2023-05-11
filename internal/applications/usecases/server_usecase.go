package usecases

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/services"
)

type ServerUseCase struct {
	serverService services.ServerService
	diskService   services.DiskService
}

func NewServerUseCase(su services.ServerService, ds services.DiskService) ServerUseCase {
	return ServerUseCase{
		serverService: su,
		diskService:   ds,
	}
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

func (u *ServerUseCase) GetServerDisk(ctx context.Context, serverId string) (*entities.ServerDisk, error) {
	server, err := u.serverService.GetServerById(ctx, serverId)
	if err != nil {
		return nil, err
	}
	disks, err := u.diskService.GetAllDiskByServerId(ctx, serverId)
	if err != nil {
		return nil, err
	}

	sd := entities.ServerDisk{
		Id:         server.Id,
		GroupId:    server.GroupId,
		ServerName: server.ServerName,
		Location:   server.Location,
		Status:     server.Status,
		Memory:     server.Memory,
		Ip:         server.Id,
		Disk:       disks,
	}

	return &sd, nil
}
