package services

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/repository"
	"github.com/elsyarif/pms-api/pkg/uid"
)

type InspectionService struct {
	inspectionRepo repository.InspectionRepository
	idGenerator    uid.NanoGenerator
}

func NewInspectionService(ir repository.InspectionRepository, uid uid.NanoGenerator) InspectionService {
	return InspectionService{
		inspectionRepo: ir,
		idGenerator:    uid,
	}
}

func (s *InspectionService) CreateInspection(ctx context.Context, inspection entities.InspectionRequest) (*string, error) {
	inspection.Id = s.idGenerator.NanoId("inspect")

	for _, i := range inspection.InspectionDetail {
		i.Id = s.idGenerator.NanoId("i-detail")

		for _, j := range i.InspectionDisk {
			j.Id = s.idGenerator.NanoId("i-disk")
		}
	}

	err := s.inspectionRepo.AddInspection(ctx, inspection)
	if err != nil {
		return nil, err
	}

	return &inspection.Id, nil
}
