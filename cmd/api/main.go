package main

import (
	"BaseApi/internal/config"
	"BaseApi/internal/database"
	"BaseApi/internal/logger"
	"BaseApi/internal/server"
	service2 "BaseApi/internal/service"
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
	service := service2.NewMusicService(db)

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
