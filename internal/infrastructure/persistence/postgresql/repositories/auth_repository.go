package repositories

import (
	"context"
	"errors"
	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/elsyarif/pms-api/internal/domain/repository"
	"github.com/elsyarif/pms-api/pkg/helper/log"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AuthRepositoryPostgres struct {
	DB *sqlx.DB
}

func NewAuthRepositoryPostgres(db *sqlx.DB) repository.AuthRepository {
	return &AuthRepositoryPostgres{
		DB: db,
	}
}

func (a *AuthRepositoryPostgres) AddToken(ctx context.Context, request entities.RefreshTokenRequest) error {
	query := "INSERT INTO authentications VALUES ($1)"

	tx, err := a.DB.Beginx()
	if err != nil {
		return err
	}

	auth := entities.Authentication{Token: request.RefreshToken}
	result, err := tx.ExecContext(ctx, query, auth.Token)
	if err != nil {
		log.Error("exec add token error", logrus.Fields{"error": err})
		_ = tx.Rollback()
		return err
	}

	if row, err := result.RowsAffected(); err == nil && row > 0 {
		_ = tx.Commit()
		return nil
	}

	return err
}

func (a *AuthRepositoryPostgres) CheckAvailabilityToken(ctx context.Context, request entities.RefreshTokenRequest) error {
	query := "SELECT token FROM authentications WHERE token = $1"
	token := entities.Authentication{}

	tx, err := a.DB.Beginx()
	if err != nil {
		return err
	}

	err = tx.GetContext(ctx, &token, query, request.RefreshToken)
	if err != nil {
		log.Error("CheckAvailabilityToken error", logrus.Fields{"error": err})
		return errors.New("token tidak ditemukan")
	}

	return nil
}

func (a *AuthRepositoryPostgres) DeleteToken(ctx context.Context, request entities.RefreshTokenRequest) error {
	query := "DELETE FROM authentications WHERE token = $1"

	tx, err := a.DB.Beginx()
	if err != nil {
		return err
	}

	auth := entities.Authentication{Token: request.RefreshToken}
	result, err := tx.ExecContext(ctx, query, auth.Token)
	if err != nil {
		log.Error("exec delete token error", logrus.Fields{"error": err})
		_ = tx.Rollback()
		return err
	}

	if row, err := result.RowsAffected(); err == nil && row > 0 {
		_ = tx.Commit()
		return nil
	}

	return nil
}
