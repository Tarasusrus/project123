package main

import (
	"BaseApi/internal/config"
	"BaseApi/internal/database"
	"BaseApi/internal/server"
	"context"
	"log"
)

func main() {
	// Загрузка конфигурации
	cfg := loadConfig()

	// connections (database)
	_, err := database.NewGORM(cfg.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	var _ = context.Background()

	httpHandler := server.NewHandler(&cfg.AppConfig)
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
