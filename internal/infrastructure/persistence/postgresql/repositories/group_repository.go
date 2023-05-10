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

type GroupRepositoryPostgres struct {
	DB *sqlx.DB
}

func NewGroupRepositoryPostgres(db *sqlx.DB) repository.GroupRepository {
	return &GroupRepositoryPostgres{
		DB: db,
	}
}

func (g *GroupRepositoryPostgres) AddGroup(ctx context.Context, group entities.Group) error {
	query := "INSERT INTO groups VALUES ($1, $2, $3)"

	tx, err := g.DB.Beginx()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, query, group.Id, group.Name, group.Description)
	if err != nil {
		log.Error("exec add group error", logrus.Fields{"error": err.Error()})
		_ = tx.Rollback()
		return err
	}

	if row, err := result.RowsAffected(); err == nil && row > 0 {
		_ = tx.Commit()
		return nil
	}

	return err
}

func (g *GroupRepositoryPostgres) GetAllGroup(ctx context.Context, name string) (*[]entities.Group, error) {
	query := "SELECT * FROM groups WHERE name like $1"
	var groups []entities.Group

	tx, err := g.DB.Beginx()
	if err != nil {
		return nil, err
	}

	err = tx.SelectContext(ctx, &groups, query, "%"+name+"%")
	if err != nil {
		log.Error("GetAllGroup error", logrus.Fields{"error": err.Error()})
		return nil, errors.New("group not found")
	}

	if len(groups) < 1 {
		groups = []entities.Group{}
	}

	return &groups, nil
}

func (g *GroupRepositoryPostgres) GetGroupById(ctx context.Context, groupId string) (*entities.Group, error) {
	query := "SELECT * FROM groups WHERE id = $1"
	group := entities.Group{}

	tx, err := g.DB.Beginx()
	if err != nil {
		return nil, err
	}

	err = tx.GetContext(ctx, &group, query, groupId)
	if err != nil {
		log.Error("GetGroupById error", logrus.Fields{"error": err.Error()})
		return nil, errors.New("group not found")
	}

	return &group, err
}

func (g *GroupRepositoryPostgres) EditGroup(ctx context.Context, groupId string, group entities.Group) error {
	query := "UPDATE groups SET name = $1, description = $2 WHERE id = $3"

	tx, err := g.DB.Beginx()
	if err != nil {
		return nil
	}
	result, err := tx.ExecContext(ctx, query, group.Name, group.Description, groupId)
	if err != nil {
		_ = tx.Rollback()
		log.Error("EditGroup error", logrus.Fields{"error": err.Error()})
		return err
	}

	if row, err := result.RowsAffected(); err == nil && row > 0 {
		_ = tx.Commit()
		return nil
	}

	return err
}

func (g *GroupRepositoryPostgres) DeleteGroup(ctx context.Context, groupId string) error {
	query := "DELETE FROM groups WHERE id = $1"

	tx, err := g.DB.Beginx()
	if err != nil {
		return nil
	}

	result, err := tx.ExecContext(ctx, query, groupId)
	if err != nil {
		_ = tx.Rollback()
		return errors.New("DeleteGroup error :" + err.Error())
	}

	if row, err := result.RowsAffected(); err == nil && row > 0 {
		_ = tx.Commit()
		return nil
	}

	return err
}
