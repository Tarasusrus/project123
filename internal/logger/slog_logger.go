package logger

import (
	"log/slog"
)

// Реализация интерфейса Logger на базе slog
type SlogLogger struct {
	logger *slog.Logger
}

// Создаем новый SlogLogger
func NewLogger(baseLogger *slog.Logger) Logger {
	return &SlogLogger{
		logger: baseLogger,
	}
}

func (sl *SlogLogger) Log(message string, level Level, ctx ...any) {
	switch level {
	case Debug:
		sl.Debug(message, ctx...)
	case Info:
		sl.Info(message, ctx...)
	case Warn:
		sl.Warn(message, ctx...)
	case Error:
		sl.Error(message, ctx...)
	case Fatal:
		sl.Fatal(message, ctx...)
	case Panic:
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
