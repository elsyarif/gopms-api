package services

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/repository"
	"github.com/elsyarif/pms-api/pkg/uid"
)

type DiskService struct {
	diskRepo    repository.DiskRepository
	idGenerator uid.NanoGenerator
}

func NewDiskService(dr repository.DiskRepository, uid uid.NanoGenerator) DiskService {
	return DiskService{
		diskRepo:    dr,
		idGenerator: uid,
	}
}

func (s *DiskService) CreateDisk(ctx context.Context, disk entities.Disk) (*entities.Disk, error) {
	disk.Id = s.idGenerator.NanoId("disk")
	err := s.diskRepo.AddDisk(ctx, disk)
	if err != nil {
		return nil, err
	}

	return &disk, nil
}

func (s *DiskService) GetAllDiskByServerId(ctx context.Context, serverId string) (*[]entities.Disk, error) {
	return s.diskRepo.GetAllDiskByServerId(ctx, serverId)
}

func (s *DiskService) GetDiskById(ctx context.Context, diskId string) (*entities.Disk, error) {
	return s.diskRepo.GetDiskById(ctx, diskId)
}

func (s *DiskService) EditDisk(ctx context.Context, diskId string, disk entities.Disk) error {
	_, err := s.diskRepo.GetDiskById(ctx, diskId)
	if err != nil {
		return err
	}
	return s.diskRepo.EditDisk(ctx, diskId, disk)
}

func (s *DiskService) DeleteDisk(ctx context.Context, diskId string) error {
	_, err := s.diskRepo.GetDiskById(ctx, diskId)
	if err != nil {
		return err
	}
	return s.diskRepo.DeleteDisk(ctx, diskId)
}
