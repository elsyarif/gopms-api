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

	for i1, i := range inspection.InspectionDetail {
		inspection.InspectionDetail[i1].Id = s.idGenerator.NanoId("i-detail")

		for i2, _ := range i.InspectionDisk {
			inspection.InspectionDetail[i1].InspectionDisk[i2].Id = s.idGenerator.NanoId("i-disk")
		}
	}

	err := s.inspectionRepo.AddInspection(ctx, inspection)
	if err != nil {
		return nil, err
	}

	return &inspection.Id, nil
}

func (s *InspectionService) GetInspectionByGroupId(ctx context.Context, groupId string, periodStart string, periodEnd string) (*entities.InspectionResponse, error) {
	return s.inspectionRepo.GetInspection(ctx, groupId, periodStart, periodEnd)
}
