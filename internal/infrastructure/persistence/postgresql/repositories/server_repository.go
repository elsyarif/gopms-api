package repositories

import (
	"context"
	"errors"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/repository"
	"github.com/elsyarif/pms-api/pkg/helper/log"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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
	query := "INSERT INTO servers VALUES ($1, $2, $3, $4, $5, $6, $7)"

	tx, err := s.DB.Beginx()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, query, server.Id, server.GroupId, server.ServerName, server.Location, server.Status, server.Memory, server.Ip)
	if err != nil {
		_ = tx.Rollback()
		log.Error("AddServer error", logrus.Fields{"error": err.Error()})
		return err
	}

	if row, err := result.RowsAffected(); err == nil && row > 0 {
		_ = tx.Commit()
		return nil
	}

	return err
}

func (s *ServerRepositoryPostgres) VerifyServerGroup(ctx context.Context, serverId, groupId string) error {
	query := "SELECT * FROM servers where id = $1 AND group_id = $1"
	server := entities.Server{}

	tx, err := s.DB.Beginx()
	if err != nil {
		return err
	}

	err = tx.GetContext(ctx, &server, query, serverId, groupId)
	if err != nil {
		return errors.New("akses tidak sah")
	}

	return nil
}

func (s *ServerRepositoryPostgres) GetAllServerByGroup(ctx context.Context, groupId string) (*[]entities.Server, error) {
	query := "SELECT * FROM servers WHERE group_id = $1"
	var servers []entities.Server

	tx, err := s.DB.Beginx()
	if err != nil {
		return nil, err
	}

	err = tx.SelectContext(ctx, &servers, query, groupId)
	if err != nil {
		return nil, err
	}

	if len(servers) < 1 {
		servers = []entities.Server{}
	}

	return &servers, nil
}

func (s *ServerRepositoryPostgres) GetServerById(ctx context.Context, serverId string) (*entities.Server, error) {
	query := "SELECT * FROM servers WHERE id = $1"
	server := entities.Server{}

	tx, err := s.DB.Beginx()
	if err != nil {
		return nil, err
	}

	err = tx.GetContext(ctx, &server, query, serverId)
	if err != nil {
		return nil, errors.New("server not found")
	}

	return &server, nil
}

func (s *ServerRepositoryPostgres) EditServer(ctx context.Context, serverId string, server entities.Server) error {
	query := "UPDATE servers set server_name = $1, location = $2, status = $3, memory = $4, ip= $5 WHERE id = $6"

	tx, err := s.DB.Beginx()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, query, server.ServerName, server.Location, server.Status, server.Memory, server.Ip, serverId)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if row, err := result.RowsAffected(); err == nil && row > 0 {
		_ = tx.Commit()
		return nil
	}

	return err
}

func (s *ServerRepositoryPostgres) DeleteServer(ctx context.Context, serverId string) error {
	query := "DELETE FROM servers WHERE id = $1"

	tx, err := s.DB.Beginx()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, query, serverId)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if row, err := result.RowsAffected(); err == nil && row > 0 {
		_ = tx.Commit()
		return nil
	}

	return err
}
