package repository

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
)

type InspectionRepository interface {
	AddInspection(ctx context.Context, i entities.InspectionRequest) error
	GetInspection(ctx context.Context, groupId string, periodStart string, periodEnd string) (*entities.InspectionResponse, error)
}
