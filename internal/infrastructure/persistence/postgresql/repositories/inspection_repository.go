package repositories

import (
	"context"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/repository"
	"github.com/jmoiron/sqlx"
)

type InspectionRepositoryPostgres struct {
	DB *sqlx.DB
}

func NewInspectionRepositoryPostgres(db *sqlx.DB) repository.InspectionRepository {
	return &InspectionRepositoryPostgres{
		DB: db,
	}
}

func (ir *InspectionRepositoryPostgres) AddInspection(ctx context.Context, i entities.InspectionRequest) error {
	q1 := "INSERT INTO inspections VALUES ($1, $2, $3, $4, $5, $6, $7)"
	q2 := "INSERT INTO inspection_server VALUES ($1, $2, $3, $4, $5, $6)"
	q3 := "INSERT INTO inspection_disk VALUES ($1, $2, $3, $4, $5, $6)"

	tx, err := ir.DB.Beginx()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, q1, i.Id, i.GroupId, i.GroupName, i.Date, i.UserBy, i.PeriodStart, i.PeriodEnd)
	if err != nil {
		return err
	}

	for _, id := range i.InspectionDetail {
		result, err = tx.ExecContext(ctx, q2, id.Id, i.Id, id.ServerId, id.ServerName, id.CpuUsage, id.MemoryUsage)
		if err != nil {
			return err
		}

		for _, ids := range id.InspectionDisk {
			result, err = tx.ExecContext(ctx, q3, ids.Id, i.Id, id.Id, ids.DiskId, ids.DiskName, ids.DiskUsage)
			if err != nil {
				return err
			}
		}
	}

	if row, err := result.RowsAffected(); err != nil && row <= 0 {
		_ = tx.Rollback()
		return err
	}

	_ = tx.Commit()
	return nil
}

func (ir *InspectionRepositoryPostgres) GetInspection(ctx context.Context, groupId string, periodStart string, periodEnd string) (*entities.InspectionResponse, error) {
	q1 := "SELECT * FROM inspections WHERE group_id = $1 AND period_start = $2 AND period_end = $3"
	q2 := "SELECT * FROM inspection_server WHERE inspection_id = $1"
	q3 := "SELECT * FROM inspection_disk WHERE inspection_id = $1"

	response := entities.InspectionResponse{}
	inspect := entities.Inspection{}
	var insServer []entities.InspectionServer
	var insDisk []entities.InspectionDisk

	tx, err := ir.DB.Beginx()
	if err != nil {
		return nil, err
	}

	err = tx.GetContext(ctx, &inspect, q1, groupId, periodStart, periodEnd)
	if err != nil {
		return nil, err
	}
	tx.SelectContext(ctx, &insServer, q2, inspect.Id)
	tx.SelectContext(ctx, &insDisk, q3, inspect.Id)

	response.Id = inspect.Id
	response.GroupId = inspect.GroupId
	response.GroupName = inspect.GroupName
	response.Date = inspect.Date
	response.UserBy = inspect.UserBy
	response.PeriodStart = inspect.PeriodStart
	response.PeriodEnd = inspect.PeriodEnd

	for i, ss := range insServer {
		response.InspectionDetail = append(response.InspectionDetail, entities.InspectionRequestDetail{
			Id:          ss.Id,
			ServerId:    ss.ServerId,
			ServerName:  ss.ServerName,
			CpuUsage:    ss.CpuUsage,
			MemoryUsage: ss.MemoryUsage,
		})
		for _, dd := range insDisk {
			if dd.InspectionServerId == ss.Id {
				response.InspectionDetail[i].InspectionDisk = append(response.InspectionDetail[i].InspectionDisk, entities.InspectionDisk{
					Id:                 dd.Id,
					InspectionId:       dd.InspectionId,
					InspectionServerId: dd.InspectionServerId,
					DiskId:             dd.DiskId,
					DiskName:           dd.DiskName,
					DiskUsage:          dd.DiskUsage,
				})
			}
		}
	}
	return &response, nil
}
