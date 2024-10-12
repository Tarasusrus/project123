package main

import (
	"BaseApi/internal/config"
	"BaseApi/internal/database"
	"BaseApi/internal/logger"
	"BaseApi/internal/server"
	"BaseApi/internal/service"
	"context"
	"log"
)

// @title Music API
// @version 1.0
// @description Это API для управления музыкой.
// @host localhost:8080
// @BasePath /api/v1

func main() {
	// Загрузка конфигурации
	cfg := loadConfig()

	baseLogger := logger.SetUpLogger(
		cfg.AppConfig.Mode,
		cfg.LogConfig.Level,
		cfg.LogConfig.Format,
		cfg.LogConfig.AddSource)
	logSlog := logger.NewLogger(baseLogger)

	logSlog.Info("logger start")
	db, err := database.NewGORM(cfg.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	logSlog.Info("db start")

	if err := db.RunMigrations(); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}
	logSlog.Info("Migrations ok")

	var _ = context.Background()
	service := service.NewMusicService(db)

	httpHandler := server.NewHandler(&cfg.AppConfig, logSlog, service)
	if err = httpHandler.Serve(); err != nil {
		log.Fatal(err)
	}

	logSlog.Info("httpHandler ok")

}

func loadConfig() *config.Env {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	return cfg

}

// todo написать доки
// todo по-тестить
