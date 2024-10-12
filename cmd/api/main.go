package main

import (
	"BaseApi/internal/config"
	"BaseApi/internal/database"
	"BaseApi/internal/server"
	"BaseApi/tools/logger"
	"context"
	"log"
)

func main() {
	// Загрузка конфигурации
	cfg := loadConfig()

	// slogLogger
	logSlog := logger.SetUpLogger(
		cfg.AppConfig.Mode,
		cfg.LogConfig.Level,
		cfg.LogConfig.Format,
		cfg.LogConfig.AddSource)

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

// todo Добавить миграцию
// todo написать контроллеры
// todo написать логику
// todo написать репозиторий
// todo написать доки
// todo по-тестить
