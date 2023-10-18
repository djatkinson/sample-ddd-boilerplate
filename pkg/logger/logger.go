package logger

import (
	"context"
	"go.elastic.co/apm/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitializeLogger() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	Logger, _ = config.Build()
}

func Ctx(ctx context.Context) *zap.Logger {
	log := Logger
	traceID := apm.TransactionFromContext(ctx).TraceContext().Trace
	return log.With(zap.String("trace_id", traceID.String()))
}
