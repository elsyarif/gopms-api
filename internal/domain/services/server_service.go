package services

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/repository"
	"github.com/elsyarif/pms-api/pkg/uid"
)

type ServerService struct {
	serverRepo  repository.ServerRepository
	idGenerator uid.NanoGenerator
}

func NewServerService(sr repository.ServerRepository, uid uid.NanoGenerator) ServerService {
	return ServerService{
		serverRepo:  sr,
		idGenerator: uid,
	}
}

func (s *ServerService) CreateServer(ctx context.Context, server entities.Server) (*entities.Server, error) {
	server.Id = s.idGenerator.NanoId("server")
	err := s.serverRepo.AddServer(ctx, server)
	if err != nil {
		return nil, err
	}

	return &server, nil
}

func (s *ServerService) GetAllServer(ctx context.Context, groupId string) (*[]entities.Server, error) {
	return s.serverRepo.GetAllServerByGroup(ctx, groupId)
}

func (s *ServerService) GetServerById(ctx context.Context, serverId string) (*entities.Server, error) {
	return s.serverRepo.GetServerById(ctx, serverId)
}

func (s *ServerService) EditServer(ctx context.Context, serverId string, server entities.Server) error {
	_, err := s.serverRepo.GetServerById(ctx, serverId)
	if err != nil {
		return err
	}
	return s.serverRepo.EditServer(ctx, serverId, server)
}

func (s *ServerService) DeleteServer(ctx context.Context, serverId string) error {
	_, err := s.serverRepo.GetServerById(ctx, serverId)
	if err != nil {
		return err
	}

	return s.serverRepo.DeleteServer(ctx, serverId)
}
