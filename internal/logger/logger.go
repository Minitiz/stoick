package logger

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type correlationIdType int

const (
	requestIdKey correlationIdType = iota
	sessionIdKey
)

var Logger *zap.Logger

func init() {
	atom := zap.NewAtomicLevel()
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	Logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))

}

// Logger returns a zap logger with as much context as possible
func Set(ctx context.Context) *zap.Logger {
	newLogger := Logger
	if ctx != nil {
		if ctxRqId, ok := ctx.Value(requestIdKey).(string); ok {
			newLogger = newLogger.With(zap.String("rqId", ctxRqId))
		}
		if ctxSessionId, ok := ctx.Value(sessionIdKey).(string); ok {
			newLogger = newLogger.With(zap.String("sessionId", ctxSessionId))
		}
	}
	return newLogger
}
