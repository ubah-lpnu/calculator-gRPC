package logger

import (
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func initLogger() {
	logger, _ = zap.NewDevelopment()
}

func GetLogger() *zap.Logger {
	if logger == nil {
		initLogger()
	}
	return logger
}
