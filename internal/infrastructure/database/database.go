package database

import (
	"fmt"
	"github.com/elsyarif/pms-api/pkg/config"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"time"
)

func LoadConfig() error {
	err := envconfig.Process("app", &config.Conf)
	if err != nil {
		return err
	}
	return nil
}

func NewConnectPostgres() (*sqlx.DB, error) {
	ds := fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=disable", config.Conf.DBUser, config.Conf.DBPass, config.Conf.DBName)
	db, err := sqlx.Open("postgres", ds)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
