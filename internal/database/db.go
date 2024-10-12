package database

import (
	"BaseApi/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `mapstructure:"DB_HOST" default:"localhost"`
	Port     string `mapstructure:"DB_PORT" default:"5432"`
	User     string `mapstructure:"DB_USER" default:"postgres"`
	DBName   string `mapstructure:"DB_NAME" default:"postgres"`
	Password string `mapstructure:"DB_PASSWORD" default:"pass"`
	SSLMode  string `mapstructure:"SSL_MODE" default:"disable"`
}

type Database struct {
	client *gorm.DB
}

func NewGORM(cfg Config) (*Database, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Database{client: db}, nil

}

// RunMigrations выполняет миграции для всех моделей
func (d *Database) RunMigrations() error {
	return d.client.AutoMigrate(
		&models.Song{},
	)
}

func (d *Database) GetClient() *gorm.DB {
	return d.client
}

func (d *Database) Close() error {
	db, err := d.client.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
