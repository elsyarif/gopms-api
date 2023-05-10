package postgresql

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func Insert(s *sqlx.DB, ctx context.Context, query string, args ...any) error {
	tx, err := s.Beginx()
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, query, args)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if row, err := result.RowsAffected(); err == nil && row > 0 {
		_ = tx.Commit()
		return nil
	}

	return err
}

func SelectAll(s *sqlx.DB, ctx context.Context, query string, dest interface{}, args ...any) error {
	tx, err := s.Beginx()
	if err != nil {
		return err
	}

	err = tx.SelectContext(ctx, &dest, query, args)
	if err != nil {
		return err
	}

	return nil
}
