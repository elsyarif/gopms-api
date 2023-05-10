package repositories

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/repository"
	"github.com/elsyarif/pms-api/internal/infrastructure/persistence/postgresql"
	"github.com/jmoiron/sqlx"
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

	err := postgresql.Insert(d.DB, ctx, query, disk)
	if err != nil {
		return err
	}

	return nil
}

func (d *DiskRepositoryPostgres) GetAllDiskByServerId(ctx context.Context, serverId string) (*[]entities.Disk, error) {
	query := "SELECT * FROM disks WHERE server_id = $1"
	var disks []entities.Disk

	err := postgresql.SelectAll(d.DB, ctx, query, disks, serverId)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &disk, nil
}

func (d *DiskRepositoryPostgres) EditDisk(ctx context.Context, diskId string, disk entities.Disk) error {
	query := "UPDATE FROM disks SET name = $1 , total = $2 WHERE id = $3"

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
	query := "DELETE disks WHERE id = $1"

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
