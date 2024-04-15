// logger.go
package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger inicializa o logger com as configurações específicas.
func InitLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	logger, _ := config.Build()

	// Defere a sincronização do logger
	defer logger.Sync()

	return logger
}
