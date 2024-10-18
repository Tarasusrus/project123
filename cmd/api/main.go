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

//TODO 1. У группы может быть много песен и хранить данные в куче в одной таблице - это нарушение нормализации БД
//TODO 2. Не стоит забывать про индексы при создании БД
//TODO 3. Дату релиза все же стоит хранить в соответствующем типе, а не в числе
//TODO 4. Не реализована интеграция с внешним АПИ
//TODO 5. Фильтрация не реализована
//TODO 6. Пагинация не реализована
//TODO 7. debug-логов нет
