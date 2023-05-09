package repositories

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/repository"
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
	//TODO implement me
	panic("implement me")
}

func (d *DiskRepositoryPostgres) GetAllDiskByServerId(ctx context.Context, serverId string) (*[]entities.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DiskRepositoryPostgres) GetDiskById(ctx context.Context, diskId string) (*entities.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DiskRepositoryPostgres) EditDisk(ctx context.Context, diskId string, disk entities.Disk) error {
	//TODO implement me
	panic("implement me")
}

func (d *DiskRepositoryPostgres) DeleteDisk(ctx context.Context, diskId string) error {
	//TODO implement me
	panic("implement me")
}
