package usecases

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/services"
)

type InspectionUseCase struct {
	inspectionService services.InspectionService
}

func NewInspectionUseCase(is services.InspectionService) InspectionUseCase {
	return InspectionUseCase{
		inspectionService: is,
	}
}

func (u *InspectionUseCase) AddInspection(ctx context.Context, i entities.InspectionRequest) (*string, error) {
	return u.inspectionService.CreateInspection(ctx, i)
}
