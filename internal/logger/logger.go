package logger

import (
	"go.uber.org/zap"
)

func New() *zap.Logger {
	var opts []zap.Option
	zConfig := zap.NewDevelopmentConfig()
	opts = append(opts, zap.AddStacktrace(zap.ErrorLevel))
	logger, err := zConfig.Build(opts...)
	if err != nil {
		panic(err)
	}
	return logger
}
