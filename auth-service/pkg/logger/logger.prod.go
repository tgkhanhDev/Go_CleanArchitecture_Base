//go:build prod
// +build prod

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

func InitLogger() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"app.log", "stdout"}
	config.ErrorOutputPaths = []string{"error.log"}
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	log, err = config.Build()
	if err != nil {
		panic(err)
	}
}

func GetLogger() *zap.Logger {
	// Apply singleton
	once.Do(InitLogger)
	return log
}
