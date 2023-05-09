package repositories

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/repository"
	"github.com/jmoiron/sqlx"
)

type ServerRepositoryPostgres struct {
	DB *sqlx.DB
}

func NewServerRepositoryPostgres(db *sqlx.DB) repository.ServerRepository {
	return &ServerRepositoryPostgres{
		DB: db,
	}
}

func (s *ServerRepositoryPostgres) AddServer(ctx context.Context, server entities.Server) error {
	//TODO implement me
	panic("implement me")
}

func (s *ServerRepositoryPostgres) VerifyServerGroup(ctx context.Context, groupId string) error {
	//TODO implement me
	panic("implement me")
}

func (s *ServerRepositoryPostgres) GetAllServerByGroup(ctx context.Context, groupId string) (*[]entities.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerRepositoryPostgres) GetServerById(ctx context.Context, serverId string) (*entities.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerRepositoryPostgres) EditServer(ctx context.Context, serverId string, server entities.Server) error {
	//TODO implement me
	panic("implement me")
}

func (s *ServerRepositoryPostgres) DeleteServer(ctx context.Context, serverId string) error {
	//TODO implement me
	panic("implement me")
}
