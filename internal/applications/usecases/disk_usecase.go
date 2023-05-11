package usecases

import (
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/services"
	"golang.org/x/net/context"
)

type DiskUseCase struct {
	diskService services.DiskService
}

func NewDiskUseCase(du services.DiskService) DiskUseCase {
	return DiskUseCase{
		diskService: du,
	}
}

func (u *DiskUseCase) AddDisk(ctx context.Context, disk entities.Disk) (*entities.Disk, error) {
	return u.diskService.CreateDisk(ctx, disk)
}

func (u *DiskUseCase) GetDiskById(ctx context.Context, diskId string) (*entities.Disk, error) {
	return u.diskService.GetDiskById(ctx, diskId)
}

func (u *DiskUseCase) EditDisk(ctx context.Context, diskId string, disk entities.Disk) error {
	return u.diskService.EditDisk(ctx, diskId, disk)
}

func (u *DiskUseCase) DeleteDisk(ctx context.Context, diskId string) error {
	return u.diskService.DeleteDisk(ctx, diskId)
}
