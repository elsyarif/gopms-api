package config

import "time"

type config struct {
	DBUser          string        `default:"postgres"`
	DBPass          string        `default:"postgres"`
	DBHost          string        `default:"localhost"`
	DBPort          string        `default:"5432"`
	DBName          string        `default:"authapi"`
	DBTimeout       time.Duration `default:"10s"`
	Host            string        `default:"127.0.0.1"`
	Port            uint          `default:"5000"`
	ShutdownTimeout time.Duration `default:"20s"`
	AccessTokenKey  string        `default:"secret"`
	RefreshTokenKey string        `default:"secret"`
	AccessTokenAge  time.Duration `default:"100ns"`
}

var Conf config
