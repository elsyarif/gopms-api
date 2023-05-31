package usecases

import (
	"context"
	"github.com/elsyarif/pms-api/internal/applications/report"
	"github.com/elsyarif/pms-api/internal/domain/services"
)

type ReportUseCase struct {
	InspectService services.InspectionService
}

func NewReportUseCase(is services.InspectionService) ReportUseCase {
	return ReportUseCase{
		InspectService: is,
	}
}

func (u *ReportUseCase) ExportExcel(ctx context.Context, groupId string, periodStart string, periodEnd string) error {
	inspections, err := u.InspectService.GetInspectionByGroupId(ctx, groupId, periodStart, periodEnd)
	if err != nil {
		return err
	}

	err = report.InspectionExcel(inspections)
	if err != nil {
		return err
	}

	return nil
}
