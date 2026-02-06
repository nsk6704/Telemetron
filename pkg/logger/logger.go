package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func Init(level string) error {
	var config zap.Config

	if level == "debug" {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	var err error
	Log, err = config.Build()
	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	if Log != nil {
		return Log.Sync()
	}
	return nil
}
