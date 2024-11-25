package loggers

import (
	"test-sms-2-pro/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger(cfg config.App) {
	var config zap.Config
	if cfg.Env == "dev" {
		config = zap.NewDevelopmentConfig()

	} else {
		config = zap.NewProductionConfig()
	}
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""
	log, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger = log
	defer logger.Sync()
}

func Info(msg string, field ...zapcore.Field) {
	logger.Info(msg, field...)
}
func Error(msg string, field ...zapcore.Field) {
	logger.Error(msg, field...)
}
func Fatal(msg string, field ...zapcore.Field) {
	logger.Fatal(msg, field...)
}
