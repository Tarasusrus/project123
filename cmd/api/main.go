package main

import (
	"BaseApi/internal/config"
	"BaseApi/internal/database"
	baseLogger "BaseApi/internal/logger"
	"BaseApi/internal/server"
	"context"
	"log"
	"log/slog"
)

func main() {
	// Загрузка конфигурации
	cfg := loadConfig()

	// logger
	logger := baseLogger.SetUpLogger(cfg.AppConfig.Mode)
	logger.Info("started", slog.String("APP MODEz", cfg.AppConfig.Mode))
	logger.Debug("Debug enabled")

	// connections (database)
	_, err := database.NewGORM(cfg.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	var _ = context.Background()

	httpHandler := server.NewHandler(&cfg.AppConfig, logger)
	if err = httpHandler.Serve(); err != nil {
		log.Fatal(err)
	}

}

func loadConfig() *config.Env {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	return cfg

}
