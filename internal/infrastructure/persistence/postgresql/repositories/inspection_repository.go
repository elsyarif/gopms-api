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
