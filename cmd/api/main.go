// @title Базовое бекенд решение
// @version 1.0
// @description Простой сервер.
// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"BaseApi/internal/config"
	"BaseApi/internal/database"
	"BaseApi/internal/logger"
	"BaseApi/internal/server"
	"context"
	"log"
)

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

	// connections (database)
	_, err := database.NewGORM(cfg.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	var _ = context.Background()

	httpHandler := server.NewHandler(&cfg.AppConfig, logSlog)
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
