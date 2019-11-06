package main

import (
	"fmt"

	"github.com/8bitdogs/log"
	"github.com/antonmashko/envconf"
)

type DBConfig struct {
	Host     string `env:"DB_HOST" default:"localhost"`
	Port     int    `env:"DB_PORT" default:"5432"`
	User     string `env:"DB_USER" default:"postgres"`
	Password string `env:"DB_PASSWORD" default:"pqpass"`
	Name     string `env:"DB_NAME" default:"postgres"`
	SSLMode  string `env:"DB_SSL_MODE" default:"disable"`
}

type config struct {
	Server struct {
		Addr string `env:"SERVER_ADDR" default:":8080"`
	}
	DatabaseDriveName string `default:"postgres"`
	Database          DBConfig
}

func (db DBConfig) ConnectionString() string {
	// NOTE: this connection string is for postgres. It may not work on other database
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.Name, db.SSLMode)
}

func parse() (*config, error) {
	var cfg config
	envconf.SetLogger(log.DefaultPrinters().InfoPrinter())
	err := envconf.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
