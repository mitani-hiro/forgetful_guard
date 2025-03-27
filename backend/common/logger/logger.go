package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

// InitLogger ロギング初期化.
func InitLogger() {
	Logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
	Info("Log initialization was successful.")
}

// Info Infoレベルのログ出力.
func Info(msg string, args ...any) {
	Logger.Info(msg, args...)
}

// Error Errorレベルのログ出力.
func Error(msg string, err any) {
	Logger.Error(msg, "error", err)
}
