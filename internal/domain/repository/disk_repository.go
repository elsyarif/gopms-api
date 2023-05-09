package repository

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
)

type DiskRepository interface {
	AddDisk(ctx context.Context, disk entities.Disk) error
	GetAllDiskByServerId(ctx context.Context, serverId string) (*[]entities.Disk, error)
	GetDiskById(ctx context.Context, diskId string) (*entities.Disk, error)
	EditDisk(ctx context.Context, diskId string, disk entities.Disk) error
	DeleteDisk(ctx context.Context, diskId string) error
}
