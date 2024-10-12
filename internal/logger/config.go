package logger

import (
	"log"
	"log/slog"
	"os"
)

type LogConfig struct {
	Level     string `mapstructure:"level"`
	Format    string `mapstructure:"format"`
	AddSource string `mapstructure:"add_source"`
}

// SetUpLogger - основная функция для настройки логгера
func SetUpLogger(env, level, format, addSource string) *slog.Logger {
	var logLevel slog.Level

	addSourceFlag := false
	if addSource == "true" {
		addSourceFlag = true
	}

	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	var logOutput *os.File
	var err error
	if env == "local" {
		logOutput = os.Stdout
	} else {
		// Все кроме local пишем в файл
		logOutput, err = os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Println("Не удалось открыть файл для логов: " + err.Error() + ". Логи будут записываться в консоль.")
			logOutput = os.Stdout
		}
	}

	var handler slog.Handler
	switch format {
	case "json":
		handler = slog.NewJSONHandler(logOutput, &slog.HandlerOptions{
			Level:     logLevel,
			AddSource: addSourceFlag,
		})
	case "text":
		handler = slog.NewTextHandler(logOutput, &slog.HandlerOptions{
			Level:     logLevel,
			AddSource: addSourceFlag,
		})
	default:
		handler = slog.NewTextHandler(logOutput, &slog.HandlerOptions{
			Level:     logLevel,
			AddSource: false,
		})
	}
	return slog.New(handler)
}
