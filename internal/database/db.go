package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `mapstructure:"DB_HOST" default:"localhost"`
	Port     string `mapstructure:"DB_PORT" default:"5432"`
	User     string `mapstructure:"DB_USER" default:"seleroad"`
	DBName   string `mapstructure:"DB_NAME" default:"wbadmin"`
	Password string `mapstructure:"DB_PASSWORD" default:"seleroad"`
	SSLMode  string `mapstructure:"SSL_MODE" default:"disable"`
}

func NewGORM(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil

}
