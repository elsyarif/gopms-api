package repositories

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/repository"
	"github.com/jmoiron/sqlx"
)

type GroupRepositoryPostgres struct {
	DB *sqlx.DB
}

func NewGroupRepositoryPostgres(db *sqlx.DB) repository.GroupRepository {
	return &GroupRepositoryPostgres{
		DB: db,
	}
}

func (g *GroupRepositoryPostgres) AddGroup(ctx context.Context, group entities.Group) error {
	//TODO implement me
	panic("implement me")
}

func (g *GroupRepositoryPostgres) GetAllGroup(ctx context.Context, name string) (*[]entities.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GroupRepositoryPostgres) GetGroupById(ctx context.Context, groupId string) (*entities.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GroupRepositoryPostgres) EditGroup(ctx context.Context, groupId string, group entities.Group) error {
	//TODO implement me
	panic("implement me")
}

func (g *GroupRepositoryPostgres) DeleteGroup(ctx context.Context, groupId string) error {
	//TODO implement me
	panic("implement me")
}
