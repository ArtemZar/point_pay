package utils

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitializeLogger() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	defer logger.Sync() // flushes buffer, if any
	Logger = logger.Sugar()

	return nil
}
