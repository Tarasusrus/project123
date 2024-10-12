package logger

type Logger interface {
	Log(message string, level Level, ctx ...any)
	Debug(message string, ctx ...any)
	Info(message string, ctx ...any)
	Warn(message string, ctx ...any)
	Error(message string, ctx ...any)
	Fatal(message string, ctx ...any)
	Panic(message string, ctx ...any)
	GetLogger() any
}

type Level int

const (
	Debug Level = -4
	Info  Level = 0
	Warn  Level = 4
	Error Level = 8
	Fatal Level = 12
	Panic Level = 16
)
