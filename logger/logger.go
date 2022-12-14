package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(service string) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncoderConfig.CallerKey = zapcore.OmitKey
	config.DisableStacktrace = true
	config.InitialFields = map[string]any{
		"service": service,
	}

	log, err := config.Build()
	if err != nil {
		return nil, err
	}

	return log.Sugar(), nil
}
