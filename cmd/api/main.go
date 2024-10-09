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
	logSlog := setupLogger(cfg)
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

func setupLogger(cfg *config.Env) *logger.SlogLogger {
	// Инициализируем логгер на основе конфигурации
	baseLogger := logger.SetUpLogger(cfg.AppConfig.Mode, cfg.LogConfig.Level, cfg.LogConfig.Format, cfg.LogConfig.AddSource)
	return logger.NewSlogLogger(baseLogger)
}

// todo добавить логирования БД
// todo Добавить миграцию
// todo написать контроллеры
// todo написать логику
// todo написать репозиторий
// todo написать доки
// todo по-тестить
