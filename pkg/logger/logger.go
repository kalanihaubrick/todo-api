package logger

import (
	"log"
	"log/slog"
	"os"
)

var Logger *slog.Logger

func init() {
	handler := slog.NewTextHandler(os.Stdout, nil)
	Logger = slog.New(handler)
}

func Info(msg string, keysAndValues ...interface{}) {
	Logger.Info(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	Logger.Error(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	Logger.Error(msg, keysAndValues...)
	log.Fatal(msg)
}
