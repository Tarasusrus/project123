package config

import (
	"BaseApi/internal/database"
	"BaseApi/internal/server"
	"github.com/spf13/viper"
)

type Env struct {
	DBConfig  database.Config
	AppConfig server.AppCfg
}

func Init() (*Env, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Env

	if err := viper.UnmarshalKey("pg_config", &cfg.DBConfig); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("app_config", &cfg.AppConfig); err != nil {
		return nil, err
	}
	return &cfg, nil
}
