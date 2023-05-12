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

type DiskRepositoryPostgres struct {
	DB *sqlx.DB
}

func NewDiskRepositoryPostgres(db *sqlx.DB) repository.DiskRepository {
	return &DiskRepositoryPostgres{
		DB: db,
	}
}

func (d *DiskRepositoryPostgres) AddDisk(ctx context.Context, disk entities.Disk) error {
	query := "INSERT INTO disks VALUES ($1, $2, $3, $4)"

	tx, err := d.DB.Beginx()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, query, disk.Id, disk.ServerId, disk.Name, disk.Total)
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

func (d *DiskRepositoryPostgres) GetAllDiskByServerId(ctx context.Context, serverId string) (*[]entities.Disk, error) {
	query := "SELECT * FROM disks WHERE server_id = $1"
	var disks []entities.Disk

	tx, err := d.DB.Beginx()
	if err != nil {
		return nil, err
	}

	err = tx.SelectContext(ctx, &disks, query, serverId)
	if err != nil {
		return nil, err
	}

	if len(disks) < 1 {
		disks = []entities.Disk{}
	}

	return &disks, nil
}

func (d *DiskRepositoryPostgres) GetDiskById(ctx context.Context, diskId string) (*entities.Disk, error) {
	query := "SELECT * FROM disks WHERE id = $1"
	disk := entities.Disk{}

	tx, err := d.DB.Beginx()
	if err != nil {
		return nil, err
	}

	err = tx.GetContext(ctx, &disk, query, diskId)
	if err != nil {
		return nil, errors.New("disk not found")
	}

	return &disk, nil
}

func (d *DiskRepositoryPostgres) EditDisk(ctx context.Context, diskId string, disk entities.Disk) error {
	query := "UPDATE disks SET name = $1 , total = $2 WHERE id = $3"

	tx, err := d.DB.Beginx()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, query, disk.Name, disk.Total, diskId)
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

func (d *DiskRepositoryPostgres) DeleteDisk(ctx context.Context, diskId string) error {
	query := "DELETE FROM disks WHERE id = $1"

	tx, err := d.DB.Beginx()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, query, diskId)
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
