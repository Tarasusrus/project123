package logger

import (
	"BaseApi/internal/logger"
	"log/slog"
)

// Реализация интерфейса Logger на базе slog
type SlogLogger struct {
	logger *slog.Logger
}

// Создаем новый SlogLogger
func NewSlogLogger(baseLogger *slog.Logger) *SlogLogger {
	return &SlogLogger{
		logger: baseLogger,
	}
}

func (sl *SlogLogger) Log(message string, level logger.Level, ctx ...any) {
	switch level {
	case logger.Debug:
		sl.Debug(message, ctx...)
	case logger.Info:
		sl.Info(message, ctx...)
	case logger.Warn:
		sl.Warn(message, ctx...)
	case logger.Error:
		sl.Error(message, ctx...)
	case logger.Fatal:
		sl.Fatal(message, ctx...)
	case logger.Panic:
		sl.Panic(message, ctx...)
	}
}

func (sl *SlogLogger) Debug(message string, ctx ...any) {
	sl.logger.Debug(message, ctx...)
}

func (sl *SlogLogger) Info(message string, ctx ...any) {
	sl.logger.Info(message, ctx...)
}

func (sl *SlogLogger) Warn(message string, ctx ...any) {
	sl.logger.Warn(message, ctx...)
}

func (sl *SlogLogger) Error(message string, ctx ...any) {
	sl.logger.Error(message, ctx...)
}

func (sl *SlogLogger) Fatal(message string, ctx ...any) {
	sl.logger.Error(message, ctx...)
}

func (sl *SlogLogger) Panic(message string, ctx ...any) {
	sl.logger.Error(message, ctx...)
}

// Метод для получения конкретного slog.Logger
func (sl *SlogLogger) GetLogger() any {
	return sl.logger
}
