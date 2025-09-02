package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var (
	log  *zap.Logger
	once sync.Once // Singleton Pattern
	err  error
)

func GetLogger() *zap.Logger {
	// Apply singleton
	once.Do(initLogger)
	return log
}

func initLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"app.log", "stdout"}
	config.ErrorOutputPaths = []string{"error.log"}
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	log, err = config.Build()
	if err != nil {
		panic(err)
	}
}
